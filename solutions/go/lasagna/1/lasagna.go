package lasagna
const OvenTime = 40
func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}
func PreparationTime(numberOfLayers int) int {
	timeOfLayser := 2 
	return timeOfLayser * numberOfLayers
}
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return (numberOfLayers * 2) + actualMinutesInOven
}