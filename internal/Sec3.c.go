package internal

/*
 * figures out nx and ny
 *   res = resolution and component flags table 3.3
 *   scan = scan mode table 3.4
 *
 * Use this code rather than .h files for nx, ny, res and scan
 */
/*
 int get_nxny(unsigned char **sec, int *nx, int *ny, unsigned int *npnts, int *res, int *scan) {
	unsigned int nxx, nyy;
	int err;

	err =  get_nxny_(sec, &nxx, &nyy, npnts, res, scan);
	*nx = nxx == 0 ? -1 : nxx;
	*ny = nyy == 0 ? -1 : nyy;
	return err;
 }
*/
func get_nxny(sec [][]unsigned_char, nx *int, ny *int, npnts *unsigned_int, res *int, scan *int, n_variable_dim *int, variable_dim *[]int, raw_variable_dim *[]int) (err error) {
	var nxx, nyy unsigned_int

	err = get_nxny_(sec, &nxx, &nyy, npnts, res, scan, n_variable_dim, variable_dim, raw_variable_dim)
	//*nx = nxx == 0 ? -1 : nxx;
	if nxx == 0 {
		*nx = -1
	} else {
		*nx = int(nxx)
	}
	//*ny = nyy == 0 ? -1 : nyy;
	if nyy == 0 {
		*ny = -1
	} else {
		*ny = int(nyy)
	}
	return err
}

