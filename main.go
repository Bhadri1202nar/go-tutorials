package main

import "fmt"

func main() {
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of thing1 is: %p", &thing1)
	square(&thing1)

	fmt.Printf("\nThe  updated memory location of thing1 is: %p", &thing1)
	fmt.Printf("\nThe value of thing1 is: %v", thing1)

	var p *int32 = new(int32)
	var i int32
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	slice[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)
	fmt.Printf("The value p points is: %v", *p)
	fmt.Printf("\nThe value if i is: %v", i)
	p = &i
	*p = 1
	fmt.Printf("The value p points is: %v", *p)
	fmt.Printf("\nThe value if i is: %v", i)

}
func square(thing2 *[5]float64) {
	fmt.Printf("The value thing2 points to: %v\n", *thing2)
	fmt.Printf("The address thing2 points to (same as &thing1): %p\n", thing2)
	fmt.Printf("The address of thing2 variable itself: %p\n", &thing2)
	for i := range thing2 {
		thing2[i] *= thing2[i]

	}

}
