//doesn’t support inheritance. It uses composition.
package main
import "fmt"

type Engine struct {
    Power int
}

type Car struct {
    Engine // embedded struct
    Brand  string
}

func main() {
    c := Car{
        Engine: Engine{Power: 150},
        Brand:  "Toyota",
    }

    fmt.Println(c.Brand, c.Power)
}
