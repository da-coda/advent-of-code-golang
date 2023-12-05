package main

import (
	"testing"
)

func BenchmarkRun(b *testing.B) {
	for n := 0; n < b.N; n++ {
		run("calibration_document.txt")
	}
}

func BenchmarkRunAsWorkerPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runAsWorkerPool("calibration_document.txt")
	}
}
