package main

import "fmt"

func main() {
	fmt.Print("Enter Score: ")
	score := 0
	fmt.Scanln(&score)

	fmt.Println("Your Grade is", gladeCal(score))
}

func gladeCal(score int) string {
	if score >= 80 {
		return "A"
	} else if score >= 70 {
		return "B"
	} else if score >= 60 {
		return "C"
	} else if score >= 50 {
		return "D"
	} else {
		return "F"
	}
}

func gladeCalSw(score int)String{
	switch score{
		case score >= 80{

		}

	}


}