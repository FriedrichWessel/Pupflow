package main

import (
	"math"
	"testing"
)

const (
	INPUT_MIN = V_MIN
	INPUT_MID = 0
	INPUT_MAX = 32767
	ERROR = 0.1
)

func TestPositiveSymmetricRemap(t *testing.T) {
	var sco SceneObject

	sco = SceneObject {
		ValueMap: [2]float64{-2, 2},
		Midpoint: 0,
	}
	if v := sco.RemapValue(INPUT_MIN); math.Abs(v - (-2)) > ERROR {
		t.Fatalf("f(MIN) = %f, expected -2", v)
	}
	if v := sco.RemapValue(INPUT_MAX); math.Abs(v - (2)) > ERROR {
		t.Fatalf("f(MAX) = %f, expected 2", v)
	}
	if v := sco.RemapValue(INPUT_MID); math.Abs(v - (0)) > ERROR {
		t.Fatalf("f(MID) = %f, expected 0", v)
	}
}

func TestNegativeSymmetricRemap(t *testing.T) {
	var sco SceneObject

	sco = SceneObject {
		ValueMap: [2]float64{2, -2},
		Midpoint: 0,
	}
	if v := sco.RemapValue(INPUT_MIN); math.Abs(v - (2)) > ERROR {
		t.Fatalf("f(MIN) = %f, expected 2", v)
	}
	if v := sco.RemapValue(INPUT_MAX); math.Abs(v - (-2)) > ERROR {
		t.Fatalf("f(MAX) = %f, expected -2", v)
	}
	if v := sco.RemapValue(INPUT_MID); math.Abs(v - (0)) > ERROR {
		t.Fatalf("f(MID) = %f, expected 0", v)
	}
}

func TestPositiveAsymmetricRemap(t *testing.T) {
	var sco SceneObject

	sco = SceneObject {
		ValueMap: [2]float64{-2, 2},
		Midpoint: 1,
	}
	if v := sco.RemapValue(INPUT_MIN); math.Abs(v - (-2)) > ERROR {
		t.Fatalf("f(MIN) = %f, expected -2", v)
	}
	if v := sco.RemapValue(INPUT_MAX); math.Abs(v - (2)) > ERROR {
		t.Fatalf("f(MAX) = %f, expected 2", v)
	}
	if v := sco.RemapValue(INPUT_MID); math.Abs(v - (1)) > ERROR {
		t.Fatalf("f(MID) = %f, expected 1", v)
	}
}

func TestNegativeAsymmetricRemap(t *testing.T) {
	var sco SceneObject

	sco = SceneObject {
		ValueMap: [2]float64{2, -2},
		Midpoint: 1,
	}
	if v := sco.RemapValue(INPUT_MIN); math.Abs(v - (2)) > ERROR {
		t.Fatalf("f(MIN) = %f, expected 2", v)
	}
	if v := sco.RemapValue(INPUT_MAX); math.Abs(v - (-2)) > ERROR {
		t.Fatalf("f(MAX) = %f, expected -2", v)
	}
	if v := sco.RemapValue(INPUT_MID); math.Abs(v - (1)) > ERROR {
		t.Fatalf("f(MID) = %f, expected 1", v)
	}
}
