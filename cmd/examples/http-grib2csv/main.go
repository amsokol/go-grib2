package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/amsokol/go-grib2"
)

func main() {
	log.Println("Downloading data from HTTP server...")
	res, err := http.Get("http://nomads.ncep.noaa.gov/cgi-bin/filter_gfs_0p25_1hr.pl?file=gfs.t00z.pgrb2.0p25.f001&lev_10_m_above_ground=on&var_UGRD=on&var_VGRD=on&subregion=&leftlon=-10&rightlon=19&toplat=60&bottomlat=35.7&dir=%2Fgfs.2017072000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Data downloaded")

	csvfile := "./gfs.t00z.pgrb2.0p25.f001.csv"

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

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
