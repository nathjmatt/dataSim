package main

import "time"

func main() {

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

	genAndSendPackets()

	time.Sleep(3 * time.Second)
}

func genAndSendPackets() {
	sendingIPAddr := "127.0.0.1"
	portNumber := uint(8080)
	s := NewSender(sendingIPAddr, portNumber)

	senderChan := make(chan []byte, 10)
	defer close(senderChan)
	go s.Start(senderChan)

	for i := 0; i < 30; i++ {
		byteData := generateTwoCyclePacketBytes()
		time.Sleep(33 * time.Millisecond)
		senderChan <- byteData
	}

}
