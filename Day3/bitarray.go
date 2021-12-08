package main

type BitArray []bool

func NewBitArray(str string) BitArray {
	bitArray := make(BitArray, 0, len(str))
	for _, b := range str {
		switch b {
		case '0':
			bitArray = append(bitArray, false)
		case '1':
			bitArray = append(bitArray, true)
		}
	}
	return bitArray
}

func (ba BitArray) ToDecimal() int {
	sum := 0
	for i := range ba {
		current := ba[len(ba)-1-i]
		if current {
			sum += IntPow(2, i)
		}
	}
	return sum
}

func (ba BitArray) ToBinary() string {
	str := ""
	for _, b := range ba {
		if b {
			str += "1"
		} else {
			str += "0"
		}
	}
	return str
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
