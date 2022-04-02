package hello

import (
	"fmt"
)

func Hello(name string, index int, num int) {
	fmt.Printf("这是%s的第%d/%d个go语言程序，在单独的hello函数打印", name, index, num)
}
