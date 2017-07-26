package internal

/*
 * uint4_missing
 * if missing return 0
 * uint4_missing is only used in Sec3.c where an undefined nx/ny == 0 is a good responce
 */
/*
 unsigned int uint4_missing(unsigned const char *p) {
    int t;

    t = p[0];
    t = t << 8 | p[1];
    t = t << 8 | p[2];
    t = t << 8 | p[3];

    if (t == 0xffffffff) return 0;
    return t;
}
*/
func uint4_missing(p []unsigned_char) unsigned_int {
	var t int

	t = int(p[0])
	t = t<<8 | int(p[1])
	t = t<<8 | int(p[2])
	t = t<<8 | int(p[3])

	if t == 0xffffffff {
		return 0
	}
	return unsigned_int(t)
}

/*
int int2(unsigned const char *p) {
	int i;
	if (p[0] & 0x80) {
		i = -(((p[0] & 0x7f) << 8) + p[1]);
	}
	else {
		i = (p[0] << 8) + p[1];
	}
	return i;
}
*/
func int2(p []unsigned_char) int {
	var i int
	if (p[0] & 0x80) != 0 {
		i = -(((int(p[0]) & 0x7f) << 8) + int(p[1]))
	} else {
		i = (int(p[0]) << 8) + int(p[1])
	}
	return i
}

/*
unsigned int uint2(unsigned char const *p) {
	return (p[0] << 8) + p[1];
}
*/
func uint2(p []unsigned_char) unsigned_int {
	return (unsigned_int(p[0]) << 8) + unsigned_int(p[1])
}

//
// floating point values are often represented as int * power of 10
//
func scaled2flt(scale_factor int, scale_value int) float {
	if scale_factor == 0 {
		return float(scale_value)
	}
	if scale_factor < 0 {
		return float(double(scale_value) * Int_Power(10.0, -scale_factor))
	}
	return float(double(scale_value) / Int_Power(10.0, scale_factor))
}

func scaled2dbl(scale_factor int, scale_value int) double {
	if scale_factor == 0 {
		return double(scale_value)
	}
	if scale_factor < 0 {
		return double(scale_value) * Int_Power(10.0, -scale_factor)
	}
	return double(scale_value) / Int_Power(10.0, scale_factor)
}
