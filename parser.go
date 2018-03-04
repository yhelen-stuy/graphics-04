package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseFile(filename string, t *Matrix, e *Matrix, image *Image) error {
	f, err := os.Open(filename)
	if err != nil {
		return errors.New("Couldn't open file")
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		switch c := strings.TrimSpace(scanner.Text()); c {
		case "line":
			scanner.Scan()
			arg := strings.TrimSpace(scanner.Text())
			args := strings.Fields(arg)
			if len(args) != 6 {
				fmt.Errorf("Scale: Incorrect # of args. Got: %d, expected: 6\n", len(args))
			}
			fargs := numerize(args)
			fmt.Printf("unfinished line")
			fmt.Println(fargs)
		case "ident":
			t.Ident()
		case "scale":
			scanner.Scan()
			arg := strings.TrimSpace(scanner.Text())
			args := strings.Fields(arg)
			if len(args) != 3 {
				fmt.Errorf("Scale: Incorrect # of args. Got: %d, expected: 3\n", len(args))
			}
			fargs := numerize(args)
			scale := MakeScale(fargs[0], fargs[1], fargs[2])
			t, _ = t.Mult(scale)
		case "translate":
			scanner.Scan()
			arg := strings.TrimSpace(scanner.Text())
			args := strings.Fields(arg)
			if len(args) != 3 {
				fmt.Errorf("Scale: Incorrect # of args. Got: %d, expected: 3\n", len(args))
			}
			fargs := numerize(args)
			translate := MakeTranslate(fargs[0], fargs[1], fargs[2])
			t, _ = t.Mult(translate)
		case "rotate":
			fmt.Println("unfinished rotate")
		case "apply":
			e.Mult(t)
		case "display":
			image.DrawLines(e, Color{r: 0, b: 0, g: 255})
			image.Display()
		case "save":
			scanner.Scan()
			arg := strings.TrimSpace(scanner.Text())
			args := strings.Fields(arg)
			if len(args) != 1 {
				fmt.Errorf("Scale: Incorrect # of args. Got: %d, expected: 3\n", len(args))
			}
			image.DrawLines(e, Color{r: 0, b: 0, g: 255})
			image.SavePPM(args[0])
		case "quit":
			break
		default:
			fmt.Println(c)
		}
	}
	return nil
}

// Error handling?
func numerize(args []string) []float64 {
	fargs := make([]float64, len(args))
	for i, val := range args {
		fargs[i], _ = strconv.ParseFloat(val, 64)
	}
	return fargs
}
