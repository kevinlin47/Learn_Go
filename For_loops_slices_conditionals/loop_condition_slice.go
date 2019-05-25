package main

import "fmt"

func main() {
	sum := 0

	//Standard for loop
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum) //Summation from 0 to 9

	//Go's while loop
	sum = 0
	i := 0

	for i < 10 {
		sum += i
		i++
	}

	fmt.Println(sum) //Summation from 0 to 9

	digitsToSum := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sum = 0

	//For each loop
	for i, d := range digitsToSum {
		sum += d
		fmt.Println(i)
	}

	fmt.Println(sum)

	/********** Decision Making ************/
	myAge := 24
	hisAge := 19

	checkDrinkingAge(myAge)
	checkDrinkingAge(hisAge)

	age1 := 16
	age2 := 17
	age3 := 18
	age4 := 15

	drivingPrivilage(age1)
	drivingPrivilage(age2)
	drivingPrivilage(age3)
	drivingPrivilage(age4)

	drivingPrivilage2(age1)
	drivingPrivilage2(age2)
	drivingPrivilage2(age3)
	drivingPrivilage2(age4)

	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		} else {
			continue
		}
	}

	i = 0

	for {
		fmt.Scan(&i)

		if i == 7 {
			break
		}
	}

	fmt.Println("Exit")
}

func checkDrinkingAge(a int) {
	fmt.Printf("You are %d year(s) old\n", a)

	if a < 21 {
		fmt.Println("No drink for you")
	} else {
		fmt.Println("Enjoy your drink")
	}

	fmt.Println()
}

func drivingPrivilage(a int) {
	if a >= 18 {
		fmt.Println("Full Drivers License")
	} else if a == 17 {
		fmt.Println("Propationary License")
	} else if a == 16 {
		fmt.Println("Permit")
	} else {
		fmt.Println("Not Allowed To Drive")
	}
}

func drivingPrivilage2(a int) {
	switch {
	case a >= 18:
		fmt.Println("Full Drivers License")
	case a == 17:
		fmt.Println("Propationary License")
	case a == 16:
		fmt.Println("Permit")
	default:
		fmt.Println("Not Allowed To Drive")
	}
}
