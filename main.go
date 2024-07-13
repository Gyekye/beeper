package main

import (
	
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	fmt.Printf("Playing sound for %v\n", scanner.Bytes())
	// 	go player()
	// }
	// fmt.Println("Scanner done")
	// strokesBuf := make(chan bool, len(Strokes))
	keysEvents, err := keyboard.GetKeys(100)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	fmt.Println("Press ESC to quit Beeper")
	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}
		go player()
		if event.Key == keyboard.KeyEsc {
			break
		}
	}
}

func player() {
	f, err := os.Open("./clicks/0.wav")
	if err != nil {
		log.Fatal(err)
	}
	
	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/100))
	speaker.Play(streamer)
}