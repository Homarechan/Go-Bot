package datas

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var mainURL = []byte("http://www.ekidata.jp/api/")

// Pref is struct to parse the returned data by GetPrefData
type Pref struct {
	Pref struct {
		Code int    `xml:"code"`
		Name string `xml:"name"`
	} `xml:"pref"`
	Lines []struct {
		Code int    `xml:"line_cd"`
		Name string `xml:"line_name"`
	} `xml:"line"`
}

// Line is struct to parse the returned data by GetLineData
type Line struct {
	Line struct {
		Code       int     `xml:"line_cd"`
		Name       string  `xml:"line_name"`
		Longtitude float64 `xml:"line_lon"`
		Latitude   float64 `xml:"line_lat"`
		Zoom       int     `xml:"line_zoom"`
	} `xml:"line"`
	Stations []struct {
		Code       int     `xml:"station_cd"`
		GroupCode  int     `xml:"station_g_cd"`
		Name       string  `xml:"station_name"`
		Longtitude float64 `xml:"lon"`
		Latitude   float64 `xml:"lat"`
	} `xml:"station"`
}

// Station is struct to parse the returned data by GetStationData
type Station struct {
	Station struct {
		PrefCode   int     `xml:"pref_cd"`
		LineCode   int     `xml:"line_cd"`
		LineName   string  `xml:"line_name"`
		Code       int     `xml:"station_cd"`
		GroupCode  int     `xml:"station_g_cd"`
		Name       string  `xml:"station_name"`
		Longtitude float64 `xml:"lon"`
		Latitude   float64 `xml:"lat"`
	} `xml:"station"`
}

// Group is struct to parse the returned data by GetGroupData
type Group struct {
	Station struct {
		LineCode   int     `xml:"line_cd"`
		LineName   string  `xml:"line_name"`
		Code       int     `xml:"station_cd"`
		GroupCode  int     `xml:"station_g_cd"`
		Name       string  `xml:"station_name"`
		Longtitude float64 `xml:"lon"`
		Latitude   float64 `xml:"lat"`
	} `xml:"station"`
	GroupStations []struct {
		PrefCode int    `xml:"pref_cd"`
		LineCode int    `xml:"line_cd"`
		LineName string `xml:"line_name"`
		Code     int    `xml:"station_cd"`
		Name     string `xml:"station_name"`
	} `xml:"station_g"`
}

// Join is struct to parse the returned data by GetJoinData
type Join struct {
	StationJoins []struct {
		Code1       int     `xml:"station_cd1"`
		Code2       int     `xml:"station_cd2"`
		Name1       string  `xml:"station_name1"`
		Name2       string  `xml:"station_name2"`
		Longtitude1 float64 `xml:"lon1"`
		Latitude1   float64 `xml:"lat1"`
		Longtitude2 float64 `xml:"lon2"`
		Latitude2   float64 `xml:"lat2"`
	} `xml:"station_join"`
}

// GetPrefData returns data by prefecture
func GetPrefData(prefCode int) (Pref, error) {
	body, err := getData([]byte("p/"), prefCode)

	var inventory Pref

	if err != nil {
		return inventory, err
	}

	err = parse(body, &inventory)

	if err != nil {
		return inventory, err
	}

	return inventory, nil
}

// GetLineData returns data by prefecture
func GetLineData(lineCode int) (Line, error) {
	body, err := getData([]byte("l/"), lineCode)

	var inventory Line

	if err != nil {
		return inventory, err
	}

	err = parse(body, &inventory)

	if err != nil {
		return inventory, err
	}

	return inventory, nil
}

// GetStationData returns data by prefecture
func GetStationData(stationCode int) (Station, error) {
	body, err := getData([]byte("s/"), stationCode)

	var inventory Station

	if err != nil {
		return inventory, err
	}

	err = parse(body, &inventory)

	if err != nil {
		return inventory, err
	}

	return inventory, nil
}

// GetGroupData returns data by prefecture
func GetGroupData(stationCode int) (Group, error) {
	body, err := getData([]byte("g/"), stationCode)

	var inventory Group

	if err != nil {
		return inventory, err
	}

	err = parse(body, &inventory)

	if err != nil {
		return inventory, err
	}

	return inventory, nil
}

// GetJoinData returns data by prefecture
func GetJoinData(lineCode int) (Join, error) {
	body, err := getData([]byte("n/"), lineCode)

	var inventory Join

	if err != nil {
		return inventory, err
	}

	err = parse(body, &inventory)

	if err != nil {
		return inventory, err
	}

	return inventory, nil
}

func getData(endPoint []byte, code int) ([]byte, error) {
	url := make([]byte, 0, 10)
	url = mainURL
	url = append(url, endPoint...)
	url = append(url, strconv.Itoa(code)...)
	url = append(url, ".xml"...)

	res, err := http.Get(string(url))
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func parse(data []byte, v interface{}) error {
	return xml.Unmarshal(data, &v)
}
