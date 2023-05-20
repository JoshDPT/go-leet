package main

func validParenthesis(s string) bool {
	dict := make(map[string]string)
	dict["}"] = "{"
	dict["]"] = "["
	dict[")"] = "("

	stk := ByteStack{}

	for _, c := range s {
		ch := string(c)

		switch ch {

		case "{", "[", "(":
			stk.Push(byte(c))

		case "}", "]", ")":
			if stk.IsEmpty()|| string(stk.Peek()) != dict[ch] {
				return false

			} else {
				stk.Pop()

			}

		}
	}
	return stk.IsEmpty()
}
