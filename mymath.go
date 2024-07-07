package mymath

import "math"

// Sqrt - функция для вычисления квадратного корня из числа
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Ceil - функция для округления числа вверх
func Ceil(x float64) float64 {
	return math.Ceil(x)
}

// Floor - функция для округления числа вниз
func Floor(x float64) float64 {
	return math.Floor(x)
}

// Pow - функция для возведения числа в степень
func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

// Max - функция для нахождения максимального из двух чисел
func Max(x, y float64) float64 {
	return math.Max(x, y)
}

// Min - функция для нахождения минимального из двух чисел
func Min(x, y float64) float64 {
	return math.Min(x, y)
}

func Abs(x float64) float64 {
	return math.Abs(x)
}

func Acos(x float64) float64 {
	return math.Acos(x)
}

func Acosh(x float64) float64 {
	return math.Acosh(x)
}

func Asin(x float64) float64 {
	return math.Asin(x)
}

func Asinh(x float64) float64 {
	return math.Asinh(x)
}

func Atan(x float64) float64 {
	return math.Atan(x)
}

func Atanh(x float64) float64 {
	return math.Atanh(x)
}

func Cbrt(x float64) float64 {
	return math.Cbrt(x)
}
