package internal

const NCEP = 7
const JMA1 = 34
const JMA2 = 35

// #define INT1(a)   ((a & 0x80) ? - (int) (a & 127) : (int) (a & 127))
func INT1(a unsigned_char) int {
	if (a & 0x80) != 0 {
		return -int(a & 127)
	}
	return int(a & 127)
}

// #define UINT2(a,b) ((int) ((a << 8) + (b)))
func UINT2(a, b unsigned_char) int {
	return (int(a) << 8) + int(b)
}

// #define GB2_Center(sec)			UINT2(sec[1][5], sec[1][6])
func GB2_Center(sec [][]unsigned_char) int {
	return UINT2(sec[1][5], sec[1][6])
}

// #define GB2_Sec3_npts(sec)		uint4(sec[3]+6)
func GB2_Sec3_npts(sec [][]unsigned_char) unsigned_int {
	return uint4(sec[3][6:])
}

// #define GDS_Scan_y(scan)		((scan & 64) == 64)
func GDS_Scan_y(scan int) bool {
	return (scan & 64) == 64
}

// #define GDS_Scan_staggered_storage(scan)	(((scan) & (1)) != 0)
func GDS_Scan_staggered_storage(scan int) bool {
	return ((scan) & (1)) != 0
}

// #define GDS_LatLon_basic_ang(gds)	int4(gds+38)
func GDS_LatLon_basic_ang(gds []unsigned_char) int {
	return int4(gds[38:])
}

// #define GDS_LatLon_sub_ang(gds)		int4(gds+42)
func GDS_LatLon_sub_ang(gds []unsigned_char) int {
	return int4(gds[42:])
}

// #define GDS_LatLon_lat1(gds)		int4(gds+46)
func GDS_LatLon_lat1(gds []unsigned_char) int {
	return int4(gds[46:])
}

// #define GDS_LatLon_lon1(gds)		uint4(gds+50)
func GDS_LatLon_lon1(gds []unsigned_char) unsigned_int {
	return uint4(gds[50:])
}

// #define GDS_LatLon_lat2(gds)		int4(gds+55)
func GDS_LatLon_lat2(gds []unsigned_char) int {
	return int4(gds[55:])
}

// #define GDS_LatLon_lon2(gds)		uint4(gds+59)
func GDS_LatLon_lon2(gds []unsigned_char) unsigned_int {
	return uint4(gds[59:])
}

// #define GDS_LatLon_dlon(gds)		int4(gds+63)
func GDS_LatLon_dlon(gds []unsigned_char) int {
	return int4(gds[63:])
}

// #define GDS_LatLon_dlat(gds)		int4(gds+67)
func GDS_LatLon_dlat(gds []unsigned_char) int {
	return int4(gds[67:])
}

// #define GDS_Scan_row_rev(scan)		((scan & 16) == 16)
func GDS_Scan_row_rev(scan int) bool {
	return (scan & 16) == 16
}

// #define GDS_Scan_x(scan)		((scan & 128) == 0)
func GDS_Scan_x(scan int) bool {
	return (scan & 128) == 0
}

// #define GB2_ProdDefTemplateNo(sec)	(UINT2(sec[4][7],sec[4][8]))
func GB2_ProdDefTemplateNo(sec [][]unsigned_char) int {
	return UINT2(sec[4][7], sec[4][8])
}

// #define GB2_Discipline(sec)		((int) (sec[0][6]))
func GB2_Discipline(sec [][]unsigned_char) int {
	return int(sec[0][6])
}

// #define GB2_MasterTable(sec)		((int) (sec[1][9]))
func GB2_MasterTable(sec [][]unsigned_char) int {
	return int(sec[1][9])
}

// #define GB2_LocalTable(sec)		((int) (sec[1][10]))
func GB2_LocalTable(sec [][]unsigned_char) int {
	return int(sec[1][10])
}

// #define GB2_ParmCat(sec)		(sec[4][9])
func GB2_ParmCat(sec [][]unsigned_char) int {
	return int(sec[4][9])
}

// #define GB2_ParmNum(sec)		(sec[4][10])
func GB2_ParmNum(sec [][]unsigned_char) int {
	return int(sec[4][10])
}

// #define GB2_Subcenter(sec)		UINT2(sec[1][7], sec[1][8])
func GB2_Subcenter(sec [][]unsigned_char) int {
	return UINT2(sec[1][7], sec[1][8])
}

// #define GDS_RotLatLon_sp_lat(gds)	(int4(gds+72))
func GDS_RotLatLon_sp_lat(gds []unsigned_char) int {
	return int4(gds[72:])
}

// #define GDS_RotLatLon_sp_lon(gds)	(uint4(gds+76))
func GDS_RotLatLon_sp_lon(gds []unsigned_char) unsigned_int {
	return uint4(gds[76:])
}

// #define GDS_RotLatLon_rotation(gds)	(int4(gds+80))
func GDS_RotLatLon_rotation(gds []unsigned_char) unsigned_int {
	return uint4(gds[80:])
}

