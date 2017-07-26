package internal

func regular2ll(sec [][]unsigned_char, lat *[]double, lon *[]double) error {

	var basic_ang, sub_ang int
	var units, dlat, dlon, lat1, lat2, lon1, lon2 double
	var e, w, n, s, dx, dy double

	var i, j unsigned_int
	var llat, llon []double
	var gds []unsigned_char
	var nnx, nny unsigned_int
	var nres, nscan int
	var nnpnts unsigned_int

	var n_variable_dim int
	var variable_dim, raw_variable_dim []int

	get_nxny_(sec, &nnx, &nny, &nnpnts, &nres, &nscan, &n_variable_dim, &variable_dim, &raw_variable_dim)
	gds = sec[3]

	if nny < 1 {
		return fprintf( /*stderr, */ "Sorry code does not handle variable ny yet")
	}

	/*
		   if ((*lat = (double *) malloc(((size_t) nnpnts) * sizeof(double))) == NULL) {
			   return fatal_error("regular2ll memory allocation failed","");
		   }
	*/
	*lat = make([]double, nnpnts, nnpnts)
	/*
		   if ((*lon = (double *) malloc(((size_t) nnpnts) * sizeof(double))) == NULL) {
			   return fatal_error("regular2ll memory allocation failed","");
		   }
	*/
	*lon = make([]double, nnpnts, nnpnts)

	/* now figure out the grid coordinates mucho silly grib specification */

	basic_ang = GDS_LatLon_basic_ang(gds)
	sub_ang = GDS_LatLon_sub_ang(gds)
	if basic_ang != 0 {
		units = double(basic_ang) / double(sub_ang)
	} else {
		units = 0.000001
	}

	dlat = double(GDS_LatLon_dlat(gds)) * units
	dlon = double(GDS_LatLon_dlon(gds)) * units
	lat1 = double(GDS_LatLon_lat1(gds)) * units
	lat2 = double(GDS_LatLon_lat2(gds)) * units
	lon1 = double(GDS_LatLon_lon1(gds)) * units
	lon2 = double(GDS_LatLon_lon2(gds)) * units

	if lon1 < 0.0 || lon2 < 0.0 {
		return fatal_error("BAD grid definition lon < zero", "")
	}
	if lon1 > 360.0 || lon2 > 360.0 {
		return fatal_error("BAD grid definition lon >= 360", "")
	}
	if lat1 < -90.0 || lat2 < -90.0 || lat1 > 90.0 || lat2 > 90.0 {
		return fatal_error("BAD grid definition lat", "")
	}

	/* find S latitude and dy */
	if GDS_Scan_y(nscan) {
		s = lat1
		n = lat2
	} else {
		s = lat2
		n = lat1
	}
	if s > n {
		return fatal_error("lat-lon grid: lat1 and lat2 inconsistent with scan order", "")
	}

	if nny > 1 {
		dy = (n - s) / double(nny-1)
		if nres&16 != 0 { /* lat increment is valid */
			if fabs(dy-dlat) > 0.001 {
				return fatal_error("lat-lon grid: dlat is inconsistent", "")
			}
		}
	} else {
		dy = 0.0
	}
	// fprintf(stderr,">>> geo:  dy %lf dlat %lf nres %d has dy %d has dx %d\n", dy, dlat, nres, nres & 16, nres & 32);

	/* find W latitude and dx */

	if GDS_Scan_row_rev(nscan) && (nny%2 == 0) && ((nres & 32) == 0) {
		//fatal_error("grib GDS ambiguity", "")
	}

	if GDS_Scan_x(nscan) {
		w = lon1
		e = lon2
		if GDS_Scan_row_rev(nscan) && ((nres & 32) == 0) {
			e = lon1 + (double(nnx)-1)*dlon
		}
	} else {
		w = lon2
		e = lon1
		if GDS_Scan_row_rev(nscan) && ((nres & 32) == 0) {
			w = lon1 - (double(nnx)-1)*dlon
		}
	}

	if e <= w {
		e += 360.0
	}
	if e-w > 360.0 {
		e -= 360.0
	}
	if w < 0 {
		w += 360.0
		e += 360.0
	}

	/* lat-lon should be in a WE:SN order */

	if nnx > 0 && nny > 0 { /* non-thinned, potentially staggered grid */
		/* put x[] and y[] values in lon[] and lat[] */
		llat = *lat
		llon = *lon
		err := stagger(sec, nnpnts, llon, llat, &n_variable_dim, &variable_dim, &raw_variable_dim)
		if err != nil {
			return fatal_error_wrap(err, "geo: stagger problem", "")
		}

		if nnx > 1 {
			dx = (e - w) / (double(nnx) - 1)
			dx = fabs(dx)
			if (nres & 32) != 0 { /* lon increment is valid */
				if fabs(dx-fabs(dlon)) > 0.001 {
					return fatal_error("lat-lon grid: dlon is inconsistent", "")
				}
			}
		} else {
			dx = 0.0
		}
		dy = fabs(dy)

		for j = 0; j < nnpnts; j++ {
			llon[j] = lon1 + llon[j]*dx
			// llon[j] = llon[j] >= 360.0 ? llon[j] - 360.0 : llon[j];
			if llon[j] >= 360.0 {
				llon[j] = llon[j] - 360.0
			}

			// llon[j] = llon[j] < 0.0 ? llon[j] + 360.0 : llon[j];
			if llon[j] < 0.0 {
				llon[j] = llon[j] + 360.0
			}
			llat[j] = lat1 + llat[j]*dy
		}
		return nil
	}

	/* must be thinned grid */

	llat = *lat
	llatIndex := 0
	/* quasi-regular grid */
	for j = 0; j < nny; j++ {
		for i = 0; i < unsigned_int(variable_dim[j]); i++ {
			llat[llatIndex] = s + double(j)*dy
			llatIndex++
		}
	}

	llon = *lon
	llonIndex := 0
	/* quasi-regular grid */
	for j = 0; j < nny; j++ {
		dx = (e - w) / double(variable_dim[j]-1)
		for i = 0; i < unsigned_int(variable_dim[j]); i++ {
			//*llon++ = w + i*dx >= 360.0 ? w + i*dx - 360.0: w + i*dx;
			if w+double(i)*dx >= 360.0 {
				llon[llonIndex] = w + double(i)*dx - 360.0
			} else {
				llon[llonIndex] = w + double(i)*dx
			}
			llonIndex++
		}
	}
	return nil
} /* end regular2ll() */
