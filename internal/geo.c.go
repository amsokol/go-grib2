package internal

const M_PI = 3.14159265358979323846

const M_PI_4 = 0.78539816339744830962 /* pi/4 */

const M_PI_2 = 1.57079632679489661923 /* pi/2 */

func todegrees(x double) double {
	return x * (180.0 / M_PI)
}

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

func rot_regular2ll(sec [][]unsigned_char, lat *[]double, lon *[]double) error {

	var gds []unsigned_char
	var units double
	var tlon, tlat []double
	var sp_lat, sp_lon, angle_rot double
	var sin_a, cos_a double
	var basic_ang, sub_ang int
	var i, npnts unsigned_int
	var a, b, r, pr, gr, pm, gm, glat, glon double

	/* get the lat-lon coordinates in rotated frame of referencee */
	err := regular2ll(sec, lat, lon)
	if err != nil {
		return fatal_error_wrap(err, "Failed to execute regular2ll")
	}

	gds = sec[3]
	npnts = GB2_Sec3_npts(sec)

	basic_ang = GDS_LatLon_basic_ang(gds)
	sub_ang = GDS_LatLon_sub_ang(gds)
	if basic_ang != 0 {
		units = double(basic_ang) / double(sub_ang)
	} else {
		units = 0.000001
	}

	sp_lat = double(GDS_RotLatLon_sp_lat(gds)) * units
	sp_lon = double(GDS_RotLatLon_sp_lon(gds)) * units
	angle_rot = double(GDS_RotLatLon_rotation(gds)) * units

	a = (M_PI / 180.0) * (90.0 + sp_lat)
	b = (M_PI / 180.0) * sp_lon
	r = (M_PI / 180.0) * angle_rot

	sin_a = sin(a)
	cos_a = cos(a)

	tlat = *lat
	tlon = *lon
	tlat_index := 0
	tlon_index := 0
	for i = 0; i < npnts; i++ {
		pr = (M_PI / 180.0) * tlat[tlat_index]
		gr = -(M_PI / 180.0) * tlon[tlon_index]
		pm = asin(cos(pr) * cos(gr))
		gm = atan2(cos(pr)*sin(gr), -sin(pr))
		glat = (180.0 / M_PI) * (asin(sin_a*sin(pm) - cos_a*cos(pm)*cos(gm-r)))
		glon = -(180.0 / M_PI) * (-b + atan2(cos(pm)*sin(gm-r), sin_a*cos(pm)*cos(gm-r)+cos_a*sin(pm)))
		tlat[tlat_index] = glat
		tlat_index++
		tlon[tlon_index] = glon
		tlon_index++
	}
	return nil
}

