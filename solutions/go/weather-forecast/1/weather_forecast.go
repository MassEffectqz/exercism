//Package weather is designed to calculate weather conditions in the city.
package weather

// var (
// 	CurrentCondition string
// 	CurrentLocation  string
// )

//Forecast function displays current weather information.
func Forecast(city, condition string) string {
    // this another comment
	CurrentLocation, CurrentCondition := city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
