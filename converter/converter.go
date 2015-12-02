package converter

import (
	"../processor"
	"strings"
	"regexp"
	"strconv"
	"errors"
	"fmt"
)

var (
	blanks = regexp.MustCompile(`\s+`)
	nums = regexp.MustCompile(`^\d+$`)
	loopStart = regexp.MustCompile(`<loop\s(\d+)\s*>`)
)

func TextToOperations(text *string) *[]processor.Operation {
	body := strings.Replace(*text, "\r\n", "\n", -1)
	body = strings.Replace(body, "\r", "\n", -1)
	lines := strings.Split(body, "\n")
	return LinesToOperations(&lines)
}

func LinesToOperations(lines *[]string) *[]processor.Operation {
	lastState := State{}
	operations := make([]processor.Operation, 0, len(*lines))
	buffer := make([]processor.Operation, 0, len(*lines))
	var count int

	for _, v := range *lines {
		if loopStart.MatchString(v) {
			count, _ = strconv.Atoi(loopStart.FindStringSubmatch(v)[1])
			operations = append(operations, buffer...)
			buffer = make([]processor.Operation, 0, len(*lines))
			continue
		} else if strings.Index(v, "</loop>") == 0 {
			for i := 0; i < count; i++ {
				operations = append(operations, buffer...)
			}
			buffer = make([]processor.Operation, 0, len(*lines))
			continue
		}

		state, err := lineToState(&v)
		if err != nil {
			fmt.Println(err)
			continue
		}

		operation := processor.Operation{}
		if lastState.direction != state.direction {
			operation.Direction = state.direction
		}

		if lastState.lp != state.lp {
			if state.lp {
				operation.Lp = 1
			} else {
				operation.Lp = -1
			}
		}
		if lastState.mp != state.mp {
			if state.mp {
				operation.Mp = 1
			} else {
				operation.Mp = -1
			}
		}
		if lastState.hp != state.hp {
			if state.hp {
				operation.Hp = 1
			} else {
				operation.Hp = -1
			}
		}
		if lastState.lk != state.lk {
			if state.lk {
				operation.Lk = 1
			} else {
				operation.Lk = -1
			}
		}
		if lastState.mk != state.mk {
			if state.mk {
				operation.Mk = 1
			} else {
				operation.Mk = -1
			}
		}
		if lastState.hk != state.hk {
			if state.hk {
				operation.Hk = 1
			} else {
				operation.Hk = -1
			}
		}
		if lastState.pause != state.pause {
			if state.pause {
				operation.Pause = 1
			} else {
				operation.Pause = -1
			}
		}
		if lastState.save != state.save {
			if state.save {
				operation.Save = 1
			} else {
				operation.Save = -1
			}
		}
		if lastState.reload != state.reload {
			if state.reload {
				operation.Reload = 1
			} else {
				operation.Reload = -1
			}
		}
		operation.Frames = state.frames

		buffer = append(buffer, operation)
		lastState = *state
	}

	operations = append(operations, buffer...)

	return &operations
}

func lineToState(line *string) (*State, error) {
	if line == nil || len(*line) == 0 {
		return nil, errors.New("line is nil or empty")
	}

	pair := blanks.Split(*line, 2)
	if len(pair) != 2 {
		return nil, errors.New("line does not contain any blanks")
	}

	keys := strings.Split(pair[0], ",")
	frames, err := strconv.Atoi(pair[1])
	if err != nil {
		return nil, errors.New("second value must be integers")
	}

	state := State{}
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
		} else if v == "pause" {
			state.pause = true
		} else if v == "save" {
			state.save = true
		} else if v == "reload" {
			state.reload = true
		}
	}

	return &state, nil
}
