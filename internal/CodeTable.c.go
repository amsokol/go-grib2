package internal

func code_table_1_2(sec [][]unsigned_char) int {
	var p []unsigned_char
	p = code_table_1_2_location(sec)
	if p == nil {
		return -1
	}
	return int(p[0])
}

func code_table_1_2_location(sec [][]unsigned_char) []unsigned_char {
	return sec[1][11:]
}

/*
int code_table_3_1(unsigned char **sec) {
    return  (int) uint2(sec[3]+12);
}
*/
func code_table_3_1(sec [][]unsigned_char) int {
	return int(uint2(sec[3][12:]))
}

func code_table_3_2_location(sec [][]unsigned_char) []unsigned_char {
	var grid_def, center int
	grid_def = code_table_3_1(sec)

	if grid_def < 50 {
		return sec[3][14:]
	}

	switch grid_def {
	case 90, 110, 130, 140, 204, 1000, 1100:
		return sec[3][14:]
	default:
	}

	center = GB2_Center(sec)
	if center == NCEP {
		if grid_def == 32768 || (grid_def == 32769) {
			return sec[3][14:]
		}
	}
	if (center == JMA1) || (center == JMA2) {
		if grid_def == 40110 {
			return sec[3][14:]
		}
	}
	return nil
}

func code_table_4_0(sec [][]unsigned_char) int {
	return GB2_ProdDefTemplateNo(sec)
}

func code_table_4_4(sec [][]unsigned_char) int {
	var p []unsigned_char
	p = code_table_4_4_location(sec)
	if p == nil {
		return -1
	}
	return int(p[0])
}

func code_table_4_4_location(sec [][]unsigned_char) []unsigned_char {
	var pdt, center, n int
	pdt = GB2_ProdDefTemplateNo(sec)
	center = GB2_Center(sec)

	switch pdt {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 32, 60, 61, 51, 91, 1000, 1001, 1002:
		return sec[4][17:]
	case 40, 41, 42, 43:
		return sec[4][19:]
	case 44, 45, 46, 47:
		return sec[4][30:]
	case 48:
		return sec[4][41:]
	case 52:
		return sec[4][20:]
	case 57:
		n = number_of_mode(sec)
		if n <= 0 || n == 65535 {
			return nil
		}
		return sec[4][26+5*n:]
	case 50008:
		if center == JMA1 || center == JMA2 {
			return sec[4][17:]
		}
		return nil
	case 50009:
		if center == JMA1 || center == JMA2 {
			return sec[4][17:]
		}
		return nil
	case 50011:
		if center == JMA1 || center == JMA2 {
			return sec[4][17:]
		}
		return nil
	}
	return nil
}

func code_table_4_5a_location(sec [][]unsigned_char) ([]unsigned_char, error) {
	var pdt, center, n int
	pdt = GB2_ProdDefTemplateNo(sec)
	center = GB2_Center(sec)

	if center == JMA1 || center == JMA2 {
		switch pdt {
		case 50008:
			return sec[4][22:], nil
		case 50009:
			return sec[4][22:], nil
		case 51020, 51021, 51022, 51122, 52020:
			return nil, nil
		case 50010, 50011:
			return sec[4][22:], nil
		default:
		}
	}

	switch pdt {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 51, 60, 61, 91, 1100, 1101:
		return sec[4][22:], nil
	case 40, 41, 42, 43:
		return sec[4][24:], nil
	case 44:
		return sec[4][33:], nil
	case 45, 46, 47:
		return sec[4][35:], nil
	case 48:
		return sec[4][46:], nil
	case 52: // validation
		return sec[4][25:], nil
	case 57:
		n = number_of_mode(sec)
		if n <= 0 || n == 65535 {
			return nil, fatal_error_i("PDT 4.57 bad number of mode %d", n)
		}
		return sec[4][31+5*n:], nil
	case 20, 30, 31, 32, 1000, 1001, 1002, 254:
		return nil, nil
	default:
		return nil, fprintf("code_table_4.5a: product definition template #%d not supported\n", pdt)
	}
}

func code_table_4_5b_location(sec [][]unsigned_char) ([]unsigned_char, error) {
	var pdt, center, n int
	pdt = GB2_ProdDefTemplateNo(sec)
	center = GB2_Center(sec)

	if center == JMA1 || center == JMA2 {
		switch pdt {
		case 50008:
			return sec[4][28:], nil
		case 50009:
			return sec[4][28:], nil
		case 51020, 51021, 51022, 51122:
			return nil, nil
		case 50010:
		case 50011:
			return sec[4][28:], nil
		default:
		}
	}

	switch pdt {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 51, 60, 61, 91, 1100, 1101:
		return sec[4][28:], nil
	case 40, 41, 42, 43:
		return sec[4][30:], nil
	case 44:
		return sec[4][39:], nil
	case 45, 46, 47:
		return sec[4][41:], nil
	case 48:
		return sec[4][52:], nil
	case 52:
		return nil, nil
	case 57:
		n = number_of_mode(sec)
		if n <= 0 || n == 65535 {
			return nil, fatal_error_i("PDT 4.57 bad number of mode %d", n)
		}
		return sec[4][37+5*n:], nil
	case 20, 30, 31, 32, 1000, 1001, 1002, 254:
		return nil, nil
	default:
		return nil, fprintf("code_table_4.5b: product definition template #%d not supported\n", pdt)
	}
}

func code_table_5_0(sec [][]unsigned_char) int {
	return int(uint2(sec[5][9:]))
}

func code_table_6_0(sec [][]unsigned_char) int {
	return int(sec[6][5])
}
