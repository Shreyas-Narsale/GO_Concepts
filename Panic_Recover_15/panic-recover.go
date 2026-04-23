func example() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered1 :", r)
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered2 :", r)
		}
	}()

	panic("panic 1")

	// This line will never execute
	panic("panic 2")
}

here output: Recovered2 : panic
