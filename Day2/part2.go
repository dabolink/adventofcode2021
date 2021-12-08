package main

func CalculateCourseWithAim(commands []Command) int {
	depth := 0
	horizontal := 0
	aim := 0

	for _, command := range commands {
		switch command.Direction {
		case Forward:
			horizontal += command.Distance
			depth += aim * command.Distance
		case Down:
			aim += command.Distance
		case Up:
			aim -= command.Distance
		}
	}
	return depth * horizontal
}
