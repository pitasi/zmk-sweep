# AGENTS.md

Personal [ZMK](https://zmk.dev) firmware config for a **Ferris Sweep** (36-key
split, 34 used) running on two `nice_nano_v2` controllers over Bluetooth. The
layout is inspired by [Seniply](https://stevep99.github.io/seniply/) and borrows
home-row-mod ideas from Miryoku. It also ships `sweeptrain`, a small Go TUI for
drilling the layout into muscle memory.

## Layout

- **Repo structure**
  - `config/cradio.keymap` — the keymap (devicetree). This is the file you edit
    for almost every layout change.
  - `config/cradio.conf` — Kconfig tweaks (BT power, deep sleep, debounce).
  - `config/west.yml` — pins ZMK to a specific tag (currently `v0.3`).
  - `build.yaml` — declares the two shields (`cradio_left` / `cradio_right`).
  - `.github/workflows/build.yml` — CI builds firmware via ZMK's
    `build-user-config.yml`; download the `.uf2` artifacts from the run.
  - `README.md` — human-facing docs with ASCII diagrams of every layer. **Keep
    it in sync when you change the keymap.**
  - `sweeptrain/` — Go practice tool (see below).

## Keymap conventions (`config/cradio.keymap`)

- **Target OS is macOS (ABC/US layout).** Accents and rare symbols are produced
  via Option/AltGr dead keys (e.g. `RA(GRAVE)` then a vowel). Don't assume a
  Linux/Windows layout when adding characters.
- **Layers** (see the `#define`s at the top): `BASE 0`, `SYM 1`, `EXT 2`,
  `FNC 3`, `SYM2 4`, `ACCENT 5`, `SETTINGS 6`.
- **Home-row mods** use positional hold-tap (`hml` / `hmr` behaviors), GACS
  order: left `A/S/D/F` = Gui/Alt/Ctrl/Shift, right `J/K/L/'` mirror it. They
  only trigger as mods when the *next* key is on the opposite hand or a thumb
  (`hold-trigger-key-positions` + `require-prior-idle-ms`), so same-hand rolls
  stay plain letters. Ctrl/Cmd shortcuts have no dedicated keys — you hold the
  opposite-hand mod and press the letter.
- **Key positions** matter: `KEYS_L`, `KEYS_R`, `THUMBS` (0–33) drive the
  positional hold-taps and combos. Rows are 3×10 + 4 thumbs (30,31 left /
  32,33 right).
- **Layer access**: right-outer thumb holds `SYM` only (no tap output);
  left-outer thumb taps Esc / holds `EXT`. Enter is `EXT` + physical `M`.
  `SYM2`, `ACCENT`, `SETTINGS` use thumb **combos** (see the `combos` node).
  If you move thumb keys, update the combo `key-positions`.
- Each layer's `bindings` block is formatted as an aligned grid mirroring the
  physical layout — preserve that alignment when editing.

## Building & flashing

- Push to GitHub; the Actions workflow builds `cradio_left` and `cradio_right`
  `.uf2` files. Flash each half by putting the `nice_nano_v2` into bootloader
  (double-tap reset, or the `&bootloader` key on the SETTINGS layer).
- To pin a different ZMK version, edit `revision` in `config/west.yml` and the
  workflow `@v0.3` ref together.
- There is no local ZMK toolchain assumed here; rely on CI unless the user asks
  to set up a local `west` build.

## sweeptrain (Go practice tool)

- `cd sweeptrain && go build && ./sweeptrain` — a raw-terminal drill that shows
  a character and checks what you type. Weights toward keys you miss.
- Flags: `-layer base|sym|sym2|accent` to drill one layer; `-mods` to drill
  `Ctrl+<letter>` home-row-mod chords.
- The layer legends live in `sweeptrain/keymap.go` (the `layers` map). **These
  are a hand-maintained copy of the keymap** — if you change letters/symbols in
  `cradio.keymap`, mirror the change here so drills stay accurate.
- Go 1.22; only dep is `golang.org/x/term`. Tests: `go test ./...`.
- The built binary `sweeptrain/sweeptrain` is gitignored.

## When changing the layout

Treat these three as one unit — updating one usually means updating the others:
1. `config/cradio.keymap` (the actual change)
2. `README.md` ASCII diagrams + notes
3. `sweeptrain/keymap.go` legends (if letters/symbols moved)
