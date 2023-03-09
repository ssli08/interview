package main

import "fmt"

var queue []map[int]int

// put key, value to the  queue
func put(k, v int) {
	if cap(queue) >= 1001 {

	}
	for _, m := range queue {
		if _, ok := m[k]; ok {
			fmt.Println("exist")
			return
		}
	}
	q := append(queue, map[int]int{})
	fmt.Println(q)

	return
}
func get(k int) int {
	for _, m := range queue {
		if val, ok := m[k]; ok {
			return -1
		} else {
			return val
		}
	}
	return -1
}
func del(k int) int {
	for _, m := range queue {
		if val, ok := m[k]; ok {
			delete(m, k)
			return val
		} else {
			return -1
		}
	}
	return -1
}
func main() {

}
