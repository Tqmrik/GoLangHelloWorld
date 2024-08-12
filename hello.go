package main

func bulkSend(numMessages int) float64 {
	var b float64 = 0
	for i := 0; i < numMessages; i++ {
		b += 1 + 0.01*float64(i)
	}
	return b
}
