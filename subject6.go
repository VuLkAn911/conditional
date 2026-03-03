package main

import "fmt"

func main() {
	// sub 1
	fmt.Println(5 == 5)                 // true
	fmt.Println(10 != 3)                // true
	fmt.Println(7 > 12)                 //false
	fmt.Println(15 < 20)                // true
	fmt.Println(8 >= 8)                 // true
	fmt.Println(6 <= 4)                 // false
	fmt.Println((10 > 5) && (3 < 1))    // false
	fmt.Println((10 > 5) || (3 < 1))    // true
	fmt.Println(!(5 == 5))              //false
	fmt.Println(!(7 < 3))               // true
	fmt.Println(true && false)          // false
	fmt.Println(true || false)          // true
	fmt.Println((4+6 == 10) && (9 > 2)) // true
	fmt.Println((12/3 == 4) || (8 < 5)) // true

	// sub 2

	age := 25
	hasTicket := true

	canEnter := (age >= 18) && hasTicket
	fmt.Printf("Can enter: %t\n", canEnter)

	//sub 3

	var isLoggedIn bool = true
	var isAdmin bool = false

	hasAccess := isAdmin && isLoggedIn || isLoggedIn && !isAdmin
	fmt.Printf("Has access: %t\n", hasAccess)
}
