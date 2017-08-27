package grib2

import (
	"encoding/binary"
	"io/ioutil"
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

func loadSections(data []byte) ([][][]byte, error) {
	messages := make([][][]byte, 0)

	for len(data) > 0 {

		bSections := make([][]byte, 8)

		if string(data[0:4]) != "GRIB" {
			return messages, errors.New("First 4 bytes of raw data must be 'GRIB'")
		}

		sectionSize := int(16)
		endOfSections := false

		// section0
		currentSection := data[:16]
		bSections[0] = currentSection

		// get message length in section0
		messageSize := int(binary.BigEndian.Uint64(currentSection[8:]))
		message := data[:messageSize]

		// jump over section0
		message = message[16:]
		data = data[messageSize:]

		// Retrieve all the 7 other sections in the message
		for !endOfSections {
			if string(message[0:4]) == "7777" { // this is the signature of section8 (end of sections)
				message = message[4:]
				endOfSections = true
			} else {
				sectionSize = int(binary.BigEndian.Uint32(message[0:4]))
				currentSection = message[0:sectionSize]
				message = message[sectionSize:]
				sectionType := currentSection[4]
				bSections[sectionType] = currentSection
			}
		}

		messages = append(messages, bSections)

	}

	return messages, nil
}

// LoadGrib load all grib data from a file
func LoadGrib(filename string) ([]GRIB2, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return Read(data)
}

// Read reads raw GRIB2 files and return slice of structured GRIB2 data
func Read(data []byte) ([]GRIB2, error) {
	if data == nil {
		return nil, errors.New("Raw data is nil")
	}

	if len(data) < 4 {
		return nil, errors.New("Raw data should be 4 bytes at least")
	}

	sectionGroups, err := loadSections(data)
	if err != nil {
		return nil, err
	}

	gribs := make([]GRIB2, 0)

	for _, sections := range sectionGroups {
		grib := GRIB2{
			Values: []Value{},
		}

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
		if err != nil {
			return gribs, err
		}
		c := len(lon)
		v := make([]Value, c, c)
		for i := 0; i < c; i++ {
			v[i].Longitude = lon[i]
			v[i].Latitude = lat[i]
			v[i].Value = raw[i]
		}

		grib.Values = append(grib.Values, v...)

		gribs = append(gribs, grib)
	}

	return gribs, nil
}
