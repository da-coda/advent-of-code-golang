package main

import (
	"testing"
)

func BenchmarkRun(b *testing.B) {
	for n := 0; n < b.N; n++ {
		run("calibration_document.txt")
	}
}

func TestCalcCalibrationValue(t *testing.T) {
	lines := map[string]int{
		"kone1ptnkjhks65sixrsseight":                   18,
		"9hfjjmgrzntssjpxcvbzpvmqzgsd54twonine":        99,
		"7kndzrhvcnstgfxjlff9twoninervrknsffmfzmdhtth": 79,
	}
	for line, expected := range lines {
		actual := calcCalibrationValue(line)
		if expected != actual {
			t.Errorf("Expected %d, got %d for line %s", expected, actual, line)
		}
	}
}
