package day2

import (
	"advent/utils"
	"strings"
)

func Solve() int {
	lines, _ := utils.ReadLines("./day2/input.txt")

	var x, y, aim int

	for _, v := range lines {
		split := strings.Split(v, " ")
		switch split[0] {
		case "down":
			aim += utils.Atoi((split[1]))
		case "up":
			aim -= utils.Atoi((split[1]))
		case "forward":
			num := utils.Atoi((split[1]))
			x += num
			y += (aim * num)
		}
	}
	return (x * y)
}
