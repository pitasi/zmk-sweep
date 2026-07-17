package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"

	"golang.org/x/term"
)

type stat struct {
	attempts, errors int
	total            time.Duration // time to first correct answer
	correct          int
}

func main() {
	layerFlag := flag.String("layer", "", "drill only this layer (base|sym|sym2|accent)")
	flag.Parse()

	set := buildDrillSet()
	if *layerFlag != "" {
		if _, ok := layers[*layerFlag]; !ok {
			fmt.Fprintf(os.Stderr, "unknown layer %q\n", *layerFlag)
			os.Exit(1)
		}
		var filtered []entry
		for _, e := range set {
			if e.layer == *layerFlag {
				filtered = append(filtered, e)
			}
		}
		set = filtered
	}

	fd := int(os.Stdin.Fd())
	old, err := term.MakeRaw(fd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	stats := map[rune]*stat{}
	defer func() {
		term.Restore(fd, old)
		summary(stats)
	}()

	fmt.Print("sweeptrain — type the shown character on your sweep. Esc or Ctrl-C to quit.\r\n\r\n")
	in := bufio.NewReader(os.Stdin)
	for {
		e := pick(set, stats)
		st := stats[e.r]
		if st == nil {
			st = &stat{}
			stats[e.r] = st
		}
		fmt.Printf("type: %c  ", e.r)
		start := time.Now()
		for {
			r, _, err := in.ReadRune()
			if err != nil || r == 27 || r == 3 { // Esc / Ctrl-C
				fmt.Print("\r\n")
				return
			}
			st.attempts++
			if r == e.r {
				st.correct++
				st.total += time.Since(start)
				fmt.Print("✓\r\n")
				break
			}
			st.errors++
			fmt.Printf("✗ (%c)  hint — layer %s: %s\r\n", r, e.layer, layerHint[e.layer])
			fmt.Print(renderLayer(e.layer, e.pos))
			fmt.Printf("type: %c  ", e.r)
		}
	}
}

// pick weights toward keys you've missed: weight = errors + 1.
func pick(set []entry, stats map[rune]*stat) entry {
	total := 0
	for _, e := range set {
		if st := stats[e.r]; st != nil {
			total += st.errors
		}
		total++
	}
	n := rand.Intn(total)
	for _, e := range set {
		w := 1
		if st := stats[e.r]; st != nil {
			w += st.errors
		}
		if n < w {
			return e
		}
		n -= w
	}
	return set[len(set)-1]
}

func summary(stats map[rune]*stat) {
	attempts, errors := 0, 0
	type row struct {
		r  rune
		st *stat
	}
	var rows []row
	for r, st := range stats {
		attempts += st.attempts
		errors += st.errors
		if st.errors > 0 || st.correct > 0 {
			rows = append(rows, row{r, st})
		}
	}
	if attempts == 0 {
		return
	}
	fmt.Printf("\n%d attempts, %.0f%% accuracy\n", attempts, 100*float64(attempts-errors)/float64(attempts))
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].st.errors != rows[j].st.errors {
			return rows[i].st.errors > rows[j].st.errors
		}
		return avg(rows[i].st) > avg(rows[j].st)
	})
	fmt.Println("worst keys:")
	for i, rw := range rows {
		if i == 5 {
			break
		}
		fmt.Printf("  %c  %d misses, avg %.1fs\n", rw.r, rw.st.errors, avg(rw.st).Seconds())
	}
}

func avg(st *stat) time.Duration {
	if st.correct == 0 {
		return 0
	}
	return st.total / time.Duration(st.correct)
}
