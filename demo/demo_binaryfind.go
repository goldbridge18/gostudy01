package main

import "fmt"

func BinaryFind(arr *[6]int,leftindex int,rightindex int,findval int){

	if leftindex >rightindex {
		fmt.Println("not found")
		return
	}
	middle := (leftindex+rightindex)/2

	if (*arr)[middle] > findval {
		BinaryFind(arr,leftindex,middle - 1,findval)
	} else if  (*arr)[middle] < findval{

		BinaryFind(arr,middle + 1,rightindex,findval)

	}else {
		fmt.Println("find valuse's index   is ",middle)
	}

}

func runFind(){

	arr := [6]int{11,13,14,25,27,88}

	BinaryFind(&arr,0,len(arr)-1,205)
}
