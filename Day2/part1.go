package main

func CalculateCourse(commands []Command) int {
	depth := 0
	horizontal := 0

	for _, command := range commands {
		switch command.Direction {
		case Forward:
			horizontal += command.Distance
		case Down:
			depth += command.Distance
		case Up:
			depth -= command.Distance
		}
	}
	return depth * horizontal
}
