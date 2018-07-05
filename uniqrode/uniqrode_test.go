package uniqrode

import (
    "fmt"
	"testing"
)

var bits_int = [][]int {
    {0, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 1},
    {1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0, 1},
}

func int2bool(bits_int [][]int) [][]bool {
    res := make([][]bool, len(bits_int))
    for y, row := range bits_int {
        res[y] = make([]bool, len(bits_int[y]))
        for x, _ := range row {
            if bits_int[y][x] == 0 {
                res[y][x] = false
            } else {
                res[y][x] = true
            }
        }
    }
    return res
}

func testEq2d(a, b [][]bool) bool {

    if a == nil && b == nil {
        return true;
    }

    if a == nil || b == nil {
        return false;
    }

    if len(a) != len(b) {
        return false
    }

    for y, row := range a {
        if len(b[y]) != len(a[y]) {
            return false
        }
        for x, val := range row {
            if b[y][x] != val {
                return false
            }
        }
    }

    return true
}

func Test_int2bool(t *testing.T) {

    test01_int := [][]int {
        {0, 1, 1, 0},
        {1, 1, 0, 0},
    }

    test01_bool := [][]bool {
        {false, true, true, false},
        {true, true, false, false},
    }

    if ! testEq2d(int2bool(test01_int), test01_bool) {
        t.Fatal("Test_int2bool failed")
    }
}

func TestMode01(t *testing.T) {
    data := int2bool(bits_int)
    u, err := New(1, false, &data)
    if err != nil {
        t.Fatal("New returned err. mode=1, invert=false")
    }

    expected := "    ██████    ████  ████\n████    ████████  ██  ██\n"
    if u.Draw() != expected {
        fmt.Println(u.Draw())
        t.Fatal("Draw returned err. mode=1, invert=false")
    }

    u.Invert = true

    expected =  "████      ████    ██    \n    ████        ██  ██  \n"
    if u.Draw() != expected {
        fmt.Println(u.Draw())
        t.Fatal("Draw returned err. mode=1, invert=true")
    }
}

func TestMode02(t *testing.T) {

    data := int2bool(bits_int)
    u, err := New(2, false, &data)
    if err != nil {
        t.Fatal("New returned err. mode=2, invert=false")
    }

    expected := "▄▄▀▀█▄▄█▀▄▀█\n"
    if u.Draw() != expected {
        fmt.Println(u.Draw())
        t.Fatal("Draw returned err. mode=2, invert=false")
    }

    u.Invert = true

    expected =  "▀▀▄▄ ▀▀ ▄▀▄ \n"
    if u.Draw() != expected {
        fmt.Println(u.Draw())
        t.Fatal("Draw returned err. mode=2, invert=true")
    }
}

func TestMode03(t *testing.T) {

    data := int2bool(bits_int)
    u, err := New(3, false, &data)
    if err != nil {
        t.Fatal("New returned err. mode=3, invert=false")
    }

    expected := "▄▀▙▟▚▜\n"
    if u.Draw() != expected {
        fmt.Println(u.Draw())
        t.Fatal("Draw returned err. mode=3, invert=false")
    }

    u.Invert = true

    expected =  "▀▄▝▘▞▖\n"
    if u.Draw() != expected {
        fmt.Println(u.Draw())
        t.Fatal("Draw returned err. mode=3, invert=true")
    }
}
