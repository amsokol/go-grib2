package internal

import (
	"fmt"
	"math"
	"time"
	"unsafe"

	"github.com/pkg/errors"
)

type unsigned_char byte
type char byte
type unsigned_int uint
type double float64
type float float32
type size_t unsigned_int

/*
unsigned int uint4(unsigned const char *p) {
    return ((p[0] << 24) + (p[1] << 16) + (p[2] << 8) + p[3]);
}
*/
func uint4(p []unsigned_char) unsigned_int {
	return ((unsigned_int(p[0]) << 24) + (unsigned_int(p[1]) << 16) + (unsigned_int(p[2]) << 8) + unsigned_int(p[3]))
}

/*
int int4(unsigned const char *p) {
	int i;
	if (p[0] & 0x80) {
		i = -(((p[0] & 0x7f) << 24) + (p[1] << 16) + (p[2] << 8) + p[3]);
	}
	else {
		i = (p[0] << 24) + (p[1] << 16) + (p[2] << 8) + p[3];
	}
	return i;
}
*/
func int4(p []unsigned_char) int {
	var i int
	if (p[0] & 0x80) != 0 {
		i = -(((int(p[0]) & 0x7f) << 24) + (int(p[1]) << 16) + (int(p[2]) << 8) + int(p[3]))
	} else {
		i = (int(p[0]) << 24) + (int(p[1]) << 16) + (int(p[2]) << 8) + int(p[3])
	}
	return i
}

func fabs(v double) double {
	return double(math.Abs(float64(v)))
}

func ldexp(frac double, exp int) double {
	return double(math.Ldexp(float64(frac), exp))
}

func pow(x, y double) double {
	return double(math.Pow(float64(x), float64(y)))
}

func sqrt(x double) double {
	return double(math.Sqrt(float64(x)))
}

func Int_Power(x double, y int) double {
	return double(math.Pow(float64(x), float64(y)))
}

func exp(v double) double {
	return double(math.Exp(float64(v)))
}

func sin(v double) double {
	return double(math.Sin(float64(v)))
}

func cos(v double) double {
	return double(math.Cos(float64(v)))
}

func asin(v double) double {
	return double(math.Asin(float64(v)))
}

func atan2(x, y double) double {
	return double(math.Atan2(float64(x), float64(y)))
}

func log(v double) double {
	return double(math.Log(float64(v)))
}

func tan(v double) double {
	return double(math.Tan(float64(v)))
}

func atan(v double) double {
	return double(math.Atan(float64(v)))
}

func fatal_error(format string, args ...interface{}) error {
	return errors.Errorf(format, args)
}

func fatal_error_ii(format string, args ...interface{}) error {
	return errors.Errorf(format, args)
}

func fatal_error_i(format string, args ...interface{}) error {
	return errors.Errorf(format, args)
}

func fatal_error_wrap(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args)
}

func fprintf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}

func sprintf(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func LatLon(sec [][]byte, lon *[]float64, lat *[]float64) error {
	g_sec := *(*[][]unsigned_char)(unsafe.Pointer(&sec))
	g_lon := *(**[]double)(unsafe.Pointer(&lon))
	g_lat := *(**[]double)(unsafe.Pointer(&lat))

	return get_latlon(g_sec, g_lon, g_lat)
}

func UnpackData(sec [][]byte) ([]float32, error) {
	g_sec := *(*[][]unsigned_char)(unsafe.Pointer(&sec))

	ndata := GB2_Sec3_npts(g_sec)
	g_data := make([]float, ndata)
	err := unpk_grib(g_sec, g_data)
	if err != nil {
		return nil, err
	}

	return *(*[]float32)(unsafe.Pointer(&g_data)), nil
}

func RefTime(sec [][]byte) time.Time {
	var year, month, day, hour, minute, second int

	g_sec := *(*[][]unsigned_char)(unsafe.Pointer(&sec))
	reftime(g_sec, &year, &month, &day, &hour, &minute, &second)

	return time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)
}

func VerfTime(sec [][]byte) (time.Time, error) {
	var year, month, day, hour, minute, second int

	g_sec := *(*[][]unsigned_char)(unsafe.Pointer(&sec))

	err := verftime(g_sec, &year, &month, &day, &hour, &minute, &second)
	if err != nil {
		return time.Now(), errors.Wrapf(err, "Failed to run verftime")
	}

	return time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC), nil
}

func GetInfo(sec [][]byte) (name string, desc string, unit string, err error) {
	g_sec := *(*[][]unsigned_char)(unsafe.Pointer(&sec))

	err = getName(g_sec, &name, &desc, &unit)
	if err != nil {
		return "", "", "", errors.Wrapf(err, "Failed to execute getName")
	}
	return name, desc, unit, nil
}

func GetLevel(sec [][]byte) (level string, err error) {
	g_sec := *(*[][]unsigned_char)(unsafe.Pointer(&sec))

	err = f_lev(g_sec, &level)
	if err != nil {
		return "", errors.Wrapf(err, "Failed to execute f_lev")
	}
	return level, nil
}
