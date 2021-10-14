package services

import (
	"fmt"
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
