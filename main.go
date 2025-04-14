package main

import "fmt"

type gasEngine struct {
	mpg     uint32
	gallons uint32
}
type electricEngine struct {
	mpkwh uint32
	kwh   uint32
}

func (e gasEngine) milesLeft() uint32 {
	return e.gallons * e.mpg
}
func (e electricEngine) milesLeft() uint32 {
	return e.mpkwh * e.kwh
}

type engine interface {
	milesLeft() uint32
}

func canReach(e engine, miles uint32) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")

	} else {
		fmt.Println("Need to fuelup first")
	}

}

/*type owner struct{
	name string

}*/
func main() {
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15}
	fmt.Println(myEngine.milesLeft())
	var myE_Engine electricEngine = electricEngine{mpkwh: 15, kwh: 10}
	fmt.Println(myE_Engine.milesLeft())
	fmt.Println(myEngine.mpg, myEngine.gallons)

}
