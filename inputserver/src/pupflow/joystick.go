package main

import (
	"github.com/banthar/Go-SDL/sdl"
)

const (
	V_MAX = 32767
	V_MIN = -32768
)

var (
	Smoothing = 0
)

func openAllJoysticks() ([]*sdl.Joystick) {
	r := make([]*sdl.Joystick, 0, 10)
	for i := 0; i < sdl.NumJoysticks(); i++ {
		r = append(r, sdl.JoystickOpen(i))
	}
	return r
}

func getJoystickState(js []*sdl.Joystick) []int16 {
	r := make([]int16, 0)
	for _, j := range js {
		for i := 0; i < j.NumAxes(); i++ {
			r = append(r, int16(j.GetAxis(i)))
		}
		for i := 0; i < j.NumButtons(); i++ {
			v := int16(V_MIN);
			if j.GetButton(i) {
				v = V_MAX
			}
			r = append(r, v)
		}
		for i := 0; i < j.NumHats(); i++ {
			r = append(r, extractHorizontalAxis(j.GetHat(i)))
			r = append(r, extractVerticalAxis(j.GetHat(i)))
		}
	}
	for i := range r {
		r[i] = smooth(r[i])
	}
	return r
}

func smooth(v int16) int16 {
		return (v >> uint(Smoothing)) << uint(Smoothing)
}

const (
	TOP = 1 << iota
	RIGHT
	BOTTOM
	LEFT
)

func extractVerticalAxis(hat int8) int16 {
	if hat & BOTTOM != 0 {
		return V_MAX
	}
	if hat & TOP != 0 {
		return V_MIN
	}
	return 0
}

func extractHorizontalAxis(hat int8) int16 {
	if hat & RIGHT != 0 {
		return V_MAX
	}
	if hat & LEFT != 0 {
		return V_MIN
	}
	return 0
}
