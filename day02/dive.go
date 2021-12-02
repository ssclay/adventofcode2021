//https://adventofcode.com/2021/day/2
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var (
		i, depth, horizontal, aim int
		// commands = []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	)

	// Read in file as an array of strings
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		//Do nothing, but you have to use err.
	}
	commands := strings.Split(string(content), "\n")

	for i = 0; i < len(commands); i++ {
		// Use string contains for direction then grab the value
		// Cheat: Data set does not contain any double digit values so just grab the last character as the value.

		// Do weird things to seperate the int out from the string slice because 5 != 53, which it would if it wasn't converted from a byte to an int.
		command := commands[i]

		// Byte things: when using len(command)-1 by itself, returns 53, with the "unneeded index len(command) it returns 5"
		value := command[len(command)-1 : len(command)]
		delta, err := strconv.Atoi(value)

		if err != nil {
			//Do nothing, but you have to use err.
		}

		if strings.Contains(commands[i], "forward") == true {
			// If forward add to horizontal
			horizontal += delta
			depth += (aim * delta)
		}

		if strings.Contains(commands[i], "down") == true {
			// If down add to aim
			aim += delta
		}

		if strings.Contains(commands[i], "up") == true {
			// If up subtract from aim
			aim -= delta
		}

	}

	// Multiply final horizontal and depth together
	fmt.Println("horizontal:", horizontal)
	fmt.Println("depth:", depth)
	fmt.Println("aim:", aim)
	fmt.Println("Final Answer:", depth*horizontal)
}
