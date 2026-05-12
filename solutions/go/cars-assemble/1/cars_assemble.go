package cars
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) / 100 * successRate
}
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	result := float64(productionRate) / 100 * float64(successRate) / 60
	return int(result)
}
func CalculateCost(carsCount int) uint { 
	count1 := carsCount / 10 
	count2 := carsCount % 10   
	count1 = count1 * 95000
	count2 = count2 * 10000
	carsCount = count1 + count2
    return uint(carsCount)
}