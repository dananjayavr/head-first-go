package main

import (
	"fmt"
	"github.com/dananjayavr/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record() // <- type assertion here
	} else {
		fmt.Println("Attempting to Record...")
		fmt.Println("Player was not a TapeRecorder")
	}

}
func main() {
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	var player Player = gadget.TapePlayer{}
	fmt.Println("Test #1")
	TryOut(player)
	playList(player, mixtape)
	player = gadget.TapeRecorder{}
	fmt.Println("Test #2")
	TryOut(player)
	playList(player, mixtape)

	//var testDevice gadget.TapeRecorder = player.(gadget.TapeRecorder)
	//TryOut(testDevice)
	//TryOut(player)
}
