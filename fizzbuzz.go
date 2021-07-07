package fizzbuzz

/*var ops = []func(int) (string, bool){
	func(num int) (string, bool) {
		if num%3 == 0 && num%5 == 0 {
			return "fizz buzz", true
		}
		return "", false
	},
	func(num int) (string, bool) {
		if num%3 == 0 {
			return "fizz", true
		}
		return "", false
	},
	func(num int) (string, bool) {
		if num%5 == 0 {
			return "buzz", true
		}
		return "", false
	},
}*/

/*func Fizzbuzz(num int) (string, bool) {
	for _, op := range ops {
		result, ok := op(num)
		if ok {
			return result, ok
		}
	}
	return "", false
}*/

// Fizzbuzz takes a number, and returns either a string and boolean
// true, or an empty string and boolean false
//
// Callers are supposede to print the original number if false si
// returned
func Fizzbuzz(x int) (string, bool) {
	if x%3 == 0 && x%5 == 0 {
		return "Fizz Buzz", true
	}
	if x%3 == 0 {
		return "Fizz", true
	}
	if x%5 == 0 {
		return "Buzz", true
	}
	return "", false
}
