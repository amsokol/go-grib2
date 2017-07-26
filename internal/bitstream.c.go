package internal

var ones = []int{0, 1, 3, 7, 15, 31, 63, 127, 255}

/*
 * void rd_bitstream_flt
 *   rd_bitstream_flt() is like rd_bitstream() except that returns a float instead of int
 */
func rd_bitstream_flt(p []unsigned_char, offset int, u []float, n_bits int, n int) error {

	var tbits unsigned_int
	var i, t_bits, new_t_bits int

	// not the best of tests

	if INT_MAX <= 2147483647 && n_bits > 31 {
		return fatal_error_i("rd_bitstream: n_bits is %d", n_bits)
	}

	if offset < 0 || offset > 7 {
		return fatal_error_i("rd_bitstream_flt: illegal offset %d", offset)
	}

	if n_bits == 0 {
		for i = 0; i < n; i++ {
			u[i] = 0.0
		}
		return nil
	}

	t_bits = 8 - offset
	p_index := 0
	tbits = unsigned_int(p[p_index]) & unsigned_int(ones[t_bits])
	p_index++

	for i = 0; i < n; i++ {

		for n_bits-t_bits >= 8 {
			t_bits += 8
			tbits = (tbits << 8) | unsigned_int(p[p_index])
			p_index++
		}

		if n_bits > t_bits {
			new_t_bits = 8 - (n_bits - t_bits)
			u[i] = float(int(tbits<<unsigned_int(n_bits-t_bits) | (unsigned_int(p[p_index]) >> unsigned_int(new_t_bits))))
			t_bits = new_t_bits
			tbits = unsigned_int(p[p_index]) & unsigned_int(ones[t_bits])
			p_index++
		} else if n_bits == t_bits {
			u[i] = float(tbits)
			tbits = 0
			t_bits = 0
		} else {
			t_bits -= n_bits
			u[i] = float(tbits >> unsigned_int(t_bits))
			tbits = tbits & unsigned_int(ones[t_bits])
		}
	}
	return nil
}
