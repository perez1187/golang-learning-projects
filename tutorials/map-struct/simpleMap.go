package main

import "fmt"

// declaring a struct
type Address struct {
	Name    string
	city    string
	Pincode int
}

func SimpleMap() {
	// Creating struct instances
	a1 := Address{"Pam", "Mumbai", 2200}
	a2 := Address{Name: "Ram", city: "Delhi", Pincode: 2400}
	a3 := Address{Name: "Sam", city: "Lucknow", Pincode: 1070}

	// Declaring and initialising using map literals
	sample := map[Address]int{a1: 1, a2: 1, a3: 1}
	for str, val := range sample {
		fmt.Println(str, val)
	}

	sample[Address{Name: "SRam", city: "Delhjhi", Pincode: 2400}] = 4

	// You can also access a struct
	// field while using a loop
	for str := range sample {
		fmt.Println(str.Name)
	}

	delete(sample, Address{Name: "Ram", city: "Delhi", Pincode: 2400})
	fmt.Println("second")

	for str := range sample {
		fmt.Println(str.Name)
	}

	// checking if key exist in map
	value, check := sample[Address{"Pama", "Mumbai", 2200}]
	fmt.Println("Is the key present:", check)
	fmt.Println("Value of the key:", value)

	_, check2 := sample[Address{"Pam", "Mumbai", 2200}]
	fmt.Println("Is the key present:", check2)

	if check2 == true {
		fmt.Println("yeees")
	} else {
		fmt.Println("nieee")
	}

}
