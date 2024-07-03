package mathutil

func Min[T int8 | int16 | int32 | int | int64 | float32 | float64](a, b T) T {
	if a < b {
		return a
	}

	return b
}

func Max[T int8 | int16 | int32 | int | int64 | float32 | float64](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func Div[T int8 | int16 | int32 | int | int64 | float32 | float64](a, b T) T {
	if b == 0 {
		return 0
	}

	return a / b
}

func Mul[T int8 | int16 | int32 | int | int64 | float32 | float64](a, b T) T {
	return a * b
}

func Abs[T int8 | int16 | int32 | int | int64 | float32 | float64](a T) T {
	if a < 0 {
		return -a
	}

	return a
}
