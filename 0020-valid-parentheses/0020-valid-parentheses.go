// To solve the problem we need to match pairs of parentheses taking into account

// the number of parentheses,
// their direction (open or closed) and
// their order.
// To be able to do it we have to reverse the order of opening parenthesis temporarily to match them to their closing peers. A data structure that would able to "reverse" the parentheses' order satisfying this requirement is stack.

// We traverse the string from left to right and check:

// If we see an opening parenthesis we put it to stack
// If it’s the closing one, then it must correspond to the parenthesis at the top of our stack, so we check if it is true and we remove this pair of parentheses if it's so. If not, we return false and stop the function’s execution immediately.
// At the end, iff we have empty stack, we have valid string thus return true.

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
    
	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}

	for _, r := range s {
		if _, ok := pairs[r]; ok {
			stack = append(stack, r)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != r {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}