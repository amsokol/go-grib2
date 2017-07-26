package internal

/*
 * get the name information
 *
 * if inv_out, name, desc, unit == NULL, not used
 *
 * v1.0 Wesley Ebisuzaki 2006
 * v1.1 Wesley Ebisuzaki 4/2007 netcdf support
 * v1.2 Wesley Ebisuzaki 4/2007 multiple table support
 * v1.3 Wesley Ebisuzaki 6/2011 make parameter cat >= 192 local
 * v1.4 Wesley Ebisuzaki 2/2012 fixed search_gribtab for local tables
 * v1.5 Wesley Ebisuzaki 4/2013 gribtab -> gribtable, added user_gribtable
 */
func getName(sec [][]unsigned_char, name *string, desc *string, unit *string) error {

	var discipline /*center, mastertab, localtab,*/, parmcat, parmnum int
	var pdt int
	var p *gribtable_s
	var p_unit string

	/*
		if user_gribtable != nil {
			p = search_gribtable(user_gribtable, sec)
		}
	*/

	p, err := search_gribtable(gribtable, sec)
	if err != nil {
		return fatal_error_wrap(err, "Failed to execute search_gribtable")
	}

	p_unit = "unit"
	if p != nil {
		p_unit = p.unit
		pdt = code_table_4_0(sec)
		if pdt == 5 || pdt == 9 {
			p_unit = "prob"
		}
	}

	if p != nil {
		*name = p.name
		*desc = p.desc
		*unit = p_unit
	} else {
		discipline = GB2_Discipline(sec)
		// center = GB2_Center(sec)
		// mastertab = GB2_MasterTable(sec)
		// localtab = GB2_LocalTable(sec)
		parmcat = GB2_ParmCat(sec)
		parmnum = GB2_ParmNum(sec)

		*name = sprintf("var%d_%d_%d", discipline, parmcat, parmnum)
		*desc = "desc"
		*unit = p_unit
	}
	return nil
}

func search_gribtable(gts []gribtable_s, sec [][]unsigned_char) (*gribtable_s, error) {
	var discipline, center, mastertab, localtab, parmcat, parmnum int
	var use_local_table bool
	// static int count = 0;

	discipline = GB2_Discipline(sec)
	center = GB2_Center(sec)
	mastertab = GB2_MasterTable(sec)
	localtab = GB2_LocalTable(sec)
	parmcat = GB2_ParmCat(sec)
	parmnum = GB2_ParmNum(sec)

	// use_local_table = (mastertab == 255) ? 1 : 0;
	use_local_table = mastertab == 255

	if (parmnum >= 192 && parmnum <= 254) || (parmcat >= 192 && parmcat <= 254) || (discipline >= 192 && discipline <= 254) {
		use_local_table = true
	}

	if use_local_table && localtab == 0 {
		//return nil, fprintf("**** ERROR: local table = 0 is not allowed, set to 1 ***\n")
		localtab = 1
	}
	if use_local_table && localtab == 255 {
		return nil, fatal_error("local gribtable is undefined (255)", "")
	}

	if !use_local_table {
		for _, p := range gts {
			if discipline == p.disc && (mastertab >= p.mtab_low) && (mastertab <= p.mtab_high) &&
				parmcat == p.pcat && parmnum == p.pnum {
				return &p, nil
			}
		}
	} else {
		//	printf(">> cname local find: disc %d center %d localtab %d pcat %d pnum %d\n", discipline, center, localtab, parmcat, parmnum);
		for _, p := range gts {
			if discipline == p.disc && center == p.cntr && localtab == p.ltab && parmcat == p.pcat && parmnum == p.pnum {
				return &p, nil
			}
		}
	}
	return nil, nil
}
