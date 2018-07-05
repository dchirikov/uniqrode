/* Copyright (C) 2018 Dmitry Chirikov
 *
 * This software may be modified and distributed under the terms
 * of the MIT license.  See the LICENSE file for details.
 */

// Mapper class. Convert bit (bool) representation to ASCII-art pictures

package uniqrode

import (
    "fmt"
    "bytes"
    "errors"
)

// UniQRode object
type UniQRode struct {

    // should we draw QR-code in negative
    Invert bool

    mode asciiMapping
    data *[][]bool
}

// Stores map table and "resolution" for this table
type asciiMapping struct {

    // "resolution" of the resulting picture
    // represents how many bits HxW be used to get single char
    // depends on the mode
    stepx int
    stepy int

    // exact mapping uses for drawing
    mapTable *map[int]string
}

// Several modes to draw QR code
var modes = map[int]asciiMapping {
    1: {stepx: 1, stepy: 1, mapTable: &u_symbol_mappings_1x1},
    2: {stepx: 1, stepy: 2, mapTable: &u_symbol_mappings_1x2},
    3: {stepx: 2, stepy: 2, mapTable: &u_symbol_mappings_2x2},
}

// Map for chars. Main part will construct integer from bit mask, i.e.
// {{0, 0}, {1, 1}} will become 3, and {{0, 1}, {1, 1}} - 7
var u_symbol_mappings_1x1 = map[int]string {
    0: "  ",             // Empty block
    1: "\u2588\u2588",   // Full block
}

var u_symbol_mappings_1x2 = map[int]string {
    0: " ",        // | | Empty block
    1: "\u2584",   // |▄| Lower half block
    2: "\u2580",   // |▀| Upper half block
    3: "\u2588",   // |█| Full block
}

var u_symbol_mappings_2x2 = map[int]string {
    0:  " ",        // | | Empty block
    1:  "\u2597",   // |▗| Quadrant lower right
    2:  "\u2596",   // |▖| Quadrant lower left
    3:  "\u2584",   // |▄| Lower half block
    4:  "\u259D",   // |▝| Quadrant upper right
    5:  "\u2590",   // |▐| Right half block
    6:  "\u259E",   // |▞| Quadrant upper right and lower left
    7:  "\u259F",   // |▟| Quadrant upper right and lower left and lower right
    8:  "\u2598",   // |▘| Quadrant upper left
    9:  "\u259A",   // |▚| Quadrant upper left and lower right
    10: "\u258C",   // |▌| Left half block
    11: "\u2599",   // |▙| Quadrant upper left and lower left and lower right
    12: "\u2580",   // |▀| Upper half block
    13: "\u259C",   // |▜| Quadrant upper left and upper right and lower right
    14: "\u259B",   // |▛| Quadrant upper left and upper right and lower left
    15: "\u2588",   // |█| Full block
}

// Tiny helper function. Converts bool to int.
func btou(b bool) int {
    if b {
        return 1
    }
    return 0
}

// Constructor for UniQRode object.
func New(mode int, invert_colors bool, data *[][]bool) (*UniQRode, error) {
    m, present := modes[mode];

    // worth to check first if we have the mode we are going to draw
    if !present {
        msg := fmt.Sprintf("No such mode - %d", mode)
        return nil, errors.New(msg)
    }

    u := &UniQRode {
        Invert: invert_colors,
        mode: m,
        data: data,
    }

    return u, nil
}

// The only public method available for UniQRode
// scans through data cutting patches size stepx * stepy.
// Then converts every patch to a symbol based on
// which mapping table was chosen in constructor.
func (u *UniQRode) Draw() string {
    var buf bytes.Buffer
    lenY := len(*u.data)
    for y := 0; y <= lenY - u.mode.stepy; y += u.mode.stepy {
        lenX := len((*u.data)[y])
        for x := 0; x <= lenX - u.mode.stepx; x += u.mode.stepx {
            patch := u.getPatch(x, y)
            buf.WriteString(u.convertChar(&patch))
        }
        buf.WriteString("\n")
    }
    return buf.String()
}

// Returns patch sized stepx * stepy on position startX, startY.
func (u *UniQRode) getPatch(startX int, startY int) [][]bool {
    res := make([][]bool, u.mode.stepy)
    for y := 0; y < u.mode.stepy; y++ {
        res[y] = make([]bool, u.mode.stepx)
        for x := 0; x < u.mode.stepx; x++ {
            res[y][x] = (*u.data)[startY + y][startX + x]
        }
    }
    return res
}

// Convert patch to single int.
// Shifting bits, starting from bottom-right corner.
func (u *UniQRode) convertChar(pixels *[][]bool) string {
    symbol_key := 0;
    // flatten array and convert the result to int
    for _, row := range (*pixels) {
        for _, val := range row {
            // "bool != bool" is XOR
            symbol_key = (symbol_key << 1) + btou(u.Invert != val)
        }
    }
    return (*u.mode.mapTable)[symbol_key]
}