func mercator2ll(sec [][]unsigned_char, lat *[]double, lon *[]double) error {

	var dx, dy, lat1, lat2, lon1, lon2 double
	var llat, llon []double
	var i, j unsigned_int
	var dlon, circum double

	var n, s, e, w, tmp, error_ double
	var gds []unsigned_char

	var nnx, nny unsigned_int
	var nres, nscan int
	var nnpnts unsigned_int

	var n_variable_dim int
	var variable_dim, raw_variable_dim []int

	get_nxny_(sec, &nnx, &nny, &nnpnts, &nres, &nscan, &n_variable_dim, &variable_dim, &raw_variable_dim)
	gds = sec[3]

	dy = GDS_Mercator_dy(gds)
	dx = GDS_Mercator_dx(gds)
	lat1 = GDS_Mercator_lat1(gds)
	lat2 = GDS_Mercator_lat2(gds)
	lon1 = GDS_Mercator_lon1(gds)
	lon2 = GDS_Mercator_lon2(gds)

	if lon1 < 0.0 || lon2 < 0.0 || lon1 > 360.0 || lon2 > 360.0 {
		return fatal_error("BAD GDS lon", "")
	}
	if lat1 < -90.0 || lat2 < -90.0 || lat1 > 90.0 || lat2 > 90.0 {
		return fatal_error("BAD GDS lat", "")
	}

	if GDS_Mercator_ori_angle(gds) != 0.0 {
		return fprintf("cannot handle non-zero mercator orientation angle %f\n",
			double(GDS_Mercator_ori_angle(gds)))
	}

	if nnx < 1 || nny < 1 {
		return fprintf("Sorry geo/mercator code does not handle variable nx/ny yet\n")
	}

	/*
			if ((*lat = (double *) malloc(((size_t) nnpnts) * sizeof(double))) == NULL) {
		        fatal_error("mercator2ll memory allocation failed","");
		    }
		    if ((*lon = (double *) malloc(((size_t) nnpnts) * sizeof(double))) == NULL) {
		        fatal_error("mercator2ll memory allocation failed","");
		    }
	*/
	*lat = make([]double, nnpnts, nnpnts)
	*lon = make([]double, nnpnts, nnpnts)

	/* now figure out the grid coordinates mucho silly grib specification */

	/* find S and N latitude */
	if GDS_Scan_y(nscan) {
		s = lat1
		n = lat2
	} else {
		s = lat2
		n = lat1
	}
	if s > n {
		fatal_error("Mercator grid: lat1 and lat2", "")
	}

	/* find W and E longitude */

	if ((nscan & 16) == 16) && (nny%2 == 0) && ((nres & 32) == 0) {
		fatal_error("grib GDS ambiguity", "")
	}

	if ((nscan & 16) == 16) && (nny%2 == 0) {
		fatal_error("more code needed to decode GDS", "")
	}

	if GDS_Scan_x(nscan) {
		w = lon1
		e = lon2
	} else {
		w = lon2
		e = lon1
	}
	if e <= w {
		e += 360.0
	}

	llat = *lat
	llon = *lon

	dlon = (e - w) / double(nnx-1)
	radius, err := radius_earth(sec)
	if err != nil {
		return fatal_error_wrap(err, "Failed to execute radius_earth")
	}
	circum = 2.0 * M_PI * radius * cos(double(GDS_Mercator_latD(gds))*(M_PI/180.0))
	dx = dx * 360.0 / circum

	// dlon should be almost == to dx
	// replace dx by dlon to get end points to match

	if dx != 0.0 {
		error_ = fabs(dx-dlon) / fabs(dx)
		if error_ >= 0.001 {
			return fprintf("\n*** Mercator grid error: inconsistent d-longitude, radius and grid domain\n*** d-longitude from grid domain %f (used), d-longitude from dx %f (not used)\n", dlon, dx)
		}
		dx = dlon
	}

	s = log(tan((45 + s/2) * M_PI / 180))
	n = log(tan((45 + n/2) * M_PI / 180))
	dy = (n - s) / (double(nny) - 1)

	llat_index := 0
	for j = 0; j < nny; j++ {
		tmp = (atan(exp(s+double(j)*dy))*180/M_PI - 45) * 2
		for i = 0; i < nnx; i++ {
			llat[llat_index] = tmp
			llat_index++
		}
	}

	for j = 0; j < nnx; j++ {
		if w+double(j)*dx >= 360.0 {
			llon[j] = w + double(j)*dx - 360.0
		} else {
			llon[j] = w + double(j)*dx
		}
	}
	for j = nnx; j < nnpnts; j++ {
		llon[j] = llon[j-nnx]
	}
	return nil
} /* end mercator2ll() */

