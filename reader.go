package grib2

import (
	"encoding/binary"
	"time"

	"github.com/amsokol/go-grib2/internal"
	"github.com/pkg/errors"
)

// GRIB2 is simplified GRIB2 file structure
type GRIB2 struct {
	RefTime     time.Time
	VerfTime    time.Time
	Name        string
	Description string
	Unit        string
	Level       string
	Values      []Value
}

// Value is data item of GRIB2 file
type Value struct {
	Longitude float64
	Latitude  float64
	Value     float32
}

// Read reads raw GRIB2 files and return slice of structured GRIB2 data
func Read(data []byte) ([]GRIB2, error) {
	if data == nil {
		return nil, errors.New("Raw data is nil")
	}

	dlen := len(data)

	if dlen < 4 {
		return nil, errors.New("Raw data should be 4 bytes at least")
	}

	gribs := []GRIB2{}

	start := 0
	eod := false
	for !eod {
		if string(data[0:4]) != "GRIB" {
			return nil, errors.New("First 4 bytes of raw data must be 'GRIB'")
		}

		grib := GRIB2{
			Values: []Value{},
		}

		sections := [][]byte{nil, nil, nil, nil, nil, nil, nil, nil}

		size := 16
		sections[0] = data[start : start+size]
		start += size

		prv := -1
		cur := 0
		eof := false
		for !eof {
			prv = cur
			if prv == 7 {
				// block is read -> export data to values

				grib.RefTime = internal.RefTime(sections)

				var err error
				grib.VerfTime, err = internal.VerfTime(sections)
				if err != nil {
					return nil, errors.Wrapf(err, "Failed to get VerfTime")
				}

				grib.Name, grib.Description, grib.Unit, err = internal.GetInfo(sections)
				if err != nil {
					return nil, errors.Wrapf(err, "Failed to GetInfo")
				}

				grib.Level, err = internal.GetLevel(sections)
				if err != nil {
					return nil, errors.Wrapf(err, "Failed to GetLevel")
				}

				var lon, lat []float64
				err = internal.LatLon(sections, &lon, &lat)
				if err != nil {
					return nil, errors.Wrapf(err, "Failed to get longitude and latitude")
				}
				raw, err := internal.UnpackData(sections)
				c := len(lon)
				v := make([]Value, c, c)
				for i := 0; i < c; i++ {
					v[i].Longitude = lon[i]
					v[i].Latitude = lat[i]
					v[i].Value = raw[i]
				}

				grib.Values = append(grib.Values, v...)

				sections[2] = nil
				sections[3] = nil
				sections[4] = nil
				sections[5] = nil
				sections[6] = nil
				sections[7] = nil

				if string(data[start:start+4]) == "7777" {
					eof = true
					size = 4
				}
			} else {
				size = int(binary.BigEndian.Uint32(data[start:]))
				cur = int(data[start+4])
				sections[cur] = data[start : start+size]
			}
			start += size
		}

		gribs = append(gribs, grib)

		if start == dlen {
			eod = true
		}
	}

	return gribs, nil
}
