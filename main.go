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
		data[i] = r.Intn(100000000)
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
	var mu sync.Mutex

	rem := len(data) % CHUNKS

	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		go func(chunkInd int) {
			defer wg.Done()

			start := chunkInd * chunkSize
			end := start + chunkSize

			if chunkInd == CHUNKS-1 && rem > 0 {
				end += rem
			}

			chunk := data[start:end]
			if len(chunk) == 0 {
				return
			}

			chunkMax := chunk[0]
			for j := 1; j < len(chunk); j++ {
				if chunk[j] > chunkMax {
					chunkMax = chunk[j]
				}
			}

			mu.Lock()
			res[chunkInd] = chunkMax
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	return maximum(res)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)

	data := generateRandomElements(SIZE)

	if len(data) == 0 {
		fmt.Println("Ошибка: массив пустой")
		return
	}

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
