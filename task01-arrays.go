package homework

func average(input [15]float32) (result float32) {
	var sum float32 = 0.0
	N := 0.0
	for _, i := range input {
		sum += i
		if i != 0.0 {
			N++
		}
	}
	result = sum / float32(N)
	return
}
