package main

func EstAutodesc2(n string) bool {
	var test [10]int
	var nombre = n
	for i := 0; i < len(nombre); i++ {
		val := int(nombre[i] - '0')
		test[val] += 1
	}

	for i := 0; i < len(n); i++ {
		val := int(n[i] - '0')
		if val != test[i] {
			return false
		}

	}
	return true
}
