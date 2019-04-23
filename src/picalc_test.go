package main

import "testing"

func TestCalculatePi(t *testing.T) {
    var pi float64
    pi = calculatePi(20000000)
    if pi < 3.1415926 || pi >= 3.1415927 {
       t.Errorf("Value %.10f is incorrect", pi)
    }
}