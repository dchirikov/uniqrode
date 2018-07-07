/* Copyright (C) 2018 Dmitry Chirikov
 *
 * This software may be modified and distributed under the terms
 * of the MIT license.  See the LICENSE file for details.
 */

package main

import (
    "os"
    "fmt"
    "bufio"
    "io"
    "flag"
    ascii_mapper "github.com/dchirikov/uniqrode/ascii_mapper"
    qrcode "github.com/skip2/go-qrcode"
)

var recovery_level_default = 2
var mode_default = 2
var inverse_default = false

func showHelp() {
    fmt.Println("Usage: uniqrode [-level=1..4] [-mode=1..3] [-inverse]")
}

// Read input from pipe stdin
func getInput() string {
    fi, err := os.Stdin.Stat()
    if err != nil {
        panic(err)
    }

    if fi.Mode() & os.ModeNamedPipe == 0 {
        fmt.Println("Nothing to draw")
        os.Exit(0)
    }

    reader := bufio.NewReader(os.Stdin)
    var input []rune
    for {
		chunk, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		input = append(input, chunk)
	}

    return string(input)
}

// Map to get recovery level for generated QR code from passes
// CLI argument
var recoveryLevels = map[int]qrcode.RecoveryLevel {
    1: qrcode.Low,
    2: qrcode.Medium,
    3: qrcode.High,
    4: qrcode.Highest,
}

func main() {

    argNeedHelpPtr := flag.Bool("h", false, "help")
    argRecoveryLevelPtr := flag.Int(
        "level", recovery_level_default, "Recovery level")
    argDrawModePtr := flag.Int(
        "mode", mode_default, "Draw mode")
    argInversePtr := flag.Bool(
        "inverse", inverse_default, "Draw inverse")

    // Parse arguments and check if they are valid
    flag.Parse()

    if *argNeedHelpPtr {
        showHelp()
        os.Exit(0)
    }

    var q *qrcode.QRCode

    recoveryLevel, present := recoveryLevels[*argRecoveryLevelPtr];

    if !present {
        showHelp()
        os.Exit(1)
    }

    // Create QR code
    q, err := qrcode.New(getInput(), recoveryLevel)

    if err != nil {
        panic (err)
    }

    // Draw ASCII picture
    bits := q.Bitmap()

    u, err := ascii_mapper.New(*argDrawModePtr, !(*argInversePtr), &bits)

    if err != nil {
        panic (err)
    }

    fmt.Printf("%s", u.Draw())
}
