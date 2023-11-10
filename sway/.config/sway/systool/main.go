package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func DoBrightness(operator byte, amount int) {
	file, err := os.OpenFile("/sys/class/backlight/intel_backlight/brightness", os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	
	data, err := io.ReadAll(file)
	currentStr := strings.TrimSpace(string(data))
	current, err := strconv.Atoi(currentStr)
	fmt.Println(err)
	file.Seek(0,0)
	if operator == '-' {
		amount *= -1
	}

	file.WriteString(strconv.Itoa(current + amount))
}

func main () {
	thing := os.Args[1] // this is like brighness or something
	param := os.Args[2]

	operator := param[0]
	amount, err  := strconv.Atoi(param[1:])
	if err != nil {
		panic(err)
	}
	switch thing {
	case "brightness":
		DoBrightness(operator, amount)
	}

}