func polar2ll(sec [][]unsigned_char, llat *[]double, llon *[]double) error {

	var lat, lon []double
	var gds []unsigned_char

	var dx, dy, orient, de, de2, dr, tmp, xp, yp, h, lat1, lon1, dr2 double
	var di, dj, LatD double
	var ix, iy unsigned_int
	var nnx, nny unsigned_int
	var nres, nscan int
	var nnpnts unsigned_int
	var nx size_t

	var n_variable_dim int
	var variable_dim, raw_variable_dim []int

	get_nxny_(sec, &nnx, &nny, &nnpnts, &nres, &nscan, &n_variable_dim, &variable_dim, &raw_variable_dim)

	gds = sec[3]

	if nnx < 1 || nny < 1 {
		return fprintf("Sorry code does not handle variable nx/ny yet\n")
	}
	nx = size_t(nnx) /* size_t, multiplications will not overflow */

	/*
			if ((*llat = (double *) malloc(((size_t) nnpnts) * sizeof(double))) == NULL) {
		        fatal_error("polar2ll memory allocation failed","");
		    }
		    if ((*llon = (double *) malloc(((size_t) nnpnts) * sizeof(double))) == NULL) {
		        fatal_error("polar2ll memory allocation failed","");
			}
	*/
	*llat = make([]double, nnpnts, nnpnts)
	*llat = make([]double, nnpnts, nnpnts)

	lat = *llat
	lon = *llon

	/* based on iplib */

	lat1 = GDS_Polar_lat1(gds) * (M_PI / 180)
	lon1 = GDS_Polar_lon1(gds)
	orient = GDS_Polar_lov(gds)
	LatD = GDS_Polar_lad(gds)

	lon1 *= (M_PI / 180.0)
	orient *= (M_PI / 180.0)

	dy = GDS_Polar_dy(gds)
	dx = GDS_Polar_dx(gds)

	h = 1.0
	if GDS_Polar_sps(gds) {
		h = -1.0
		/* added 12/19/2008 WNE sps checkout */
		orient -= M_PI
	}

	// removed 12/11    if (! (GDS_Scan_x(nscan))) dx = -dx;
	// removed 12/11    if (! (GDS_Scan_y(nscan))) dy = -dy;

	/* 60 probably becomes something else in grib2 */
	/* vsm: from comment to grib2 polar template:
	   "Grid length is in units of 10-3 m at the latitude specified by LaD"
	    do use GDS_Polar_lad(gds) instead of 60?
	    Do use fabs for southern hemisphere?
	*/

	radius, err := radius_earth(sec)
	if err != nil {
		return fatal_error_wrap(err, "Failed to execute radius_earth")
	}
	de = (1.0 + sin(fabs(LatD)*(M_PI/180.0))) * radius
	dr = de * cos(lat1) / (1 + h*sin(lat1))

	xp = -h * sin(lon1-orient) * dr / dx
	yp = cos(lon1-orient) * dr / dy

	// added 12/11
	if !(GDS_Scan_y(nscan)) {
		yp = yp - double(nny) + 1
	}
	if !(GDS_Scan_x(nscan)) {
		xp = xp - double(nnx) + 1
	}

	de2 = de * de
	//#pragma omp parallel for private(iy,ix,di,dj,dr2,tmp)
	for iy = 0; iy < nny; iy++ {
		for ix = 0; ix < nnx; ix++ {
			di = (double(ix) - xp) * dx
			dj = (double(iy) - yp) * dy
			dr2 = di*di + dj*dj
			if dr2 < de2*1e-6 {
				lon[ix+iy*unsigned_int(nx)] = 0.0
				lat[ix+iy*unsigned_int(nx)] = h * 90.0
			} else {
				tmp = (orient + h*atan2(di, -dj)) * (180.0 / M_PI)
				if tmp < 0.0 {
					tmp += 360.0
				}
				if tmp > 360.0 {
					tmp -= 360.0
				}
				lon[ix+iy*unsigned_int(nx)] = tmp
				lat[ix+iy*unsigned_int(nx)] = h * asin((de2-dr2)/(de2+dr2)) * (180.0 / M_PI)
			}
		}
	}
	return nil
}

