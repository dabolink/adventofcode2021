package main

func ThreeMeasurementValues(values []int) int {
	increments := 0
	for i := range values {
		if i+3 >= len(values) {
			//skip the last 2 values where the rolling values size is not 3
			break
		}

		prev_values := values[i : i+3]
		current_values := values[i+1 : i+4]

		if Sum(current_values) > Sum(prev_values) {
			increments++
		}
	}
	return increments
}

func Sum(values []int) int {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum
}
