package services

import (
	"fmt"
	"regexp"
	"strings"
)

type Vm struct {
	ID    string
	Name  string
	State string
}

func ListRunningVms() ([]Vm, error) {
	vms := make([]Vm, 0, 3)
	out, err := CmdExecStrOutput("virsh", "-c", "qemu:///system", "list")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(out, "\n")
	for i, l := range lines {
		if i <= 1 {
			continue
		}
		items := strings.Fields(l)
		if len(items) == 3 {
			vm := Vm{
				ID:    strings.TrimSpace(items[0]),
				Name:  strings.TrimSpace(items[1]),
				State: strings.TrimSpace(items[2]),
			}
			vms = append(vms, vm)
		}

	}
	fmt.Println(vms)
	return vms, nil
}

type UsbDevice struct {
	ID     string
	Name   string
	Bus    string
	Device string
}

func ListUsbs() ([]UsbDevice, error) {
	usbs := make([]UsbDevice, 0, 3)
	out, err := CmdExecStrOutput("lsusb")
	if err != nil {
		return nil, err
	}
	reg := regexp.MustCompile(`Bus (\d+) Device (\d+): ID (\S+) (.*)`)
	lines := strings.Split(out, "\n")
	for _, l := range lines {
		groups := reg.FindStringSubmatch(l)
		if len(groups) > 0 {
			usb := UsbDevice{
				Bus:    groups[1],
				Device: groups[2],
				ID:     groups[3],
				Name:   groups[4],
			}
			usbs = append(usbs, usb)
		}
	}
	return usbs, nil
}
