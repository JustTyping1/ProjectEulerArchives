package main

func main() {
	var fibonaccis []int

	var running = 0

	fibonaccis = append(fibonaccis, 1)
	fibonaccis = append(fibonaccis, 2)

	for fibonaccis[len(fibonaccis)-1] < 4000000 {
		fibonaccis = append(fibonaccis, ((fibonaccis[len(fibonaccis)-1]) + (fibonaccis[len(fibonaccis)-2])))
	}

	for _, element := range fibonaccis {
		if element%2 == 0 {
			running += element
		}
	}

	println(running)
}
