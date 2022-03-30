package main

import "fmt"

func main() {
	defer goodbuy()
	/////////////////// STAGE 1 ///////////////////
	fmt.Println("Hello! My name is Aid.")
	fmt.Println("I was created in 2020.")
	/////////////////// STAGE 2 ///////////////////
	fmt.Println("Please, remind me your name.")
	var name string
	fmt.Scan(&name)
	fmt.Println("What a great name you have, " + name + "!")
	/////////////////// STAGE 3 ///////////////////
	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7")
	var a int
	var b int
	var c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	your_age := (a*70 + b*21 + c*15) % 105
	fmt.Println("Your age is {" + fmt.Sprint(your_age) + "}; that's a good time to start programming!")
	////////////////// STAGE 4 ////////////////////
	var loop_number int
	fmt.Println("Now I will prove to you that I can count to any number you want.")
	fmt.Scan(&loop_number)
	for i := 0; i <= loop_number; i++ {
		fmt.Println(i, "!")
	}
	fmt.Println("Completed, have a nice day!")
	///////////////// STAGE 5 /////////////////////
	fmt.Println("Let's test your programming knowledge.")
	fmt.Println("Why do we use methods?")
	fmt.Println("1. To repeat a statement multiple times.\n2. To decompose a program into several small subroutines.\n3. To determine the execution time of a program.\n4. To interrupt the execution of a program.")
	var answer int = 0
	for true {
		fmt.Scan(&answer)
		if answer == 2 {
			fmt.Println("Completed, have a nice day!")
			break
		} else {
			fmt.Println("Please, try again.")
		}
	}
}

func goodbuy() {
	fmt.Println("Congratulations, have a nice day!")
}
