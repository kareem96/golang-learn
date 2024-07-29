package main

import (
	"fmt"
	"sort"
)

type User struct{
	Name string
	Age int
}

type UserSlice[] User

func(s UserSlice) Len() int{
	return len(s)
}

func (s UserSlice) Less(i, j int) bool{
	return s[i].Age < s[j].Age
}

func(s UserSlice) Swap(i, j int){
	// temp := s[i]
	// s[i] = s[j]
	// s[j] = temp

	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"kareem", 1},
		{"abdul", 10},
		{"melayu", 12},
		{"abd", 3},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}