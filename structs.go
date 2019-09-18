package main

import "fmt"

// Monitor - struct for defining a monitor
type Monitor struct {
	Resolution string
	Connector  string
	Value      float64
}

func main() {
	monitor := Monitor{Value: 240.99, Resolution: "1080p", Connector: "HDMI"}
	showFields(monitor)
}

func showFields(m Monitor) {
	fmt.Println("Resolution: ", m.Resolution, " Connector: ", m.Connector, " Value: ", m.Value)
}
