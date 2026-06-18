package main

import (
	"github.com/fatih/color"
	"hexlet-go/greeting"
	"fmt"
)

func main() {
	fmt.Println(greeting.Hello())
	color.Blue("Prints %s in blue.", "text")
}
