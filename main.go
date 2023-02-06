package main

import (
	"fmt"
	"sim/config"
	"sync"
	"time"
)

func main() {

	senderChan := make(chan []byte, 10)

	s := NewSender(config.Config.Dest.IP, config.Config.Dest.Port)

	var wg sync.WaitGroup
	wg.Add(2)
	go s.Start(senderChan, &wg)
	go createPackets(config.Config.Runtime.PacketsToSend, senderChan, &wg)
	wg.Wait()

}

func createPackets(packetsToSend int, senderChan chan<- []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Creating %d packets...\n", packetsToSend)
	for i := 0; i < packetsToSend; i++ {
		byteData := generateTwoCyclePacketBytes(uint32(i + 1000))
		time.Sleep(33 * time.Millisecond)
		senderChan <- byteData
	}
	fmt.Printf("Done creating %d packets.\n", packetsToSend)
	close(senderChan)

}

func plotSinCosWave() {
	// Number of samples to generate
	numSamples := 64

	// Sampling frequency
	samplingFreq := 1920.0

	// Wave frequency
	freq := 60.0

	// Amplitude of sin wave
	amplitude := 2.0

	// Generate sin wave
	sinWave := generateSinWave(numSamples, samplingFreq, freq, amplitude)

	// Generate cos wave
	cosWave := generateCosWave(numSamples, samplingFreq, freq, amplitude)

	plotFloat64(sinWave, "sin.png")
	plotFloat64(cosWave, "cos.png")
}
