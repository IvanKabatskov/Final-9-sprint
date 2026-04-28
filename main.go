package main

import (
	"fmt"
	"math/rand"
	"os"
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
		fmt.Fprintf(os.Stderr, "Недопустимый размер слайса: %d\n", size)
		return nil
	}
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Int()
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		fmt.Fprintf(os.Stderr, "Максимальное число не определено, получен пустой слайс\n")
		return 0
	}
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var wg sync.WaitGroup
	if len(data)/CHUNKS <= 0 {
		fmt.Fprintf(os.Stderr, "Слишком короткий слайс, максимально возможное кол-во частей %d\n", len(data))
		return maximum(data)
	}
	step := len(data) / CHUNKS
	maxNums := make([]int, CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		start := i * step
		end := start + step
		if i == CHUNKS-1 {
			end = len(data)
		}
		chunk := data[start:end]
		go func(i int, chunk []int) {
			maxNums[i] = maximum(chunk)
			wg.Done()
		}(i, chunk)
	}
	wg.Wait()
	return maximum(maxNums)
}
func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)
	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
