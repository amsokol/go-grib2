package internal

/*
int flag_table_3_3(unsigned char **sec) {
    unsigned char *p;
    p = flag_table_3_3_location(sec);
    if (p == NULL) return -1;
    return (int) *p;
}
*/
func flag_table_3_3(sec [][]unsigned_char) int {
	var p []unsigned_char
	p = flag_table_3_3_location(sec)
	if p == nil {
		return -1
	}
	return int(p[0])
}

/*
unsigned char *flag_table_3_3_location(unsigned char **sec) {
    int grid_template, center;
    unsigned char *gds;

    grid_template = code_table_3_1(sec);
    center = GB2_Center(sec);
    gds = sec[3];
    switch (grid_template) {
        case 0:
        case 1:
        case 2:
        case 3:
        case 40:
        case 41:
        case 42:
        case 43:
	case 140:
        case 204:
              return gds+54; break;
        case 4:
        case 5:
        case 10:
        case 12:
        case 20:
        case 30:
        case 31:
        case 90:
        case 110:
              return gds+46; break;
	case 32768:
		if (center == NCEP) return gds+54;
		return NULL;
	case 32769:
		if (center == NCEP) return gds+54;
		return NULL;
	case 40110:
		if ((center == JMA1) || (center == JMA2)) return gds+46;
		return NULL;
		break;
        default: break;
    }
    return NULL;
}
*/

func flag_table_3_3_location(sec [][]unsigned_char) []unsigned_char {
	var grid_template, center int
	var gds []unsigned_char

	grid_template = code_table_3_1(sec)
	center = GB2_Center(sec)
	gds = sec[3]
	switch grid_template {
	case 0, 1, 2, 3, 40, 41, 42, 43, 140, 204:
		return gds[54:]
	case 4, 5, 10, 12, 20, 30, 31, 90, 110:
		return gds[46:]
	case 32768:
		if center == NCEP {
			return gds[54:]
		}
		return nil
	case 32769:
		if center == NCEP {
			return gds[54:]
		}
		return nil
	case 40110:
		if (center == JMA1) || (center == JMA2) {
			return gds[46:]
		}
		return nil
	default:
	}
	return nil
}

/*
int flag_table_3_4(unsigned char **sec) {
    unsigned char *p;
    p = flag_table_3_4_location(sec);
    if (p == NULL) return -1;
    return (int) *p;
}
*/
func flag_table_3_4(sec [][]unsigned_char) int {
	var p []unsigned_char
	p = flag_table_3_4_location(sec)
	if p == nil {
		return -1
	}
	return int(p[0])
}

/*

unsigned char *flag_table_3_4_location(unsigned char **sec) {
    int grid_template, center;
    unsigned char *gds;

    gds = sec[3];
    grid_template = code_table_3_1(sec);
    center = GB2_Center(sec);

    switch (grid_template) {
        case 0:
        case 1:
        case 2:
        case 3:
        case 40:
        case 41:
        case 42:
        case 43:
                 return gds+71; break;
	case 4:
	case 5:
                 return gds+47; break;
        case 10:
        case 12:
                 return gds+59; break;
        case 20: return gds+64; break;
        case 30:
        case 31: return gds+64; break;
        case 50:
        case 51:
        case 52:
        case 53:
                 /* spectral modes don't have scan order /
                 return NULL; break;
        case 90:
        case 140:
		  return gds+63; break;
        case 110: return gds+56; break;
        case 190:
	case 120: return gds+38; break;
	case 204: return gds+71; break;
        case 1000: return gds+50; break;
	case 32768:
		if (center == NCEP) return gds+71;
		return NULL;
		break;
	case 32769:
		if (center == NCEP) return gds+71;
		return NULL;
		break;
	case 40110:
		if ((center == JMA1) || (center == JMA2)) return gds+56;
		return NULL;
		break;
        default: break;
    }
    return NULL;
}
*/

func flag_table_3_4_location(sec [][]unsigned_char) []unsigned_char {
	var grid_template, center int
	var gds []unsigned_char

	gds = sec[3]
	grid_template = code_table_3_1(sec)
	center = GB2_Center(sec)

	switch grid_template {
	case 0, 1, 2, 3, 40, 41, 42, 43:
		return gds[71:]
	case 4, 5:
		return gds[47:]
	case 10, 12:
		return gds[59:]
	case 20:
		return gds[64:]
	case 30, 31:
		return gds[64:]
	case 50, 51, 52, 53:
		/* spectral modes don't have scan order */
		return nil
	case 90, 140:
		return gds[63:]
	case 110:
		return gds[56:]
	case 190, 120:
		return gds[38:]
	case 204:
		return gds[71:]
	case 1000:
		return gds[50:]
	case 32768:
		if center == NCEP {
			return gds[71:]
		}
		return nil
	case 32769:
		if center == NCEP {
			return gds[71:]
		}
		return nil
	case 40110:
		if (center == JMA1) || (center == JMA2) {
			return gds[56:]
		}
		return nil
	default:
	}
	return nil
}
