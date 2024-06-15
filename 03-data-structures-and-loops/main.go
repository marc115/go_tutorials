package main

import "fmt"

func main() {
	var intArr [3]int32 = [3]int32{1, 2, 3} // array
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	var intSlice []int32 = []int32{4, 5, 6} // slice (like an array but more functional)
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...) // go has the spread operator. go is therefore superior
	fmt.Println(intSlice)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Eva"])   //If the key doesn't exist it will return 0
	var age, ok = myMap2["Adam"] // You can also get another value to see if the key was found
	if ok {
		fmt.Println("Adam's age is", age)
	}

	delete(myMap2, "Adam") //deletes a key-value pair
	myMap2["Eva"] = 20

	for name := range myMap2 {
		fmt.Println("Name:", name) //when iterating through maps, no order is preserved
	}

	var i int = 0
	for i < 10 { //similar functionality to that of a while loop
		fmt.Println(i)
		i += 1
	}

	for i := 0; i < 10; i++ { //a more traditional for loop
		fmt.Println(i)
	}
}
