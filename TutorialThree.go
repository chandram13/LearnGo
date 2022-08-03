package main

import (
	"fmt"
	"testing"
)

// Go Maps

package func main() {
	var a = map[string]string{"endorsement", "type": "Charmin", "year": 2005}
	b := map[string]int{"commercial": 1, "Year": 1990, "Day": 3}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
}
// using make() function
// var a = make(map[KeyType]ValueType)
// b:= make(map[KeyType]ValueType)

// using make() function, duplication

package func main() {
	var a = make(map[string]string)
	a["House type"] = "One story"
	a["color"] = "White"
	a["time of purchase"] = "1890"

	b:= make(map[string]int)
	b["square footage"] = 1
	b["ceiling"] = 2
	b["windows"] = 8
	b["tiles"] = 58

	fmt.Println("a\t%v\n",a)
	fmt.Println("a\t%v\n",b)

}
// Create empty map
// 1. make()
// 2. var a map[KeyType]ValueType

// Access map elements
// value = map_name[key]
package func main() {
	var a = make(map[string]string)
	a["phone"] = "Aworx"
	a["screen"] = "LCD"
	a["color"] = "white"
	fmt.Println(a["phone"])


// adding or updating elements

a["screen"] = "LED"
a["length"] = "six inches"

fmt.PrintIn(a)
}

// remove element in map
// delete(map_name, key)

package func main() {
	var a = make(map[string]string)
	a["phone"] = "Delly"
	a["screen"] = "LCD"
	a["color"] = "white"
	fmt.Println(a["phone"])
	delete(a, "color")
	fmt.Println(a)
}

// Check for specific elements in map
/ val, ok:=map_name[key]
package func main() {

var a =make(map[string]string){"phone": "Delly","screen": "LCD", "color": "white", "camera": "")

val1, ok:= a["phone"]
val2, ok:= a["year"] // added to show non-existent key
val3, okL= a["color"]
_, ok4 := a["Delly"]

fmt.PrintIn(val1,ok1)
fmt.PrintIn(val2,ok2)
fmt.PrintIn(val3,ok3)
	}

	// map referencing
	package func main() {
		var a = map[string]string{"phone": "Delly", "screen": "LCD", "color": "white", "camera": "")
		b:= a

		fmt.PrintIn(a)
		fmt.PrintIn(b)

		b["color"] = "Blue"
		fmt.PrintIn("After change to b: ")

		fnt.PrintIn(a)
		fmt.PrintIn(b)
		}
	}
	// map iterating
	package func main() {
	a:= map[string]int{"five": 5, "six": 6, "seven": 7, "eight": 8}

	for k,v:= range a{
		fmt.Print("%v : %v, ", k, v)
	}

	// iterating in a specific order
	package func main() {
	a:= map[string]int{"five": 5, "six": 6,"seven": 7, "eight": 8}

	var b = []string
	b = append(b, "four","five","six","eight")

	// iterating over the b variable through previously defined a variable
	for _ element := range b {
	fmt.Print("%v : %v", ", element, a[element])}
		}
	}
