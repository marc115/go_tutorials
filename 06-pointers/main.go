package main

import "fmt"

func main() {
	var p *int32 = new(int32) //points to the memory address of a 32 bit integer
	var i int32
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\n The value of i is: %v", i)
	*p = 10 // this changes the value p is pointing to
	p = &i  // now p points to the memory address of i
	*p = 1  //therefore, if I change the value p is pointing to, i's value is also changed

	var slice = []int32{1, 2, 3} //internally, slices point to an array in memory
	var slice2 = make([]int32, len(slice))
	copy(slice2, slice) //go has a method to copy one slice to the other
	slice2[2] = 4       // so now we can modify the second slice without affecting the original
	fmt.Println(slice, slice2)
}
