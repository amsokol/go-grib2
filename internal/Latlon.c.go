package internal

/*
int get_latlon(unsigned char **sec, double **lon, double **lat) {

    int grid_template;

    if (*lat != NULL) {
        free(*lat);
        free(*lon);
        *lat = *lon = NULL;
    }

    grid_template = code_table_3_1(sec);
    if (grid_template == 0) {
        regular2ll(sec, lat, lon);
    }
    else if (grid_template == 1) {		// rotated lat-lon
        rot_regular2ll(sec, lat, lon);
    }
    else if (grid_template == 10) {
        mercator2ll(sec, lat, lon);
    }
    else if (grid_template == 20) {
        polar2ll(sec, lat, lon);
    }
    else if (grid_template == 30) {
        lambert2ll(sec, lat, lon);
    }
    else if (grid_template == 40) {
        gauss2ll(sec, lat, lon);
    }
    else if (grid_template == 90) {
        space_view2ll(sec, lat, lon);
    }
    else if (grid_template == 130) {
        irr_grid2ll(sec, lat, lon);
    }

    return 0;
}
*/
func get_latlon(sec [][]unsigned_char, lon *[]double, lat *[]double) error {

	var grid_template int

	grid_template = code_table_3_1(sec)
	if grid_template == 0 {
		return regular2ll(sec, lat, lon)
	} else if grid_template == 1 { // rotated lat-lon
		return rot_regular2ll(sec, lat, lon)
	} else if grid_template == 10 {
		return mercator2ll(sec, lat, lon)
	} else if grid_template == 20 {
		return polar2ll(sec, lat, lon)
	} else if grid_template == 30 {
		return lambert2ll(sec, lat, lon)
	} else if grid_template == 40 {
		// TODO: port gauss2ll
		// gauss2ll(sec, lat, lon)
	} else if grid_template == 90 {
		// TODO: port space_view2ll
		// space_view2ll(sec, lat, lon)
	} else if grid_template == 130 {
		// TODO: port irr_grid2ll
		// irr_grid2ll(sec, lat, lon)
	}

	return fatal_error("Unsupported Grid template: %d", grid_template)
}
