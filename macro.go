package main

import (
	"./converter"
	"./processor"
	"./vjoy"
	"./file"
	"./http"
	"flag"
)

func main() {
	vjoy.Acuire(1)
	defer vjoy.Close(1)

	flag.Parse()

	if flag.NArg() == 0 {
		textMode("command.txt")
	} else if flag.Arg(0) == "server" {
		serverMode()
	} else {
		textMode(flag.Arg(0))
	}
}

func textMode(path string) {
	lines := file.ReadAll(path)
	operations := converter.LinesToOperations(lines)
	processor.Process(operations)
}

func serverMode() {
	http.Server(serverFunction)
}

func serverFunction(body *string) {
	operations := converter.TextToOperations(body)
	processor.Process(operations)
}
