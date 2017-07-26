package internal

/* 1996				wesley ebisuzaki
 *
 * Unpack BDS section
 *
 * input: *bits, pointer to packed integer data
 *        *bitmap, pointer to bitmap (undefined data), NULL if none
 *        n_bits, number of bits per packed integer
 *        n, number of data points (includes undefined data)
 *        ref, scale: flt[] = ref + scale*packed_int
 * output: *flt, pointer to output array
 *        undefined values filled with UNDEFINED
 *
 * note: code assumes an integer >= 32 bits
 *
 * 7/98 v1.2.1 fix bug for bitmaps and nbit >= 25 found by Larry Brasfield
 * 2/01 v1.2.2 changed jj from long int to double
 * 3/02 v1.2.3 added unpacking extensions for spectral data
 *             Luis Kornblueh, MPIfM
 * 7/06 v.1.2.4 fixed some bug complex packed data was not set to undefined
 * 10/15 v.1.2.5 changed n and i to unsigned
 * 3/16 v.1.2.6 OpenMP
 * 6/16 v.1.2.7 faster OpenMP and optimization
 */

var mask = []int{0, 1, 3, 7, 15, 31, 63, 127, 255}
var shift = []double{1.0, 2.0, 4.0, 8.0, 16.0, 32.0, 64.0, 128.0, 256.0}

/*
 void unpk_0(float *flt, unsigned char *bits0, unsigned char *bitmap0,
	 int n_bits, unsigned int n, double ref, double scale, double dec_scale) {

	 unsigned char *bits, *bitmap;

	 int c_bits, j_bits, nthreads;
	 unsigned int map_mask, bbits, i, j, k, n_missing, ndef, di;
	 double jj;

	 ref = ref * dec_scale;
	 scale = scale * dec_scale;
	 bits = bits0;
	 bitmap = bitmap0;

	 bbits = 0;

	 /* assume integer has 32+ bits */
/* optimized code for n_bits <= 25bits /
	 if (n_bits <= 25) {
		 n_missing = bitmap ? missing_points(bitmap0, n) : 0;
	 ndef = n - n_missing;

	 // 1-cpu: rd_bitstream_flt(bits0, 0, flt+n_missing, n_bits, ndef);
	 // 1-cpu: for (j = 0; j < ndef; j++) flt[j+n_missing] = ref + scale*flt[j+n_missing];

 #pragma omp parallel private(i,j,k)
	 {
 #pragma omp single
		 {
			 nthreads = omp_get_num_threads();
			 di = (ndef + nthreads - 1) / nthreads;
				 di = ((di + 7) | 7) ^ 7;
		 }
 #pragma omp for
		 for (i = 0; i < ndef; i += di) {
			 k  = ndef - i;
			 if (k > di) k = di;
			 rd_bitstream_flt(bits0 + (i/8)*n_bits, 0, flt+n_missing+i, n_bits, k);
			 for (j = i+n_missing; j < i+k+n_missing; j++) {
			 flt[j] = ref + scale*flt[j];
			 }
		 }
	 }
 /*
 #pragma omp parallel for private(i,j,k)
	 for (i = 0; i < ndef; i += CACHE_LINE_BITS) {
		 k  = ndef - i;
		 if (k > CACHE_LINE_BITS) k = CACHE_LINE_BITS;
		 rd_bitstream_flt(bits0 + (i/8)*n_bits, 0, flt+n_missing+i, n_bits, k);
		 for (j = i+n_missing; j < i+k+n_missing; j++) {
		 flt[j] = ref + scale*flt[j];
		 }
	 }
 /

	 if (n_missing != 0) {
		 j = n_missing;
		 for (i = 0; i < n; i++) {
		 /* check bitmap /
		 if ((i & 7) == 0) bbits = *bitmap++;
		 if (bbits & 128) {
			 flt[i] = flt[j++];
		 }
		 else {
			 flt[i] = UNDEFINED;
		 }
		 bbits = bbits << 1;
		 }
		 }
	 }
	 else {
	 /* older unoptimized code, not often used /
		 c_bits = 8;
		 map_mask = 128;
		 while (n-- > 0) {
		 if (bitmap) {
			 j = (*bitmap & map_mask);
			 if ((map_mask >>= 1) == 0) {
			 map_mask = 128;
			 bitmap++;
			 }
			 if (j == 0) {
			 *flt++ = UNDEFINED;
			 continue;
			 }
		 }

		 jj = 0.0;
		 j_bits = n_bits;
		 while (c_bits <= j_bits) {
			 if (c_bits == 8) {
			 jj = jj * 256.0  + (double) (*bits++);
			 j_bits -= 8;
			 }
			 else {
			 jj = (jj * shift[c_bits]) + (double) (*bits & mask[c_bits]);
			 bits++;
			 j_bits -= c_bits;
			 c_bits = 8;
			 }
		 }
		 if (j_bits) {
			 c_bits -= j_bits;
			 jj = (jj * shift[j_bits]) + (double) ((*bits >> c_bits) & mask[j_bits]);
		 }
		 *flt++ = ref + scale*jj;
		 }
	 }
	 return;
 }
*/

