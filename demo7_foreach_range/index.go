package main

import "fmt"

func main()  {
	courses := []string {"android","ios","react"};
	for index, item := range courses{
		fmt.Printf("%d %s\n",index ,item)
	}

	for _, item := range courses {
		fmt.Printf("%s\n" ,item)
	}
}
