package main

func GetMostCommon(input []BitArray) BitArray {
	if len(input) == 0 {
		return BitArray{}
	}
	return GetByteArray(input, MostCommonSelection, true)
}

func MostCommonSelection(zeroes int, ones int) bool {
	if zeroes == ones {
		return true
	} else if zeroes > ones {
		return false
	} else {
		return true
	}
}

func Selection(input Input, index int, selectionFunc func(zeroes int, ones int) bool) bool {
	zeroes := 0
	ones := 0
	for _, ba := range input {
		switch ba[index] {
		case false:
			zeroes++
		case true:
			ones++
		}
	}
	return selectionFunc(zeroes, ones)
}

func LeastCommonSelection(zeroes int, ones int) bool {
	if zeroes == ones {
		return false
	}
	if zeroes > ones {
		return true
	} else {
		return false
	}
}

func GetLeastCommon(input []BitArray) BitArray {
	return GetByteArray(input, LeastCommonSelection, false)
}

func GetByteArray(input []BitArray, selectionFunc func(zeroes int, ones int) bool, prefer bool) BitArray {
	output := make(BitArray, len(input[0]))

	for i := range output {
		output[i] = Selection(input, i, selectionFunc)
	}
	return output
}