// #define GDS_Mercator_dy(gds)		((uint4(gds+68))*0.001)
func GDS_Mercator_dy(gds []unsigned_char) double {
	return double(uint4(gds[68:])) * 0.001
}

// #define GDS_Mercator_dx(gds)		((uint4(gds+64))*0.001)
func GDS_Mercator_dx(gds []unsigned_char) double {
	return double(uint4(gds[64:])) * 0.001
}

// #define GDS_Mercator_lat1(gds)		(int4(gds+38)*0.000001)
func GDS_Mercator_lat1(gds []unsigned_char) double {
	return double(int4(gds[38:])) * 0.000001
}

// #define GDS_Mercator_lat2(gds)		(int4(gds+51)*0.000001)
func GDS_Mercator_lat2(gds []unsigned_char) double {
	return double(int4(gds[51:])) * 0.000001
}

// #define GDS_Mercator_lon1(gds)		(uint4(gds+42)*0.000001)
func GDS_Mercator_lon1(gds []unsigned_char) double {
	return double(uint4(gds[42:])) * 0.000001
}

// #define GDS_Mercator_lon2(gds)		(uint4(gds+55)*0.000001)
func GDS_Mercator_lon2(gds []unsigned_char) double {
	return double(uint4(gds[55:])) * 0.000001
}

// #define GDS_Mercator_ori_angle(gds)	(uint4(gds+60)*0.000001)
func GDS_Mercator_ori_angle(gds []unsigned_char) double {
	return double(uint4(gds[60:])) * 0.000001
}

// #define GDS_Mercator_latD(gds)		(int4(gds+47)*0.000001)
func GDS_Mercator_latD(gds []unsigned_char) double {
	return double(int4(gds[47:])) * 0.000001
}

// #define GDS_Polar_lat1(gds)		(int4(gds+38)*0.000001)
func GDS_Polar_lat1(gds []unsigned_char) double {
	return double(int4(gds[38:])) * 0.000001
}

// #define GDS_Polar_lon1(gds)		(uint4(gds+42)*0.000001)
func GDS_Polar_lon1(gds []unsigned_char) double {
	return double(uint4(gds[42:])) * 0.000001
}

// #define GDS_Polar_lov(gds)		(uint4(gds+51)*0.000001)
func GDS_Polar_lov(gds []unsigned_char) double {
	return double(uint4(gds[51:])) * 0.000001
}

// #define GDS_Polar_lad(gds)		(int4(gds+47)*0.000001)
func GDS_Polar_lad(gds []unsigned_char) double {
	return double(int4(gds[47:])) * 0.000001
}

// #define GDS_Polar_dx(gds)		(uint4(gds+55)*0.001)
func GDS_Polar_dx(gds []unsigned_char) double {
	return double(uint4(gds[55:])) * 0.001
}

// #define GDS_Polar_dy(gds)		(uint4(gds+59)*0.001)
func GDS_Polar_dy(gds []unsigned_char) double {
	return double(uint4(gds[59:])) * 0.001
}

// #define GDS_Polar_sps(gds)		((gds[63] & 128) == 128)
func GDS_Polar_sps(gds []unsigned_char) bool {
	return (gds[63] & 128) == 128
}

// #define GDS_Lambert_dy(gds)		(int4(gds+59) * 0.001)
func GDS_Lambert_dy(gds []unsigned_char) double {
	return double(int4(gds[59:])) * 0.001
}

// #define GDS_Lambert_dx(gds)		(int4(gds+55) * 0.001)
func GDS_Lambert_dx(gds []unsigned_char) double {
	return double(int4(gds[55:])) * 0.001
}

// #define GDS_Lambert_La1(gds)		(int4(gds+38) * 0.000001)
func GDS_Lambert_La1(gds []unsigned_char) double {
	return double(int4(gds[38:])) * 0.000001
}

// #define GDS_Lambert_Lo1(gds)		(int4(gds+42) * 0.000001)
func GDS_Lambert_Lo1(gds []unsigned_char) double {
	return double(int4(gds[42:])) * 0.000001
}

// #define GDS_Lambert_Lov(gds)		(int4(gds+51) * 0.000001)
func GDS_Lambert_Lov(gds []unsigned_char) double {
	return double(int4(gds[51:])) * 0.000001
}

// #define GDS_Lambert_Latin1(gds)		(int4(gds+65) * 0.000001)
func GDS_Lambert_Latin1(gds []unsigned_char) double {
	return double(int4(gds[65:])) * 0.000001
}

// #define GDS_Lambert_Latin2(gds)		(int4(gds+69) * 0.000001)
func GDS_Lambert_Latin2(gds []unsigned_char) double {
	return double(int4(gds[69:])) * 0.000001
}

// #define GDS_Lambert_LatD(gds)		(int4(gds+47) * 0.000001)
func GDS_Lambert_LatD(gds []unsigned_char) double {
	return double(int4(gds[47:])) * 0.000001
}
