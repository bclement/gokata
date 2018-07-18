package munge

import "testing"
import "regexp"

func TestWeather(t *testing.T) {
	pattern := `^\s+([0-9]+)[^0-9]+([0-9]+)[^0-9]+([0-9]+)`
	run(t, "weather.dat", pattern, "14")
}

func TestFootball(t *testing.T) {
	pattern := `^\s+\d+\.\s([a-zA-Z_]+).*\s(\d+)\s+-\s+(\d+)`
	run(t, "football.dat", pattern, "Aston_Villa")
}

func run(t *testing.T, path string, pattern string, expected string) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		t.Errorf("pattern %s", err)
	}
	res, err := SmallDiff(path, re)
	if err != nil {
		t.Errorf("problem %s", err)
	}
	if res != expected {
		t.Errorf("expected '%s' got '%s'", expected, res)
	}
}
