
/// arr copy by value
func change(s [3]int) {
	s[0] = 100
}

func main() {
	arr := [3]int{1, 2, 3}
	change(arr)
	fmt.Println(arr) // [1 2 3]
}


/// vs sli copy by refernce


func change(s []int) {
	s[0] = 100
}

func main() {
	arr := []int{1, 2, 3}
	change(arr)
	fmt.Println(arr) // [100 2 3]
}
