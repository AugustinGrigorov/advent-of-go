package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type opMode int64

const (
	positionMode  opMode = 0
	immediateMode opMode = 1
	relativeMode  opMode = 2
)

type coordinates struct {
	x, y int64
}

type arcade struct {
	screen                 map[string]int64
	instructionsExecuted   int
	latestX, latestY       int64
	minX, minY, maxX, maxY int64
	score                  int64
	paddleCoordinates      coordinates
	ballCoordinates        coordinates
}

var mode = flag.String("mode", "bot", "Determines the mode the game should be played in")
var speed = flag.String("speed", "fast", "Determines the bot gamespeed")

func main() {
	flag.Parse()
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	inputString := string(input)
	inputArray := strings.Split(strings.Trim(inputString, "\n"), ",")
	output := make(map[int64]int64, len(inputArray))
	for index, opcode := range inputArray {
		output[int64(index)], err = strconv.ParseInt(opcode, 10, 64)
		if err != nil {
			panic(err)
		}
	}

	output[0] = 2

	arc := arcade{
		screen: make(map[string]int64),
	}

	executeIndex(0, 0, output, &arc)
	fmt.Println(arc.score)
}

func (a *arcade) executeCommand(command int64) {
	switch a.instructionsExecuted % 3 {
	case 0:
		a.latestX = command
		if command > a.maxX {
			a.maxX = command
		}
		if command < a.minX {
			a.minX = command
		}
	case 1:
		a.latestY = command
		if command > a.maxY {
			a.maxY = command
		}
		if command < a.minY {
			a.minY = command
		}
	case 2:
		if a.latestX == -1 && a.latestY == 0 {
			a.score = command
		} else {
			if command == 4 {
				a.ballCoordinates = coordinates{x: a.latestX, y: a.latestY}
			}
			if command == 3 {
				a.paddleCoordinates = coordinates{x: a.latestX, y: a.latestY}
			}
			a.screen[fmt.Sprintf("%v:%v", a.latestX, a.latestY)] = command
		}
	}
	a.instructionsExecuted++
}

func (a *arcade) renderCurrentState() {
	for y := a.maxY; y >= a.minY; y-- {
		fmt.Println()
		for x := a.minX; x < a.maxX; x++ {
			fmt.Print(interpretBlock(a.screen[fmt.Sprintf("%v:%v", x, y)]))
		}
	}
}

func (a *arcade) getDirection() int64 {
	if *mode == "player" {
		fmt.Println()
		var i int64
		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			panic("invalid input")
		}
		return i
	} else {
		if a.paddleCoordinates.x < a.ballCoordinates.x {
			a.paddleCoordinates.x++
			return 1
		}
		if a.paddleCoordinates.x > a.ballCoordinates.x {
			a.paddleCoordinates.x--
			return -1
		}
		return 0
	}
}

func interpretBlock(blockData int64) string {
	switch blockData {
	case 0:
		return " "
	case 1:
		return "|"
	case 2:
		return "#"
	case 3:
		return "-"
	case 4:
		return "o"
	default:
		panic("unsupported block")
	}
}

func executeIndex(index, relativeBase int64, instructions map[int64]int64, a *arcade) map[int64]int64 {
	operator := instructions[index]
	switch operator % 100 {
	case 1:
		flags := getFlags(operator)
		paramIndexes := getParameterIndexes(index, relativeBase, instructions, flags, 3)
		instructions[paramIndexes[2]] = instructions[paramIndexes[0]] + instructions[paramIndexes[1]]
		return executeIndex(index+4, relativeBase, instructions, a)
	case 2:
		flags := getFlags(operator)
		paramIndexes := getParameterIndexes(index, relativeBase, instructions, flags, 3)
		instructions[paramIndexes[2]] = instructions[paramIndexes[0]] * instructions[paramIndexes[1]]
		return executeIndex(index+4, relativeBase, instructions, a)
	case 3:
		a.renderCurrentState()
		if *mode == "bot" {
			if *speed == "slow" {
				time.Sleep(250 * time.Millisecond)
			}
		}
		input := a.getDirection()
		flags := getFlags(operator)
		paramIndexes := getParameterIndexes(index, relativeBase, instructions, flags, 1)
		instructions[paramIndexes[0]] = input
		return executeIndex(index+2, relativeBase, instructions, a)
	case 4:
		flags := getFlags(operator)
		paramIndexes := getParameterIndexes(index, relativeBase, instructions, flags, 1)
		a.executeCommand(instructions[paramIndexes[0]])
		return executeIndex(index+2, relativeBase, instructions, a)
	case 5:
		flags := getFlags(operator)
		paramIndexes := getParameterIndexes(index, relativeBase, instructions, flags, 2)
		if instructions[paramIndexes[0]] > 0 {
			return executeIndex(instructions[paramIndexes[1]], relativeBase, instructions, a)
		}
		return executeIndex(index+3, relativeBase, instructions, a)
	case 6:
		flags := getFlags(operator)
		paramIndexes := getParameterIndexes(index, relativeBase, instructions, flags, 2)
		if instructions[paramIndexes[0]] == 0 {
			return executeIndex(instructions[paramIndexes[1]], relativeBase, instructions, a)
		}
		return executeIndex(index+3, relativeBase, instructions, a)
	case 7:
		flags := getFlags(operator)
		paramIndexes := getParameterIndexes(index, relativeBase, instructions, flags, 3)
		if instructions[paramIndexes[0]] < instructions[paramIndexes[1]] {
			instructions[paramIndexes[2]] = 1
		} else {
			instructions[paramIndexes[2]] = 0
		}
		return executeIndex(index+4, relativeBase, instructions, a)
	case 8:
		flags := getFlags(operator)
		paramIndexes := getParameterIndexes(index, relativeBase, instructions, flags, 3)
		if instructions[paramIndexes[0]] == instructions[paramIndexes[1]] {
			instructions[paramIndexes[2]] = 1
		} else {
			instructions[paramIndexes[2]] = 0
		}
		return executeIndex(index+4, relativeBase, instructions, a)
	case 9:
		flags := getFlags(operator)
		paramIndexes := getParameterIndexes(index, relativeBase, instructions, flags, 1)
		return executeIndex(index+2, relativeBase+instructions[paramIndexes[0]], instructions, a)
	case 99:
		return instructions
	}
	return instructions
}

func getParameterIndexes(index, relativeBase int64, instructions map[int64]int64, flags []opMode, requiredParams int64) (params []int64) {
	for i := int64(0); i < requiredParams; i++ {
		params = append(params, getParameterIndex(index+i+1, relativeBase, instructions, flags[i]))
	}
	return
}

func getParameterIndex(parameterIndex int64, relativeBase int64, instructions map[int64]int64, mode opMode) int64 {
	switch mode {
	case positionMode:
		return instructions[parameterIndex]
	case immediateMode:
		return parameterIndex
	case relativeMode:
		return instructions[parameterIndex] + relativeBase
	}
	return 0
}

func getFlags(operator int64) []opMode {
	return []opMode{
		opMode(operator / 100 % 10),
		opMode(operator / 1000 % 10),
		opMode(operator / 10000 % 10),
	}
}
