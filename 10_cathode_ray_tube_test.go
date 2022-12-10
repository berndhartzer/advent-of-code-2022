package aoc

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

type cpu struct {
	xRegister     int
	cycles        int
	postCycleHook func(cycles int, xRegister int) bool
	preCycleHook  func(cycles int, xRegister int) bool
}

func (c *cpu) execute(instruction string) bool {
	exit := c.runCycle()
	if exit {
		return exit
	}

	if instruction == "noop" {
		return false
	}

	parts := strings.Fields(instruction)
	n, err := strconv.Atoi(parts[1])
	if err != nil {
		panic("failed to convert string to int")
	}

	exit = c.runCycle()
	if exit {
		return exit
	}

	c.xRegister += n
	return false
}

func (c *cpu) runCycle() bool {
	exit := c.preCycleHook(c.cycles, c.xRegister)
	if exit {
		return exit
	}

	c.cycles++

	exit = c.postCycleHook(c.cycles, c.xRegister)
	return exit
}

func cathodeRayTubePartOne(input []string) int {
	checkCycles := []int{20, 60, 100, 140, 180, 220}
	signals := []int{}

	preCycleHook := func(cycles int, xRegister int) bool {
		return false
	}

	postCycleHook := func(cycles int, xRegister int) bool {
		if cycles != checkCycles[0] {
			return false
		}

		signals = append(signals, cycles*xRegister)

		checkCycles = checkCycles[1:]

		if len(checkCycles) != 0 {
			return false
		}

		return true
	}

	cpu := &cpu{
		xRegister:     1,
		preCycleHook:  preCycleHook,
		postCycleHook: postCycleHook,
	}

	for _, instruction := range input {
		exit := cpu.execute(instruction)
		if exit {
			break
		}
	}

	total := signals[0]
	for i := 1; i < len(signals); i++ {
		total += signals[i]
	}

	return total
}

func cathodeRayTubePartTwo(input []string) int {
	screen := [240]string{}

	preCycleHook := func(cycles int, xRegister int) bool {
		if cycles >= len(screen) {
			return true
		}

		// translate sprite vertical position according to screen of 40 x 6
		spritePos := xRegister + ((cycles / 40) * 40)

		drawPixel := "."
		if spritePos == cycles || spritePos-1 == cycles || spritePos+1 == cycles {
			drawPixel = "#"
		}

		screen[cycles] = drawPixel

		return false
	}

	postCycleHook := func(cycles int, xRegister int) bool {
		return false
	}

	cpu := &cpu{
		xRegister:     1,
		preCycleHook:  preCycleHook,
		postCycleHook: postCycleHook,
	}

	for _, instruction := range input {
		exit := cpu.execute(instruction)
		if exit {
			break
		}
	}

	rows := [][]string{
		screen[0:40],
		screen[40:80],
		screen[80:120],
		screen[120:160],
		screen[160:200],
		screen[200:240],
	}
	for _, row := range rows {
		fmt.Println(row)
	}

	return 0
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

	t.Run("part two", func(t *testing.T) {
		tests := map[string]testConfig{
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, cathodeRayTubePartTwo)
	})
}
