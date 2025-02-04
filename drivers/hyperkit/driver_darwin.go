package hyperkit

import (
	"github.com/densityops/machine/libmachine/drivers"
)

const (
	defaultMemory = 8192
	defaultCPU    = 4
)

type Driver struct {
	*drivers.VMDriver
	Cmdline       string
	UUID          string
	VpnKitSock    string
	VpnKitUUID    string
	VSockPorts    []string
	VmlinuzPath   string
	InitrdPath    string
	KernelCmdLine string
	HyperKitPath  string
	QcowToolPath  string
	VMNet         bool
	BootromPath   string
}

func NewDriver(hostName, storePath string) *Driver {
	return &Driver{
		VMDriver: &drivers.VMDriver{
			BaseDriver: &drivers.BaseDriver{
				MachineName: hostName,
				StorePath:   storePath,
			},
			Memory: defaultMemory,
			CPU:    defaultCPU,
		},
	}
}
