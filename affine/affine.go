package affine

import (
	"math"
	"strings"
)

func Encrypt(text string, a int, b int, useSpaces bool) string {
	//abc := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", " "}

	mod := 27

	abcLetters := map[int]string{0: "a",
		1:  "b",
		2:  "c",
		3:  "d",
		4:  "e",
		5:  "f",
		6:  "g",
		7:  "h",
		8:  "i",
		9:  "j",
		10: "k",
		11: "l",
		12: "m",
		13: "n",
		14: "o",
		15: "p",
		16: "q",
		17: "r",
		18: "s",
		19: "t",
		20: "u",
		21: "v",
		22: "w",
		23: "x",
		24: "y",
		25: "z",
		26: " "}

	abcPositions := map[string]int{"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
		"f": 5,
		"g": 6,
		"h": 7,
		"i": 8,
		"j": 9,
		"k": 10,
		"l": 11,
		"m": 12,
		"n": 13,
		"o": 14,
		"p": 15,
		"q": 16,
		"r": 17,
		"s": 18,
		"t": 19,
		"u": 20,
		"v": 21,
		"w": 22,
		"x": 23,
		"y": 24,
		"z": 25,
		" ": 26,
	}

	text = strings.ToLower(text)
	if !useSpaces {
		mod = 26
		text = strings.ReplaceAll(text, " ", "")
	}

	encrypted := ""

	for i := 0; i < len(text); i++ {
		encrypted += abcLetters[int(math.Floor(float64(abcPositions[string(text[i])])*float64(a)+float64(b)))%mod]
	}

	return encrypted
}
