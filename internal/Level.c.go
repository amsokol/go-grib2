package internal

/* code table 4.5 */

var level_table = []string{
	/* 0 */ "reserved",
	/* 1 */ "surface",
	/* 2 */ "cloud base",
	/* 3 */ "cloud top",
	/* 4 */ "0C isotherm",
	/* 5 */ "level of adiabatic condensation from sfc",
	/* 6 */ "max wind",
	/* 7 */ "tropopause",
	/* 8 */ "top of atmosphere",
	/* 9 */ "sea bottom",
	/* 10 */ "entire atmosphere",
	/* 11 */ "cumulonimbus base",
	/* 12 */ "cumulonimbus top",
	/* 13 */ "lowest level %g%% integrated cloud cover",
	/* 14 */ "level of free convection",
	/* 15 */ "convection condensation level",
	/* 16 */ "level of neutral buoyancy",
	/* 17 */ "reserved",
	/* 18 */ "reserved",
	/* 19 */ "reserved",
	/* 20 */ "%g K level",
	/* 21 */ "reserved",
	/* 22 */ "reserved",
	/* 23 */ "reserved",
	/* 24 */ "reserved",
	/* 25 */ "reserved",
	/* 26 */ "reserved",
	/* 27 */ "reserved",
	/* 28 */ "reserved",
	/* 29 */ "reserved",
	/* 30 */ "reserved",
	/* 31 */ "reserved",
	/* 32 */ "reserved",
	/* 33 */ "reserved",
	/* 34 */ "reserved",
	/* 35 */ "reserved",
	/* 36 */ "reserved",
	/* 37 */ "reserved",
	/* 38 */ "reserved",
	/* 39 */ "reserved",
	/* 40 */ "reserved",
	/* 41 */ "reserved",
	/* 42 */ "reserved",
	/* 43 */ "reserved",
	/* 44 */ "reserved",
	/* 45 */ "reserved",
	/* 46 */ "reserved",
	/* 47 */ "reserved",
	/* 48 */ "reserved",
	/* 49 */ "reserved",
	/* 50 */ "reserved",
	/* 51 */ "reserved",
	/* 52 */ "reserved",
	/* 53 */ "reserved",
	/* 54 */ "reserved",
	/* 55 */ "reserved",
	/* 56 */ "reserved",
	/* 57 */ "reserved",
	/* 58 */ "reserved",
	/* 59 */ "reserved",
	/* 60 */ "reserved",
	/* 61 */ "reserved",
	/* 62 */ "reserved",
	/* 63 */ "reserved",
	/* 64 */ "reserved",
	/* 65 */ "reserved",
	/* 66 */ "reserved",
	/* 67 */ "reserved",
	/* 68 */ "reserved",
	/* 69 */ "reserved",
	/* 70 */ "reserved",
	/* 71 */ "reserved",
	/* 72 */ "reserved",
	/* 73 */ "reserved",
	/* 74 */ "reserved",
	/* 75 */ "reserved",
	/* 76 */ "reserved",
	/* 77 */ "reserved",
	/* 78 */ "reserved",
	/* 79 */ "reserved",
	/* 80 */ "reserved",
	/* 81 */ "reserved",
	/* 82 */ "reserved",
	/* 83 */ "reserved",
	/* 84 */ "reserved",
	/* 85 */ "reserved",
	/* 86 */ "reserved",
	/* 87 */ "reserved",
	/* 88 */ "reserved",
	/* 89 */ "reserved",
	/* 90 */ "reserved",
	/* 91 */ "reserved",
	/* 92 */ "reserved",
	/* 93 */ "reserved",
	/* 94 */ "reserved",
	/* 95 */ "reserved",
	/* 96 */ "reserved",
	/* 97 */ "reserved",
	/* 98 */ "reserved",
	/* 99 */ "reserved",
	/* 100 */ "%g mb",
	/* 101 */ "mean sea level",
	/* 102 */ "%g m above mean sea level",
	/* 103 */ "%g m above ground",
	/* 104 */ "%g sigma level",
	/* 105 */ "%g hybrid level",
	/* 106 */ "%g m underground",
	/* 107 */ "%g K isentropic level",
	/* 108 */ "%g mb above ground",
	/* 109 */ "PV=%g (Km^2/kg/s) surface",
	/* 110 */ "reserved",
	/* 111 */ "%g Eta level",
	/* 112 */ "reserved",
	/* 113 */ "logarithmic hybrid level",
	/* 114 */ "snow level",
	/* 115 */ "reserved",
	/* 116 */ "reserved",
	/* 117 */ "mixed layer depth",
	/* 118 */ "hybrid height level",
	/* 119 */ "hybrid pressure level",
	/* 120 */ "reserved",
	/* 121 */ "reserved",
	/* 122 */ "reserved",
	/* 123 */ "reserved",
	/* 124 */ "reserved",
	/* 125 */ "reserved",
	/* 126 */ "reserved",
	/* 127 */ "reserved",
	/* 128 */ "reserved",
	/* 129 */ "reserved",
	/* 130 */ "reserved",
	/* 131 */ "reserved",
	/* 132 */ "reserved",
	/* 133 */ "reserved",
	/* 134 */ "reserved",
	/* 135 */ "reserved",
	/* 136 */ "reserved",
	/* 137 */ "reserved",
	/* 138 */ "reserved",
	/* 139 */ "reserved",
	/* 140 */ "reserved",
	/* 141 */ "reserved",
	/* 142 */ "reserved",
	/* 143 */ "reserved",
	/* 144 */ "reserved",
	/* 145 */ "reserved",
	/* 146 */ "reserved",
	/* 147 */ "reserved",
	/* 148 */ "reserved",
	/* 149 */ "reserved",
	/* 150 */ "%g generalized vertical height coordinate",
	/* 151 */ "soil level %g",
	/* 152 */ "reserved",
	/* 153 */ "reserved",
	/* 154 */ "reserved",
	/* 155 */ "reserved",
	/* 156 */ "reserved",
	/* 157 */ "reserved",
	/* 158 */ "reserved",
	/* 159 */ "reserved",
	/* 160 */ "%g m below sea level",
	/* 161 */ "%g m below water surface",
	/* 162 */ "lake or river bottom",
	/* 163 */ "bottom of sediment layer",
	/* 164 */ "bottom of thermally active sediment layer",
	/* 165 */ "bottom of sediment layer penetrated by thermal wave",
	/* 166 */ "maxing layer",
	/* 167 */ "bottom of root zone",
	/* 168 */ "reserved",
	/* 169 */ "reserved",
	/* 170 */ "reserved",
	/* 171 */ "reserved",
	/* 172 */ "reserved",
	/* 173 */ "reserved",
	/* 174 */ "top surface of ice on sea, lake or river",
	/* 175 */ "top surface of ice, und snow on sea, lake or river",
	/* 176 */ "bottom surface ice on sea, lake or river",
	/* 177 */ "deep soil",
	/* 178 */ "reserved",
	/* 179 */ "top surface of glacier ice and inland ice",
	/* 180 */ "deep inland or glacier ice",
	/* 181 */ "grid tile land fraction as a model surface",
	/* 182 */ "grid tile water fraction as a model surface",
	/* 183 */ "grid tile ice fraction on sea, lake or river as a model surface",
	/* 184 */ "grid tile glacier ice and inland ice fraction as a model surface",
	/* 185 */ "reserved",
	/* 186 */ "reserved",
	/* 187 */ "reserved",
	/* 188 */ "reserved",
	/* 189 */ "reserved",
	/* 190 */ "reserved",
	/* 191 */ "reserved",
}

