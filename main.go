package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = r.Int()
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	res := make([]int, CHUNKS)
	var wg sync.WaitGroup

	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}

		chunk := data[start:end]

		go func(chunk []int, index int) {
			defer wg.Done()

			res[index] = maximum(chunk)
		}(chunk, i)
	}
	wg.Wait()

	return maximum(res)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)

	data := generateRandomElements(SIZE)

	fmt.Println("\nИщем максимальное значение в один поток")

	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)

	start = time.Now()
	maxParallel := maxChunks(data)
	elapsed = time.Since(start).Milliseconds()

	fmt.Printf("\nМаксимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	if max == maxParallel {
		fmt.Println("Результаты совпадают")
	} else {
		fmt.Println("Результаты не совпадают")
	}
}
