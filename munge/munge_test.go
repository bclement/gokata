package munge

import "bufio"
import "os"
import "testing"
import "regexp"

func TestWeather(t *testing.T){
    re,_ := regexp.Compile(`^\s+([0-9]+)[^0-9]+([0-9]+)[^0-9]+([0-9]+)`)
    file, _ := os.Open("weather.dat")
    reader := bufio.NewReader(file)
    for {
        line, err := reader.ReadString('\n')
        if err != nil{
            t.Errorf("%s", err)
            break
        }
        match := re.FindStringSubmatch(line)
        for _, val := range match{
            t.Logf("'%s'", val)
        }
    }
}
