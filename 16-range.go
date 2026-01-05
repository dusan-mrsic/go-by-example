package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5}
	for i, v := range s {
		fmt.Println("index: ", i, "value: ", v)
	}

	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println("total sum: ", sum)

	m := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range m {
		fmt.Println(k, "->", v)
	}

	for k := range m {
		fmt.Println("key: ", k)
	}

	for i, v := range "go" {
		fmt.Println("index: ", i, "value: ", v)
	}
}
