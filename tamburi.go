package main

import (
	"github.com/opensussex/wave_tamburi"
	"os"
	"time"
	"math"
	"math/rand"
)

const (
	PCM = 1
)

func main() {
	fd, err := os.OpenFile("test.wav", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		println(err.Error())
		return
	}
	defer fd.Close()

	file := wave.CreateFile(PCM, 1, 44100, 8)
	//file.AppendSine(440, 5000, 20)
	file.AppendByte(DrumSynth(2, 0.9995, 70.0, 0.002, 1.0, 0.9998, 0.5));
	err = file.Write(fd)
	if err != nil {
		println(err.Error())
	}
}


func DrumSynth( amp, decay, freq, freqdecay, noise, noisedecay, noisefilter float64 ) []byte{
	var i, datalen uint32
	var byte_data []byte
	datalen = 44100

	rand.Seed( time.Now().UTC().UnixNano())
	var f float64
	for i = 0; i < datalen; i++ {
		data := amp * math.Sin( 2.0 * math.Pi * freq * float64(i) / float64(datalen) )
		f = (noisefilter) * f + (1.0-noisefilter) * noise * ( float64( rand.Intn(100))  / 100.0 - 0.5)
		data += f

		noise *= noisedecay
		amp *= decay
		freq -= freqdecay

		if( noise < .0 ){
			noise = .0
		}
		if( amp < .0 ){
			amp = .0
		}
		if( freq < .0 ){
			freq = .0
		}
		if( data < -1.0 ){
			data = -1.0
		}
		if( data >  1.0 ) {
			data =  1.0
		}	
		byte_data = append(byte_data,byte(data))
	}
	return byte_data
}