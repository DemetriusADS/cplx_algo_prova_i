package main

import (
	"fmt"
	"time"

	"github.com/DemetriusADS/cplx_algo_prova_i/machine"
	"github.com/DemetriusADS/cplx_algo_prova_i/machine/sensors/temperature"
	"github.com/DemetriusADS/cplx_algo_prova_i/machine/sensors/volume"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	// text, _ := reader.ReadString('\n')
	// fmt.Println(text)

	tempSensor := temperature.NewTemperatureSensor()
	volSensor := volume.NewVolumeSensor()

	m := machine.NewMachine("machine_1", volSensor, tempSensor)

	time.Sleep(10 * time.Second)

	currentData := m.Read()

	for _, data := range currentData {
		fmt.Printf("{Temperature: %2f Timestamp: %s}\n", data.Temperature.Value, data.Temperature.Time)
		fmt.Printf("{Volume: %2f Timestamp: %s}\n", data.Volume.Value, data.Volume.Time)
	}
}
