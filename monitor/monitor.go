package monitor

import (
	"fmt"

	"github.com/DemetriusADS/cplx_algo_prova_i/machine"
	"github.com/DemetriusADS/cplx_algo_prova_i/machine/sensors/temperature"
	"github.com/DemetriusADS/cplx_algo_prova_i/machine/sensors/volume"
)

type Monitor struct {
	maquinas [10]*machine.Machine
}

func NewMonitor() *Monitor {
	maquinas := [10]*machine.Machine{}
	for i := 0; i < 10; i++ {
		volumeSensor := volume.NewVolumeSensor()
		temperatureSensor := temperature.NewTemperatureSensor()
		newMachine := machine.NewMachine("machine_"+fmt.Sprint(i), volumeSensor, temperatureSensor)
		maquinas[i] = newMachine
	}
	return &Monitor{
		maquinas: maquinas,
	}
}

func (m *Monitor) Start() {
	var command string
	for {
		m.menu()
		fmt.Scanln(&command)
		fmt.Print("\033[H\033[2J")
		switch command {
		case "1":
			fmt.Println("Listando maquinas...")
			m.listarMaquinas()
			continue
		case "2":
			fmt.Println("Listando as métricas coletadas...")
			m.listarMetricas()
			continue
		case "3":
			fmt.Println("Ordenando as métricas por data (DESC)...")
			m.ordenarMaquinas()
		case "4":
			m.calibrarMaquinas()
			continue
		case "5":
			fmt.Println("Saindo...")
			return
		}
	}
}

func (m *Monitor) menu() {
	fmt.Println("\nDigite o comando desejado:")
	fmt.Println("1 - Listar maquinas")
	fmt.Println("2 - Listar as métricas coletadas")
	fmt.Println("3 - Ordenar as métricas por data (DESC)")
	fmt.Println("4 - Corrigir a temperatura das maquinas")
	fmt.Println("5 - Sair")
}

func (m *Monitor) listarMaquinas() {
	fmt.Printf("%-10s %-15s\n", "Nome", "Está Online?")
	for _, maquina := range m.maquinas {
		fmt.Printf("%-10s %-15v\n", maquina.Name, maquina.IsOn)
	}
}

func (m *Monitor) listarMetricas() {
	for _, maquina := range m.maquinas {

		fmt.Printf("\n\nMaquina: %s\n", maquina.Name)
		fmt.Printf("%-10s | %-15v | %-15s | %-15v\n", "Temperatura (C)", "Timestamp", "Volume (L)", "Timestamp")
		metrics := maquina.Read()
		for _, metric := range metrics {
			fmt.Printf("%-10f | %-15v | %-15f | %-15v\n", metric.Temperature.Value, metric.Temperature.Time, metric.Volume.Value, metric.Volume.Time)
		}
	}
}

func (m *Monitor) calibrarMaquinas() {
	for _, maquina := range m.maquinas {
		maquina.FixTemperature()
	}
}

func (m *Monitor) ordenarMaquinas() {
	for _, maquina := range m.maquinas {
		maquina.BubbleSortDescending()
	}
}

/**
* TODO List:
* 1 - Inicializar 10 maquinas
* 2 - Gerar um menu interativo para o usuario
* 3 - Menu deve conter: Listar maquinas, Listas as métricas de cada maquina monitorada, ordenacao crescente dos dados das maquinas
* por fim, um algoritmo O(N^2) ou N^3.
**/
