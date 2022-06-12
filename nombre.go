package main

import "strconv"

type nombre [10]uint8

func ItoNombre(i int) nombre {
	var n nombre
	s := strconv.Itoa(i)
	for ix := 0; ix < len(s); ix++ {
		n[10-len(s)+ix] = s[ix] - '0'
	}
	return n
}

func (n nombre) int() int {
	i := 0
	for j := 0; j < len(n); j++ {
		i = i*10 + int(n[j])
	}
	return i
}

func (n *nombre) inc() {
	i := len(n) - 1
	for {
		n[i]++
		if n[i] < 10 {
			return
		}
		n[i] = 0
		i--
	}
}

func (n nombre) estAutodescriptif() bool {
	i := 0
	for i < len(n) {
		if n[i] != 0 {
			break
		}
		i++
	}
	if i == len(n) {
		return false
	}
	premierNonZero := i
	j := 0
	var compteurs nombre
	for i < len(n) {
		compteurs[j+int(n[i])]++
		i++
	}

	j = 0
	for i = premierNonZero; i < len(n); i++ {
		if n[i] != compteurs[j] {
			return false
		}
		j++
	}

	return true
}
