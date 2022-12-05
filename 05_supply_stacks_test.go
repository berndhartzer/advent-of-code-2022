package aoc

import (
	"fmt"
	"strconv"
	"testing"
	"time"
	"regexp"
)

type stack struct {
	values []string
}

func (s *stack) push(item string) {
	s.values = append(s.values, item)
}

func (s *stack) pop() string {
	idx := len(s.values) - 1
	last := s.values[idx]
	s.values = s.values[:idx]
	return last
}

func (s *stack) peek() string{
	idx := len(s.values) - 1
	return s.values[idx]
}

func (s *stack) echo() {
	for _, v := range s.values {
		fmt.Println(v)
	}
}

func supplyStacksPartOne(input []string) string {
	stacks := []*stack{}
	stackMap := map[int]int{}

	endStacksIdx := 0
	for inputIdx, line := range input {
		if line == "" {
			endStacksIdx = inputIdx
			break
		}
	}

	// parse stacks by looping backwards through input
	for k := endStacksIdx-1; k >= 0; k-- {
		line := input[k]

		// loop through columns
		// this relies on the fact that the input stacks are padded with spaces
		for i := 1; i < len(line); i += 4 {
			if k == endStacksIdx-1 {
				stacks = append(stacks, &stack{})
				stackMap[i] = len(stacks)-1
				continue
			}

			if string(line[i]) != " " {
				stacks[stackMap[i]].push(string(line[i]))
			}
		}
	}

	// regex to find numbers
	re := regexp.MustCompile("[0-9]+")

	// loop through instructions
	for i := endStacksIdx+1; i < len(input); i++ {
		nums := re.FindAllString(input[i], -1)

		numToMove, err := strconv.Atoi(nums[0])
		if err != nil {
			panic("failed to convert string to int")
		}
		from, err := strconv.Atoi(nums[1])
		if err != nil {
			panic("failed to convert string to int")
		}
		to, err := strconv.Atoi(nums[2])
		if err != nil {
			panic("failed to convert string to int")
		}

		for j := 0; j < numToMove; j++ {
			fromStack := stacks[from-1]
			toStack := stacks[to-1]

			item := fromStack.pop()
			toStack.push(item)
		}
	}

	tops := ""
	for _, v := range stacks {
		tops += v.peek()
	}

	return tops
}

func TestDayFive(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  string
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/05.txt")
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
	}

	runTests := func(t *testing.T, tests map[string]testConfig, fn func([]string) string) {
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
					"    [D]    ",
					"[N] [C]    ",
					"[Z] [M] [P]",
					" 1   2   3 ",
					"",
					"move 1 from 2 to 1",
					"move 3 from 1 to 3",
					"move 2 from 2 to 1",
					"move 1 from 1 to 2",
				},
				expected: "CMZ",
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, supplyStacksPartOne)
	})
}
