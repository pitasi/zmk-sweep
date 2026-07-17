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
