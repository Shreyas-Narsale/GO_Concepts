package main

import "fmt"

/*
No arguments
No return values
Cannot be called manually
Automatically executed by Go runtime

Execution order:
1. Package-level variables initialized
2. init() functions executed
3. main() starts

init across multiple files in same package
a.go
b.go
Each file can have its own init()
Execution order:
Based on file compilation order (not strictly guaranteed → avoid dependency)

******
init across packages 

main → imports pkgA → imports pkgB
Execution order:
pkgB init()
pkgA init()
main init()
main()

Dependency order (bottom-up)

*/
func init() {
    fmt.Println("init called")
}

// multiple init functions in the same file
// top-to-bottom order
func init() {
    fmt.Println("init 1")
}

func init() {
    fmt.Println("init 2")
}

var a = setup()

func init() {
    fmt.Println("init")
}
/* order 
setup() runs first
then init()
*/
func main() {
    fmt.Println("main called")
}
