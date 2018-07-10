package prime

import (
	"fmt"
)

//Prime function tells you the given input is prime or not
func Prime() {
	fmt.Print("Enter the number : ")
	var p int
	fmt.Scanf("%d", &p)
	if p < 2 || p%2 == 0 {
		fmt.Println(p, "Not prime number")
	} else if p == 2 {
		fmt.Println(p, "Only even prime number")
	} else {
		for i := 3; i < p; i++ {
			if p%i == 0 {
				fmt.Println(p, "Not prime numer")
				break
			}
		}
		fmt.Println(p, "Prime numer")
	}
}
