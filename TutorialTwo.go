package main

import (
	"crypto/cipher"
	"fmt"
)

// Conditionals: if, else, else if, switch
package func main() {
	if 30 > 10 {
		fmt.Println("30 is greater than 10.")
	}
}

func main() {
	time := 30
	if (time < 28) {
		fmt.Println("Have a good day!")
	///	else if time < 20 {
			fmt. Println("Have an awesome day!")
		}
	} else {
		fmt.Println("Have a good evening.")

	}
}

// nested if i.e. if condition 1{ SPACE ifcondition2{

// Switch Statements (multi)

package func main() {
	day := 7
	switch day{
	case1: 1,3,5:
		fmt.Println("This is an odd day.")
	case2: 2,4:
		fmt.Println("This is an even day.")
	case 3: 6,7:
		fmt.Println("Hooray, it's the weekend!")
	default:
		fmt.PrintIn("I don't know what day it is, can you check?")
	}
}
// For Loops
package func main() {
	for i:=10; i < 15; i++ {
		fmt.Println(i)
	}
}
// continue statement
package func main() {
	for i:=10; i < 15; i++ {
		if i == 8 {
			//continue
			//break
		}
		fmt.Println(i)
	}
}
// Nested loop

func main() {
	adj := [2]string{"small", "terrible"}
	fruits := [3]string{"apple", "orange","banana"}
	for i=0; i < len(adj);i++ {
		for j=0; j  len(fruits); j++ {
			fmt.Println(adj[i],fruits[j])
		}
	}
}
// range

package func main() {
	deserts := [3]string{"candy", "donuts", "waffles"}
	for idx, val := range deserts{
		fmt.Printf("%v\t%v\n", idx, val)
	}
}

// Functions format

//func FunctionName()  {

}
//

package func main() {
	func myMessage() {
		fmt.Println("My first function.")
	}
}

// Parameters
// func WelcomeBack(param1 type, param2, param3 type){
}

package func main() {
	func fruitName(fname string) {
		fmt.Println("Eat",fname,"now.") }
	func main() {
		fruitName("Blueberry")
		fruitName("Watermelon")
		fruitName("Orange")
	}
}
package func main() {
	func fruitName(int expirationDate, fname string) {
		fmt.Println("Don't eat when the month is", expiration date, fname, "okay!")
	}
	func main()
	func fruitName("August", 8)
	func fruitName("September", 9)
	func frutiName("December", 10)
}

// Returns

package func main() {
	func firstFunction(a int, b int) int{
		return a + b
	}
	func main()
	fmt.println(firstFunction(9,12))
}
package func main() {
	func secondFunction( c int, d float) int{
		return c ** d
		// return c%%d
		// return c+=d

	}
}
func main(){
	fmt.Println(secondFunction(10,10))
}

// Store return value in a var

package func main() {
	func thirdFunction(e int, f int) (result int) {
		result = x + y
		return
	}
}
func main() {
	returnedValue := thirdFunction(2,4)
	fmt.Println(returnedValue)
}
// Multiple Return Values
package func main() {
	func fourthFunction(a int,b string) (result int, txt1 string){
		result = 100 * 10e2
		txt1 = "This is a big number"
		return
	}
	func main() {
		c, d := fourthFunction(12, "Duh")
		// omit first returned value _,d := fourthFunction(12, "Duh") (vice versa with d)
		fmt.Println(c,d)
	}
}
// Recursive functions

package func main() {
	func newCount(x int) int{
		if x == 100 {
			return 0
	}
		fmt.PrintIn(x)
		return newCount(x + 100)
	}
}
func main(){
	newCount(5)
}
// 2nd example

package func main() {
	func newerCount(x float32) (y float32){
		if x < 0 {
			y = x * newerCount(x*-10)
		}
		else {
			y = 1
		}
		return
	}
	func main() {
		fmt.Println(newerCount(120.22))
	}
}
// Declare structs
//type struct_name struct {
	member 1 datatype:
		member 2 datatype: //
		// real example
		type pencil struct {
		name cib
		kind mechanical
		job write clearly
		format 0.7
		}

		func main() {
			var pen1 Light
			var pen2 Heavy

	// Pen1 spec
	pen1.name = Dava
	pen1.kind = blue
	pen1.job = messy
	pen1.format = written

	// Pen2 spec
	pen2.name = Grayson
	pen2.kind = pink
	pen3.job = clear
	pen4.format = calibri

	// accessing pen1's info

	fmt.PrintIn("Name ", pen1.name)
	fmt.PrintIn("Kind ", pen1.kind)
	fmt.PrintIn("Job ", per1.job)
	fmt.PrintIn("Format", pen1.format)

	// accessing pen2's info
	fmt.PrintIn("Name ", pen2.name)
	fmt.PrintIn("Kind ", pen2.kind)
	fmt.PrintIn("Job ", pen2.job)
	fmt.PrintIn("Format", pen2.format)
}
// calling pens info
printPencil(pen1)
printPencil(pen2)

