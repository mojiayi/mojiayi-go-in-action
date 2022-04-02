package main

import (
	"flag"
	"fmt"
	"mojiayi-grpc-in-action/hello"
)

var name string
var num int
var indexArgs = flag.Int("index", 0, "序号")
var container = []string{"zero", "one", "two"}

func init() {
	flag.StringVar(&name, "name", "nouser", "问候人名称")
	flag.IntVar(&num, "num", 0, "数量")
	flag.Parse()
}

func main() {
	hello.Hello(name, *indexArgs, num)
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	fmt.Printf("The element is %q.\n", container[1])
	if num > 5 {

	}
}

func compare(a int, b int) string {
	if a > b {
		return "a>b"
	} else {
		return "a<=b"
	}
}

func example(x int) int {
	if x == 0 {
		return 5
	} else {
		return x
	}
}
