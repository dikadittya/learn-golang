package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func TaskA(wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(5 * time.Second)
	fmt.Println("Start Process Task A")
}

func TaskB(wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(3 * time.Second)
	fmt.Println("Start Process Task B")
}

func TaskC(wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(2 * time.Second)
	fmt.Println("Start Process Task C")
}

func main() {
	runtime.GOMAXPROCS(3) // mengacu pada jumlah core CPU ex:3
	// runtime.GOMAXPROCS(runtime.NumCPU()) // pakai semua CPU yg ada

	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(3)
	go TaskA(&wg)
	go TaskB(&wg)
	go TaskC(&wg)

	wg.Wait()
	duration := time.Since(start)
	fmt.Println("Exec time: ", duration.Milliseconds())
}
