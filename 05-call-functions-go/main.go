package main

// ✋ make this function callable from host
//
//export add
func add(x int, y int) int {
	return x + y
}

//export hello
func hello(name string) string {
	return "hello " + name
}

func main() {
	// fmt.Println("Hello World")
	// argsWithoutProg := os.Args[1:]
	// r := add(strconv.Atoi(argsWithoutProg[0]), strconv.Atoi(argsWithoutProg[1]))
}
