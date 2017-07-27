package internal

import (
	"time"
)

func get_time(p []unsigned_char, year *int, month *int, day *int, hour *int, minute *int, second *int) {
	*year = (int(p[0]) << 8) | int(p[1])
	*month = int(p[2])
	*day = int(p[3])
	*hour = int(p[4])
	*minute = int(p[5])
	*second = int(p[6])
}

func add_time(year *int, month *int, day *int, hour *int, minute *int, second *int, dtime int, unit int) {
	if dtime == 0 {
		return
	}
	add_dt(year, month, day, hour, minute, second, dtime, unit)
}

func add_dt(year *int, month *int, day *int, hour *int, minute *int, second *int, dtime int, unit int) {

	if unit == 255 || unit == -1 {
		return
	} // no valid time unit
	if dtime == 0xffffffff {
		return
	} // missing dtime

	t := time.Date(*year, time.Month(*month), *day, *hour, *minute, *second, 0, time.UTC)

	switch unit {
	case YEAR:
		t = t.AddDate(dtime, 0, 0)
	case DECADE:
		t = t.AddDate(10*dtime, 0, 0)
	case CENTURY:
		t = t.AddDate(100*dtime, 0, 0)
	case NORMAL:
		t = t.AddDate(30*dtime, 0, 0)
	case MONTH:
		t = t.AddDate(0, dtime, 0)
	case DAY:
		t = t.AddDate(0, 0, dtime)
	case HOUR12:
		t = t.Add(time.Duration(dtime) * time.Hour * 12)
	case HOUR6:
		t = t.Add(time.Duration(dtime) * time.Hour * 6)
	case HOUR3:
		t = t.Add(time.Duration(dtime) * time.Hour * 3)
	case HOUR:
		t = t.Add(time.Duration(dtime) * time.Hour)
	case MINUTE:
		t = t.Add(time.Duration(dtime) * time.Minute)
	case SECOND:
		t = t.Add(time.Duration(dtime) * time.Second)
	}
	*year = t.Year()
	*month = int(t.Month())
	*day = t.Day()
	*hour = t.Hour()
	*minute = t.Minute()
	*second = t.Second()
}
