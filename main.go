package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var myint int8
	for i := 0; i < 129; i++ {
		myint += 1
	}
	fmt.Println(myint)

	var men uint8
	men = 5
	var women int16
	women = 6
	var people int
	people = int(men) + int(women)
	fmt.Println(people)

	var maxFloat32 float32
	maxFloat32 = 16777216
	fmt.Println(maxFloat32 == maxFloat32+10) // you would typically expect this to return false
	// it returns true
	fmt.Println(maxFloat32 + 10)      // 16777216
	fmt.Println(maxFloat32 + 2000000) // 16777216

	days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	fmt.Println(days[0]) // prints 'monday'
	fmt.Println(days[5]) // prints 'saturday'
	weekdays := days[0:5]
	fmt.Println(weekdays)

	youtubeSubscribers := map[string]int{
		"TutorialEdge":     2240,
		"MKBHD":            6580350,
		"Fun Fun Function": 171220,
	}

	fmt.Println(youtubeSubscribers["MKBHD"])

	type Person struct {
		name string
		age  int
	}

	// our Team struct
	type Team struct {
		name    string
		players [2]Person
	}

	// declaring an empty 'Team'
	var myTeam Team
	fmt.Println(myTeam)

	players := [...]Person{{name: "Forrest", age: 25}, {name: "Gordon"}}
	// declaring a team with players
	celtic := Team{name: "Celtic FC", players: players}
	fmt.Println(celtic)

	fullname := myfunction("Alex", "Hu")
	fmt.Println(fullname)

	myFunc := addOne()
	fmt.Println(myFunc()) // 2
	fmt.Println(myFunc()) // 3
	fmt.Println(myFunc()) // 4
	fmt.Println(myFunc()) // 5
}

func myfunction(firstname string, lastname string) string {
	fullname := firstname + " " + lastname
	return fullname
}

func addOne() func() int {
	var x int
	return func() int {
		x++
		return x + 1
	}
}
