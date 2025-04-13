package main

import "fmt"

func main() {
	intArr := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", intArr)
	intSlice := []int{4, 5, 6}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	MyMap := make(map[string]int)
	fmt.Println(MyMap)
	myMap := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap)
	myMap1 := map[string]int{"BHadri": 21, "Barath": 25}
	fmt.Println(myMap1["BHadri"])
	delete(myMap1, "BHadri")
	fmt.Println(myMap1)
	mymap2 := map[string]int{"Bha": 21, "Bara": 25, "Mad": 51, "Ana": 54}
	fmt.Println(mymap2)
	for name, age := range mymap2 {
		fmt.Printf("Name: %v, Age: %v", name, age)

	}

}
