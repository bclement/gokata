package munge

import "bufio"
import "os"
import "regexp"
import "strconv"
import "math"
import "io"

func SmallDiff(path string, re *regexp.Regexp) (string, error){
    file, err := os.Open(path)
    if ( err != nil){
        return "", err
    }
    reader := bufio.NewReader(file)
    var rval string
    min := math.Inf(1)
    for {
        line, err := reader.ReadString('\n')
        if err != nil{
            if (err != io.EOF){
                return "", err
            }
            break
        }
        match := re.FindStringSubmatch(line)
        if len(match) < 4{
            continue
        }
        floats := [2]float64{0.0,0.0}
        for i := range floats{
            val, err := strconv.ParseFloat(match[i+2], 64)
            floats[i] = val
            if err != nil{
                return "", err
            }
        }
        diff := math.Abs(floats[0]-floats[1])
        if diff < min{
            min = diff
            rval = match[1]
        }
    }
    return rval, nil
}
