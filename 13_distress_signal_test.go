package aoc

import (
	"encoding/json"
	"fmt"
	"sort"
	"testing"
	"time"
)

func inOrder(left, right any) int {
	leftNum, leftIsNum := left.(float64)
	rightNum, rightIsNum := right.(float64)
	if leftIsNum && rightIsNum {
		switch {
		case leftNum < rightNum:
			return 1
		case leftNum == rightNum:
			return 0
		default:
			return -1
		}
	}

	leftArr, leftIsArray := left.([]any)
	rightArr, rightIsArray := right.([]any)
	if leftIsArray && rightIsNum {
		rightArr = []any{rightNum}
	}
	if rightIsArray && leftIsNum {
		leftArr = []any{leftNum}
	}

	for i := 0; i < len(leftArr) && i < len(rightArr); i++ {
		ordered := inOrder(leftArr[i], rightArr[i])
		if ordered != 0 {
			return ordered
		}
	}

	switch leftLen := len(leftArr); {
	case leftLen == len(rightArr):
		return 0
	case leftLen > len(rightArr):
		return -1
	default:
		return 1
	}
}

func distressSignalPartOne(input []string) int {
	packetIdx := 1
	total := 0

	for i := 0; i < len(input); i += 3 {
		var left, right []any
		err := json.Unmarshal([]byte(input[i]), &left)
		if err != nil {
			panic("error unmarshaling")
		}
		err = json.Unmarshal([]byte(input[i+1]), &right)
		if err != nil {
			panic("error unmarshaling")
		}

		ordered := inOrder(left, right)
		if ordered == 1 {
			total += packetIdx
		}

		packetIdx++
	}

	return total
}

func distressSignalPartTwo(input []string) int {
	allPackets := [][]any{}

	for i := 0; i < len(input); i += 3 {
		var left, right []any
		err := json.Unmarshal([]byte(input[i]), &left)
		if err != nil {
			panic("error unmarshaling")
		}
		err = json.Unmarshal([]byte(input[i+1]), &right)
		if err != nil {
			panic("error unmarshaling")
		}

		allPackets = append(allPackets, left, right)
	}

	allPackets = append(
		allPackets,
		[]any{[]any{float64(2)}},
		[]any{[]any{float64(6)}},
	)

	sort.Slice(allPackets, func(i, j int) bool {
		return inOrder(allPackets[i], allPackets[j]) == 1
	})

	total := 1
	for i, p := range allPackets {
		if len(p) != 1 {
			continue
		}
		innerArray, ok := p[0].([]any)
		if !ok {
			continue
		}
		if len(innerArray) != 1 {
			continue
		}
		innerItem, ok := innerArray[0].(float64)
		if !ok {
			continue
		}

		if innerItem == 2 || innerItem == 6 {
			total *= i + 1
		}
	}

	return total
}

func TestDayThirteen(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/13.txt")
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
					"[1,1,3,1,1]", // true
					"[1,1,5,1,1]",
					"",
					"[[1],[2,3,4]]", // true
					"[[1],4]",
					"",
					"[9]", // false
					"[[8,7,6]]",
					"",
					"[[4,4],4,4]", // true
					"[[4,4],4,4,4]",
					"",
					"[7,7,7,7]", // false
					"[7,7,7]",
					"",
					"[]", // true
					"[3]",
					"",
					"[[[]]]", // false
					"[[]]",
					"",
					"[1,[2,[3,[4,[5,6,7]]]],8,9]", // false
					"[1,[2,[3,[4,[5,6,0]]]],8,9]",
				},
				expected: 13,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, distressSignalPartOne)
	})

	t.Run("part two", func(t *testing.T) {
		tests := map[string]testConfig{
			"test_1": {
				input: []string{
					"[1,1,3,1,1]",
					"[1,1,5,1,1]",
					"",
					"[[1],[2,3,4]]",
					"[[1],4]",
					"",
					"[9]",
					"[[8,7,6]]",
					"",
					"[[4,4],4,4]",
					"[[4,4],4,4,4]",
					"",
					"[7,7,7,7]",
					"[7,7,7]",
					"",
					"[]",
					"[3]",
					"",
					"[[[]]]",
					"[[]]",
					"",
					"[1,[2,[3,[4,[5,6,7]]]],8,9]",
					"[1,[2,[3,[4,[5,6,0]]]],8,9]",
				},
				expected: 140,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, distressSignalPartTwo)
	})
}