func f_lev(sec [][]unsigned_char, inv_out *string) error {

	var level_type1, level_type2 int
	var val1, val2 float
	var undef_val1, undef_val2 int
	var center, subcenter int

	center = GB2_Center(sec)
	subcenter = GB2_Subcenter(sec)

	err := fixed_surfaces(sec, &level_type1, &val1, &undef_val1, &level_type2, &val2, &undef_val2)
	if err != nil {
		return fatal_error_wrap(err, "Failed to execute fixed_surfaces")
	}

	/*
		if mode > 1 {
			if undef_val1 == 0 {
				sprintf(inv_out, "lvl1=(%d,%lg) ", level_type1, val1)
			} else {
				sprintf(inv_out, "lvl1=(%d,missing) ", level_type1)
			}
			inv_out += strlen(inv_out)

			if undef_val2 == 0 {
				sprintf(inv_out, "lvl2=(%d,%lg):", level_type2, val2)
			} else {
				sprintf(inv_out, "lvl2=(%d,missing):", level_type2)
			}
			inv_out += strlen(inv_out)
		}
	*/
	level2(level_type1, undef_val1, val1, level_type2, undef_val2, val2, center, subcenter, inv_out)
	return nil
}

/*
 * level2 is for layers
 */

func level2(type1 int, undef_val1 int, value1 float, type2 int, undef_val2 int, value2 float, center int, subcenter int, inv_out *string) {
	if type1 == 100 && type2 == 100 {
		*inv_out += sprintf("%g-%g mb", value1/100, value2/100)
	} else if type1 == 102 && type2 == 102 {
		*inv_out += sprintf("%g-%g m above mean sea level", value1, value2)
	} else if type1 == 103 && type2 == 103 {
		*inv_out += sprintf("%g-%g m above ground", value1, value2)
	} else if type1 == 104 && type2 == 104 {
		*inv_out += sprintf("%g-%g sigma layer", value1, value2)
	} else if type1 == 105 && type2 == 105 {
		*inv_out += sprintf("%g-%g hybrid layer", value1, value2)
	} else if type1 == 106 && type2 == 106 {
		/* sprintf(inv_out,"%g-%g m below ground",value1/100,value2/100); removed 1/2007 */
		*inv_out += sprintf("%g-%g m below ground", value1, value2)
	} else if type1 == 107 && type2 == 107 {
		*inv_out += sprintf("%g-%g K isentropic layer", value1, value2)
	} else if type1 == 108 && type2 == 108 {
		*inv_out += sprintf("%g-%g mb above ground", value1/100, value2/100)
	} else if type1 == 160 && type2 == 160 {
		*inv_out += sprintf("%g-%g m below sea level", value1, value2)
	} else if type1 == 161 && type2 == 161 {
		*inv_out += sprintf("%g-%g m ocean layer", value1, value2)
	} else if type1 == 1 && type2 == 8 {
		*inv_out += sprintf("atmos col") // compatible with wgrib
	} else if type1 == 9 && type2 == 1 {
		*inv_out += sprintf("ocean column")
	} else if center == NCEP && type1 == 235 && type2 == 235 {
		*inv_out += sprintf("%g-%gC ocean isotherm layer", value1/10, value2/10)
	} else if center == NCEP && type1 == 236 && type2 == 236 { // obsolete
		*inv_out += sprintf("%g-%g m ocean layer", value1*10, value2*10)
	} else if type1 == 255 && type2 == 255 {
		*inv_out += sprintf("no_level")
	} else {
		level1(type1, undef_val1, value1, center, subcenter, inv_out)
		if type2 != 255 {
			*inv_out += " - "
			level1(type2, undef_val2, value2, center, subcenter, inv_out)
		}
	}
}

