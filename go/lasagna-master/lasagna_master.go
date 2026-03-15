package lasagnamaster

// PreparationTime returns the time it takes to complete a lasagna with given number of layers.
// When 0 time per layer is given, assume 2.
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		return len(layers) * 2
	}
	return len(layers) * timePerLayer
}

// Quantities accepts a slice of layers where 'noodles' and 'sauce' can be present
// returns a tuple with total noodles: int, sauce: float64.  each noodle is 50g, sauce 0.2 L.
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, v := range layers {
		switch v {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// AddSecretIngredient replaces the last item in myIngredients with the last item in friendIngredients
func AddSecretIngredient(friendIngredients []string, myIngredients []string) {
	myIngredients[len(myIngredients)-1] = friendIngredients[len(friendIngredients)-1]
}

// ScaleRecipe returns a slice representing the amounts for different numbers of portions while keeping the original slice unaltered.
// Note that the slice: 'amounts' includes enough for 2 portions.
func ScaleRecipe(amounts []float64, portions int) []float64 {
	scaledRecipe := make([]float64, len(amounts))
	copy(scaledRecipe, amounts)
	for i := 0; i < len(amounts); i++ {
		scaledRecipe[i] = (amounts[i] / 2) * float64(portions)
	}
	return scaledRecipe
}
