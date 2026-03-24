package cars

const CarCost = 10000
const groupCost = 95000
const groupSize = 10

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var unGroupedCars = carsCount % groupSize
	var unGroupedTotalCost = unGroupedCars * CarCost
	var groupedCarTotalCost = (carsCount / groupSize) * groupCost

	return uint(unGroupedTotalCost) + uint(groupedCarTotalCost)
}
