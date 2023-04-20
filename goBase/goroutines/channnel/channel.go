package main

func main() {
	send := make(chan int)
	recive := make(chan int, 3)

	go func(){
		defer close(send)
		for i := 0; i < 10; i++ {
			send <- i
		}
	}()

	go func() {
		defer close(recive)
		for i := range send {
			recive <- i * i
		}
	}()

	for i := range recive {
		println(i)
	}
}