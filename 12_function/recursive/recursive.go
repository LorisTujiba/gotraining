package recursive

/*==============================
Recursive
============================
A Function that call itself
over and over again
*/

func Factorial(input int) int {
	if input == 1 {
		return 1
	}
	return input * Factorial(input-1)
}
