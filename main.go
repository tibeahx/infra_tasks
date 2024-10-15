package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

//можно использовать https://go.dev/play/p/w0s82mkTK_s

func main() {
	td := testData{}
	wg := &sync.WaitGroup{}

	generate(wg, 100, &td)

	wg.Wait() //ждем пока остальные горутины завершат работу, иначе мы получим либо дедлок либо просто ничего
	fmt.Println(len(td.phones))
}

type testData struct {
	phones []int
	mu     sync.Mutex
}

func (td *testData) add(wg *sync.WaitGroup) {
	if td == nil {
		log.Println("got nil td, exiting...")
		wg.Done() // если td nil, просто выходим из функции и сигнализируем что горутина отработала
		return
	}

	td.mu.Lock()         //лочим мьютекс перед аппендом так как аппенд является записывающей операцией
	defer td.mu.Unlock() //деферим анлок мьютекса чтобы гарантировано перед ретерном разлочится

	td.phones = append(td.phones, randPhone())
	wg.Done() //функция отработала, значит сигнализируем что горутина отработала
}

func generate(wg *sync.WaitGroup, n int, td *testData) *testData {
	if wg == nil {
		log.Println("got nil wg, exiting...")
	}
	if td == nil {
		log.Println("got nil td, exiting...")
	}
	if n <= 0 { //добавил проверку, чтобы просто так не вызывать функцию, так как все равно нет смысла генерить с нулевым n
		return nil
	}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go td.add(wg)
	}
	return td
}

func randPhone() int {
	return 89000000000 + rand.Intn(999999999-100000000) + 100000000
}
