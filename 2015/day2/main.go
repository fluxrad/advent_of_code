package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Present is a Christmas present, lovingly crafted and wrapped by Santa's elves.
type Present struct {
	Length int
	Width  int
	Height int
}

// WrappingPaper returns the amount of wrapping paper necessary to wrap the Present
func (p Present) WrappingPaper() int {
	surfaces := []int{
		p.Length * p.Width,
		p.Width * p.Height,
		p.Height * p.Length,
	}

	// Find the smallest surface for the extra
	extra := surfaces[0]
	for _, surface := range surfaces {
		if extra > surface {
			extra = surface
		}
	}

	return (2*(surfaces[0]+surfaces[1]+surfaces[2]) + extra)
}

// Ribbon returns the amount of ribbon necessary to make the present look
// beautiful, including amount of material necessary for a bow
func (p Present) Ribbon() int {
	bow := p.Length * p.Width * p.Height

	// find two short sides
	sides := []int{p.Length, p.Width, p.Height}
	sort.Ints(sides)

	// 2x the two shortest sides, plus a bow)
	return (2*(sides[0]+sides[1]) + bow)
}

// NewPresent returns a pointer to a new present from a string "LxWxH"
func NewPresent(s string) *Present {
	lwh := strings.Split(s, "x")

	p := &Present{
		Length: measure(lwh[0]),
		Width:  measure(lwh[1]),
		Height: measure(lwh[2]),
	}

	return p
}

func measure(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatalf("Couldn't open file %s", err)
	}
	defer f.Close()

	totalWrappingPaper := 0
	totalRibbon := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		p := NewPresent(scanner.Text())
		totalWrappingPaper += p.WrappingPaper()
		totalRibbon += p.Ribbon()
	}

	fmt.Println("Total wrapping paper needed is ", totalWrappingPaper, " square feet")
	fmt.Println("Total ribbon needed is ", totalRibbon, " feet")
}
