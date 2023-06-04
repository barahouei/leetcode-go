// Package validparentheses validates parentheses and brackets.
package validparentheses

var openToClose = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

var closeToOpen = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

// stack represents a stack of items.
type stack struct {
	items []rune
}

// push adds a bracket to the stack.
func (s *stack) push(char rune) {
	s.items = append(s.items, char)
}

// pop returns the top item of the stack.
func (s *stack) pop() rune {
	length := len(s.items) - 1

	top := s.items[length]
	s.items = s.items[:length]

	return top
}

// IsValid returns true if open and closed brackets are in order and correspond.
func IsValid(s string) bool {
	var brackets stack

	for _, char := range s {
		if _, ok := openToClose[char]; ok {
			brackets.push(char)
		}

		if _, ok := closeToOpen[char]; ok {
			if len(brackets.items) == 0 {
				return false
			}

			top := brackets.pop()

			if top != closeToOpen[char] {
				return false
			}
		}
	}

	return len(brackets.items) == 0
}
