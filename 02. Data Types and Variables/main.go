package main

import "fmt"

func main() {
	fmt.Println("___________Data types in Go____________")

	var num int = 5
	fmt.Printf("num is type of %T \n", num)

	var numUin8 uint8 = 255 // 0 to 255
	fmt.Printf("numUin8 is type of %T \n", numUin8)

	var pi float32 = 3.14
	fmt.Printf("pi is type of %T \n", pi)
	fmt.Println(pi)

	var pi64 float64 = 3.14159265358979323846264338327950288419716939937510
	fmt.Printf("pi is type of %T \n", pi64)
	fmt.Println(pi64)

	var name string = "Abhishek"
	fmt.Printf("name is type of %T \n", name)

	var isLoggedIn bool = true
	fmt.Printf("isLoggedIn is type of %T \n", isLoggedIn)

	// implicit type
	var implicitVar = "WUBBA LUBBA DUB DUB"
	fmt.Printf("var is type of %T \n", implicitVar)

	// constant
	const username string = "Abhishek"
	fmt.Println(username)
	const repoUrl = "https://github.com/Abhishekbagdiya01/Go-lang-learning"
	fmt.Println(repoUrl)

}
