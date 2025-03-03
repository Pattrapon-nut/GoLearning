package main

import "fmt"

func main()  {
fnFor();
fnWhile();
fnWhileUsingBreak();
}

func fnFor()  {

	for i := 0; i <= 10; i++ {
		fmt.Printf("for i %d\n", i)
	}
}

func fnWhile()  {
	var index int = 0;
	for index < 5{
		index++;
		fmt.Printf("index %d\n", index)
	}
}

	func fnWhileUsingBreak()  {
		var index int = 0;
		for true{
			if index > 5 {
				break
			}
			
			fmt.Printf("while break=index %d\n", index)
			index++;
		}
}