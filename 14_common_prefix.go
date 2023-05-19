package main

func commonPrefix (arr []string) string {
	common := ""
	base := arr[0]

	// iterate though each character of BASE
	for i := 0; i < len(base); i++ {

		// create a flag variable and initialize it to true
		isCommon := true

		// iterate through each word of the array
		for j := 0; j < len(arr); j++ {

			// conditional if base at i NOT equal that word at i
			if(arr[j][i] != base[i]){

				// change flag to false
				isCommon = false
			}
		}
		if isCommon == false {

			// if it is false, exit loop immediately
			break
		}
		
		// increment common string by that char if flag is true
		common += string(base[i])
	}
	return common
}