package main

import (
	"errors"
	"fmt"
)

func functions() {
	fmt.Println("getFullName:", getFullName("Phm Duc", "Huy"))

	valid, msg := validateAge(18)
	fmt.Println("validateAge:", valid, msg)

	id, name, email, err := fetchUser(1)
	if err != nil {
		fmt.Println("fetchUser error:", err)
	} else {
		fmt.Println("fetchUser result:", id, name, email, err)
	}

	fmt.Println("processData (valid):", processData(""))
	fmt.Println("processData (invalid):", processData("some data"))

	fmt.Println("findMax:", findMax())

	sum, avg, err := calculateStats([]int{1, 2, 3, 4, 5})
	if err != nil {
		fmt.Println("calculateStats error:", err)
	} else {
		fmt.Println("calculateStats result:", sum, avg)
	}
}

/*
Q1: Basic Function

	function getFullName(firstName, lastName) {
		return firstName + " " + lastName;
	}
*/
func getFullName(firstName, lastName string) string {
	return firstName + " " + lastName
}

/*
Q2: Multiple Return Values
*/
func validateAge(age int) (bool, string) {
	if age < 0 {
		return false, "invalid age"
	} else if age < 18 {
		return false, "too young"
	} else {
		return true, "valid"
	}
}

/*
Q3: Error Handling
*/
func fetchUser(idInput int) (id int, name string, email string, err error) {
	if idInput <= 0 {
		err = errors.New("invalid user ID")
	}

	id = idInput
	name = "User " + fmt.Sprint(id)
	email = "user" + fmt.Sprint(idInput) + "@example.com"

	return
}

/*
Q4: Fix the Error Handling
*/
func processData(data string) string {
	result, err := parseData(data)

	if err != nil {
		return "Error: " + err.Error()
	}

	return result
}

func parseData(data string) (string, error) {
	if data == "" {
		return "", errors.New("empty data")
	}
	return "processed: " + data, nil
}

/*
Q5: Variadic Function
*/
func findMax(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

/*
Q6: Named Returns
*/
func calculateStats(numbers []int) (sum int, average float64, err error) {
	if len(numbers) == 0 {
		return 0, 0, errors.New("empty slice")
	}

	for _, num := range numbers {
		sum += num
	}

	average = float64(sum) / float64(len(numbers))
	return sum, average, nil
}
