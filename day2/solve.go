package day2

import (
	"advent/utils"
	"strconv"
	"strings"
)

func Solve() int {
	lines, _ := utils.ReadLines("./day2/input.txt")

	var x, y, aim int

	for _, v := range lines {
		split := strings.Split(v, " ")
		switch split[0] {
		case "down":
			num, _ := strconv.Atoi((split[1]))
			aim += num
		case "up":
			num, _ := strconv.Atoi((split[1]))
			aim -= num
		case "forward":
			num, _ := strconv.Atoi((split[1]))
			x += num
			y += (aim * num)
		}
	}
	return (x * y)
}