func unpk_0(flt []float, bits0 []unsigned_char, bitmap0 []unsigned_char, n_bits int, n unsigned_int, ref double, scale double, dec_scale double) error {

	var bits, bitmap []unsigned_char

	var c_bits, j_bits, nthreads int
	var map_mask, bbits, i, j, k, n_missing, ndef, di unsigned_int
	var jj double

	ref = ref * dec_scale
	scale = scale * dec_scale
	//bits = bits0
	bitmap = bitmap0

	bbits = 0

	/* assume integer has 32+ bits */
	/* optimized code for n_bits <= 25bits */
	if n_bits <= 25 {
		// n_missing = bitmap ? missing_points(bitmap0, n) : 0;
		if bitmap != nil {
			n_missing = missing_points(bitmap0, n)
		} else {
			n_missing = 0
		}
		ndef = n - n_missing

		// 1-cpu: rd_bitstream_flt(bits0, 0, flt+n_missing, n_bits, ndef);
		// 1-cpu: for (j = 0; j < ndef; j++) flt[j+n_missing] = ref + scale*flt[j+n_missing];

		//#pragma omp parallel private(i,j,k)
		{
			//#pragma omp single
			{
				nthreads = 1 //omp_get_num_threads();
				di = (ndef + unsigned_int(nthreads) - 1) / unsigned_int(nthreads)
				di = ((di + 7) | 7) ^ 7
			}
			//#pragma omp for
			for i = 0; i < ndef; i += di {
				k = ndef - i
				if k > di {
					k = di
				}
				err := rd_bitstream_flt(bits0[int(i/8)*n_bits:], 0, flt[n_missing+i:], n_bits, int(k))
				if err != nil {
					return fatal_error_wrap(err, "Failed to execute rd_bitstream_flt")
				}
				for j = i + n_missing; j < i+k+n_missing; j++ {
					flt[j] = float(ref + scale*double(flt[j]))
				}
			}
		}
		/*
		   #pragma omp parallel for private(i,j,k)
		    for (i = 0; i < ndef; i += CACHE_LINE_BITS) {
		   	 k  = ndef - i;
		   	 if (k > CACHE_LINE_BITS) k = CACHE_LINE_BITS;
		   	 rd_bitstream_flt(bits0 + (i/8)*n_bits, 0, flt+n_missing+i, n_bits, k);
		   	 for (j = i+n_missing; j < i+k+n_missing; j++) {
		   	 flt[j] = ref + scale*flt[j];
		   	 }
		    }
		*/

		if n_missing != 0 {
			j = n_missing
			bitmap_index := 0
			for i = 0; i < n; i++ {
				/* check bitmap */
				if (i & 7) == 0 {
					bbits = unsigned_int(bitmap[bitmap_index])
					bitmap_index++
				}
				if (bbits & 128) != 0 {
					flt[i] = flt[j]
					j++
				} else {
					flt[i] = UNDEFINED
				}
				bbits = bbits << 1
			}
		}
	} else {
		/* older unoptimized code, not often used */
		c_bits = 8
		map_mask = 128
		flt_index := 0
		bitmap_index := 0
		bits_index := 0
		for n > 0 {
			n--
			if bitmap != nil {
				j = unsigned_int(bitmap[bitmap_index]) & map_mask
				//if ((map_mask >>= 1) == 0) {
				map_mask = map_mask >> 1
				if map_mask == 0 {
					map_mask = 128
					bitmap_index++
				}
				if j == 0 {
					flt[flt_index] = UNDEFINED
					flt_index++
					continue
				}
			}

			jj = 0.0
			j_bits = n_bits
			for c_bits <= j_bits {
				if c_bits == 8 {
					jj = jj*256.0 + double(bits[bits_index])
					bits_index++
					j_bits -= 8
				} else {
					jj = (jj * shift[c_bits]) + double(int(bits[bits_index])&mask[c_bits])
					bits_index++
					j_bits -= c_bits
					c_bits = 8
				}
			}
			if j_bits != 0 {
				c_bits -= j_bits
				jj = (jj * shift[j_bits]) + double((int(bits[bits_index])>>unsigned_int(c_bits))&mask[j_bits])
			}
			flt[flt_index] = float(ref + scale*jj)
			flt_index++
		}
	}
	return nil
}
