package techpalace
import (
	"fmt"
	"strings"
)
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stars := strings.Repeat("*", numStarsPerLine)
	return fmt.Sprintf("%v\n%v\n%v", stars, welcomeMsg, stars)
}
func CleanupMessage(oldMsg string) string {
	oldMsg = strings.Replace(oldMsg, "*", "", len(oldMsg))
	oldMsg = strings.TrimSpace(oldMsg)
	return oldMsg
}