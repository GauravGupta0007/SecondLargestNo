//Program to find Factorial of number
package main
import "fmt"


func main(){    

	var n int
	fmt.Print("Enter size of array: ")
    fmt.Scan(&n)  
	fmt.Println("Enter array elements")
	var arr = make([]int, n)
	
	for i:=0;i<n;i++ {
		fmt.Scan(&arr[i])
	}

	ans:=arr[0]

	m:=arr[0]

	for i:=0;i<n;i++ {
		m = max(m, arr[i])
	}

	for i:=0;i<n;i++ {
		if(arr[i] == m) {
			var temp = arr[i];
			arr[i] = arr[n-1]
			arr[n-1] = temp
			break
		}
	}

	for i:=0;i<n-1;i++ {
		ans = max(ans, arr[i])
	}

	fmt.Println("Second Largest no : ", ans)

}