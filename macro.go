package main

import (
	"./converter"
	"./processor"
	"./vjoy"
	"./file"
	"./http"
	"flag"
	"log"
)

var (
	vjoy1Enabled bool
	vjoy2Enabled bool
)

func main() {
	err := vjoy.Acuire(1)
	if err != nil {
		println("vJoy[1] is not enabled.")
	} else {
		vjoy1Enabled = true
		println("vJoy[1] is acquired.")
	}
	defer vjoy.Close(1)

	err = vjoy.Acuire(2)
	if err != nil {
		println("vJoy[2] is not enabled.")
	} else {
		vjoy2Enabled = true
		println("vJoy[2] is acquired.")
	}
	defer vjoy.Close(2)

	if !vjoy1Enabled && !vjoy2Enabled {
		log.Fatal("No vJoy is enabled.")
	}

	if flag.NArg() == 0 {
		serverMode()
	} else {
		textMode(flag.Arg(0))
	}
}

func textMode(path string) {
	lines := file.ReadAll(path)
	operations := converter.LinesToOperations(lines)

	if vjoy1Enabled {
		processor.ProcessSingle(1, operations)
	} else {
		processor.ProcessSingle(2, operations)
	}
}

func serverMode() {
	http.Server(serverFunction)
}

func serverFunction(player1 *string, player2 *string) string {
	if !vjoy1Enabled && len(*player1) > 0 {
		return "Commands for player1 are sent but vjoy1 is disabled."
	}
	if !vjoy2Enabled && len(*player2) > 0 {
		return "Commands for player2 are sent but vjoy2 is disabled."
	}

	if len(*player1) > 0 && len(*player2) > 0 {
		log.Println("Run player1 & player2")
		operations1 := converter.TextToOperations(player1)
		operations2 := converter.TextToOperations(player2)
		processor.ProcessPair(1, operations1, 2, operations2)
	} else if len(*player1) > 0 {
		log.Println("Run player1 only")
		operations := converter.TextToOperations(player1)
		processor.ProcessSingle(1, operations)
	} else if len (*player2) > 0 {
		log.Println("Run player2 only")
		operations := converter.TextToOperations(player2)
		processor.ProcessSingle(2, operations)
	}

	return ""
}
