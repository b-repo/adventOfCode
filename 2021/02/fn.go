package _2

import (
	"strconv"
	"strings"
)

func ReadCommandsFromData(data []string) []Command {
	commands := make([]Command, 0, (len(data)/2)+1)

	for i := 0; i < len(data); i++ {
		command := strings.Split(data[i], " ")
		power, _ := strconv.Atoi(command[1])
		commands = append(commands, Command{
			Direction: command[0],
			Power:     power,
		})
	}

	return commands
}

func CalculateFinalePosition(commands []Command) (int, int) {
	var position Position

	for i := 0; i < len(commands); i++ {
		switch commands[i].Direction {
		case "forward":
			position.Hzt += commands[i].Power
		case "down":
			position.Depth += commands[i].Power
		case "up":
			position.Depth -= commands[i].Power
		}
	}

	return position.Hzt, position.Depth
}

func CalculateFinalePositionWithAim(commands []Command) (int, int) {
	position := Position{
		Hzt:   0,
		Depth: 0,
		Aim:   0,
	}

	for i := 0; i < len(commands); i++ {
		switch commands[i].Direction {
		case "forward":
			position.Hzt += commands[i].Power
			position.Depth += position.Aim * commands[i].Power
		case "down":
			position.Aim += commands[i].Power
		case "up":
			position.Aim -= commands[i].Power
		}
	}

	return position.Hzt, position.Depth
}
