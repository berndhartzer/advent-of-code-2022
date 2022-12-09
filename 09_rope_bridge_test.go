package aoc

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

type ropeKnot struct {
	x, y int
}

func ropeBridgePartOne(input []string) int {
	head := &ropeKnot{}
	tail := &ropeKnot{}
	visited := map[string]bool{}

	for _, line := range input {
		dist, err := strconv.Atoi(line[2:])
		if err != nil {
			panic("failed to convert string to int")
		}

		xMove := 0
		yMove := 0
		switch line[0] {
		case 'U':
			yMove = -1
		case 'D':
			yMove = 1
		case 'L':
			xMove = -1
		case 'R':
			xMove = 1
		}

		for i := 0; i < dist; i++ {
			head.x += xMove
			head.y += yMove

			xDist := head.x - tail.x
			yDist := head.y - tail.y

			switch {
			case (abs(xDist) > 0 && abs(yDist) > 0) && (abs(xDist) > 1 || abs(yDist) > 1):
				// diagonal move
				if xDist > 0 {
					tail.x += 1
				} else {
					tail.x -= 1
				}
				if yDist > 0 {
					tail.y += 1
				} else {
					tail.y -= 1
				}

			case xDist > 1:
				tail.x += 1
			case xDist < -1:
				tail.x -= 1
			case yDist > 1:
				tail.y += 1
			case yDist < -1:
				tail.y -= 1
			}

			tailPos := fmt.Sprintf("%d,%d", tail.x, tail.y)
			visited[tailPos] = true
		}
	}

	return len(visited)
}

func ropeBridgePartTwo(input []string) int {
	rope := []*ropeKnot{
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
	}
	visited := map[string]bool{}

	for _, line := range input {
		dist, err := strconv.Atoi(line[2:])
		if err != nil {
			panic("failed to convert string to int")
		}

		xMove := 0
		yMove := 0
		switch line[0] {
		case 'U':
			yMove = -1
		case 'D':
			yMove = 1
		case 'L':
			xMove = -1
		case 'R':
			xMove = 1
		}

		for i := 0; i < dist; i++ {
			rope[0].x += xMove
			rope[0].y += yMove

			for j := 1; j < len(rope); j++ {
				prevKnot := rope[j-1]
				thisKnot := rope[j]

				xDist := prevKnot.x - thisKnot.x
				yDist := prevKnot.y - thisKnot.y

				switch {
				case (abs(xDist) > 0 && abs(yDist) > 0) && (abs(xDist) > 1 || abs(yDist) > 1):
					// diagonal move
					if xDist > 0 {
						thisKnot.x += 1
					} else {
						thisKnot.x -= 1
					}
					if yDist > 0 {
						thisKnot.y += 1
					} else {
						thisKnot.y -= 1
					}

				case xDist > 1:
					thisKnot.x += 1
				case xDist < -1:
					thisKnot.x -= 1
				case yDist > 1:
					thisKnot.y += 1
				case yDist < -1:
					thisKnot.y -= 1
				}

			}

			tailPos := fmt.Sprintf("%d,%d", rope[9].x, rope[9].y)
			visited[tailPos] = true
		}
	}

	return len(visited)
}

func TestDayNine(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/09.txt")
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
					"R 4",
					"U 4",
					"L 3",
					"D 1",
					"R 4",
					"D 1",
					"L 5",
					"R 2",
				},
				expected: 13,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, ropeBridgePartOne)
	})

	t.Run("part two", func(t *testing.T) {
		tests := map[string]testConfig{
			"test_1": {
				input: []string{
					"R 4",
					"U 4",
					"L 3",
					"D 1",
					"R 4",
					"D 1",
					"L 5",
					"R 2",
				},
				expected: 1,
			},
			"test_2": {
				input: []string{
					"R 5",
					"U 8",
					"L 8",
					"D 3",
					"R 17",
					"D 10",
					"L 25",
					"U 20",
				},
				expected: 36,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, ropeBridgePartTwo)
	})
}
