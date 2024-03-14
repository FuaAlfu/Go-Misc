import (
	"fmt"
	"math"
	"strings"

)

/*
The function checks for integer overflow to 
ensure the result stays within the 32-bit signed integer range.
*/

func myAtoi(str string) int {
	// Trim leading and trailing whitespaces
	str = strings.TrimSpace(str)

	// Return 0 if the string is empty
	if str == "" {
		return 0
	}

	// Define variables for the result and the sign of the number
	var result int
	sign := 1

	// Check for the sign character (+ or -)
	if str[0] == '+' || str[0] == '-' {
		// Update the sign and move to the next character
		if str[0] == '-' {
			sign = -1
		}
		str = str[1:]
	}

	// Iterate through the string and convert characters to integers
	for _, char := range str {
		// Break if the character is not a digit
		if char < '0' || char > '9' {
			break
		}

		// Convert the character to an integer
		digit := int(char - '0')

		// Check for integer overflow
		if result > math.MaxInt32/10 || 
		(result == math.MaxInt32/10 && digit > math.MaxInt32%10) {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		// Update the result by appending the digit
		result = result*10 + digit
	}

	// Apply the sign to the result
	result *= sign

	return result
}

func main() {
	str := "   -42"
	result := myAtoi(str)
	fmt.Println(result) // Output: -42
}