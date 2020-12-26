package utils

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

type AddressDetail struct {
	City         string `json:"city"`
	CityCode     int    `json:"city_code"`
	Province     string `json:"province"`
	District     string `json:"district"`
	Street       string `json:"street"`
	StreetNumber string `json:"street_number"`
}

type Point struct {
	X string `json:"x"`
	Y string `json:"y"`
}

type Content struct {
	Address       string        `json:"address"`
	AddressDetail AddressDetail `json:"address_detail"`
	Point         Point         `json:"point"`
}

type Response struct {
	Address string  `json:"address"`
	Content Content `json:"content"`
	Status  int     `json:"status"`
}

func GetLocationByIp(ip string) (Response, error) {
	url := "http://api.map.baidu.com/location/ip?ak=%s&ip=%s&coor=bd09ll"
	res, err := http.Get(fmt.Sprintf(url, viper.GetString("baidu.key"), ip))
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}
	var temp = Response{}
	fmt.Println(string(data))
	// 部分时候，data返回类型不是这样的
	err = json.Unmarshal(data, &temp)
	if err != nil {
		fmt.Println(err)
		return Response{}, nil
	}
	return temp, nil
}
