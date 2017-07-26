package internal

func stat_proc_verf_time(sec [][]unsigned_char, year *int, month *int, day *int, hour *int, minute *int, second *int) int {
	var stat_proc_time []unsigned_char

	stat_proc_time = stat_proc_verf_time_location(sec)

	if stat_proc_time != nil {
		get_time(stat_proc_time, year, month, day, hour, minute, second)
		return 0
	}
	*year = 0
	*month = 0
	*day = 0
	*hour = 0
	*minute = 0
	*second = 0

	return 1
}

/*
 * returns location of statistical time processing section
 *  ie location of overall time
 *
 */
func stat_proc_verf_time_location(sec [][]unsigned_char) []unsigned_char {
	var i, j, center, nb int
	i = code_table_4_0(sec)
	center = GB2_Center(sec)
	j = 0
	if i == 8 {
		j = 34
	} else if i == 9 {
		j = 47
	} else if i == 10 {
		j = 35
	} else if i == 11 {
		j = 37
	} else if i == 12 {
		j = 36
	} else if i == 13 {
		j = 68
	} else if i == 14 {
		j = 64
	} else if i == 34 {
		nb = int(sec[4][22])
		j = 26 + 11*nb
	} else if i == 42 {
		j = 36
	} else if i == 43 {
		j = 39
	} else if i == 46 {
		j = 47
	} else if i == 47 {
		j = 50
	} else if i == 61 {
		j = 44
	} else if i == 50008 && (center == JMA1 || center == JMA2) {
		j = 34
	}
	if j == 0 {
		return nil
	}
	return sec[4][j:]
}

/*
 * returns number of particle size distributions used by template 4.57
 */
func number_of_mode(sec [][]unsigned_char) int {
	var pdt int
	pdt = code_table_4_0(sec)
	if pdt == 57 {
		return int(uint2(sec[4][13:]))
	}
	return -1
}

/*
 * forecast_time_in_units
 *
 * v1.1 4/2015:  allow forecast time to be a signed quantity
 *      old: return unsigned value, ! code_4_4, return 0xffffffff;
 *      new: return signed value    ! code_4_4, return 0
 */
func forecast_time_in_units(sec [][]unsigned_char) int {

	var code_4_4 []unsigned_char
	var pdt int

	code_4_4 = code_table_4_4_location(sec)
	if code_4_4 != nil {
		// silly WMO codes group
		pdt = code_table_4_0(sec)
		if pdt != 44 {
			return int4(code_4_4[1:])
		}
		return int2(code_4_4[1:])
	}
	return 0
	// return 0xffffffff;
}

func fixed_surfaces(sec [][]unsigned_char, type1 *int, surface1 *float, undef_val1 *int, type2 *int, surface2 *float, undef_val2 *int) error {

	var p1, p2 []unsigned_char
	*undef_val1 = 1
	*undef_val2 = 1
	*surface1 = UNDEFINED
	*surface2 = UNDEFINED
	*type1 = 255
	*type2 = 255

	p1, err := code_table_4_5a_location(sec)
	if err != nil {
		fatal_error_wrap(err, "Failed to execute code_table_4_5a_location")
	}
	p2, err = code_table_4_5b_location(sec)
	if err != nil {
		fatal_error_wrap(err, "Failed to execute code_table_4_5b_location")
	}

	if p1 != nil && p1[0] != 255 {
		*type1 = int(p1[0])
		if p1[1] != 255 {
			if p1[2] != 255 || p1[3] != 255 || p1[4] != 255 || p1[5] != 255 {
				*undef_val1 = 0
				*surface1 = scaled2flt(INT1(p1[1]), int4(p1[2:]))
			}
		}
	}
	if p2 != nil && p2[0] != 255 {
		*type2 = int(p2[0])
		if p2[1] != 255 {
			if p2[2] != 255 || p2[3] != 255 || p2[4] != 255 || p2[5] != 255 {
				*undef_val2 = 0
				*surface2 = scaled2flt(INT1(p2[1]), int4(p2[2:]))
			}
		}
	}
	return nil
}
