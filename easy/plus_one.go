package main

import "fmt"

func plusOne(digits []int) []int {
	carry := false
	for i := len(digits) - 1; i >= 0; i-- {
		// if it's last digit add 1 if the result is 10 then set digit to 0 and carry to true
		// if not the last digit and carry set add 1 and reset carry
		// if this is first digit and is 9 and carry is set then append and set digit to 0
		switch {
		case i == 0 && carry && digits[i] == 9:
			digits[i] = 0
			digits = append([]int{1}, digits...)
			break
		case i == 0 && len(digits) == 1 && digits[i] == 9:
			digits[i] = 0
			digits = append([]int{1}, digits...)
			break
		case i == len(digits)-1:
			digits[i] += 1
			if digits[i] == 10 {
				digits[i] = 0
				carry = true
			}
			break
		case i != len(digits)-1 && carry:
			digits[i] += 1
			carry = false
			if digits[i] == 10 {
				digits[i] = 0
				carry = true
			}
			break
		}
	}
	return digits
}

func main() {
	//one := plusOne([]int{1, 2, 3})
	//one := plusOne([]int{9, 9})
	one := plusOne([]int{9, 9, 9})
	//one := plusOne([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0})
	//one := plusOne([]int{9})
	//one := plusOne([]int{9})
	fmt.Printf("%d\n", one)

}
