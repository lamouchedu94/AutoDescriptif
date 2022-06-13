package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sort"
	"strconv"
	"sync"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

func main() {
	var err error
	nb := 1210
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

	th := runtime.NumCPU()
	lockval := sync.Mutex{}
	val := []int{}
	inter, dernier := interval(nb, th)
	var wg sync.WaitGroup
	deb := time.Now()
	for i := 0; i < th; i++ {
		wg.Add(1)
		go func(i int) {
			r := segment(i*inter, inter)
			lockval.Lock()
			val = append(val, r...)
			lockval.Unlock()
			wg.Done()
		}(i)

	}
	wg.Add(1)
	go func(i int) {
		r := segment(i*inter, dernier+1)
		lockval.Lock()
		val = append(val, r...)
		lockval.Unlock()
		wg.Done()
	}(th)
	wg.Wait()
	fin := time.Now()
	sort.Ints(val)
	//fmt.Println(inter, "inter ", mult, "mult ")
	fmt.Println(val)
	fmt.Println(fin.Sub(deb))

}

func segment(debut int, longueur int) []int {
	resultat := []int{}
	//fmt.Println(n)
	n := ItoNombre(debut)
	j := longueur
	for j > 0 {
		if n.estAutodescriptif() {
			resultat = append(resultat, n.int())
		}
		n.inc()
		j--
	}
	return resultat

}

func interval(nb int, th int) (int, int) {
	res := nb / th
	diff := nb - (res * th)
	return res, diff

}
