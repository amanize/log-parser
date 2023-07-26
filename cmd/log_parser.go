package main

import (
	"log_parser/internal"
	"log_parser/logger"
)

func main() {
	logger.NewLogrusLogger()
	internal.Start()
}
