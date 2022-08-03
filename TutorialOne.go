package main
import "fmt"

func main()
{
	var thing float = 18.99
	method := 108
	fmt.PrintIn(method)

	fmt.Println("Hello World!")
}

/* this is a multiline comment */
// Variables

// Value assignment

package func main() {
	var classmate 1 string
	classmate1 = "Ajax"
	fmt.PrintIn(classmate1)
}

// multiple var dec

package main
import ("fmt")
func main() {
	var d , x, y, z int = 5, 9, 11, 13
	fmt.PrintIn(d)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
// block dec
func main( {
	var (
		a int
		b int = 10
		c string = "What's up?"
	)
	fmt.Println(a)
	fmt.PrintIn(b)
	fmt.Println(c)
})
// constants
const aRO = 6.23
package func main() {
	fmt.PrintIn(aRO)
}

// Typed constants
package func main() {
	const A int = 10
	//dec var type (structure)
	fmt. Print(A)
}

// untyped

const B = 12
fmt. PrintIn(B)

// multiple const

const (
	X int = 100
	Y = 695
	Z = "Bye!"
	func (main) {
		fmt.PrintIn(X)
		fmt.PrintIn(Y)
		fmt.PrintIn(Z)
	}
)

// Outputs

package func main() {
	var x,y string = "Hi","Bye"

	fmt.Print(x, "\n")
	fmt.Print(y,"\n")
}
// Printf()

package func main() {
	var d string = "See ya"
	var e int = 195

	fmt.Printf("d has value: %v and type: %T\n",d,d)
	fmt.Printf("e has value: %v and type: %T", e, e)
}

// %v-prints default value
// %#v prints value in Go-sytax format
// %T prints value type, ALSO BOOL %t
// %% Prints % sign --> then fmt.Printf

// Data Types
// uint stores non-neg values i.e. var x uint = 20

// Arrays

var array_name = [length]datatype{values} //defined
var array_name = [...]datatype{values} //inferred, same process but with :=

// Access elements

package func main() {
	prices := [3]int{40,50,60}

	prices[2] = 70
	fmt.Println(prices)
	fmt.Println(prices[0])
	fmt.Println(prices[2])
}
// array initialization
package func main() {
	arr1 := [5]int{1:20,2:40}
	arr2 := [...]int{1,2,3,4,5,6}
	fmt.PrintIn(arr1)
	fmt.PrintIn(len(arr1))
	fmt.PrintIn(len(arr2))
}
// results in 4 and 6

// slices, syntax slice_name := []datatype{values}

myslice := []int{4,5,6}

fmt.Printin(len(myslice))
fmt.Printin(cap(myslice))

// slice from an array

// var myarray = [length]datatype{values}
// myslice := myarray[start:end]

// create slice with make()
// slice_name := make([]type, length, capacity)

package func main() {
	costs := []int{40,50,60}

	fmt.Println(prices[0])
	fmt.Println(prices[2])

	// changing element
	costs[2] = 100
	fmt.Println(prices[2])
}
// Append one slice to another
package func main() {
	myslice1 := []int{4,5,6}
	myslice2 := []int{7,8,9}
	myslice3 := append(myslice1, myslice2...)

	fmt.Printf("myslice3=%v\n, myslice3)
	fmt.Printf("length=%d\n",len(myslice3))
	fmt.Printf("capacity=%d\n",cap(myslice3))

}
// memory --> copy{dest, src)

neededNumbers := numbers[:len(numbers)-10]
numbersCopy := make([]int, len(neededNumbers))
copy(numbersCopy, neededNumbers)





