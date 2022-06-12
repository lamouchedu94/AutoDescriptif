package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sort"
	"strconv"
	"sync"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

func main() {
	var err error
	nb := 3000000
	flag.Parse()
	if flag.NArg() > 0 {
		nb, err = strconv.Atoi(flag.Arg(0))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	th := 1

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
