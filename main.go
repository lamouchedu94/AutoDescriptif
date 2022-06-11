package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"
)

func main() {
	th := 1
	nb := 3000000
	lockval := sync.Mutex{}

	val := []int{}
	inter := interval(nb, th)
	var wg sync.WaitGroup
	deb := time.Now()
	for i := 0; i < th; i++ {
		wg.Add(1)
		go func(i int) {
			r := segment(i, inter)
			lockval.Lock()
			val = append(val, r...)
			lockval.Unlock()
			wg.Done()
		}(i)

	}
	wg.Wait()
	fin := time.Now()
	sort.Ints(val)
	//fmt.Println(inter, "inter ", mult, "mult ")
	fmt.Println(val)
	fmt.Println(fin.Sub(deb))
}

func segment(i int, inter int) []int {
	a := []int{}
	for j := inter * i; j < inter*(i+1); j++ {
		if EstAutodesc2(strconv.Itoa(j)) {
			a = append(a, j)
		}

	}
	return a

}

func interval(nb int, th int) int {
	return nb / th
}
