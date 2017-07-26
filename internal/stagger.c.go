package internal

/*
 * stagger fills x[] and y[], lo1 and la1 have X==0 and Y==0
 *
 * assumed_npnts is number if grid points that the calling program thinks is right
 *  this is for error checking.  use -1 if don't know
 *
 * to a grid transform:
 *   setup grid transform (proj4 library for example(
 *   call stagger() to get the x() and y() values of the grid
 *     transform x() and y() to lon() and lat()
 *
 * like many programs, stagger requires grid to be on we:sn order
 */
func stagger(sec [][]unsigned_char, assumed_npnts unsigned_int, x []double, y []double, n_variable_dim *int, variable_dim *[]int, raw_variable_dim *[]int) error {
	var nx, ny, res, scan int
	var npnts unsigned_int
	var nnx, nx_even, nx_odd, nx2 int
	var x0, y0, dx_offset, dx_offset_even, dx_offset_odd, dy_offset double
	var i, ix, iy, n unsigned_int

	var reduced_grid, dx_off_odd, dx_off_even, dy_off int
	var dx, dy, even int

	get_nxny(sec, &nx, &ny, &npnts, &res, &scan, n_variable_dim, variable_dim, raw_variable_dim)
	if scan == -1 {
		return fatal_error("scan == -1")
	}
	if output_order != wesn {
		return fatal_error("output_order != wesn")
	}
	if nx < 1 || ny < 1 {
		return fatal_error("nx < 1 || ny < 1")
	}

	/* get stagger bits */
	dx_off_odd = int((unsigned_int(scan) >> 3) & 1)
	dx_off_even = int((unsigned_int(scan) >> 2) & 1)
	dy_off = int((unsigned_int(scan) >> 1) & 1)
	reduced_grid = int(unsigned_int(scan) & 1)

	// dx =  (scan & 128) ? -1 : 1;
	if (scan & 128) != 0 {
		dx = -1
	} else {
		dx = 1
	}
	// dy =  (scan & 64) ? 1 : -1;
	if (scan & 64) != 0 {
		dy = 1
	} else {
		dy = -1
	}

	if reduced_grid != 0 && dy_off != 0 {
		ny--
	}

	if dy < 0 && ((ny % 2) == 0) { // swap even and odd rows if ns to sn and even number of rows
		i = unsigned_int(dx_off_odd)
		dx_off_odd = dx_off_even
		dx_off_even = int(i)
	}

	// dx_offset_odd  = reduced_grid ? 0.5 * dx_off_odd  : 0.5 * dx_off_odd  * dx;
	if reduced_grid != 0 {
		dx_offset_odd = 0.5 * double(dx_off_odd)
	} else {
		dx_offset_odd = 0.5 * double(dx_off_odd) * double(dx)
	}
	// dx_offset_even = reduced_grid ? 0.5 * dx_off_even : 0.5 * dx_off_even * dx;
	if reduced_grid != 0 {
		dx_offset_even = 0.5 * double(dx_off_even)
	} else {
		dx_offset_even = 0.5 * double(dx_off_even) * double(dx)
	}
	// dy_offset = reduced_grid ? 0.5 * dy_off : 0.5 * dy_off * dy;
	if reduced_grid != 0 {
		dy_offset = 0.5 * double(dy_off)
	} else {
		dy_offset = 0.5 * double(dy_off) * double(dy)
	}

	nx_odd = nx - (dx_off_odd & reduced_grid)
	nx_even = nx - (dx_off_even & reduced_grid)
	nx2 = nx_odd + nx_even

	//fprintf(stderr, "stagger: dx_off_odd %lf dx_off_even %lf dy_off %lf  reduced_grid %d nx=%d %d\n",
	//    dx_offset_odd, dx_offset_even, dy_offset, reduced_grid, nx_odd,nx_even);
	//fprintf(stderr,"dx_off_odd %d reduced_grid %d, and %d\n", dx_off_odd , reduced_grid, dx_off_odd & reduced_grid);
	//fprintf(stderr,"dx_off_even %d reduced_grid %d, and %d\n", dx_off_even , reduced_grid, dx_off_even & reduced_grid);

	// number of grid points
	n = unsigned_int((ny/2)*nx_even + ((ny+1)/2)*nx_odd)

	// check to number of points
	if assumed_npnts != n {
		return fatal_error_ii("stagger: program error think npnts=%d assumed npnts=%d", n, int(assumed_npnts))
	}
	if n != GB2_Sec3_npts(sec) {
		return fatal_error_ii("stagger: program error think npnts=%d, Sec3 gives %d", n, GB2_Sec3_npts(sec))
	}

	if x == nil || y == nil {
		return fatal_error("x == nil || y == nil")
	}

	/* return X[] and Y[] relative to the first grid point but on a we:sn grid */

	// x0 = (dx > 0) ? 0.0 : 1.0 - (double) nx;
	if dx > 0 {
		x0 = 0.0
	} else {
		x0 = 1.0 - double(nx)
	}
	//y0 = (dy > 0) ? 0.0 : 1.0 - (double) ny;
	if dy > 0 {
		y0 = 0.0
	} else {
		y0 = 1.0 - double(ny)
	}

	for iy = 0; iy < unsigned_int(ny); iy++ {
		// even = iy % 2;		// first row is odd .. iy % 2 == 0
		even = int(iy & 1) // first row is odd
		//i = even ?  nx2*(iy >> 1) + nx_odd : nx2*(iy >> 1);
		if even != 0 {
			i = unsigned_int(nx2)*(iy>>1) + unsigned_int(nx_odd)
		} else {
			i = unsigned_int(nx2) * (iy >> 1)
		}
		//nnx = even ? nx_even : nx_odd;
		if even != 0 {
			nnx = nx_even
		} else {
			nnx = nx_odd
		}
		//dx_offset = even ? dx_offset_even : dx_offset_odd;
		if even != 0 {
			dx_offset = dx_offset_even
		} else {
			dx_offset = dx_offset_odd
		}
		for ix = 0; ix < unsigned_int(nnx); ix++ {
			x[i+ix] = x0 + dx_offset + double(ix)
			y[i+ix] = y0 + dy_offset + double(iy)
		}
	}

	return nil
}
