package monitor

import (
	"fmt"
	"testing"
	"time"
)

func TestNewMonitor(t *testing.T) {
	monitor := NewMonitor()
	if monitor == nil {
		t.Error("Expected a monitor")
	}
}

func TestMachineList(t *testing.T) {
	monitor := NewMonitor()
	if len(monitor.maquinas) != 10 {
		t.Error("Expected 10 machines")
	}
}

func TestMachineListNames(t *testing.T) {
	monitor := NewMonitor()
	for i, maquina := range monitor.maquinas {
		if maquina.Name != "machine_"+fmt.Sprint(i) {
			t.Error("Expected machine_", i)
		}
	}
}

func TestMachineListOnline(t *testing.T) {
	monitor := NewMonitor()
	for _, maquina := range monitor.maquinas {
		if maquina.IsOn != true {
			t.Error("Expected machine online")
		}
	}
}

func TestMachineListVolume(t *testing.T) {
	monitor := NewMonitor()
	time.Sleep(10 * time.Second)
	for _, maquina := range monitor.maquinas {
		if maquina.VolumeSensor == nil {
			t.Error("Expected volume sensor")
		}
	}
}

func TestMachineListTemperature(t *testing.T) {
	monitor := NewMonitor()
	time.Sleep(10 * time.Second)
	for _, maquina := range monitor.maquinas {
		if maquina.TemperatureSensor == nil {
			t.Error("Expected temperature sensor")
		}
	}
}

func TestMachineListMetrics(t *testing.T) {
	monitor := NewMonitor()
	time.Sleep(11 * time.Second)
	for _, maquina := range monitor.maquinas {
		if len(maquina.Metrics) != 10 {
			t.Error("Expected 10 metrics")
		}

		for _, metric := range maquina.Metrics {
			if metric.Temperature.Value == 0 {
				t.Error("Expected temperature value")
			}
			if metric.Volume.Value == 0 {
				t.Error("Expected volume value")
			}
		}
	}
}

func TestCorrigirTemperatura(t *testing.T) {
	monitor := NewMonitor()
	time.Sleep(11 * time.Second)
	monitor.calibrarMaquinas()
	for _, maquina := range monitor.maquinas {
		for _, metric := range maquina.Metrics {
			if metric.Temperature.Value < 100 {
				t.Error("Expected temperature value less than 100")
			}
		}
	}
}