/*
int get_nxny_(unsigned char **sec, unsigned int *nx, unsigned int *ny, unsigned int *npnts, int *res, int *scan) {
    int grid_template, n_var_dim, i, j, n_octets, center;
    unsigned int npoints, n;
    unsigned char *gds, *p;

    center = GB2_Center(sec);
    grid_template = code_table_3_1(sec);
    *res = flag_table_3_3(sec);
    *scan = flag_table_3_4(sec);
    gds = sec[3];

    switch (grid_template) {
        case 0:
        case 1:
        case 2:
        case 3:
        case 4:
        case 5:
        case 10:
        case 12:
        case 20:
        case 30:
        case 31:
        case 40:
        case 41:
        case 42:
        case 43:
        case 44:
        case 90:
        case 110:
        case 140:
        case 204:
                *nx = uint4_missing(gds+30); *ny = uint4_missing(gds+34); break;
        case 51:
        case 52:
        case 53:
        case 130:
        case 50: *nx = GB2_Sec3_npts(sec); *ny = 1; break;        // should calculate for from parameters

        case 120: *nx = uint4_missing(gds+14);			// nx = bin along radials, ny = num radials
 		  *ny = uint4_missing(gds+18);
		  break;
	case 32768:
	case 32769:
	    if (center == NCEP) {
		*nx = uint4_missing(gds+30);
		*ny = uint4_missing(gds+34);
		break;
	    }
            *nx = *ny = 0; break;
	case 40110:
	    if ((center == JMA1) || (center == JMA2)) {
		*nx = uint4_missing(gds+30);
		*ny = uint4_missing(gds+34);
		break;
	    }
            *nx = *ny = 0; break;
	case 50120:
	    if ((center == JMA1) || (center == JMA2)) {
		*nx = uint4_missing(gds+14);	// Nb number of point from origin
		*ny = uint4_missing(gds+18);	// Nr number of angles
		break;
	    }
            *nx = *ny = 0; break;
	default:
                *nx = *ny = 0; break;
    }

    n_var_dim = 0;
    if (*nx == 0) n_var_dim = *ny;
    if (*ny == 0) n_var_dim = *nx;
    if (*nx == 0 && *ny == 0) n_var_dim = 0;

    p = NULL;
    if (n_var_dim) {
        switch (grid_template) {
           case 0: p = gds + 72; break;
           case 1: p = gds + 84; break;
           case 2: p = gds + 84; break;
           case 3: p = gds + 96; break;
           case 10: p = gds + 72; break;
           case 40: p = gds + 72; break;
           case 41: p = gds + 84; break;
           case 42: p = gds + 84; break;
           case 43: p = gds + 96; break;
	   case 32768: if (GB2_Center(sec) == NCEP) p = gds + 72;
			else p = NULL;
			break;
	   case 32769: if (GB2_Center(sec) == NCEP) p = gds + 80;
			else p = NULL;
			break;
           default: p = NULL; break;
        }
    }

    /* calculate number of grid points, check with GDS /
    npoints = 0;
    if (n_var_dim) {

        if (n_variable_dim != n_var_dim) {
            if (variable_dim) free(variable_dim);
            if (raw_variable_dim) free(raw_variable_dim);
            variable_dim = (int *) malloc(n_var_dim * sizeof(int));
            raw_variable_dim = (int *) malloc(n_var_dim * sizeof(int));

            if (variable_dim == NULL || raw_variable_dim == NULL)
		fatal_error("ran out of memory","");
            n_variable_dim = n_var_dim;
        }
        n_octets = (int) gds[10];        /* number of octets per integer /
        for (i = 0; i < n_var_dim; i++) {
            for (n = j = 0; j < n_octets; j++) {
                n = (n << 8) + (int) *p++;
            }
            raw_variable_dim[i] = variable_dim[i] = (int) n;
            npoints += n;
        }

        /* convert variable_dim to SN order if needed /
        if (*nx == 0 && GDS_Scan_y(*scan) == 0 && output_order == wesn) {
            for (i = 0; i < *ny; i++) {
                variable_dim[i] = raw_variable_dim[*ny-1-i];
            }
        }
        /* convert variable_dim to NS order if needed /
        else if (*nx == 0 && GDS_Scan_y(*scan) != 0 && output_order == wens) {
            for (i = 0; i < *ny; i++) {
                variable_dim[i] = raw_variable_dim[*ny-1-i];
            }
        }
    }
    else if (*nx > 0 && *ny > 0) npoints = (unsigned) *nx * *ny;
    *npnts = GB2_Sec3_npts(sec);

    if ((*nx != 0 || *ny != 0) && GB2_Sec3_npts(sec) != npoints && GDS_Scan_staggered_storage(*scan) == 0) {
        fprintf(stderr,"two values for number of points %u (GDS) %u (calculated)\n",
                      GB2_Sec3_npts(sec), npoints);
    }

/*
    for (i = 0; i < n_var_dim; i++) {
      printf("%d ", variable_dim[i]);
    }
/

    return 0;
}
*/
func get_nxny_(sec [][]unsigned_char, nx *unsigned_int, ny *unsigned_int, npnts *unsigned_int, res *int, scan *int, n_variable_dim *int, variable_dim *[]int, raw_variable_dim *[]int) error {
	var grid_template, n_var_dim, i, j, n_octets, center int
	var npoints, n unsigned_int
	var gds, p []unsigned_char

	center = GB2_Center(sec)
	grid_template = code_table_3_1(sec)
	*res = flag_table_3_3(sec)
	*scan = flag_table_3_4(sec)
	gds = sec[3]

	switch grid_template {
	case 0, 1, 2, 3, 4, 5, 10, 12, 20, 30, 31, 40, 41, 42, 43, 44, 90, 110, 140, 204:
		*nx = uint4_missing(gds[30:])
		*ny = uint4_missing(gds[34:])
	case 51, 52, 53, 130, 50:
		*nx = GB2_Sec3_npts(sec)
		*ny = 1
		// should calculate for from parameters

	case 120:
		*nx = uint4_missing(gds[14:]) // nx = bin along radials, ny = num radials
		*ny = uint4_missing(gds[18:])
	case 32768, 32769:
		if center == NCEP {
			*nx = uint4_missing(gds[30:])
			*ny = uint4_missing(gds[34:])
		} else {
			*nx = 0
			*ny = 0
		}
	case 40110:
		if (center == JMA1) || (center == JMA2) {
			*nx = uint4_missing(gds[30:])
			*ny = uint4_missing(gds[34:])
		} else {
			*nx = 0
			*ny = 0
		}
	case 50120:
		if (center == JMA1) || (center == JMA2) {
			*nx = uint4_missing(gds[14:]) // Nb number of point from origin
			*ny = uint4_missing(gds[18:]) // Nr number of angles
		} else {
			*nx = 0
			*ny = 0
		}
	default:
		*nx = 0
		*ny = 0
	}

	n_var_dim = 0
	if *nx == 0 {
		n_var_dim = int(*ny)
	}
	if *ny == 0 {
		n_var_dim = int(*nx)
	}
	if *nx == 0 && *ny == 0 {
		n_var_dim = 0
	}

	p = nil
	if n_var_dim != 0 {
		switch grid_template {
		case 0:
			p = gds[72:]
		case 1:
			p = gds[84:]
		case 2:
			p = gds[84:]
		case 3:
			p = gds[96:]
		case 10:
			p = gds[72:]
		case 40:
			p = gds[72:]
		case 41:
			p = gds[84:]
		case 42:
			p = gds[84:]
		case 43:
			p = gds[96:]
		case 32768:
			if GB2_Center(sec) == NCEP {
				p = gds[72:]
			} else {
				p = nil
			}
		case 32769:
			if GB2_Center(sec) == NCEP {
				p = gds[80:]
			} else {
				p = nil
			}
		default:
			p = nil
		}
	}

	/* calculate number of grid points, check with GDS */
	npoints = 0
	if n_var_dim != 0 {

		if *n_variable_dim != n_var_dim {
			//if (variable_dim) free(variable_dim);
			//if (raw_variable_dim) free(raw_variable_dim);
			//variable_dim = (int *) malloc(n_var_dim * sizeof(int));
			*variable_dim = make([]int, n_var_dim, n_var_dim)
			//raw_variable_dim = (int *) malloc(n_var_dim * sizeof(int));
			*raw_variable_dim = make([]int, n_var_dim, n_var_dim)

			*n_variable_dim = n_var_dim
		}
		n_octets = int(gds[10]) /* number of octets per integer */
		for i = 0; i < n_var_dim; i++ {
			//for (n = j = 0; j < n_octets; j++) {
			n = 0
			j = 0
			pIndex := 0
			for ; j < n_octets; j++ {
				// n = (n << 8) + int(*p++);
				n = (n << 8) + unsigned_int(p[pIndex])
				pIndex++
			}
			// raw_variable_dim[i] = variable_dim[i] = (int) n;
			(*raw_variable_dim)[i] = int(n)
			(*variable_dim)[i] = int(n)
			npoints += n
		}

		/* convert variable_dim to SN order if needed */
		if *nx == 0 && GDS_Scan_y(*scan) == false && output_order == wesn {
			for i = 0; i < int(*ny); i++ {
				(*variable_dim)[i] = (*raw_variable_dim)[int(*ny)-1-i]
			}
		} else if *nx == 0 && GDS_Scan_y(*scan) != false && output_order == wens {
			/* convert variable_dim to NS order if needed */
			for i = 0; i < int(*ny); i++ {
				(*variable_dim)[i] = (*raw_variable_dim)[int(*ny)-1-i]
			}
		}
	} else if *nx > 0 && *ny > 0 {
		npoints = unsigned_int(*nx * *ny)
	}
	*npnts = unsigned_int(GB2_Sec3_npts(sec))

	if (*nx != 0 || *ny != 0) && GB2_Sec3_npts(sec) != npoints && GDS_Scan_staggered_storage(*scan) == false {
		// fprintf(stderr, "two values for number of points %u (GDS) %u (calculated)\n",
		//	GB2_Sec3_npts(sec), npoints)
		return fprintf("two values for number of points %d (GDS) %d (calculated)", GB2_Sec3_npts(sec), npoints)
	}

	/*
	   for (i = 0; i < n_var_dim; i++) {
	     printf("%d ", variable_dim[i]);
	   }
	*/

	return nil
}
