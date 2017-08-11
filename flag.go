package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	mode = flag.String("mode", "", "模式")
	port = flag.Int("port", 0, "端口")
	file = flag.String("file", "", "文件")
)

func main() {
	flag.Parse()
	fmt.Println(*mode, *port, *file)

	switch *mode {
	case "server":
		runServer(*port)
	case "client":
		runClient(*port, *file)
	default:
		log.Fatalf("未解析的模式: %s ", *mode)
	}
}
