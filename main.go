package main

import (
	"fmt"
	"math/rand"
	"os"
	"slices"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 || size > 100_000_000 {
		fmt.Fprintf(os.Stderr, "Некорректная длина слайса: %d. Требуется число от 1 до 100_000_000\n", size)
		return []int{}
	}
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(size)
	}
	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		fmt.Fprintf(os.Stderr, "Максимальное число не определено, получен пустой слайс\n")
		return 0
	}
	return slices.Max(data)
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var maxNum []int
	var wg sync.WaitGroup
	if len(data)/CHUNKS <= 0 {
		fmt.Fprintf(os.Stderr, "Слишком короткий слайс, максимально возможное кол-во частей %d\n", len(data))
		return 0
	}
	if len(data)%CHUNKS == 0 {
		maxNum = make([]int, CHUNKS)
		counter := 0
		for chunk := range slices.Chunk(data, len(data)/CHUNKS) {
			wg.Add(1)
			go func(i int, chunk []int) {
				maxNum[i] = slices.Max(chunk)
				wg.Done()
			}(counter, chunk)
			counter++
		}
	} else {
		var length int
		counter := 0
		maxNum = make([]int, CHUNKS)
		for counter < CHUNKS-1 {
			wg.Add(1)
			step := length + len(data)/CHUNKS
			chunk := data[length:step]
			go func(i int, chunk []int) {
				maxNum[i] = slices.Max(chunk)
				wg.Done()
			}(counter, chunk)
			counter++
			length = step
		}
		wg.Add(1)
		chunk := data[length:]
		go func(i int, chunk []int) {
			maxNum[i] = slices.Max(chunk)
			wg.Done()
		}(counter, chunk)
	}
	wg.Wait()
	return slices.Max(maxNum)
}
func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	slice := generateRandomElements(SIZE)
	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(slice)
	elapsed := time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(slice)
	elapsed = time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
