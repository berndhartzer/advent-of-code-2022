package aoc

import (
	"fmt"
	"testing"
	"time"
)

func treeTopTreeHousePartOne(input []string) int {
	count := 0

	// edges are visible
	count += len(input) * 2
	count += (len(input[0]) - 2) * 2

	columns := make([]string, len(input[0]))

	for y := 1; y < len(input)-1; y++ {
		for x := 1; x < len(input[y])-1; x++ {
			height := input[y][x]

			// Get cached column, or calculate
			column := func() string {
				if len(columns[x]) != 0 {
					return columns[x]
				}

				s := ""
				for j := 0; j < len(input); j++ {
					s += string(input[j][x])
				}
				columns[x] = s
				return s
			}()

			visible := true
			up := column[:y]
			for _, h := range up {
				if byte(h) >= height {
					visible = false
					break
				}
			}
			if visible {
				count++
				continue
			}

			visible = true
			down := column[y+1:]
			for _, h := range down {
				if byte(h) >= height {
					visible = false
					break
				}
			}
			if visible {
				count++
				continue
			}

			visible = true
			left := input[y][:x]
			for _, h := range left {
				if byte(h) >= height {
					visible = false
					break
				}
			}
			if visible {
				count++
				continue
			}

			visible = true
			right := input[y][x+1:]
			for _, h := range right {
				if byte(h) >= height {
					visible = false
					break
				}
			}
			if visible {
				count++
				continue
			}
		}
	}

	return count
}

func TestDayEight(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/08.txt")
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
					"30373",
					"25512",
					"65332",
					"33549",
					"35390",
				},
				expected: 21,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, treeTopTreeHousePartOne)
	})
}
