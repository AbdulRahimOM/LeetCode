package easy

func NumWaterBottles(numBottles int, numExchange int) int {
	drank := numBottles
	for numBottles >= numExchange {
		drank += numBottles / numExchange
		numBottles=numBottles/numExchange+numBottles%numExchange
	}
	return drank
}


func NumWaterBottles_InitialAttempt1(numBottles int, numExchange int) int {
	emptyBottle := numBottles
	drank := numBottles
	for emptyBottle >= numExchange {
		drank += emptyBottle / numExchange
		emptyBottle = emptyBottle/numExchange + emptyBottle%numExchange
	}
	return drank
}