/*
 * level1 is for a single level (not a layer)
 */
func level1(type_ int, undef_val int, val float, center int, subcenter int, inv_out *string) {

	var string_ string

	/* local table for NCEP */
	if center == NCEP && type_ >= 192 && type_ <= 254 {
		if type_ == 235 {
			*inv_out += sprintf("%gC ocean isotherm", val/10)
			return
		}
		if type_ == 241 {
			*inv_out += sprintf("%g in sequence", val)
		}

		switch type_ {
		case 200:
			string_ = "entire atmosphere (considered as a single layer)"
		case 201:
			string_ = "entire ocean (considered as a single layer)"
		case 204:
			string_ = "highest tropospheric freezing level"
		case 206:
			string_ = "grid scale cloud bottom level"
		case 207:
			string_ = "grid scale cloud top level"
		case 209:
			string_ = "boundary layer cloud bottom level"
		case 210:
			string_ = "boundary layer cloud top level"
		case 211:
			string_ = "boundary layer cloud layer"
		case 212:
			string_ = "low cloud bottom level"
		case 213:
			string_ = "low cloud top level"
		case 214:
			string_ = "low cloud layer"
		case 215:
			string_ = "cloud ceiling"
		case 220:
			string_ = "planetary boundary layer"
		case 221:
			string_ = "layer between two hybrid levels"
		case 222:
			string_ = "middle cloud bottom level"
		case 223:
			string_ = "middle cloud top level"
		case 224:
			string_ = "middle cloud layer"
		case 232:
			string_ = "high cloud bottom level"
		case 233:
			string_ = "high cloud top level"
		case 234:
			string_ = "high cloud layer"
		case 235:
			string_ = "ocean isotherm level (1/10 deg C)"
		case 236:
			string_ = "layer between two depths below ocean surface"
		case 237:
			string_ = "bottom of ocean mixed layer"
		case 238:
			string_ = "bottom of ocean isothermal layer"
		case 239:
			string_ = "layer ocean surface and 26C ocean isothermal level"
		case 240:
			string_ = "ocean mixed layer"
		case 241:
			string_ = "ordered sequence of data"
		case 242:
			string_ = "convective cloud bottom level"
		case 243:
			string_ = "convective cloud top level"
		case 244:
			string_ = "convective cloud layer"
		case 245:
			string_ = "lowest level of the wet bulb zero"
		case 246:
			string_ = "maximum equivalent potential temperature level"
		case 247:
			string_ = "equilibrium level"
		case 248:
			string_ = "shallow convective cloud bottom level"
		case 249:
			string_ = "shallow convective cloud top level"
		case 251:
			string_ = "deep convective cloud bottom level"
		case 252:
			string_ = "deep convective cloud top level"
		case 253:
			string_ = "lowest bottom level of supercooled liquid water layer"
		case 254:
			string_ = "highest top level of supercooled liquid water layer"
		}
		if len(string_) > 0 {
			*inv_out += sprintf(string_, val)
			return
		}
	}

	if type_ == 100 || type_ == 108 {
		val = val * 0.01
	} // Pa -> mb

	// no numeric information
	if type_ == 255 {
		return
	}

	if type_ < 192 {
		*inv_out += sprintf(level_table[type_], val)
	} else if center == NCEP {
		if undef_val == 0 {
			*inv_out += sprintf("NCEP level type %d %g", type_, val)
		} else {
			*inv_out += sprintf("NCEP level type %d", type_)
		}
	} else {
		if undef_val == 0 {
			*inv_out += sprintf("local level type %d %g", type_, val)
		} else {
			*inv_out += sprintf("local level type %d", type_)
		}
	}

	return
}
