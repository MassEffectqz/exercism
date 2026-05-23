package isogram
import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	seen := make(map[rune]bool)
	ignoredChars  := map[rune]bool{
		'!': true, 
		'?': true,
		' ': true, 
		'_': true,  
		'-': true,  
		'.': true,  
		
	}
	for _, letter := range word {
		if ignoredChars[letter] {
			continue
		}
		if seen[letter] {
			return false
		}
		seen[letter] = true
	}
	return true
}