package main

import "fmt"

type Computer interface {
	InsertIntoLightningPort()
}

func AbstractInsertIntoLightingPort(comp Computer) {
	comp.InsertIntoLightningPort()
}

type Mac struct{}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

type WindowsAdapter struct {
	win *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.win.insertIntoUSBPort()
}

func main() {
	mac := &Mac{}
	win := &Windows{}
	winAdaper := &WindowsAdapter{win: win}

	AbstractInsertIntoLightingPort(mac)
	AbstractInsertIntoLightingPort(winAdaper)
}
