package callback

/*==============================
Callback
============================
function that can be passed as an
argument.
*/


func IsAMale(inputs []string,check func(string) bool) int{

	var count int = 0

	for _, input := range inputs{
		if check(input){
			count++
		}
	}
	return count
}