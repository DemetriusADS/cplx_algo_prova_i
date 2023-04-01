package machine

import (
	"time"

	"github.com/DemetriusADS/cplx_algo_prova_i/machine/ports/sensor"
)

type Metric struct {
	Temperature sensor.SensorDTO
	Volume      sensor.SensorDTO
}

type Machine struct {
	Metrics []Metric
	Name    string
	IsOn    bool

	VolumeSensor      sensor.Sensor
	TemperatureSensor sensor.Sensor
}

func NewMachine(name string, volumeSensor, temperatureSensor sensor.Sensor) *Machine {
	machine := Machine{
		Name:              name,
		IsOn:              true,
		VolumeSensor:      volumeSensor,
		TemperatureSensor: temperatureSensor,
	}

	go machine.genData()
	return &machine
}

func (m *Machine) genData() {
	count := 0
	for {
		time.Sleep(1 * time.Second)

		temp := m.TemperatureSensor.Read()
		vol := m.VolumeSensor.Read()

		m.Metrics = append(m.Metrics, Metric{
			Temperature: *temp,
			Volume:      *vol,
		})
		count++
		if count == 10 {
			break
		}
	}
}

func (m *Machine) Read() []Metric {
	return m.Metrics
}
