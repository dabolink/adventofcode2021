package main

func (input Input) Filter(index int, filter func(index int, ba BitArray) bool) Input {
	if len(input) == 0 {
		return input
	}

	filtered := make(Input, 0)

	for _, ba := range input {
		if filter(index, ba) {
			filtered = append(filtered, ba)
		}
	}
	return filtered
}
