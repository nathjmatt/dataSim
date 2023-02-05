package main

import "math"

//	Generates and returns a slice that represents a sin wave.
//
// numSamples: is the number of samples to generate
//
// samplingFreq: sampling frequency of the device
//
// freq: frequency of the sin wave
//
// amplitude: amplitude of the sin wave
func generateSinWave(numSamples int, samplingFreq, freq, amplitude float64) []float64 {

	// Stores sin wave samples
	sinWave := make([]float64, numSamples)

	// Time step between each sample
	dt := 1.0 / samplingFreq

	// Current time of the sample
	t := dt

	for i := 0; i < numSamples; i++ {
		// Current time
		t = dt * float64(i)

		// Generate and store sample
		sinWave[i] = math.Sin(2*math.Pi*freq*t) * amplitude
	}

	return sinWave
}

//	Generates and returns a slice that represents a cos wave.
//
// numSamples: is the number of samples to generate
//
// samplingFreq: sampling frequency of the device
//
// freq: frequency of the cos wave
//
// amplitude: amplitude of the cos wave
func generateCosWave(numSamples int, samplingFreq, freq, amplitude float64) []float64 {

	// Stores cos wave samples
	cosWave := make([]float64, numSamples)

	// Time step between each sample
	dt := 1.0 / samplingFreq

	// Current time of the sample
	t := dt

	for i := 0; i < numSamples; i++ {
		// Current time
		t = dt * float64(i)

		// Generate and store sample
		cosWave[i] = math.Cos(2*math.Pi*freq*t) * amplitude
	}

	return cosWave
}
