package main

func validParenthesis(s string) bool {
	keys := []rune{'{', '(', '['}
	values := []rune{'}', ')', ']'}
	dict := make(map[rune]rune)

	for i := range keys {
		dict[keys[i]] = values[i]
	}

	stk := Stack{}

	for _, c := range s {
		if _, ok := dict[c]; ok {
			stk.Push(c)
		} else {
			if dict[stk.Peek().(rune)] == c {
				stk.Pop()
			} else {
				return false
			}
		}
	}

	return stk.IsEmpty()
}

// for i := range keys {
// 	dict[keys[i]] = values[i]
// }

// stk := &Stack{}

// for _, c := range s {
// 	if _,ok := dict[c]; ok {
// 		stk.Push(c)
// 	} else if stk.IsEmpty() {
// 		return false
// 	} else {
// 		top, _ := stk.Pop()
// 		if dict[top.(rune)] != c {
// 			return false
// 		}
// 	}
// }
