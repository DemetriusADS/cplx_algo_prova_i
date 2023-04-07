package machine

import (
	"fmt"
	"time"

	"github.com/DemetriusADS/cplx_algo_prova_i/machine/ports/sensor"
)

type Metric struct {
	Temperature sensor.SensorDTO
	Volume      sensor.SensorDTO
}

type Machine struct {
	Metrics []*Metric
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

		m.Metrics = append(m.Metrics, &Metric{
			Temperature: *temp,
			Volume:      *vol,
		})
		count++
		if count == 10 {
			break
		}
	}
}

func (m *Machine) Read() []*Metric {
	return m.Metrics
}

func (m *Machine) FixTemperature() {
	fmt.Printf("CALIBRANDO A MAQUINA %s PARA A TEMPERATURA IDEAL\n", m.Name)
	for _, metric := range m.Read() {
		now := time.Now().Format("2006-01-02 15:04:05")
		newTemp := metric.Volume.Value * 2.5
		if newTemp < 100 {
			newTemp = 100
		}
		metric.Temperature.Value = newTemp
		metric.Temperature.Time = now
		metric.Volume.Time = now
	}
}

// Aqui escolhi utilizar o algoritmo de ordenação Bubble Sort, pois ele é simples e fácil de implementar.
// O Bubble Sort é um algoritmo de ordenação simples que percorre o array várias vezes, comparando elementos adjacentes e os trocando de posição se estiverem na ordem errada.
// Sua complexidade é O(n²) por conter dois "for" aninhados.
func (m *Machine) BubbleSortDescending() {
	n := len(m.Metrics)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			date1, _ := time.Parse("2006-01-02 15:04:05", m.Metrics[j].Temperature.Time)
			date2, _ := time.Parse("2006-01-02 15:04:05", m.Metrics[j+1].Temperature.Time)

			if date1.Before(date2) {
				m.Metrics[j], m.Metrics[j+1] = m.Metrics[j+1], m.Metrics[j]
			}
		}
	}
}
