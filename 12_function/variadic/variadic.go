package variadic

/*==============================
...[type] - Variadic
============================
function that can be invoked
with zero or more arguments
for that function.
*/

//Calculate , variadic examples. using dot dot dot
func Calculate(inputs ...float64) float64 {
	var total float64

	//get the inputs
	for _, input := range inputs {
		total += input
	}
	return total
}
