package internal

//// 2009 public domain wesley ebisuzaki
////
//// note: assumption that the grib file will use 25 bits or less for storing data
////       (limit of bitstream unpacking routines)
//// note: assumption that all data can be stored as integers and have a value < INT_MAX
//
//int unpk_complex(unsigned char **sec, float *data, unsigned int ndata) {
//
//unsigned int j, n;
//int i, k, nbits, ref_group_length;
//unsigned char *p, *d, *mask_pointer;
//double ref_val,factor_10, factor_2, factor;
//float missing1, missing2;
//int n_sub_missing;
//int pack, offset;
//unsigned clocation;
//unsigned int ngroups, ref_group_width, nbit_group_width, len_last, npnts;
//int nbits_group_len, group_length_factor;
//int *group_refs, *group_widths, *group_lengths, *group_location, *group_offset, *udata;
//unsigned int *group_clocation;
//
//int m1, m2, mask, last, penultimate;
//int extra_vals[2];
//int min_val;
//int ctable_5_4, ctable_5_6,  bitmap_flag, extra_octets;
//
//
//extra_vals[0] = extra_vals[1] = 0;
//pack = code_table_5_0(sec);
//if (pack != 2 && pack != 3) return 0;
//
//p = sec[5];
//ref_val = ieee2flt(p+11);
//factor_2 = Int_Power(2.0, int2(p+15));
//factor_10 = Int_Power(10.0, -int2(p+17));
//ref_val *= factor_10;
//factor = factor_2 * factor_10;
//nbits = p[19];
//ngroups = uint4(p+31);
//bitmap_flag = code_table_6_0(sec);
//ctable_5_6 = code_table_5_6(sec);
//
//if (pack == 3 && (ctable_5_6 != 1 && ctable_5_6 != 2))
//fatal_error_i("unsupported: code table 5.6=%d", ctable_5_6);
//
//extra_octets = (pack == 2) ? 0 : sec[5][48];
//
//if (ngroups == 0) {
//if (bitmap_flag == 255) {
//for (i = 0; i < ndata; i++) data[i] = ref_val;
//return 0;
//}
//if (bitmap_flag == 0 || bitmap_flag == 254) {
//mask_pointer = sec[6] + 6;
//mask = 0;
//for (i = 0; i < ndata; i++) {
//if ((i & 7) == 0) mask = *mask_pointer++;
//data[i] = (mask & 128) ?  ref_val : UNDEFINED;
//mask <<= 1;
//}
//return 0;
//}
//fatal_error("unknown bitmap", "");
//}
//
//ctable_5_4 = code_table_5_4(sec);
//ref_group_width = p[35];
//nbit_group_width = p[36];
//ref_group_length = uint4(p+37);
//group_length_factor = p[41];
//len_last = uint4(p+42);
//nbits_group_len = p[46];
//
//npnts =  GB2_Sec5_nval(sec); 	// number of defined points
//n_sub_missing = sub_missing_values(sec, &missing1, &missing2);
//
//// allocate group widths and group lengths
//group_refs = (int *) malloc(ngroups * sizeof (unsigned int));
//group_widths = (int *) malloc(ngroups * sizeof (unsigned int));
//group_lengths = (int *) malloc(ngroups * sizeof (unsigned int));
//group_location = (int *) malloc(ngroups * sizeof (unsigned int));
//group_clocation = (unsigned int *) malloc(ngroups * sizeof (unsigned int));
//group_offset = (int *) malloc(ngroups * sizeof (unsigned int));
//udata = (int *) malloc(npnts * sizeof (unsigned int));
//if (group_refs == NULL || group_widths == NULL || group_lengths ==
//NULL || udata == NULL) fatal_error("com unpack error","");
//
//// read any extra values
//d = sec[7]+5;
//min_val = 0;
//if (extra_octets) {
//extra_vals[0] = uint_n(d,extra_octets);
//d += extra_octets;
//if (ctable_5_6 == 2) {
//extra_vals[1] = uint_n(d,extra_octets);
//d += extra_octets;
//}
//min_val = int_n(d,extra_octets);
//d += extra_octets;
//}
//
//if (ctable_5_4 != 1) fatal_error_i("internal decode does not support code table 5.4=%d",
//ctable_5_4);
//
//#pragma omp parallel
//{
//#pragma omp sections
//{
//
//
//#pragma omp section
//{
//// read the group reference values
//rd_bitstream(d, 0, group_refs, nbits, ngroups);
//}
//
//
//#pragma omp section
//{
//int i;
//// read the group widths
//
//rd_bitstream(d+(nbits*ngroups+7)/8,0,group_widths,nbit_group_width,ngroups);
//for (i = 0; i < ngroups; i++) group_widths[i] += ref_group_width;
//}
//
//
//#pragma omp section
//{
//int i;
//// read the group lengths
//
//if (ctable_5_4 == 1) {
//rd_bitstream(d+(nbits*ngroups+7)/8+(ngroups*nbit_group_width+7)/8,
//0,group_lengths, nbits_group_len, ngroups-1);
//
//for (i = 0; i < ngroups-1; i++) {
//group_lengths[i] = group_lengths[i] * group_length_factor + ref_group_length;
//}
//group_lengths[ngroups-1] = len_last;
//}
//}
//
//}
//
//
//#pragma omp single
//{
//d += (nbits*ngroups + 7)/8 +
//(ngroups * nbit_group_width + 7) / 8 +
//(ngroups * nbits_group_len + 7) / 8;
//
//// do a check for number of grid points and size
//clocation = offset = n = j = 0;
//}
//
//#pragma omp sections
//{
//
//
//#pragma omp section
//{
//int i;
//for (i = 0; i < ngroups; i++) {
//group_location[i] = j;
//j += group_lengths[i];
//n += group_lengths[i]*group_widths[i];
//}
//}
//
//#pragma omp section
//{
//int i;
//for (i = 0; i < ngroups; i++) {
//group_clocation[i] = clocation;
//clocation = clocation + group_lengths[i]*(group_widths[i]/8) +
//(group_lengths[i]/8)*(group_widths[i] % 8);
//}
//}
//
//#pragma omp section
//{
//int i;
//for (i = 0; i < ngroups; i++) {
//group_offset[i] = offset;
//offset += (group_lengths[i] % 8)*(group_widths[i] % 8);
//}
//}
//}
//}
//
//if (j != npnts) fatal_error_i("bad complex packing: n points %d",j);
//if (d + (n+7)/8 - sec[7] != GB2_Sec7_size(sec))
//fatal_error("complex unpacking size mismatch old test","");
//
//
//if (d + clocation + (offset + 7)/8 - sec[7] != GB2_Sec7_size(sec)) fatal_error("complex unpacking size mismatch","");
//
//#pragma omp parallel for private(i) schedule(static)
//for (i = 0; i < ngroups; i++) {
//group_clocation[i] += (group_offset[i] / 8);
//group_offset[i] = (group_offset[i] % 8);
//
//rd_bitstream(d + group_clocation[i], group_offset[i], udata+group_location[i],
//group_widths[i], group_lengths[i]);
//}
//
//// handle substitute, missing values and reference value
//if (n_sub_missing == 0) {
//#pragma omp parallel for private(i,k,j)
//for (i = 0; i < ngroups; i++) {
//j = group_location[i];
//for (k = 0; k < group_lengths[i]; k++) {
//udata[j++] += group_refs[i];
//}
//}
//}
//else if (n_sub_missing == 1) {
//
//#pragma omp parallel for private(i,m1,k,j)
//for (i = 0; i < ngroups; i++) {
//j = group_location[i];
//if (group_widths[i] == 0) {
//m1 = (1 << nbits) - 1;
//if (m1 == group_refs[i]) {
//for (k = 0; k < group_lengths[i]; k++) udata[j++] = INT_MAX;
//}
//else {
//for (k = 0; k < group_lengths[i]; k++) udata[j++] += group_refs[i];
//}
//}
//else {
//m1 = (1 << group_widths[i]) - 1;
//for (k = 0; k < group_lengths[i]; k++) {
//if (udata[j] == m1) udata[j] = INT_MAX;
//else udata[j] += group_refs[i];
//j++;
//}
//}
//}
//}
//else if (n_sub_missing == 2) {
//#pragma omp parallel for private(i,j,k,m1,m2)
//for (i = 0; i < ngroups; i++) {
//j = group_location[i];
//if (group_widths[i] == 0) {
//m1 = (1 << nbits) - 1;
//m2 = m1 - 1;
//if (m1 == group_refs[i] || m2 == group_refs[i]) {
//for (k = 0; k < group_lengths[i]; k++) udata[j++] = INT_MAX;
//}
//else {
//for (k = 0; k < group_lengths[i]; k++) udata[j++] += group_refs[i];
//}
//}
//else {
//m1 = (1 << group_widths[i]) - 1;
//m2 = m1 - 1;
//for (k = 0; k < group_lengths[i]; k++) {
//if (udata[j] == m1 || udata[j] == m2) udata[j] = INT_MAX;
//else udata[j] += group_refs[i];
//j++;
//}
//}
//}
//}
//
//// post processing
//
//if (pack == 3) {
//if (ctable_5_6 == 1) {
//last = extra_vals[0];
//i = 0;
//while (i < npnts) {
//if (udata[i] == INT_MAX) i++;
//else {
//udata[i++] = extra_vals[0];
//break;
//}
//}
//while (i < npnts) {
//if (udata[i] == INT_MAX) i++;
//else {
//udata[i] += last + min_val;
//last = udata[i++];
//}
//}
//}
//else if (ctable_5_6 == 2) {
//penultimate = extra_vals[0];
//last = extra_vals[1];
//
//i = 0;
//while (i < npnts) {
//if (udata[i] == INT_MAX) i++;
//else {
//udata[i++] = extra_vals[0];
//break;
//}
//}
//while (i < npnts) {
//if (udata[i] == INT_MAX) i++;
//else {
//udata[i++] = extra_vals[1];
//break;
//}
//}
//for (; i < npnts; i++) {
//if (udata[i] != INT_MAX) {
//udata[i] =  udata[i] + min_val + last + last - penultimate;
//penultimate = last;
//last = udata[i];
//}
//}
//}
//else fatal_error_i("Unsupported: code table 5.6=%d", ctable_5_6);
//}
//
//// convert to float
//
//if (bitmap_flag == 255) {
//#pragma omp parallel for schedule(static) private(i)
//for (i = 0; i < (int) ndata; i++) {
//data[i] = (udata[i] == INT_MAX) ? UNDEFINED :
//ref_val + udata[i] * factor;
//}
//}
//else if (bitmap_flag == 0 || bitmap_flag == 254) {
//n = 0;
//mask = 0;
//mask_pointer = sec[6] + 6;
//for (i = 0; i < ndata; i++) {
//if ((i & 7) == 0) mask = *mask_pointer++;
//if (mask & 128) {
//if (udata[n] == INT_MAX) data[i] = UNDEFINED;
//else data[i] = ref_val + udata[n] * factor;
//n++;
//}
//else data[i] = UNDEFINED;
//mask <<= 1;
//}
//}
//else fatal_error_i("unknown bitmap: %d", bitmap_flag);
//
//free(group_refs);
//free(group_widths);
//free(group_lengths);
//free(group_location);
//free(group_clocation);
//free(group_offset);
//free(udata);
//
//return 0;
//}
func unpk_complex(sec [][]unsigned_char, data []float, ndata unsigned_int) error {

	var extra_vals [2]unsigned_int
	extra_vals[0] = 0
	extra_vals[1] = 0
	pack := code_table_5_0(sec)
	if pack != 2 && pack != 3 {
		return nil
	}

	p := sec[5]
	ref_val := double(ieee2flt(p[11:]))
	factor_2 := Int_Power(2.0, int2(p[15:]))
	factor_10 := Int_Power(10.0, -int2(p[17:]))
	ref_val *= factor_10
	factor := factor_2 * factor_10
	nbits := int(p[19])
	ngroups := int(ieee2flt(p[31:]))

	bitmap_flag := code_table_6_0(sec)
	ctable_5_6 := code_table_5_6(sec)

	if pack == 3 && (ctable_5_6 != 1 && ctable_5_6 != 2) {
		_ = fatal_error_i("unsupported: code table 5.6=%d", ctable_5_6)
	}

	var extra_octets unsigned_char
	if pack == 2 {
		extra_octets = 0
	} else {
		extra_octets = sec[5][48]
	}

	if ngroups == 0 {
		if bitmap_flag == 255 {
			for i := 0; i < int(ndata); i++ {
				data[i] = float(ref_val)
			}
			return nil
		}

		if bitmap_flag == 0 || bitmap_flag == 254 {
			mask_pointer := sec[6][6:]
			var mask unsigned_char
			mask_pointer_index := 0
			for i := 0; i < int(ndata); i++ {
				if (i & 7) == 0 {
					mask = mask_pointer[mask_pointer_index]
					mask_pointer_index++
				}
				if (mask & 128) != 0 {
					data[i] = float(ref_val)
				} else {
					data[i] = UNDEFINED
				}
				mask <<= 1
			}
			return nil
		}

		_ = fatal_error_i("unknown bitmap")
	}

	ctable_5_4 := code_table_5_4(sec)
	ref_group_width := unsigned_int(p[35])
	nbit_group_width := unsigned_int(p[36])
	ref_group_length := unsigned_int(ieee2flt(p[37:]))
	group_length_factor := unsigned_int(p[41])
	len_last := unsigned_int(ieee2flt(p[42:]))
	nbits_group_len := unsigned_int(p[46])

	npnts := GB2_Sec5_nval(sec) // number of defined points

	var missing1, missing2 float
	n_sub_missing := sub_missing_values(sec, &missing1, &missing2)

	// allocate group widths and group lengths

	//group_refs = (int *) malloc(ngroups * sizeof (unsigned int))
	//group_widths = (int *) malloc(ngroups * sizeof (unsigned int))
	//group_lengths = (int *) malloc(ngroups * sizeof (unsigned int))
	//group_location = (int *) malloc(ngroups * sizeof (unsigned int))
	//group_clocation = (unsigned int *) malloc(ngroups * sizeof (unsigned int))
	//group_offset = (int *) malloc(ngroups * sizeof (unsigned int))
	//udata = (int *) malloc(npnts * sizeof (unsigned int))
	//if (group_refs == NULL || group_widths == NULL || group_lengths ==
	//NULL || udata == NULL) fatal_error("com unpack error","")
	//

	//var group_refs, group_widths, group_lengths, group_location, group_offset, udata []int
	//	var group_clocation []unsigned_int

	//// read any extra values
	d := sec[7][5:]
	min_val := 0
	if extra_octets != 0 {
		extra_vals[0] = uint_n(d, int(extra_octets))
		d = append(d, extra_octets)
		if ctable_5_6 == 2 {
			extra_vals[1] = uint_n(d, int(extra_octets))
			d = append(d, extra_octets)
		}
		min_val = int_n(d, int(extra_octets))
		d = append(d, extra_octets)
	}

	if ctable_5_4 != 1 {
		_ = fatal_error_i("internal decode does not support code table 5.4=%d", ctable_5_4)
	}

	group_refs := make([]int, ngroups)
	group_location := make([]int, ngroups)
	group_lengths := make([]int, ngroups)
	group_widths := make([]int, ngroups)
	group_clocation := make([]int, ngroups)
	group_offset := make([]int, ngroups)
	j := 0
	n := 0
	clocation := 0
	offset := 0
	//	#pragma omp parallel
	{
		//	#pragma omp sections
		{
			//	#pragma omp section
			{
				// read the group reference values
				err := rd_bitstream(d, 0, group_refs, int(nbits), int(ngroups))
				if err != nil {
					return err
				}
			}

			//	#pragma omp section
			{
				// read the group widths
				dp := append(d, unsigned_char(nbits*ngroups+7)/8)
				err := rd_bitstream(dp, 0, group_widths, int(nbit_group_width), int(ngroups))
				if err != nil {
					return err
				}

				for i := 0; i < ngroups; i++ {
					group_widths[i] += int(ref_group_width)
				}
			}

			//	#pragma omp section
			{
				// read the group lengths

				if ctable_5_4 == 1 {
					dp54 := append(d, unsigned_char(nbits*ngroups+7)/8)
					dp54 = append(dp54, unsigned_char(unsigned_int(ngroups)*nbit_group_width+7)/8)

					err := rd_bitstream(dp54, 0, group_lengths, int(nbits_group_len), int(ngroups-1))
					if err != nil {
						return err
					}
					for i := 0; i < ngroups-1; i++ {
						group_lengths[i] = group_lengths[i]*int(group_length_factor) + int(ref_group_length)
					}
					group_lengths[ngroups-1] = int(len_last)
				}
			}

		}
	}

	//#pragma omp single
	{

		gappend := (unsigned_int(nbits)*unsigned_int(ngroups)+7)/8 +
			(unsigned_int(ngroups)*nbit_group_width+7)/8 +
			(unsigned_int(ngroups)*nbits_group_len+7)/8

		d = append(d, unsigned_char(gappend))
		// do a check for number of grid points and size
	}

	//	#pragma omp sections
	{
		//	#pragma omp section
		{
			for i := 0; i < int(ngroups); i++ {
				group_location[i] = j
				j += group_lengths[i]
				n += group_lengths[i] * group_widths[i]
			}
		}

		//	#pragma omp section
		{

			for i := 0; i < int(ngroups); i++ {
				group_clocation[i] = clocation
				clocation = clocation + group_lengths[i]*(group_widths[i]/8) +
					(group_lengths[i]/8)*(group_widths[i]%8)
			}
		}

		//	#pragma omp section
		{

			for i := 0; i < ngroups; i++ {
				group_offset[i] = offset
				offset += (group_lengths[i] % 8) * (group_widths[i] % 8)
			}
		}
	}

	if j != int(npnts) {
		return fatal_error("bad complex packing: n points %d", j)
	}
	//SY: TODO
	//if (d + (n+7)/8 - sec[7] != GB2_Sec7_size(sec)){
	//	return fatal_error("complex unpacking size mismatch old test","")
	//}

	//if (d + clocation + (offset + 7)/8 - sec[7] != GB2_Sec7_size(sec)) fatal_error("complex unpacking size mismatch","")
	//

	udata := make([]int, npnts)
	//#pragma omp parallel for private(i) schedule(static)
	for i := 0; i < ngroups; i++ {
		group_clocation[i] += group_offset[i] / 8
		group_offset[i] = group_offset[i] % 8

		dp := append(d, unsigned_char(group_clocation[i]))
		udp := append(udata, group_location[i])
		err := rd_bitstream(dp, group_offset[i], udp, group_widths[i], group_lengths[i])
		if err != nil {
			return err
		}
	}

	// handle substitute, missing values and reference value
	if n_sub_missing == 0 {
		//	#pragma omp parallel for private(i,k,j)
		for i := 0; i < ngroups; i++ {
			j := group_location[i]
			for k := 0; k < group_lengths[i]; k++ {
				j++
				udata[j] += group_refs[i]
			}
		}
	}

	//else if
	if n_sub_missing == 1 {

		//#pragma omp parallel for private(i,m1,k,j)
		for i := 0; i < ngroups; i++ {
			j = group_location[i]
			if group_widths[i] == 0 {
				m1 := (1 << nbits) - 1
				if m1 == group_refs[i] {
					for k := 0; k < group_lengths[i]; k++ {
						j++
						udata[j] = INT_MAX
					}
				} else {
					for k := 0; k < group_lengths[i]; k++ {
						j++
						udata[j] += group_refs[i]
					}
				}
			} else {
				m1 := (1 << group_widths[i]) - 1
				for k := 0; k < group_lengths[i]; k++ {
					if udata[j] == m1 {
						udata[j] = INT_MAX
					} else {
						udata[j] += group_refs[i]
					}
					j++
				}
			}
		}
	}

	//else if
	if n_sub_missing == 2 {
		//#pragma omp parallel for private(i,j,k,m1,m2)
		for i := 0; i < ngroups; i++ {
			j = group_location[i]
			if group_widths[i] == 0 {
				m1 := (1 << nbits) - 1
				m2 := m1 - 1
				if m1 == group_refs[i] || m2 == group_refs[i] {
					for k := 0; k < group_lengths[i]; k++ {
						j++
						udata[j] = INT_MAX
					}
				} else {
					for k := 0; k < group_lengths[i]; k++ {
						j++
						udata[j] += group_refs[i]
					}
				}
			} else {
				m1 := (1 << group_widths[i]) - 1
				m2 := m1 - 1
				for k := 0; k < group_lengths[i]; k++ {
					if udata[j] == m1 || udata[j] == m2 {
						udata[j] = INT_MAX
					} else {
						udata[j] += group_refs[i]
					}
					j++
				}
			}
		}
	}

	// post processing

	if pack == 3 {
		if ctable_5_6 == 1 {
			last := extra_vals[0]
			i := 0
			for i < int(npnts) {
				if udata[i] == INT_MAX {
					i++
				} else {
					i++
					udata[i] = int(extra_vals[0])
					break
				}
			}

			for i < int(npnts) {
				if udata[i] == INT_MAX {
					i++
				} else {
					udata[i] += int(last) + min_val
					i++
					last = unsigned_int(udata[i])
				}
			}
		}

		if ctable_5_6 == 2 {
			penultimate := extra_vals[0]
			last := extra_vals[1]

			i := 0
			for i < int(npnts) {
				if udata[i] == INT_MAX {
					i++
				} else {
					i++
					udata[i] = int(extra_vals[0])
					break
				}
			}
			for i < int(npnts) {
				if udata[i] == INT_MAX {
					i++
				} else {
					i++
					udata[i] = int(extra_vals[1])
					break
				}
			}

			for i := 0; i < int(npnts); i++ {
				if udata[i] != INT_MAX {
					udata[i] = udata[i] + min_val + int(last) + int(last) - int(penultimate)
					penultimate = last
					last = unsigned_int(udata[i])
				}
			}
		} else {
			_ = fatal_error_i("Unsupported: code table 5.6=%d", ctable_5_6)
		}
	}

	// convert to float

	if bitmap_flag == 255 {
		//#pragma omp parallel for schedule(static) private(i)
		for i := 0; i < int(ndata); i++ {
			if udata[i] == INT_MAX {
				data[i] = UNDEFINED
			} else {
				data[i] = float(ref_val + double(udata[i])*factor)
			}
		}
	}

	if bitmap_flag == 0 || bitmap_flag == 254 {
		mask_pointer := sec[6][6:]
		var mask unsigned_char
		mask_pointer_index := 0
		for ii := 0; ii < int(ndata); ii++ {
			if (ii & 7) == 0 {
				mask = mask_pointer[mask_pointer_index]
				mask_pointer_index++
			}
			if (mask & 128) != 0 {
				if udata[n] == INT_MAX {
					data[ii] = UNDEFINED
				} else {
					data[ii] = float(ref_val + double(udata[n])*factor)
				}
			} else {
				data[ii] = UNDEFINED
			}
			mask <<= 1
		}
		return nil
	} else {
		_ = fatal_error_i("unknown bitmap: %d", bitmap_flag)
	}

	return nil
}
