package aoc

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

type fileTreeNode struct {
	name     string
	parent   *fileTreeNode
	children map[string]*fileTreeNode
	size     int
}

func (f *fileTreeNode) totalUnderThreshold(threshold int) int {
	total := 0

	if f.size < threshold {
		total += f.size
	}

	for _, child := range f.children {
		total += child.totalUnderThreshold(threshold)
	}

	return total
}

func noSpaceLeftOnDevicePartOne(input []string) int {
	root := &fileTreeNode{
		name:     "/",
		parent:   nil,
		children: nil,
	}
	fs := root

	for _, line := range input {
		parts := strings.Fields(line)

		if parts[0] == "$" {
			if parts[1] == "cd" {
				to := parts[2]
				switch to {
				case "/":
					// special case for first line of input
					continue

				case "..":
					fs = fs.parent

				default:
					if fs.children == nil {
						fs.children = map[string]*fileTreeNode{}
					}

					child, ok := fs.children[to]
					if !ok {
						child = &fileTreeNode{
							name:   to,
							parent: fs,
						}

						fs.children[to] = child
					}

					fs = child
				}
			}

			continue
		}

		if parts[0] == "dir" {
			continue
		}

		fileSize, err := strconv.Atoi(parts[0])
		if err != nil {
			panic("failed to convert string to int")
		}
		fs.size += fileSize

		p := fs.parent
		for {
			if p == nil {
				break
			}

			p.size += fileSize
			p = p.parent
		}
	}

	total := root.totalUnderThreshold(100000)

	return total
}

func TestDaySeven(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/07.txt")
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
					"$ cd /",
					"$ ls",
					"dir a",
					"14848514 b.txt",
					"8504156 c.dat",
					"dir d",
					"$ cd a",
					"$ ls",
					"dir e",
					"29116 f",
					"2557 g",
					"62596 h.lst",
					"$ cd e",
					"$ ls",
					"584 i",
					"$ cd ..",
					"$ cd ..",
					"$ cd d",
					"$ ls",
					"4060174 j",
					"8033020 d.log",
					"5626152 d.ext",
					"7214296 k",
				},
				expected: 95437,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, noSpaceLeftOnDevicePartOne)
	})
}
