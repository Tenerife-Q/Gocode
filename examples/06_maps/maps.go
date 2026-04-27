package main

import "fmt"

func main() {
	m := map[string]int{"alice": 30, "bob": 25}
	fmt.Println("map:", m)

	m["carol"] = 27

	if v, ok := m["dave"]; ok {
		fmt.Println("found dave", v)
	} else {
		fmt.Println("dave not found")
	}

	delete(m, "bob")
	fmt.Println("after delete:", m)
}
