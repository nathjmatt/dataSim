package main

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

}
