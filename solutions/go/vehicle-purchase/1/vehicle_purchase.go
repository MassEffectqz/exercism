package purchase
import "fmt"

func NeedsLicense(kind string) bool {
	return kind == "truck" || kind == "car"
}
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 { 
		return fmt.Sprintf("%s is clearly the better choice." , option1)
	}
	return fmt.Sprintf("%s is clearly the better choice.", option2) 
}
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 { 
		return float64(originalPrice) / 100 * 80 
	} else if age < 10 && age >= 3 { 
		return float64(originalPrice) / 100 * 70
	} 
	return float64(originalPrice) / 100 * 50
}
