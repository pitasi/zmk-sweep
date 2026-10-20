⭐ You might also like: the [Urchin](https://github.com/duckyb/urchin) keyboard ⭐

This layout is inspired by [Seniply by Stevep99](https://stevep99.github.io/seniply/)

[See it in action on YouTube!](https://youtu.be/IZ83uU0ltaE)
<div align="center">
  
  ![sweep-layout](https://user-images.githubusercontent.com/27895007/185715593-69f9f981-ae17-4788-b2a8-d1360c65622a.svg)

</div>

*This diagram was created using draw.io*  
*Click [HERE](https://viewer.diagrams.net/?tags=%7B%7D&highlight=0000ff&edit=_blank&layers=1&nav=1&title=sweep-layout.drawio#Uhttps%3A%2F%2Fdrive.google.com%2Fuc%3Fid%3D1eJrqAJkoFiEjM5GZY4EuauVtp0VhaR0j%26export%3Ddownload) to view a copy that you can edit*

## Layers

### BASE (0)
```
┌─────┬─────┬─────┬─────┬─────┐   ┌─────┬─────┬─────┬─────┬─────┐
│  q  │  w  │  e  │  r  │  t  │   │  y  │  u  │  i  │  o  │  p  │
├─────┼─────┼─────┼─────┼─────┤   ├─────┼─────┼─────┼─────┼─────┤
│  a  │  s  │  d  │  f  │  g  │   │  h  │  j  │  k  │  l  │  '  │
│  ⌘  │  ⌥  │  ⌃  │  ⇧  │     │   │     │  ⇧  │  ⌃  │  ⌥  │  ⌘  │
├─────┼─────┼─────┼─────┼─────┤   ├─────┼─────┼─────┼─────┼─────┤
│  z  │  x  │  c  │  v  │  b  │   │  n  │  m  │  ,  │  .  │  ;  │
└─────┴─────┴──┬──┴──┬──┴──┬──┘   └──┬──┴──┬──┴──┬──┴─────┴─────┘
               │ Esc │  ␣  │         │  ⌫  │     │
               │ EXT │     │         │     │ SYM │
               └─────┴─────┘         └─────┴─────┘
```
Home-row mods — **hold** a home key for the modifier shown below it (tap = letter):
`a`=⌘ `s`=⌥ `d`=⌃ `f`=⇧ · `j`=⇧ `k`=⌃ `l`=⌥ `'`=⌘. Positional + prior-idle tuned,
so same-hand rolls stay plain letters. The two ⇧ keys skip the prior-idle guard
(capitals always follow a keypress), so they hold even mid-burst.

The `EXT`/`Esc` thumb layer-tap uses the `balanced` flavor: the layer wins when
an interrupting key is pressed and released while the thumb is held.
The right-outer thumb is `SYM` only: tapping it alone types nothing, so symbol
access cannot fall back to Enter. For Enter, hold `EXT` and press the physical `M` key.

Shortcuts = hold the **opposite-hand** mod + the real letter, on this layer:
`Ctrl+c`=K+c · `Ctrl+d`=K+d · `Ctrl+f`=K+f · `Ctrl+u`=D+u · `Ctrl+n`=D+n · `Ctrl+p`=D+p ·
`⌘c`=`'`+c · `⌘v`=`'`+v · `⌘z`=`'`+z.

Thumbs: `Esc` tap / `EXT` hold · `␣` space · `⌫` backspace · `SYM` hold only

### SYM (1) — hold right-outer thumb (no tap output)
```
┌─────┬─────┬─────┬─────┬─────┐   ┌─────┬─────┬─────┬─────┬─────┐
│  1  │  2  │  3  │  4  │  5  │   │  6  │  7  │  8  │  9  │  0  │
├─────┼─────┼─────┼─────┼─────┤   ├─────┼─────┼─────┼─────┼─────┤
│  <  │  $  │  (  │  {  │  [  │   │  =  │  ?  │  '  │  -  │  *  │
├─────┼─────┼─────┼─────┼─────┤   ├─────┼─────┼─────┼─────┼─────┤
│  >  │  %  │  )  │  }  │  ]  │   │  @  │  !  │  "  │  +  │  /  │
└─────┴─────┴──┬──┴──┬──┴──┬──┘   └──┬──┴──┬──┴──┬──┴─────┴─────┘
               │ FNC │SYM2 │         │     │     │
               └─────┴─────┘         └─────┴─────┘
```

### EXT (2) — hold left-outer thumb (tap = Esc)
```
┌─────┬─────┬─────┬─────┬─────┐   ┌─────┬─────┬─────┬─────┬─────┐
│ ESC │ ⌘1  │ ⌘2  │ ⌘3  │ ⌘4  │   │HOME │PGDN │PGUP │ END │ ⌘5  │
├─────┼─────┼─────┼─────┼─────┤   ├─────┼─────┼─────┼─────┼─────┤
│  ⌘  │  ⌥  │  ⌃  │  ⇧  │AltGr│   │  ←  │  ↓  │  ↑  │  →  │ DEL │
├─────┼─────┼─────┼─────┼─────┤   ├─────┼─────┼─────┼─────┼─────┤
│     │     │     │     │     │   │C-BSP│ Ent │CAPS │     │     │
└─────┴─────┴──┬──┴──┬──┴──┬──┘   └──┬──┴──┬──┴──┬──┴─────┴─────┘
               │     │ TAB │         │ ⇧⇥  │ FNC │
               └─────┴─────┘         └─────┴─────┘
```
`Enter` is on the physical `M` key (right index finger), replacing the duplicate
EXT-layer Backspace. The dedicated base-layer Backspace thumb is unchanged.

`TAB` is the left-inner thumb, `Shift+Tab` (`⇧⇥`) the right-inner thumb — so both
directions are one press each while holding EXT.

Left-hand mods here are plain holds (no home-row needed inside a layer), so you can
chord `⇧`/`⌘`/`⌥` + arrows for selection and word/line nav. `⌘1`–`⌘5` = macOS spaces.

### FNC (3) — SYM + left-outer thumb, or EXT + right-outer thumb
```
┌─────┬─────┬─────┬─────┬─────┐   ┌─────┬─────┬─────┬─────┬─────┐
│ F1  │ F2  │ F3  │ F4  │ F5  │   │ F6  │ F7  │ F8  │ F9  │ F10 │
├─────┼─────┼─────┼─────┼─────┤   ├─────┼─────┼─────┼─────┼─────┤
│  ⌘  │  ⌥  │  ⌃  │  ⇧  │AltGr│   │ F11 │ F12 │PrtSc│     │     │
├─────┼─────┼─────┼─────┼─────┤   ├─────┼─────┼─────┼─────┼─────┤
│     │     │     │     │     │   │     │     │     │     │     │
└─────┴─────┴──┬──┴──┬──┴──┬──┘   └──┬──┴──┬──┴──┬──┴─────┴─────┘
               │     │     │         │     │     │
               └─────┴─────┘         └─────┴─────┘
```

### SYM2 (4) — hold left-inner + right-outer thumbs (combo)
```
┌─────┬─────┬─────┬─────┬─────┐   ┌─────┬─────┬─────┬─────┬─────┐
│  ^  │  _  │  £  │  €  │  #  │   │     │     │     │     │     │
├─────┼─────┼─────┼─────┼─────┤   ├─────┼─────┼─────┼─────┼─────┤
│  `  │  ~  │  \  │  |  │  &  │   │     │     │     │     │     │
├─────┼─────┼─────┼─────┼─────┤   ├─────┼─────┼─────┼─────┼─────┤
│     │     │  §  │  ç  │  °  │   │     │     │     │     │     │
└─────┴─────┴──┬──┴──┬──┴──┬──┘   └──┬──┴──┬──┴──┬──┴─────┴─────┘
               │     │     │         │     │     │
               └─────┴─────┘         └─────┴─────┘
```

### ACCENT (5) — hold both right thumbs (combo)
```
┌─────┬─────┬─────┬─────┬─────┐   ┌─────┬─────┬─────┬─────┬─────┐
│ ˋdk │ ˊdk │     │  ç  │     │   │     │     │     │     │     │
├─────┼─────┼─────┼─────┼─────┤   ├─────┼─────┼─────┼─────┼─────┤
│  à  │  ì  │  ò  │  è  │  é  │   │     │     │     │     │     │
├─────┼─────┼─────┼─────┼─────┤   ├─────┼─────┼─────┼─────┼─────┤
│     │     │     │  ù  │     │   │     │     │     │     │     │
└─────┴─────┴──┬──┴──┬──┴──┬──┘   └──┬──┴──┬──┴──┬──┴─────┴─────┘
               │     │     │         │     │     │
               └─────┴─────┘         └─────┴─────┘
```
`dk` = AltGr dead key (US-intl layout)

### SETTINGS (6) — hold both left thumbs + right-inner thumb (combo)
```
┌─────┬─────┬─────┬─────┬─────┐   ┌─────┬─────┬─────┬─────┬─────┐
│BOOT │     │     │BTCLR│BT 0 │   │BT 3 │     │UNSTK│     │BOOT │
├─────┼─────┼─────┼─────┼─────┤   ├─────┼─────┼─────┼─────┼─────┤
│     │     │     │     │BT 1 │   │BT 4 │     │     │     │     │
├─────┼─────┼─────┼─────┼─────┤   ├─────┼─────┼─────┼─────┼─────┤
│     │     │     │     │BT 2 │   │BT 5 │     │     │     │     │
└─────┴─────┴──┬──┴──┬──┴──┬──┘   └──┬──┴──┬──┴──┬──┴─────┴─────┘
               │     │     │         │     │     │
               └─────┴─────┘         └─────┴─────┘
```

### Thumb combos
| Combo | Layer |
|---|---|
| both left thumbs + right-inner thumb | SETTINGS |
| left-inner + right-outer | SYM2 |
| both right thumbs | ACCENT |
