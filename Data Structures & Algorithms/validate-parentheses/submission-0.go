func isValid(s string) bool {
	brackets := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{}
	for _, runeBracket := range s {
		value, exists := brackets[runeBracket]
		if exists {
			stack = append(stack, value)
		} else {
			if len(stack) > 0 {
				topOfStack := stack[len(stack)-1]
				if topOfStack == runeBracket {
					stack = stack[:len(stack)-1]
				} else {
					return false 
				}
			} else {
				return false
			}
		}

	}
 return len(stack) == 0
}