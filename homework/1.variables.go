package main

import (
	"fmt"
	"math"
	"strconv"
)

func variables() {
	// Q1: Variable Declaration
	// Method 1 (var with type):
	var name1 string = "Alice"
	var age1 int = 30
	var isAdmin1 bool = false
	var balance1 float32 = 100.50

	// Method 2 (var with inference):
	var name2 = "Alice"
	var age2 = 30
	var isAdmin2 = false
	var balance2 = 100.50

	// Method 3 (short declaration):
	name3, age3, isAdmin3, balance3 := "Alice", 30, false, 100.50

	fmt.Println(name1, age1, isAdmin1, balance1)
	fmt.Println(name2, age2, isAdmin2, balance2)
	fmt.Println(name3, age3, isAdmin3, balance3)

	// Q2: Fix the Code
	/*
		let username = "developer";
		const maxRetries = 3;
		let attempts = 0;
		let isConnected = false;
	*/
	username := "developer"
	const maxRetries = 3
	attempts := 0
	isConnected := false

	fmt.Println(username, maxRetries, attempts, isConnected)

	// Q3: Arrays vs Slices
	numbers := [5]int{1, 2, 3, 4, 5}
	fruitSlice := []string{"apple", "banana", "orange"}

	fmt.Println(numbers, fruitSlice)

	/*
		// Q4: Zero Values
		What are the zero values for these types? Fill in:
		int: ____
		string: ____
		bool: ____
		[]int: ____
	*/
	var zeroInt int
	var zeroString string
	var zeroBool bool
	var zeroSlice []int

	fmt.Println(zeroInt, zeroString, zeroBool, zeroSlice)

	/*
		Q5: Type Conversion
		const price: number = 19.99;
		const priceString: string = price.toString();
		const roundedPrice: number = Math.floor(price);
	*/
	priceNumber := 19.99
	priceString := strconv.FormatFloat(priceNumber, 'f', 2, 64)
	roundedPrice := math.Floor(priceNumber) // Uncomment if needed

	fmt.Println(priceNumber, priceString, roundedPrice)
}
