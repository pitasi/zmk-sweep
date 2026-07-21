package main

import "strings"

// Layer legends, 34 positions: 3 rows of 10 + 4 thumbs (30,31 left / 32,33 right).
// Single-rune legends on positions 0..29 are drillable; anything else is display-only.
var layers = map[string][34]string{
	"base": {
		"q", "w", "e", "r", "t", "y", "u", "i", "o", "p",
		"a", "s", "d", "f", "g", "h", "j", "k", "l", "'",
		"z", "x", "c", "v", "b", "n", "m", ",", ".", ";",
		"EXT", "␣", "⌫", "SYM",
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
	"sym":    "hold right-outer thumb, then the key",
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

// --- Modifier-chord drill (home-row mods) ---

// ctrlLetters are the Ctrl+<letter> chords worth drilling: common shell/vim/TUI
// bindings. All map to unambiguous control bytes (none collide with Tab/Enter/
// Esc/Backspace), so the result can be read straight off the terminal. Cmd
// chords are intentionally excluded — macOS intercepts them before stdin.
var ctrlLetters = []rune{'a', 'c', 'd', 'e', 'f', 'n', 'p', 'r', 'u', 'w'}

type modEntry struct {
	letter   rune   // letter to combine, e.g. 'u'
	code     rune   // control byte the chord produces, e.g. 21 (Ctrl+U)
	holdKey  rune   // opposite-hand home-row Ctrl to hold: 'd' or 'k'
	holdHand string // hand that holds Ctrl: "left" or "right"
}

func buildModDrillSet() []modEntry {
	var set []modEntry
	for _, r := range ctrlLetters {
		pos := basePos(r)
		if pos < 0 {
			continue
		}
		// Chord with the opposite hand's Ctrl (left Ctrl = d, right Ctrl = k).
		hold, hand := 'k', "right"
		if pos%10 >= 5 { // letter sits on the right hand
			hold, hand = 'd', "left"
		}
		set = append(set, modEntry{r, r - 'a' + 1, hold, hand})
	}
	return set
}

// basePos returns the base-layer position of a single-rune legend, or -1.
func basePos(r rune) int {
	for i := 0; i < 30; i++ {
		if layers["base"][i] == string(r) {
			return i
		}
	}
	return -1
}

// describeInput renders a received byte for feedback (control codes as Ctrl+X).
func describeInput(r rune) string {
	switch {
	case r >= 1 && r <= 26:
		return "Ctrl+" + string('a'+r-1)
	case r == 27:
		return "Esc"
	case r == ' ':
		return "Space"
	default:
		return string(r)
	}
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
