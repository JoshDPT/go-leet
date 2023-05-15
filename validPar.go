package main

func validParenthesis (s string) bool {
	dict := make(map[string]string)
	dict["}"] = "{"
	dict["]"] = "["
	dict[")"] = "("
	stk := []string{}

	for _,c := range s {
		ch := string(c)
		switch ch {
			case "{", "[" , "(":
				stk = append(stk, ch) 
			case "}", "]", ")":
				if len(stk) == 0 || stk[len(stk)-1] != dict[ch] {
					return false
				} else {
					stk = stk[:len(stk)-1]
				}
		}
	}
	return len(stk) == 0
}