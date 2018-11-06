package main

type Mode func(size int) []int

func Easy(size int) []int {
	a := make([]int, size)

	for i := 0; i < size; i++ {
		a[i] = i + 1
	}

	return a
}

func Medium(size int) []int {
	a := make([]int, size)

	for i := 0; i < size; i++ {
		a[i] = i
	}

	return a
}

func Hard(size int) []int {
	a := make([]int, 2*size + 1)

	for i := 0; i < size; i++ {
		a[i] = i + 1
	}

	for i := size; i < 2*size+1; i++ {
		a[i] = 2*size+1 - i
	}

	return a
}

func ExtraHard(size int) []int {
	a := make([]int, 2*size + 1)

	for i := 0; i < size; i++ {
		a[i] = i
	}

	for i := size; i < 2*size+1; i++ {
		a[i] = 2*size - i
	}

	return a
}