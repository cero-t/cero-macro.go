package converter

import (
	"../processor"
	"strings"
	"regexp"
	"strconv"
	"log"
)

var (
	blanks = regexp.MustCompile(`\s+`)
	nums = regexp.MustCompile(`^\d+$`)
)

func lineToState(line *string) *State {
	state := State{}

	if line == nil || len(*line) == 0 {
		return &state
	}

	pair := blanks.Split(*line, 2)
	if len(pair) != 2 {
		return &state
	}

	keys := strings.Split(pair[0], ",")
	frames, _ := strconv.Atoi(pair[1])
	state.frames = uint16(frames)

	for _, v := range keys {
		if nums.MatchString(v) {
			direction, _ := strconv.Atoi(v)
			state.direction = uint8(direction)
		} else if v == "lp" {
			state.lp = true
		} else if v == "mp" {
			state.mp = true
		} else if v == "hp" {
			state.hp = true
		} else if v == "lk" {
			state.lk = true
		} else if v == "mk" {
			state.mk = true
		} else if v == "hk" {
			state.hk = true
		}
	}

	return &state
}

func TextToOperations(text *string) *[]processor.Operation {
	body := strings.Replace(*text, "\r\n", "\n", -1)
	body = strings.Replace(body, "\r", "\n", -1)
	lines := strings.Split(body, "\n")
	return LinesToOperations(&lines)
}

func LinesToOperations(lines *[]string) *[]processor.Operation {
	lastState := State{}

	operations := make([]processor.Operation, len(*lines))

	for i, v := range *lines {
		operations[i] = processor.Operation{}
		state := *lineToState(&v)

		if lastState.direction != state.direction {
			operations[i].Direction = state.direction
		}

		if lastState.lp != state.lp {
			if state.lp {
				operations[i].Lp = 1
			} else {
				operations[i].Lp = -1
			}
		}
		if lastState.mp != state.mp {
			if state.mp {
				operations[i].Mp = 1
			} else {
				operations[i].Mp = -1
			}
		}
		if lastState.hp != state.hp {
			if state.hp {
				operations[i].Hp = 1
			} else {
				operations[i].Hp = -1
			}
		}
		if lastState.lk != state.lk {
			if state.lk {
				operations[i].Lk = 1
			} else {
				operations[i].Lk = -1
			}
		}
		if lastState.mk != state.mk {
			if state.mk {
				operations[i].Mk = 1
			} else {
				operations[i].Mk = -1
			}
		}
		if lastState.hk != state.hk {
			if state.hk {
				operations[i].Hk = 1
			} else {
				operations[i].Hk = -1
			}
		}

		operations[i].Frames = state.frames

		lastState = state
	}

	return &operations
}
