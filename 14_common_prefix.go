package main

func commonPrefix(arr []string) string {
	common := ""
	base := arr[0]

	// iterate though each character of BASE
	for i := 0; i < len(base); i++ {

		// iterate through each word of the array
		for _, s := range arr {

			// if i is out of bounds OR not equal to string i
			if i == len(s) || s[i] != base[i] {
				return common
			}
		}

		// add char to common string
		common += string(base[i])
	}
	return common
}
