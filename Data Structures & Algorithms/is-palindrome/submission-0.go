
func isPalindrome(s string) bool {

	isPalindrome := true	
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	noSpaceString := re.ReplaceAllString(strings.ReplaceAll(strings.ToLower(s)," ",""), "")
	lastIndex := len(noSpaceString) - 1
	for i := 0; i < len(noSpaceString) -1 ; i++ {
			if rune(noSpaceString[i]) !=  rune(noSpaceString[lastIndex]){
				isPalindrome = false
			} 
			lastIndex = lastIndex - 1
	}
	return isPalindrome
}
