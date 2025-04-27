package main

func main(){

	defer println("bye")
	defer println("good")

	for i := 0; i<10; i++{
		defer println(i)
	}

	println("hello")
}
	