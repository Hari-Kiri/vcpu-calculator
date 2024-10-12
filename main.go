package main

import (
	"fmt"

	goVirtQemuConnector "github.com/Hari-Kiri/govirt-qemu-connector"
	"github.com/Hari-Kiri/temboLog"
	"libvirt.org/go/libvirt"
	"libvirt.org/go/libvirtxml"
)

func main() {
	qemuConnection, errorConnectToQemu := goVirtQemuConnector.ConnectToLocalSystem()
	libvirtError, isError := errorConnectToQemu.(libvirt.Error)
	if isError {
		temboLog.FatalLogging("failed connect to qemu:", libvirtError.Message)
	}
	defer qemuConnection.Close()

	capabilitiesXml, errorGetCapabilitiesXml := qemuConnection.GetCapabilities()
	libvirtError, isError = errorGetCapabilitiesXml.(libvirt.Error)
	if isError {
		temboLog.FatalLogging("failed get host capabilities:", libvirtError.Message)
	}

	capabilities := &libvirtxml.Caps{}
	libvirtError, isError = capabilities.Unmarshal(capabilitiesXml).(libvirt.Error)
	if isError {
		temboLog.FatalLogging("failed marshall capabilities to JSON", libvirtError.Message)
	}

	threadsPerCore := capabilities.Host.CPU.Topology.Threads
	coresPerPhysicalCpu := capabilities.Host.CPU.Topology.Cores
	threadsPerPhysicalCpu := coresPerPhysicalCpu * threadsPerCore
	physicalCpuAttached := capabilities.Host.CPU.Topology.Sockets

	fmt.Print((threadsPerPhysicalCpu * coresPerPhysicalCpu) * physicalCpuAttached)
}
