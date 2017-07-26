package internal

/*
 * unpack grib -- only some formats (code table 5.0) are supported
 *
 * supported: 0 (simple), 4 (ieee), 40 (jpeg), 41(png), 42(aec)
 *
 * input:  sec[]
 *         float data[npnts]
 *
 */
func unpk_grib(sec [][]unsigned_char, data []float) error {

	var packing, bitmap_flag, nbits int
	var ndata, ii unsigned_int
	var mask_pointer []unsigned_char
	var mask unsigned_char
	var ieee, p []unsigned_char
	var tmp float
	// float reference, tmp;
	var reference double
	var bin_scale, dec_scale, b double

	packing = code_table_5_0(sec)
	// ndata = (int) GB2_Sec3_npts(sec);
	ndata = GB2_Sec3_npts(sec)
	bitmap_flag = code_table_6_0(sec)

	if bitmap_flag != 0 && bitmap_flag != 254 && bitmap_flag != 255 {
		return fatal_error("unknown bitmap", "")
	}

	if packing == 4 { // ieee
		if sec[5][11] != 1 {
			return fatal_error_i("unpk ieee grib file precision %d not supported", int(sec[5][11]))
		}

		// ieee depacking -- simple no bitmap
		if bitmap_flag == 255 {
			for ii = 0; ii < ndata; ii++ {
				data[ii] = ieee2flt_nan(sec[7][5+ii*4:])
			}
			return nil
		}
		if bitmap_flag == 0 || bitmap_flag == 254 {
			mask_pointer = sec[6][6:]
			ieee = sec[7][5:]
			ieee_index := 0
			mask = 0
			mask_pointer_index := 0
			for ii = 0; ii < ndata; ii++ {
				if (ii & 7) == 0 {
					mask = mask_pointer[mask_pointer_index]
					mask_pointer_index++
				}
				if (mask & 128) != 0 {
					data[ii] = ieee2flt_nan(ieee[ieee_index:])
					ieee_index += 4
				} else {
					data[ii] = UNDEFINED
				}
				mask <<= 1
			}
			return nil
		}
		return fatal_error("unknown bitmap", "")
	} else if packing == 0 || packing == 61 { // simple grib1 packing  61 -- log preprocessing

		p = sec[5]
		reference = double(ieee2flt(p[11:]))
		bin_scale = Int_Power(2.0, int2(p[15:]))
		dec_scale = Int_Power(10.0, -int2(p[17:]))
		nbits = int(p[19])
		b = 0.0
		if packing == 61 {
			b = double(ieee2flt(p[20:]))
		}

		if bitmap_flag != 0 && bitmap_flag != 254 && bitmap_flag != 255 {
			return fatal_error("unknown bitmap", "")
		}

		if nbits == 0 {
			tmp = float(reference * dec_scale)
			if packing == 61 {
				tmp = float(exp(double(tmp)) - b)
			} // remove log prescaling
			if bitmap_flag == 255 {
				for ii = 0; ii < ndata; ii++ {
					data[ii] = tmp
				}
				return nil
			}
			if bitmap_flag == 0 || bitmap_flag == 254 {
				mask_pointer = sec[6][6:]
				mask = 0
				mask_pointer_index := 0
				for ii = 0; ii < ndata; ii++ {
					if (ii & 7) == 0 {
						mask = mask_pointer[mask_pointer_index]
						mask_pointer_index++
					}
					// data[ii] = (mask & 128) ?  tmp : UNDEFINED;
					if (mask & 128) != 0 {
						data[ii] = tmp
					} else {
						data[ii] = UNDEFINED
					}
					mask <<= 1
				}
				return nil
			}
		}

		// mask_pointer = (bitmap_flag == 255) ? NULL : sec[6] + 6;
		if bitmap_flag == 255 {
			mask_pointer = nil
		} else {
			mask_pointer = sec[6][6:]
		}

		unpk_0(data, sec[7][5:], mask_pointer, nbits, ndata, reference,
			bin_scale, dec_scale)

		if packing == 61 { // remove log prescaling
			// #pragma omp parallel for private(ii) schedule(static)
			for ii = 0; ii < ndata; ii++ {
				if DEFINED_VAL(data[ii]) {
					data[ii] = float(exp(double(data[ii])) - b)
				}
			}
		}
		return nil
	} else if packing == 2 || packing == 3 { // complex
		return fatal_error("unpk_complex is not supported")
		// TODO: unpk_complex
		// return unpk_complex(sec, data, ndata)
	} else if packing == 200 { // run length
		return fatal_error("unpk_run_length is not supported")
		// TODO: unpk_run_length
		// return unpk_run_length(sec, data, ndata)
	}
	return fatal_error_i("packing type %d not supported", packing)
}
