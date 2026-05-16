package purchase
import "fmt"

func NeedsLicense(kind string) bool {
	return kind == "truck" || kind == "car"
}

func ChooseVehicle(option1, option2 string) string {
    bestOption := option2 
	if option1 < option2 { 
		bestOption = option1
	}
	return fmt.Sprintf("%s is clearly the better choice.", bestOption) 
}

func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 { 
		return originalPrice * 0.8
	} else if age < 10{ 
		return originalPrice * 0.7
	} 
	return originalPrice * 0.5
}
