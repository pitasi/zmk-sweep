package main

import "strings"

// Layer legends, 34 positions: 3 rows of 10 + 4 thumbs (30,31 left / 32,33 right).
// Single-rune legends on positions 0..29 are drillable; anything else is display-only.
var layers = map[string][34]string{
	"base": {
		"q", "w", "e", "r", "t", "y", "u", "i", "o", "p",
		"a", "s", "d", "f", "g", "h", "j", "k", "l", "'",
		"z", "x", "c", "v", "b", "n", "m", ",", ".", ";",
		"EXT", "⇧␣", "␣", "SYM",
	},
	"sym": {
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
		"<", "$", "(", "{", "[", "=", "?", "'", "-", "*",
		">", "%", ")", "}", "]", "@", "!", "\"", "+", "/",
		"FNC", "SY2", "", "",
	},
	"sym2": {
		"^", "_", "£", "€", "#", "", "", "", "", "",
		"`", "~", "\\", "|", "&", "", "", "", "", "",
		"", "", "§", "ç", "°", "", "", "", "", "",
		"", "", "", "",
	},
	"accent": {
		"", "", "", "", "", "", "", "", "", "",
		"à", "ì", "ò", "è", "é", "", "", "", "", "",
		"", "", "", "ù", "", "", "", "", "", "",
		"", "", "", "",
	},
}

// Dedupe priority: a char is drilled on the first layer that has it.
var layerOrder = []string{"base", "sym", "sym2", "accent"}

var layerHint = map[string]string{
	"base":   "no layer, just type it",
	"sym":    "tap right-outer thumb (sticky), then the key",
	"sym2":   "hold left-inner + right-outer thumbs, then the key",
	"accent": "hold both right thumbs, then the key",
}

type entry struct {
	r     rune
	layer string
	pos   int
}

func buildDrillSet() []entry {
	var set []entry
	seen := map[rune]bool{}
	for _, layer := range layerOrder {
		for pos := 0; pos < 30; pos++ {
			legend := layers[layer][pos]
			rs := []rune(legend)
			if len(rs) != 1 || seen[rs[0]] {
				continue
			}
			seen[rs[0]] = true
			set = append(set, entry{rs[0], layer, pos})
		}
	}
	return set
}

// renderLayer draws the layer grid with position hi highlighted (inverse video).
func renderLayer(layer string, hi int) string {
	var b strings.Builder
	cell := func(pos int) {
		legend := layers[layer][pos]
		pad := 3 - len([]rune(legend))
		if pad < 0 {
			pad = 0
		}
		s := " " + legend + strings.Repeat(" ", pad)
		if pos == hi {
			s = "\x1b[7m" + s + "\x1b[0m"
		}
		b.WriteString(s)
	}
	for row := 0; row < 3; row++ {
		for col := 0; col < 5; col++ {
			cell(row*10 + col)
		}
		b.WriteString("  ")
		for col := 5; col < 10; col++ {
			cell(row*10 + col)
		}
		b.WriteString("\r\n")
	}
	b.WriteString(strings.Repeat(" ", 12))
	cell(30)
	cell(31)
	b.WriteString("  ")
	cell(32)
	cell(33)
	b.WriteString("\r\n")
	return b.String()
}
