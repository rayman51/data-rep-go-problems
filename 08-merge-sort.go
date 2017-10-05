package main

import "fmt"


 var a[] int // declares array outside main
func main() {
	
	var n int
	fmt.Println("Please enter size of the array :")
	fmt.Println("=================================")
	fmt.Scanf("%d\n", &n)
	fmt.Printf("Please enter %d numbers only for the array :", n)
	set(n)// calls set function

	
 //fmt.Println(a)
 //fmt.Println(mergeSort(a))
}// main
func set(n int) {
	
	  a := make([]int, n)
	  for i := 0; i < n; i++ {
		  
		  fmt.Scan(&a[i])// creates the array and stores inputted numbers
		}
		fmt.Println("Original Array ")
		fmt.Println(a)
		fmt.Println("=================================")
		fmt.Println("Sorted Array ")
		fmt.Println(mergeSort(a))
	 }


func mergeSort(array []int) []int {
	
	 if len(array) <= 1 {
	  return array
	 }
	
	 middle := len(array) / 2
			
			//note how we slice the array using [:middle]
	 left := mergeSort(array[:middle])
			//next slice middle to end of array 
	 right := mergeSort(array[middle:])
	
	 return merge(left, right)
	}// mergeSort

	func merge(left, right []int) []int {
		//make an []int of size of length of left + length of right
		result := make([]int, len(left)+len(right))
		
		for i := 0; len(left) > 0 || len(right) > 0; i++ {
	   
		 if len(left) > 0 && len(right) > 0 {
		  if left[0] < right[0] {
		   result[i] = left[0]
		   left = left[1:]
		  } else {
		   result[i] = right[0]
		   right = right[1:]
		  }
		 } else if len(left) > 0 {
		  result[i] = left[0]
		  left = left[1:]
		 } else if len(right) > 0 {
		  result[i] = right[0]
		  right = right[1:]
		 }
		}
	   
		return result
	   }// merge