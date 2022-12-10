package aoc

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

type cpu struct {
	xRegister int
	cycles int
	checkCycles []int // must be sorted
	signals []int
}

func (c *cpu) execute(instruction string) (bool, int) {
	exit, signal := c.runCycle()
	if exit {
		return exit, signal
	}

	if instruction == "noop" {
		return false, 0
	}

	parts := strings.Fields(instruction)
	n, err := strconv.Atoi(parts[1])
	if err != nil {
		panic("failed to convert string to int")
	}

	exit, signal = c.runCycle()
	if exit {
		return exit, signal
	}

	c.xRegister += n
	return false, 0
}

func (c *cpu) runCycle() (bool, int) {
	c.cycles++

	if c.cycles != c.checkCycles[0] {
		return false, 0
	}

	if c.signals == nil {
		c.signals = []int{}
	}
	c.signals = append(c.signals, c.cycles*c.xRegister)

	c.checkCycles = c.checkCycles[1:]

	if len(c.checkCycles) != 0 {
		return false, 0
	}

	total := c.signals[0]
	for i := 1; i < len(c.signals); i++ {
		total += c.signals[i]
	}

	return true, total
}

func cathodeRayTubePartOne(input []string) int {
	cpu := &cpu{
		xRegister: 1,
		checkCycles: []int{20, 60, 100, 140, 180, 220},
	}

	exit := false
	signal := 0
	for _, instruction := range input {
		exit, signal = cpu.execute(instruction)
		if exit {
			break
		}
	}

	return signal
}

func TestDayTen(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/10.txt")
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
					"addx 15",
					"addx -11",
					"addx 6",
					"addx -3",
					"addx 5",
					"addx -1",
					"addx -8",
					"addx 13",
					"addx 4",
					"noop",
					"addx -1",
					"addx 5",
					"addx -1",
					"addx 5",
					"addx -1",
					"addx 5",
					"addx -1",
					"addx 5",
					"addx -1",
					"addx -35",
					"addx 1",
					"addx 24",
					"addx -19",
					"addx 1",
					"addx 16",
					"addx -11",
					"noop",
					"noop",
					"addx 21",
					"addx -15",
					"noop",
					"noop",
					"addx -3",
					"addx 9",
					"addx 1",
					"addx -3",
					"addx 8",
					"addx 1",
					"addx 5",
					"noop",
					"noop",
					"noop",
					"noop",
					"noop",
					"addx -36",
					"noop",
					"addx 1",
					"addx 7",
					"noop",
					"noop",
					"noop",
					"addx 2",
					"addx 6",
					"noop",
					"noop",
					"noop",
					"noop",
					"noop",
					"addx 1",
					"noop",
					"noop",
					"addx 7",
					"addx 1",
					"noop",
					"addx -13",
					"addx 13",
					"addx 7",
					"noop",
					"addx 1",
					"addx -33",
					"noop",
					"noop",
					"noop",
					"addx 2",
					"noop",
					"noop",
					"noop",
					"addx 8",
					"noop",
					"addx -1",
					"addx 2",
					"addx 1",
					"noop",
					"addx 17",
					"addx -9",
					"addx 1",
					"addx 1",
					"addx -3",
					"addx 11",
					"noop",
					"noop",
					"addx 1",
					"noop",
					"addx 1",
					"noop",
					"noop",
					"addx -13",
					"addx -19",
					"addx 1",
					"addx 3",
					"addx 26",
					"addx -30",
					"addx 12",
					"addx -1",
					"addx 3",
					"addx 1",
					"noop",
					"noop",
					"noop",
					"addx -9",
					"addx 18",
					"addx 1",
					"addx 2",
					"noop",
					"noop",
					"addx 9",
					"noop",
					"noop",
					"noop",
					"addx -1",
					"addx 2",
					"addx -37",
					"addx 1",
					"addx 3",
					"noop",
					"addx 15",
					"addx -21",
					"addx 22",
					"addx -6",
					"addx 1",
					"noop",
					"addx 2",
					"addx 1",
					"noop",
					"addx -10",
					"noop",
					"noop",
					"addx 20",
					"addx 1",
					"addx 2",
					"addx 2",
					"addx -6",
					"addx -11",
					"noop",
					"noop",
					"noop",
				},
				expected: 13140,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, cathodeRayTubePartOne)
	})
}
