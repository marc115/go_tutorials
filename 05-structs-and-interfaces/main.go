package main

import "fmt"

type GasEngine struct { // structs are a way of defining your own types
	milesPerGalon uint8
	gallons       uint8
	// owner         Owner
}

type ElectricEngine struct {
	milesPerKiloWattHour uint8
	KiloWattHours        uint8
}

type Engine interface { //kinda like defining a parent class
	milesLeft() uint8
}

func (engine GasEngine) milesLeft() uint8 { //this function syntax is different because this method is tied to the structure
	return engine.gallons * engine.milesPerGalon
}

func canMakeIt(engine Engine, miles uint8) { // on the other hand, this function isn't tied to the structure
	if miles <= engine.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

// type Owner struct {
// 	name string
// }

func main() {
	var myEngine GasEngine = GasEngine{milesPerGalon: 15, gallons: 15}
	var myEngine2 = struct { // this is an anonymous structure - i.e. creating structures on the go (pun intended)
		mpg     uint8
		gallons uint8
	}{25, 15}
	fmt.Println(myEngine)
	fmt.Println(myEngine2)
	fmt.Println("Total miles left in tank:", myEngine.milesLeft())
}
