package munge

import "bufio"
import "os"
import "regexp"
import "strconv"
import "math"
import "io"

/*
Opens file at path and matches lines with re.
Regexp should consist of three groups, a label and two values.
Function returns label that has smallest diff between values.
*/
func SmallDiff(path string, re *regexp.Regexp) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var rval string
	min := math.Inf(1)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				return "", err
			}
			break
		}
		matched, label, diff, err := match(line, re)
		if err != nil {
			return "", err
		}
		if !matched {
			continue
		}
		if diff < min {
			min = diff
			rval = label
		}
	}
	return rval, nil
}

/*
If line matches re, bool return is true, string is label and float is diff
*/
func match(line string, re *regexp.Regexp) (matched bool, label string,
	diff float64, err error) {
	match := re.FindStringSubmatch(line)
	if matched = len(match) > 0; !matched {
		return
	}
	label = match[1]
	floats := [2]float64{0.0, 0.0}
	for i := range floats {
		floats[i], err = strconv.ParseFloat(match[i+2], 64)
		if err != nil {
			return
		}
	}
	diff = math.Abs(floats[0] - floats[1])
	return
}
