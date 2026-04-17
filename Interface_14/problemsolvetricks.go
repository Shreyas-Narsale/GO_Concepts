🔥 Core idea

Interface = (type, value)
nil interface = (nil, nil)

🔹 Rule
    i == nil  → true ONLY if type=nil AND value=nil
    🔹 Most important cases

1. True nil
    var i interface{}
    fmt.Println(i == nil) // true


2. Typed nil (TRAP)
    var p *int = nil
    var i interface{} = p
    
    fmt.Println(i == nil) // false
    
    👉 (type=*int, value=nil) → NOT nil

3. Nil slice/map/chan in interface
    var s []int = nil
    var i interface{} = s
    
    fmt.Println(i == nil) // false


4. Nil error (VERY IMPORTANT)
    func f() error {
    	var err *MyError = nil
    	return err
    }
    err := f()
    fmt.Println(err == nil) // false err (type *MyError , value = nil)

5.Function trap
    func check(i interface{}) {
    	if i == nil {
    		fmt.Println("nil")
    	} else {
    		fmt.Println("not nil")
    	}
    }
    
    var p *int = nil
    check(p) // not nil

🔥 Memory trick
    typed nil ≠ nil interface
    
    ⚡ One-line summary
    
    If interface has a type, it is NOT nil (even if value is nil)
