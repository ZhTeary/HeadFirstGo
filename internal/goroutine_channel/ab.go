package goroutine_channel

func Greeting(myChannel chan string) {
	myChannel <- "hi"
}
func Abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}
func Def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}
