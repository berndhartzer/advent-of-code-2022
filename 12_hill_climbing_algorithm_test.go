package aoc

import (
	"fmt"
	"testing"
	"time"
)

type hillCoord struct {
	x, y int
}

func hillClimbingAlgorithmPartOne(input []string) int {
	grid := make([][]int, len(input))
	start := hillCoord{}
	end := hillCoord{}

	for y, line := range input {
		grid[y] = make([]int, len(line))

		for x, c := range line {
			switch c {
			case 'S':
				start.x, start.y = x, y
				c = 'a'
			case 'E':
				end.x, end.y = x, y
				c = 'z'
			}

			grid[y][x] = int(c)
		}
	}

	return hillClimbingPath(grid, []hillCoord{start}, end)
}

func hillClimbingAlgorithmPartTwo(input []string) int {
	grid := make([][]int, len(input))
	start := []hillCoord{}
	end := hillCoord{}

	for y, line := range input {
		grid[y] = make([]int, len(line))

		for x, c := range line {
			switch c {
			case 'S', 'a':
				start = append(start, hillCoord{x: x, y: y})
				c = 'a'
			case 'E':
				end.x, end.y = x, y
				c = 'z'
			}

			grid[y][x] = int(c)
		}
	}

	return hillClimbingPath(grid, start, end)
}

func hillClimbingPath(grid [][]int, starts []hillCoord, end hillCoord) int {
	shortest := 0

	for _, start := range starts {
		currentShortest := len(grid) * len(grid[0])
		distances := map[hillCoord]int{}
		queue := []hillCoord{start}

		for len(queue) > 0 {
			var current hillCoord
			current, queue = queue[0], queue[1:]

			if current == end && distances[current] < currentShortest {
				currentShortest = distances[current]
				break
			}

			adjacent := []hillCoord{
				{x: current.x - 1, y: current.y},
				{x: current.x + 1, y: current.y},
				{x: current.x, y: current.y - 1},
				{x: current.x, y: current.y + 1},
			}

			for _, adj := range adjacent {
				if adj.x < 0 || adj.x >= len(grid[0]) {
					continue
				}
				if adj.y < 0 || adj.y >= len(grid) {
					continue
				}

				_, ok := distances[adj]
				if ok {
					continue
				}

				if grid[adj.y][adj.x]-grid[current.y][current.x] <= 1 {
					queue = append(queue, adj)
					distances[adj] = distances[current] + 1
				}
			}
		}

		if currentShortest < shortest || shortest == 0 {
			shortest = currentShortest
		}
	}

	return shortest
}

func TestDayTwelve(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/12.txt")
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
					"Sabqponm",
					"abcryxxl",
					"accszExk",
					"acctuvwj",
					"abdefghi",
				},
				expected: 31,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, hillClimbingAlgorithmPartOne)
	})

	t.Run("part two", func(t *testing.T) {
		tests := map[string]testConfig{
			"test_1": {
				input: []string{
					"Sabqponm",
					"abcryxxl",
					"accszExk",
					"acctuvwj",
					"abdefghi",
				},
				expected: 29,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, hillClimbingAlgorithmPartTwo)
	})
}
