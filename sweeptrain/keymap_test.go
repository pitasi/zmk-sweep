package main

import (
	"strings"
	"testing"
)

func TestDrillSet(t *testing.T) {
	set := buildDrillSet()
	if len(set) < 60 {
		t.Fatalf("drill set too small: %d", len(set))
	}
	seen := map[rune]bool{}
	for _, e := range set {
		if seen[e.r] {
			t.Errorf("duplicate drill char %c", e.r)
		}
		seen[e.r] = true
		if e.pos < 0 || e.pos >= 30 {
			t.Errorf("bad position %d for %c", e.pos, e.r)
		}
		if !strings.Contains(renderLayer(e.layer, e.pos), string(e.r)) {
			t.Errorf("grid for %c (%s) missing legend", e.r, e.layer)
		}
	}
	// ' must resolve to base (dedupe priority), not sym
	for _, e := range set {
		if e.r == '\'' && e.layer != "base" {
			t.Errorf("' should drill on base, got %s", e.layer)
		}
	}
}

func TestModDrillSet(t *testing.T) {
	set := buildModDrillSet()
	if len(set) == 0 {
		t.Fatal("empty mod drill set")
	}
	for _, e := range set {
		if e.code != e.letter-'a'+1 {
			t.Errorf("%c: control code %d, want %d", e.letter, e.code, e.letter-'a'+1)
		}
		// Ctrl must be held by the hand opposite the letter.
		if e.holdKey != 'd' && e.holdKey != 'k' {
			t.Errorf("%c: hold key %c is not a Ctrl home key", e.letter, e.holdKey)
		}
		if (basePos(e.letter)%10 < 5) == (basePos(e.holdKey)%10 < 5) {
			t.Errorf("%c: hold key %c is on the same hand", e.letter, e.holdKey)
		}
		// Codes must not collide with Tab/Enter/Esc/Backspace/LF.
		switch e.code {
		case 8, 9, 10, 13, 27:
			t.Errorf("%c: ambiguous control code %d", e.letter, e.code)
		}
	}
}
