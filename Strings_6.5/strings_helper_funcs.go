var str string
	str = "sr.demo.txt"

	sli := []byte(str)
	fmt.Println(string(sli))

	isContain := strings.Contains(str, "de")
	fmt.Println(isContain)

	hasPrefix := strings.HasPrefix(str, "sr.")
	fmt.Println(hasPrefix)

	hasSuffix := strings.HasSuffix(str, ".txt")
	fmt.Println(hasSuffix)

	toUpper := strings.ToUpper(str)
	fmt.Println(toUpper)

	toLower := strings.ToLower(str)
	fmt.Println(toLower)

	str1 := "   demo  nan ** "

	trimSpace := strings.TrimSpace(str1)
	fmt.Println(trimSpace) // "demo  nan **"

	trim := strings.Trim(trimSpace, "*")
	fmt.Println(trim) // "demo  nan"

	sli1 := []string{"I", "Love", "India"}
	str3 := strings.Join(sli1, " ")
	fmt.Println(str3)

	str4 := "Java,Golang,C"
	sli4 := strings.Split(str4, ",")
	fmt.Println(sli4)

	fmt.Println(strings.Compare("abc", "abc"))
