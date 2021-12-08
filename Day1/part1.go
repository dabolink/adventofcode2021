package main

func TwoMeasurmentValues(values []int) int {
	increments := 0
	for i, current := range values {
		if i == 0 {
			//don't check the first one as we dont have the previous
			continue
		}
		prev := values[i-1]
		if current > prev {
			increments++
		}
	}
	return increments
}
