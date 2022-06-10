package main

import "fmt"

func main() {
	var arr = []int{5, 1, 33, 23, 1, -5, 1}

	fmt.Println(arr)
	result := proccess(arr)

	fmt.Println(result)
}

func proccess(array []int) []int {

	//fmt.Println(array)

	//fmt.Println(len(array))

	result := []int{}
	var countOne int

	for i := 0; i < len(array); i++ {
		//fmt.Println(array[i])
		if array[i] != 1 {
			result = append(result, array[i])
		} else {
			countOne++
		}

	}

	for j := 0; j < countOne; j++ {
		result = append(result, 1)
	}
	return result

}
