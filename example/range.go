package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum is", sum)

	for i, num := range nums {
		fmt.Println("index:", i, "value:", num)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for _, char := range "lunweiwei" {
		fmt.Println(string(char))
	}
}
