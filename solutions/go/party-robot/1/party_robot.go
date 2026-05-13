package partyrobot
import (
	"fmt"
)
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!" 
}
func HappyBirthday(name string, age int) string {
	message := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
	return message
}
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	message := fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %03d. Your table is %s, exactly %0.1f meters from here.\nYou will be sitting next to %s.", name, table, direction, distance, neighbor)
	return message
}