# GRIB2 library for Go programming language (golang)

## History

Once upon time I had a need to read the GRIB2 file in golang. I used Google search and found the following libraries:

- [https://github.com/analogic/grib](https://github.com/analogic/grib) - obsolete
- [https://github.com/brockwood/gogrib2](https://github.com/brockwood/gogrib2) - empty
- [https://github.com/nilsmagnus/grib](https://github.com/nilsmagnus/grib)

Library `nilsmagnus/grib` looked reassuringly. I downloaded [gfs.t00z.pgrb2.0p25.f001](http://nomads.ncep.noaa.gov/cgi-bin/filter_gfs_0p25_1hr.pl?file=gfs.t00z.pgrb2.0p25.f001&lev_10_m_above_ground=on&var_UGRD=on&var_VGRD=on&subregion=&leftlon=-10&rightlon=19&toplat=60&bottomlat=35.7&dir=%2Fgfs.2017072000) file from `nomads.ncep.noaa.gov`, ran `grib` export and got error `format is not supported` (((.

I decided to develop GRIB2 library by myself. I looked at `nilsmagnus/grib` and libraries for Java and Python. All libraries are attempts to understand how GRIB2 is arranged inside and to develop GRIB2 engine using appropriated programming language.

There was the reference GRIB2 implementation - [wgrib2](http://www.cpc.ncep.noaa.gov/products/wesley/wgrib2/) tool. Look at `wgrib2` source code. It's a lot of "crazy" `C` language code ))). As for me it's very difficult to understand how it works in details. There was the way to call `wgrib2` from golang application. But I rejected that due to the following reasons:

- `wgrib2` is not thread-safe. There are a lot of global read/write variables inside code.

- I prefer pure golang libraries

I saw the only one way - to port `wgrib2` to golang. I did not want to repeat approach of another libraries. I was not going to rewrite `wgrib2` to golang. I decided to port `wgrib2` engine `C` code to `golang` with minimal correction and develop simple wrapper. I used results of `wgrib2` export to CSV as the reference for testing. That approach allows:

- safe time (porting spent few days only)
- easy apply new features and patches from original `wgrib2` development stream

I used `grib2-v2.0.6c` version. Here is the result of my work. The work is on final stage.

## How-to

```bash
go get -u github.com/amsokol/go-grib2
```

It contains only one function that reads GRIB2 file:

```go
func Read(data []byte) ([]GRIB2, error)
```

where `GRIB2` is the structure with parsed data:

```go
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
```

Unlike `wgrib2` It's thread-safe.

See the following usage examples:

- [file-grib2csv](https://github.com/amsokol/go-grib2/tree/master/cmd/examples/file-grib2csv) - export GRIB2 file to CSV
- [http-grib2csv](https://github.com/amsokol/go-grib2/tree/master/cmd/examples/http-grib2csv) - export `nomads.ncep.noaa.gov` HTTP response to CSV

## Roadmap

Major work is finished. `go-grib2` is ready to use. I have a little bit to port:

- parse `longitude` and `latitude` are stored as `gauss`, `space_view`, `irr_grid`
- parse data are stored using `complex` and `run_length` algorithms

## Issues

`go-grib2` does not support `jpeg`, `png` and `aec` data package formats. The reason is `wgrib2` uses external `C` libraries to read these formats. If you have an idea how to easy port `jpeg`, `png` and `aec` please let me know.

## Need help

I need help is testing. Please test `go-grib2` with several GRIB2 files from several providers. Use `wgrib2` CSV export to validate results. Please create issues here.
