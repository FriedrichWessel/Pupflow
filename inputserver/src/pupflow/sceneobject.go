package main

import (
	"encoding/json"
	"math"
)

const (
	// MAP_MIN = 0
	// MAP_MAX = 65535
)

type SceneObject struct {
	Name string
	ValueMap [2]float64
	Axis string
	Midpoint float64
	Rotation bool
}

func (s *SceneObject) MarshalWithValue(v int16) []byte {
	m := make(map[string]interface{})
	m["Name"] = s.Name
	m["Value"] = s.RemapValue(v)
	m["Midpoint"] = s.Midpoint
	m["Rotation"] = s.Rotation
	m["Axis"] = s.Axis
	o, _ := json.Marshal(m)
	return o
}

func max(v1, v2 float64) float64 {
	if v1 > v2 {
		return v1
	}
	return v2
}

func min(v1, v2 float64) float64 {
	if v1 < v2 {
		return v1
	}
	return v2
}

// Some Invariants and/or assumptions:
// s.ValueMap[0] = s.Midpoint - dev1
// s.ValueMap[1] = s.Midpoint + dev2
func (s *SceneObject) RemapValue(v int16) float64 {
	amplitude := (float64(v) - V_MIN) / (V_MAX - V_MIN) - 0.5
	dev1 := s.Midpoint - s.ValueMap[0]
	dev2 := s.ValueMap[1] - s.Midpoint
	min_val := min(s.ValueMap[0], s.ValueMap[1])
	max_val := max(s.ValueMap[0], s.ValueMap[1])

	dev := float64(0)
	if math.Abs(dev1) > math.Abs(dev2) {
		dev = dev1
	} else {
		dev = dev2
	}

	new_min := s.Midpoint - dev
	new_max := s.Midpoint + dev

	val := s.Midpoint + amplitude * (new_max - new_min)
	if val > max_val {
		val = max_val
	}
	if val < min_val {
		val = min_val
	}
	return val
}



