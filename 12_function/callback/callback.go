package callback

/*==============================
Callback
============================
function that can be passed as an
argument.
*/

//IsAMale is exported
func IsAMale(inputs []string, check func(string) bool) int {

	var count int

	for _, input := range inputs {
		if check(input) {
			count++
		}
	}
	return count
}
