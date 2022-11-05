package main

import (
	"fmt"

	"github.com/kush-daga/readme-generator/cmd"
)

func main() {
	responses := cmd.GetReadmeComponents()
	fmt.Printf("%s", responses)
}
