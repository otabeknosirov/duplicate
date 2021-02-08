package main

import "fmt"

func main()  {
	
	var n int
	fmt.Println("Enter the length")
	fmt.Scan(&n)

	var num = make([]int,n) 

	fmt.Println("Enter the number")
	for i:=0; i < n; i++ {
         fmt.Scan(&num[i])
	}

	var hasDuplicate bool = duplicate(num)
	if(hasDuplicate){
		fmt.Println("Has Duplicate")
	}else{
		fmt.Println("Has no Duplicate")
	}
}

func duplicate(num[] int)bool{
	var hasDup bool = true
	for i:=0; i < len(num); i++{
		for j:= i + 1; j < len(num) - 1; j++{
			if(num[i] != num[j]){
				hasDup= false
			}else if(num[i] == num[j]){
				hasDup = true
				break
			}
		}
		if(hasDup == true){
			break
		}
	}
	return hasDup
}