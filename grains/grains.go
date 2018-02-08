package grains

import "errors"

func Square(square int) (uint64, error) {
	if square == 1 {
		return 1, nil
	} else if square > 1 && square <= 64 {
		x := uint64(1) << (uint64(square) - 1)
		return x, nil
	} else {
		return 0, errors.New("Error")
	}
}

func Total() (total uint64) {
	for i := 1; i <= 64; i++ {
		v, _ := Square(i)
		total += v
	}
	return
}