func lambert2ll(sec [][]unsigned_char, llat *[]double, llon *[]double) error {

	var n double
	var lat, lon []double

	var dx, dy, lat1r, lon1r, lon2d, lon2r, latin1r, latin2r double
	var lond, latd, d_lon double
	var f, rho, rhoref, theta, startx, starty double
	var nnx, nny unsigned_int
	var nres, nscan int
	var x, y, tmp double
	var gds []unsigned_char
	var latDr double
	var earth_radius double
	var j, nnpnts unsigned_int

	var n_variable_dim int
	var variable_dim, raw_variable_dim []int

	get_nxny_(sec, &nnx, &nny, &nnpnts, &nres, &nscan, &n_variable_dim, &variable_dim, &raw_variable_dim)

	if nnx < 1 || nny < 1 {
		return fprintf("Sorry code does not handle variable nx/ny yet\n")
	}

	earth_radius, err := radius_earth(sec)
	if err != nil {
		return fatal_error_wrap(err, "Failed to execute radius_earth")
	}
	gds = sec[3]
	dy = GDS_Lambert_dy(gds)
	dx = GDS_Lambert_dx(gds)
	lat1r = GDS_Lambert_La1(gds) * (M_PI / 180.0)
	lon1r = GDS_Lambert_Lo1(gds) * (M_PI / 180.0)
	lon2d = GDS_Lambert_Lov(gds)
	lon2r = lon2d * (M_PI / 180.0)
	latin1r = GDS_Lambert_Latin1(gds) * (M_PI / 180.0)
	latin2r = GDS_Lambert_Latin2(gds) * (M_PI / 180.0)

	//  fix for theta start value crossing 0 longitude
	//    if ((lon1r - lon2r) > 0) lon2r = lon2r + 2*M_PI;

	//
	// Latitude of "false origin" where scales are defined.
	// It is used to estimate "reference_R", rhoref.
	// Often latDr == latin1r == latin2r and non-modified code is true and works fine.
	// But could be different if intersection latitudes latin1r and latin2r are different.
	// Usually latDr must be latin1r <=  latDr <= latin2r, other could be strange.
	//
	latDr = GDS_Lambert_LatD(gds) * (M_PI / 180.0)

	if lon1r < 0 {
		return fatal_error("bad GDS, lon1r < 0.0", "")
	}

	if fabs(latin1r-latin2r) < 1E-09 {
		n = sin(latin1r)
	} else {
		n = log(cos(latin1r)/cos(latin2r)) /
			log(tan(M_PI_4+latin2r/2.0)/tan(M_PI_4+latin1r/2.0))
	}

	f = (cos(latin1r) * pow(tan(M_PI_4+latin1r/2.0), n)) / n

	rho = earth_radius * f * pow(tan(M_PI_4+lat1r/2.0), -n)
	// old rhoref = earth_radius * f * pow(tan(M_PI_4 + latin1r/2.0),-n);
	rhoref = earth_radius * f * pow(tan(M_PI_4+latDr/2.0), -n)

	// 2/2009 .. new code
	d_lon = lon1r - lon2r
	if d_lon > M_PI {
		d_lon -= 2 * M_PI
	}
	if d_lon < -M_PI {
		d_lon += 2 * M_PI
	}
	theta = n * d_lon
	// 2/2009 theta = n * (lon1r - lon2r);

	startx = rho * sin(theta)
	starty = rhoref - rho*cos(theta)

	/*
		if ((*llat = (double *) malloc(((size_t) nnpnts) * sizeof(double))) == NULL) {
			fatal_error("lambert2ll memory allocation failed","");
		}
		if ((*llon = (double *) malloc(((size_t) nnpnts) * sizeof(double))) == NULL) {
			fatal_error("lambert2ll memory allocation failed","");
		}
	*/
	*llat = make([]double, nnpnts, nnpnts)
	*llat = make([]double, nnpnts, nnpnts)

	lat = *llat
	lon = *llon

	/* put x[] and y[] values in lon[] and lat[] */
	if err := stagger(sec, nnpnts, lon, lat, &n_variable_dim, &variable_dim, &raw_variable_dim); err != nil {
		fatal_error_wrap(err, "geo: stagger problem")
	}

	dx = fabs(dx)
	dy = fabs(dy)

	//	#pragma omp parallel for private(j,x,y,tmp,theta,rho,lond,latd)
	for j = 0; j < nnpnts; j++ {
		y = starty + lat[j]*dy
		x = startx + lon[j]*dx
		tmp = rhoref - y
		theta = atan(x / tmp)
		rho = sqrt(x*x + tmp*tmp)
		//rho = n > 0 ? rho : -rho;
		if n > 0 {
			rho = rho
		} else {
			rho = -rho
		}
		lond = lon2d + todegrees(theta/n)
		latd = todegrees(2.0*atan(pow(earth_radius*f/rho, 1.0/n)) - M_PI_2)
		// lond = lond >= 360.0 ? lond - 360.0 : lond;
		if lond >= 360.0 {
			lond = lond - 360.0
		}
		// lond = lond < 0.0 ? lond + 360.0 : lond;
		if lond < 0.0 {
			lond = lond + 360.0
		}
		lon[j] = lond
		lat[j] = latd
	}
	return nil
} /* end lambert2ll() */
