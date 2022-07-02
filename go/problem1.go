package main

func main() {
	var runningtotal = 0
	i := 1

	for i < 1000 {
		if i%3 == 0 || i%5 == 0 {
			runningtotal += i
			println(i)
		}
		i += 1
	}
	println(runningtotal)
}
