package collatzconjecture
import "fmt"
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("n must be positive")
	}
	var iteration int 
	for ; n != 1 ;{ 
		if n % 2 == 0 {
			n = n / 2 
			iteration++
		} else {
			n = n * 3 + 1
			iteration++
		}
		
	}
	return iteration, nil
}
