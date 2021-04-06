package grams2ounces

// GramsToOunces converts weight in grammes (int) to ounces (float32)
func GramsToOunces(weightInGrammes int) float32 {
	return float32(weightInGrammes) / 28.0
}
