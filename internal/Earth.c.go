package internal

/* function to return the size of the earth */

func radius_earth(sec [][]unsigned_char) (double, error) {
	var table_3_2 int
	var radius double
	var p []unsigned_char
	var factor, value int

	p = code_table_3_2_location(sec)
	if p == nil {
		return 0, fatal_error("radius_earth: code_table 3.2 is unknown", "")
	}

	table_3_2 = int(p[0])

	/* set a default value .. not sure what to do with most values */
	radius = 6367.47 * 1000.0
	if table_3_2 == 0 {
		radius = 6367.47 * 1000.0
	} else if table_3_2 == 1 {
		factor = INT1(p[1])
		value = int4(p[2:])
		radius = scaled2dbl(factor, value)
		if radius < 6300000.0 || radius > 6400000.0 {
			return 0, fatal_error_i("radius of earth is %d m", int(radius))
		}
	} else if table_3_2 == 2 {
		radius = (6378.160 + 6356.775) * 0.5 * 1000.0
	} else if table_3_2 == 3 || table_3_2 == 7 {
		/* get major axis */
		factor = INT1(p[6])
		value = int4(p[7:])
		radius = scaled2dbl(factor, value)
		/* get minor axis */
		factor = INT1(p[11])
		value = int4(p[12:])
		radius = (radius + scaled2dbl(factor, value)) * 0.5

		/* radius in km, convert to m */
		if table_3_2 == 3 {
			radius *= 1000.0
		}

		if radius < 6300000.0 || radius > 6400000.0 {
			return 0, fatal_error_i("radius of earth is %d m", int(radius))
		}
	} else if table_3_2 == 4 {
		radius = (6378.137 + 6356.752) * 0.5 * 1000.0
	} else if table_3_2 == 5 {
		radius = (6378.137 + 6356.752) * 0.5 * 1000.0
	} else if table_3_2 == 6 {
		radius = 6371.2290 * 1000.0
	} else if table_3_2 == 8 {
		radius = 6371.200 * 1000.0
	} else if table_3_2 == 9 {
		radius = (6377563.396 + 6356256.909) * 0.5
	}

	return radius, nil
}
