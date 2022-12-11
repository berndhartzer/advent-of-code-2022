package aoc

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sort"
	"testing"
	"time"
)

type monkey struct {
	items []int
	inspectFn func(worry int) int
	throwFn func(item int)
	inspections int
}

func (m *monkey) inspectItems() {
	for _, _ = range m.items {
		m.inspect()
		m.throwItem()
	}
}

func (m *monkey) inspect() {
	m.inspections++

	worryLevel := m.items[0]

	newWorryLevel := m.inspectFn(worryLevel)

	newWorryLevel /= 3

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
	return monkeyInTheMiddle(input, 20)
}

func monkeyInTheMiddle(input []string, rounds int) int {
	monkeys := []*monkey{}

	// regex to find numbers
	re := regexp.MustCompile("[0-9]+")

	// newMonkey := &monkey{}
	// Parse input, set state
	// for _, line := range input 
	for i := 0; i < len(input); i += 7 {
		fmt.Println("===", input[i])

		newMonkey := &monkey{
			items: []int{},
		}

		itemsLine := input[i+1]
		opLine := input[i+2]
		testLine := input[i+3]
		testTrueLine := input[i+4]
		testFalseLine := input[i+5]

		fmt.Println(itemsLine)
		itemsStr := re.FindAllString(itemsLine, -1)
		// fmt.Println(itemsStr)
		for _, nStr := range itemsStr {
			n, err := strconv.Atoi(nStr)
			if err != nil {
				panic("failed to convert string to int")
			}
			newMonkey.items = append(newMonkey.items, n)
		}

		fmt.Println(opLine)
		opLineTrim := strings.TrimPrefix(opLine, "  Operation: new = ")
		// fmt.Println(opLineTrim)
		opLineSplit := strings.Fields(opLineTrim)
		// for _, opPart := range opLineSplit {
		// 	fmt.Println(opPart)
		// }
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

			if useOp == "+" {
				return a + b
			}
			return a * b
		}

		fmt.Println(testLine)
		fmt.Println(testTrueLine)
		fmt.Println(testFalseLine)
		testNumStr := re.FindString(testLine)
		testNum, err := strconv.Atoi(testNumStr)
		if err != nil {
			panic("failed to convert string to int")
		}
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
			if worry % testNum == 0 {
				target = monkeys[testTrueNum]
			} else {
				target = monkeys[testFalseNum]
			}
			target.catchItem(worry)
		}


		monkeys = append(monkeys, newMonkey)

	}


	fmt.Println(monkeys)
	fmt.Println(monkeys[0].items)

	// worryLevel := 0
	// rounds := 20

	for i := 0; i < rounds; i++ {
		fmt.Println("round:", i+1)

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
