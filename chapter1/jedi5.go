package main

//Points to ponder
//	- Underlying types
//	- Type conversion (instead of type casting).
//	- runtime package, how to get information about underlying OS and architecture.

import (
	"fmt"
	"runtime"
)

type hotdog int

var x hotdog
var y hotdog

func main() {
	fmt.Print("hello")
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = hotdog(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(runtime.GOOS)
	fmt.Print(runtime.GOARCH)

}
