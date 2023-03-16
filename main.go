package paint

// @version 0.0.4
// @license.name last updated at 2023/3/16 12:26:32

import (
	"fmt"
	"strconv"
)

type font int

const (
	Normal font = iota
	Bold
	Light
	Italic
	Underline
)

func R(args ...interface{}) string {
	return dye(0, "red", args...)
}

func G(args ...interface{}) string {
	return dye(0, "green", args...)
}

func Y(args ...interface{}) string {
	return dye(0, "yellow", args...)
}

func B(args ...interface{}) string {
	return dye(0, "blue", args...)
}

func M(args ...interface{}) string {
	return dye(0, "magenta", args...)
}

func C(args ...interface{}) string {
	return dye(0, "cyan", args...)
}

func W(args ...interface{}) string {
	return dye(0, "white", args...)
}

func HR(args ...interface{}) string {
	return dye(1, "red", args...)
}

func HG(args ...interface{}) string {
	return dye(1, "green", args...)
}

func HY(args ...interface{}) string {
	return dye(1, "yellow", args...)
}

func HB(args ...interface{}) string {
	return dye(1, "blue", args...)
}

func HM(args ...interface{}) string {
	return dye(1, "magenta", args...)
}

func HC(args ...interface{}) string {
	return dye(1, "cyan", args...)
}

func HW(args ...interface{}) string {
	return dye(1, "white", args...)
}

func LR(args ...interface{}) string {
	return dye(2, "red", args...)
}

func LG(args ...interface{}) string {
	return dye(2, "green", args...)
}

func LY(args ...interface{}) string {
	return dye(2, "yellow", args...)
}

func LB(args ...interface{}) string {
	return dye(2, "blue", args...)
}

func LM(args ...interface{}) string {
	return dye(2, "magenta", args...)
}

func LC(args ...interface{}) string {
	return dye(2, "cyan", args...)
}

func LW(args ...interface{}) string {
	return dye(2, "white", args...)
}

// Italic Red
func IR(args ...interface{}) string {
	return dye(3, "red", args...)
}

// Italic Green
func IG(args ...interface{}) string {
	return dye(3, "green", args...)
}

// Italic Yellow
func IY(args ...interface{}) string {
	return dye(3, "yellow", args...)
}

// Italic Blue
func IB(args ...interface{}) string {
	return dye(3, "blue", args...)
}

// Italic Magenta
func IM(args ...interface{}) string {
	return dye(3, "magenta", args...)
}

// Italic Cyan
func IC(args ...interface{}) string {
	return dye(3, "cyan", args...)
}

// Italic White
func IW(args ...interface{}) string {
	return dye(3, "white", args...)
}

// Underline Red
func UR(args ...interface{}) string {
	return dye(4, "red", args...)
}

// Underline Green
func UG(args ...interface{}) string {
	return dye(4, "green", args...)
}

// Underline Yellow
func UY(args ...interface{}) string {
	return dye(4, "yellow", args...)
}

// Underline Blue
func UB(args ...interface{}) string {
	return dye(4, "blue", args...)
}

// Underline Magenta
func UM(args ...interface{}) string {
	return dye(4, "magenta", args...)
}

// Underline Cyan
func UC(args ...interface{}) string {
	return dye(4, "cyan", args...)
}

// Underline White
func UW(args ...interface{}) string {
	return dye(4, "white", args...)
}

func dye(f font, color string, args ...interface{}) string {
	n := "33"
	switch color {
	case "black":
		n = "30"
	case "red":
		n = "31"
	case "green":
		n = "32"
	case "yellow":
		n = "33"
	case "blue":
		n = "34"
	case "magenta":
		n = "35"
	case "cyan":
		n = "36"
	case "white":
		n = "37"
	}
	return fmt.Sprintf("\033[%s;%sm%s\033[0m", strconv.Itoa(int(f)), n, Concat(args...))
}

func Concat(args ...interface{}) string {
	return join(" ", args...)
}

func join(separator string, args ...interface{}) (str string) {
	for i, param := range args {
		if i == len(args)-1 {
			separator = ""
		}
		str += fmt.Sprint(param) + separator
	}
	return
}
