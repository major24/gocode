package math

import (
	"testing"
	)

func TestAveragePass1(t *testing.T){
	var v float64
	v = Average([]float64{1,2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

func TestAveragePass2(t *testing.T){
	var v float64
	v = Average([]float64{1,2,4,3,0})
	if v != 2.0 {
		t.Error("Expected 2.0, got ", v)
	}
}