package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imoc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, _ := range m {
		fmt.Println(k)
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(courseName)

	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key dont exist")
	}

	fmt.Println("Delete values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

}
