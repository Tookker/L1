package main

import "fmt"

// 12.Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}

	mapS := make(map[string]struct{})
	for _, str := range strs {
		mapS[str] = struct{}{}
	}

	fmt.Println("Result = ", mapS)
}
