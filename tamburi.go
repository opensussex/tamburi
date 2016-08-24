package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/opensussex/wave_tamburi"
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
	if len(os.Args) > 1 {
		if len(os.Args) < 8 {
			fmt.Println(`Not Enough Arguments!`)
			fmt.Println(`example: 4 2 200.0 0.022 1.0 0.8998 0.3`)
			os.Exit(1)
		}
		arg1, _ := strconv.ParseFloat(os.Args[1], 64)
		arg2, _ := strconv.ParseFloat(os.Args[2], 64)
		arg3, _ := strconv.ParseFloat(os.Args[3], 64)
		arg4, _ := strconv.ParseFloat(os.Args[4], 64)
		arg5, _ := strconv.ParseFloat(os.Args[5], 64)
		arg6, _ := strconv.ParseFloat(os.Args[6], 64)
		arg7, _ := strconv.ParseFloat(os.Args[7], 64)
		file.AppendByte(DrumSynth(arg1, arg2, arg3, arg4, arg5, arg6, arg7))
	} else {
		file.AppendByte(DrumSynth(4, 2, 200.0, 0.022, 1.0, 0.8998, 0.3))
	}

	err = file.Write(fd)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Println(`Drum sound written to - test.wav`)
	}
}

func DrumSynth(amp, decay, freq, freqdecay, noise, noisedecay, noisefilter float64) []byte {
	var i, datalen uint32
	var byte_data []byte
	datalen = 44100

	rand.Seed(time.Now().UTC().UnixNano())
	var f float64
	for i = 0; i < datalen; i++ {
		data := amp * math.Sin(2.0*math.Pi*freq*float64(i)/float64(datalen))
		f = (noisefilter)*f + (1.0-noisefilter)*noise*(float64(rand.Intn(100))/100.0-0.5)
		data += f

		noise *= noisedecay
		amp *= decay
		freq -= freqdecay

		if noise < .0 {
			noise = .0
		}
		if amp < .0 {
			amp = .0
		}
		if freq < .0 {
			freq = .0
		}
		if data < -1.0 {
			data = -1.0
		}
		if data > 1.0 {
			data = 1.0
		}
		byte_data = append(byte_data, byte(data))
	}
	return byte_data
}
