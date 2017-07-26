package internal

/*
int reftime(unsigned char **sec, int *year, int *month, int *day, int *hour, int *minute, int *second)
{

    unsigned char *p;

    p = sec[1];
    *year = (p[12] << 8) | p[13];
    *month = p[14];
    *day = p[15];
    *hour = p[16];
    *minute = p[17];
    *second = p[18];

    return 0;
}
*/
func reftime(sec [][]unsigned_char, year *int, month *int, day *int, hour *int, minute *int, second *int) {
	var p []unsigned_char

	p = sec[1]
	*year = (int(p[12]) << 8) | int(p[13])
	*month = int(p[14])
	*day = int(p[15])
	*hour = int(p[16])
	*minute = int(p[17])
	*second = int(p[18])
}

/*
 * Returns the verification time: reference_time + forecast_time + statistical processing time (if any)
 * 9/2006  w. ebisuzaki
 * 1/2007  w. ebisuzaki return error code
 * 11/2007  w. ebisuzaki fixed code for non forecasts
 * 3/2008 w. ebisuzaki added code for ensemble processing
 * 4/2009 w. ebisuzaki test table 1.2 sign of reference time
 */
/*
 int verftime(unsigned char **sec, int *year, int *month, int *day, int *hour, int *minute, int *second) {

        int units, i, j;
        int dtime;
        static int error_count = 0;
        static int warning_count = 0;

        i = code_table_4_0(sec);

        // if statistically processed field, verftime is in header
        j = stat_proc_verf_time(sec, year, month, day, hour, minute, second);
        if (j == 0) return 0;

        get_time(sec[1]+12, year, month, day, hour, minute, second);

        // some products have no forecast time

        if (code_table_4_4(sec) == -1) return 0;

        // check the significance of the refernce time
        i = code_table_1_2(sec);

        if (i == 2) {
        // rt=verifying time of forecast
        // unclear what it means for time averages/accumulations
        if (warning_count == 0) {
            fprintf(stderr,"Warning: rt == vt (CodeTable 1.2)\n");
            warning_count++;
        }
        return 0;
        }
        if (i == 3) {
        // rt = observing time
        return 0;
        }

        if (i != 0 && i != 1) {
        if (error_count == 0) {
            fprintf(stderr,"verifying time: Table 1.2=%d not supported "
            " using RT=analysis/start of forecast\n", i);
            error_count++;
        }
        }

        units = code_table_4_4(sec);
        dtime = forecast_time_in_units(sec);
        if (dtime >= 0)
            return add_time(year, month, day, hour, minute, second, (unsigned int) dtime, units);
        return sub_dt(year, month, day, hour, minute, second, (unsigned int) (unsigned int) -dtime, units);
    }
*/

func verftime(sec [][]unsigned_char, year *int, month *int, day *int, hour *int, minute *int, second *int) error {
	var units, i, j int
	var dtime int

	i = code_table_4_0(sec)

	// if statistically processed field, verftime is in header
	j = stat_proc_verf_time(sec, year, month, day, hour, minute, second)
	if j == 0 {
		return nil
	}

	get_time(sec[1][12:], year, month, day, hour, minute, second)

	// some products have no forecast time

	if code_table_4_4(sec) == -1 {
		return nil
	}

	// check the significance of the reference time
	i = code_table_1_2(sec)

	if i == 2 {
		// rt=verifying time of forecast
		// unclear what it means for time averages/accumulations
		/*
			        if warning_count == 0 {
						fprintf(stderr, "Warning: rt == vt (CodeTable 1.2)\n")
						warning_count++
			        }
		*/
		return nil
	}
	if i == 3 {
		// rt = observing time
		return nil
	}

	if i != 0 && i != 1 {
		return fprintf("verifying time: Table 1.2=%d not supported using RT=analysis/start of forecast\n", i)
	}

	units = code_table_4_4(sec)
	dtime = forecast_time_in_units(sec)
	if dtime != 0 {
		add_time(year, month, day, hour, minute, second, dtime, units)
	}
	return nil
}
