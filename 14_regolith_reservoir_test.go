package aoc

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// x point where sand starts
const gridXMidPoint = 500

type reservoirCoord struct {
	isRock  bool
	hasSand bool
}

type reservoirGrid struct {
	xLen, yLen int
	yMaxRock   int
	data       [][]*reservoirCoord
}

func (g *reservoirGrid) build(input []string, width int) {
	g.xLen = width * 2
	g.yLen = width / 2 // input looks wider than deeper

	g.data = make([][]*reservoirCoord, g.yLen)
	for i := 0; i < g.yLen; i++ {
		g.data[i] = make([]*reservoirCoord, g.xLen)
	}

	for _, line := range input {
		strNums := getNumbersFromString(line)

		xFrom, err := strconv.Atoi(strNums[0])
		if err != nil {
			panic("failed to convert string to int")
		}
		yFrom, err := strconv.Atoi(strNums[1])
		if err != nil {
			panic("failed to convert string to int")
		}

		for i := 2; i < len(strNums); i += 2 {
			xTo, err := strconv.Atoi(strNums[i])
			if err != nil {
				panic("failed to convert string to int")
			}
			yTo, err := strconv.Atoi(strNums[i+1])
			if err != nil {
				panic("failed to convert string to int")
			}

			xDir, yDir := 0, 0
			xDiff := xFrom - xTo
			yDiff := yFrom - yTo

			if xDiff != 0 {
				if xDiff > 0 {
					xDir = -1
				} else {
					xDir = 1
				}
			}
			if yDiff != 0 {
				if yDiff > 0 {
					yDir = -1
				} else {
					yDir = 1
				}
			}

			for {
				g.markRock(xFrom, yFrom)

				if xFrom == xTo && yFrom == yTo {
					break
				}

				xFrom += xDir
				yFrom += yDir
			}
		}
	}
}

func (g *reservoirGrid) markRock(x, y int) {
	if x > g.xLen {
		panic("grid x too small")
	}
	if y > g.yLen {
		panic("grid y too small")
	}

	if y > g.yMaxRock {
		g.yMaxRock = y
	}

	g.data[y][x] = &reservoirCoord{
		isRock: true,
	}
}

func (g *reservoirGrid) dropSand(untilFull bool) bool {
	xSand, ySand := gridXMidPoint, 0

	for {
		if !untilFull && ySand >= g.yMaxRock {
			return true
		}

		down := g.data[ySand+1][xSand]
		if down == nil || (!down.isRock && !down.hasSand) {
			ySand++
			continue
		}

		left := g.data[ySand+1][xSand-1]
		if left == nil || (!left.isRock && !left.hasSand) {
			ySand++
			xSand--
			continue
		}

		right := g.data[ySand+1][xSand+1]
		if right == nil || (!right.isRock && !right.hasSand) {
			ySand++
			xSand++
			continue
		}

		g.data[ySand][xSand] = &reservoirCoord{
			hasSand: true,
		}

		if untilFull && ySand == 0 {
			return true
		}
		break
	}

	return false
}

func (g *reservoirGrid) addFloor() {
	for i := 0; i < len(g.data[0]); i++ {
		g.data[g.yMaxRock+2][i] = &reservoirCoord{
			isRock: true,
		}
	}
}

func regolithReservoirPartOne(input []string) int {
	grid := reservoirGrid{}
	grid.build(input, gridXMidPoint)

	total := 0
	for {
		exit := grid.dropSand(false)
		if exit {
			break
		}

		total++
	}

	return total
}

func regolithReservoirPartTwo(input []string) int {
	grid := reservoirGrid{}
	grid.build(input, gridXMidPoint)
	grid.addFloor()

	total := 0
	for {
		total++
		exit := grid.dropSand(true)
		if exit {
			break
		}
	}

	return total
}

func TestDayFourteen(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/14.txt")
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
	}

	runTests := func(t *testing.T, tests map[string]testConfig, fn func([]string) int) {
		for name, cfg := range tests {
			cfg := cfg
			t.Run(name, func(t *testing.T) {
				start := time.Now()
				output := fn(cfg.input)
				finish := time.Since(start)
				if cfg.logResult {
					t.Log(fmt.Sprintf("\nsolution:\t%v\nelapsed time:\t%s", output, finish))
					return
				}

				if output != cfg.expected {
					t.Fatalf("Incorrect output - got: %v, want: %v", output, cfg.expected)
				}
			})
		}
	}

	t.Run("part one", func(t *testing.T) {
		tests := map[string]testConfig{
			"test_1": {
				input: []string{
					"498,4 -> 498,6 -> 496,6",
					"503,4 -> 502,4 -> 502,9 -> 494,9",
				},
				expected: 24,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, regolithReservoirPartOne)
	})

	t.Run("part two", func(t *testing.T) {
		tests := map[string]testConfig{
			"test_1": {
				input: []string{
					"498,4 -> 498,6 -> 496,6",
					"503,4 -> 502,4 -> 502,9 -> 494,9",
				},
				expected: 93,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, regolithReservoirPartTwo)
	})
}
