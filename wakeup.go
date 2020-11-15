package main

import (
	"log"
	"os"
	"time"

	//	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	ogg "github.com/faiface/beep/vorbis"
)

func main() {
	filename := "./data/alarm.ogg"
	music, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer music.Close()
	streamer, format, err := ogg.Decode(music)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
	select {}

	/* Need to fix sample rate
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
	*/

}
