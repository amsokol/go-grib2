package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/amsokol/go-grib2"
)

func main() {
	grib2file := "./gfs.t00z.pgrb2.0p25.f001"
	csvfile := "./gfs.t00z.pgrb2.0p25.f001.csv"

	infile, err := os.Open(grib2file)
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()

	data, err := ioutil.ReadAll(infile)
	if err != nil {
		log.Fatal(err)
	}

	gribs, err := grib2.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Source package contains %d GRIB2 file(s)\n", len(gribs))

	outfile, err := os.Create(csvfile)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	for _, g := range gribs {
		log.Printf("Published='%s', Forecast='%s', Parameter='%s', Unit='%s', Description='%s'\n",
			g.RefTime.Format("2006-01-02 15:04:05"), g.VerfTime.Format("2006-01-02 15:04:05"), g.Name, g.Unit, g.Description)

		refTime := g.RefTime
		verfTime := g.VerfTime
		name := g.Name
		level := g.Level
		for _, v := range g.Values {
			lon := v.Longitude
			if lon > 180.0 {
				lon -= 360.0
			}

			_, err := fmt.Fprintf(outfile, "\"%s\",\"%s\",\"%s\",\"%s\",%g,%g,%g\n",
				refTime.Format("2006-01-02 15:04:05"),
				verfTime.Format("2006-01-02 15:04:05"),
				name,
				level,
				lon,
				v.Latitude,
				v.Value)
			if err != nil {
				log.Fatalln("error writing record to csv:", err)
			}
		}
	}
}
