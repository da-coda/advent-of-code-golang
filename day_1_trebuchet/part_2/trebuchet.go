package main

import (
	"advent-of-code-golang/common"
	"fmt"
	_ "net/http/pprof"
	"os"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println(run(os.Args[1]))
}

func run(path string) int {
	sum := 0
	common.ReadFile(path, func(s string) {
		value := calcCalibrationValue(s)
		sum += value
	})
	return sum
}

func calcCalibrationValue(line string) int {
	firstValue := 0
	lastValue := 0
	sm := &StateMachine{
		isFinishedState:    false,
		currState:          "",
		finishedStateValue: 0,
	}
	for _, char := range line {
		sm.TransitionState(char)
		if finished, value := sm.IsFinished(); finished {
			if firstValue == 0 {
				firstValue = value
			}
			lastValue = value
			sm.Reset()
		}
	}

	return firstValue*10 + lastValue
}

type StateMachine struct {
	isFinishedState    bool
	currState          string
	finishedStateValue int
}

func (sm *StateMachine) TransitionState(character rune) {
	if unicode.IsDigit(character) {
		sm.isFinishedState = true
		sm.finishedStateValue, _ = strconv.Atoi(string(character))
		sm.currState = ""
		return
	}
	switch sm.currState {
	case "":
		sm.currState = string(character)
	case "o":
		switch character {
		case 'n':
			sm.currState = "on"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "on":
		switch character {
		case 'e':
			sm.finishedStateValue = 1
			sm.isFinishedState = true
			sm.currState = "e"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "t":
		switch character {
		case 'w':
			sm.currState = "tw"
		case 'h':
			sm.currState = "th"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "tw":
		switch character {
		case 'o':
			sm.finishedStateValue = 2
			sm.isFinishedState = true
			sm.currState = "o"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "th":
		switch character {
		case 'r':
			sm.currState = "thr"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "thr":
		switch character {
		case 'e':
			sm.currState = "thre"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "thre":
		switch character {
		case 'e':
			sm.finishedStateValue = 3
			sm.isFinishedState = true
			sm.currState = "e"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "f":
		switch character {
		case 'o':
			sm.currState = "fo"
		case 'i':
			sm.currState = "fi"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "fo":
		switch character {
		case 'u':
			sm.currState = "fou"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "fou":
		switch character {
		case 'r':
			sm.finishedStateValue = 4
			sm.isFinishedState = true
			sm.currState = ""
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "fi":
		switch character {
		case 'v':
			sm.currState = "fiv"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "fiv":
		switch character {
		case 'e':
			sm.finishedStateValue = 5
			sm.isFinishedState = true
			sm.currState = "e"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "s":
		switch character {
		case 'i':
			sm.currState = "si"
		case 'e':
			sm.currState = "se"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "si":
		switch character {
		case 'x':
			sm.finishedStateValue = 6
			sm.isFinishedState = true
			sm.currState = ""
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "se":
		switch character {
		case 'v':
			sm.currState = "sev"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "sev":
		switch character {
		case 'e':
			sm.currState = "seve"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "seve":
		switch character {
		case 'n':
			sm.finishedStateValue = 7
			sm.isFinishedState = true
			sm.currState = "n"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "e":
		switch character {
		case 'i':
			sm.currState = "ei"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "ei":
		switch character {
		case 'g':
			sm.currState = "eig"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "eig":
		switch character {
		case 'h':
			sm.currState = "eigh"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "eigh":
		switch character {
		case 't':
			sm.finishedStateValue = 8
			sm.isFinishedState = true
			sm.currState = "t"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "n":
		switch character {
		case 'i':
			sm.currState = "ni"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "ni":
		switch character {
		case 'n':
			sm.currState = "nin"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "nin":
		switch character {
		case 'e':
			sm.finishedStateValue = 9
			sm.isFinishedState = true
			sm.currState = "e"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	default:
		sm.currState = string(character)
	}
}

func (sm *StateMachine) IsFinished() (bool, int) {
	return sm.isFinishedState, sm.finishedStateValue
}

func (sm *StateMachine) Reset() {
	sm.isFinishedState = false
	sm.finishedStateValue = 0
}
