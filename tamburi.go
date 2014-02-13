package main

import (
	"github.com/opensussex/wave_tamburi"
	"os"
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
	file.AppendDrum(2, 0.9995, 70.0, 0.002, 1.0, 0.9998, 0.5);
	err = file.Write(fd)
	if err != nil {
		println(err.Error())
	}
}
