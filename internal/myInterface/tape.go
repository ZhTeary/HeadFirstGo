package myInterface

import "fmt"

type TapePlayer struct {
	Batteries string
}
type TapeRecord struct {
	MicroPhones int
}

type Tape interface {
	Play(string)
	Stop()
}

func PlayList(device Tape, songs []string) {
	for _, s := range songs {
		device.Play(s)
	}
	device.Stop()
}

func (t TapePlayer) Play(song string) {
	fmt.Println("TapePlayer Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("TapePlayer Stop!")
}

func (t TapeRecord) Play(song string) {
	fmt.Println("TapeRecord Playing,", song)
}

func (t TapeRecord) Stop() {
	fmt.Println("TapeRecord Stop!")
}

//func main() {
//	songs := []string{"song1", "song2", "song3"}
//	myInterface.PlayList(myInterface.TapeRecord{}, songs)
//	myInterface.PlayList(myInterface.TapePlayer{}, songs)
//}
