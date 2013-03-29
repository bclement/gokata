package munge

import "testing"
import "regexp"

func TestWeather(t *testing.T){
    re, _ := regexp.Compile(`^\s+([0-9]+)[^0-9]+([0-9]+)[^0-9]+([0-9]+)`)
    res, err := SmallDiff("weather.dat", re)
    if err != nil{
        t.Errorf("problem %s", err)
    }
    if res != "14"{
        t.Errorf("result %s", res)
    }
}
