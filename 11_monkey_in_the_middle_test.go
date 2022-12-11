package aoc

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
)

type monkey struct {
	items       []int
	inspectFn   func(worry int) int
	throwFn     func(item int)
	inspections int
}

func (m *monkey) inspectItems() {
	for range m.items {
		m.inspect()
		m.throwItem()
	}
}

func (m *monkey) inspect() {
	m.inspections++

	worryLevel := m.items[0]
	newWorryLevel := m.inspectFn(worryLevel)
	m.items[0] = newWorryLevel
}

func (m *monkey) throwItem() {
	var item int
	item, m.items = m.items[0], m.items[1:]
	m.throwFn(item)
}

func (m *monkey) catchItem(item int) {
	m.items = append(m.items, item)
}

func monkeyInTheMiddlePartOne(input []string) int {
	return monkeyInTheMiddle(input, 20, true)
}

func monkeyInTheMiddlePartTwo(input []string) int {
	return monkeyInTheMiddle(input, 10000, false)
}

func getMonkeys(input []string, isPartOne bool) []*monkey {
	monkeys := []*monkey{}

	// regex to find numbers
	re := regexp.MustCompile("[0-9]+")

	allTestNums := 1

	for i := 0; i < len(input); i += 7 {
		newMonkey := &monkey{
			items: []int{},
		}

		itemsLine := input[i+1]
		opLine := input[i+2]
		testLine := input[i+3]
		testTrueLine := input[i+4]
		testFalseLine := input[i+5]

		itemsStr := re.FindAllString(itemsLine, -1)
		for _, nStr := range itemsStr {
			n, err := strconv.Atoi(nStr)
			if err != nil {
				panic("failed to convert string to int")
			}
			newMonkey.items = append(newMonkey.items, n)
		}

		opLineTrim := strings.TrimPrefix(opLine, "  Operation: new = ")
		opLineSplit := strings.Fields(opLineTrim)
		a, b := 0, 0
		aUseOld, bUseOld := false, false
		if opLineSplit[0] == "old" {
			aUseOld = true
		} else {
			aNum, err := strconv.Atoi(opLineSplit[0])
			if err != nil {
				panic("failed to convert string to int")
			}
			a = aNum
		}
		if opLineSplit[2] == "old" {
			bUseOld = true
		} else {
			bNum, err := strconv.Atoi(opLineSplit[2])
			if err != nil {
				panic("failed to convert string to int")
			}
			b = bNum
		}
		useOp := opLineSplit[1]
		newMonkey.inspectFn = func(worry int) int {
			if aUseOld {
				a = worry
			}
			if bUseOld {
				b = worry
			}

			newWorry := 0
			if useOp == "+" {
				newWorry = a + b
			} else {
				newWorry = a * b
			}

			// Must be a better way to inject this logic other than a bool flag...
			if isPartOne {
				newWorry /= 3
			} else {
				newWorry %= allTestNums
			}

			return newWorry
		}

		testNumStr := re.FindString(testLine)
		testNum, err := strconv.Atoi(testNumStr)
		if err != nil {
			panic("failed to convert string to int")
		}
		allTestNums *= testNum

		testTrueStr := re.FindString(testTrueLine)
		testTrueNum, err := strconv.Atoi(testTrueStr)
		if err != nil {
			panic("failed to convert string to int")
		}
		testFalseStr := re.FindString(testFalseLine)
		testFalseNum, err := strconv.Atoi(testFalseStr)
		if err != nil {
			panic("failed to convert string to int")
		}

		newMonkey.throwFn = func(worry int) {
			var target *monkey
			if worry%testNum == 0 {
				target = monkeys[testTrueNum]
			} else {
				target = monkeys[testFalseNum]
			}
			target.catchItem(worry)
		}

		monkeys = append(monkeys, newMonkey)
	}

	return monkeys
}

func monkeyInTheMiddle(input []string, rounds int, isPartOne bool) int {
	monkeys := getMonkeys(input, isPartOne)

	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			monkey.inspectItems()
		}
	}

	inspections := []int{}
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.inspections)
	}
	sort.Ints(inspections)

	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func TestDayEleven(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/11.txt")
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
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, monkeyInTheMiddlePartOne)
	})

	t.Run("part two", func(t *testing.T) {
		tests := map[string]testConfig{
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, monkeyInTheMiddlePartTwo)
	})
}
