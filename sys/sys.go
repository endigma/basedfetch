package sys

import (
	"fmt"
	"os/user"
	"runtime"

	"github.com/jaypipes/ghw"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type System struct {
	User        string
	Hostname    string
	OS          string
	Arch        string
	Kernel      string
	CPUVendor   string
	CPUModel    string
	GPUVendor   string
	GPUModel    string
	MemoryUsed  uint64
	MemoryTotal uint64
}

func Fetch() (System, error) {

	hostStat, err := host.Info()
	if err != nil {
		return System{}, err
	}
	cpuStat, err := cpu.Info()
	if err != nil {
		return System{}, err
	}
	memStat, err := mem.VirtualMemory()
	if err != nil {
		return System{}, err
	}
	gpuStat, err := ghw.GPU()
	if err != nil {
		return System{}, err
	}

	user, err := user.Current()
	if err != nil {
		return System{}, err
	}

	var s System

	s.User = user.Username
	s.Hostname = hostStat.Hostname
	s.MemoryTotal = memStat.Total / 1e+6
	s.MemoryUsed = memStat.Used / 1e+6
	s.Arch = hostStat.KernelArch
	s.OS = runtime.GOOS
	s.Kernel = hostStat.KernelVersion
	s.CPUVendor = cpuStat[0].VendorID
	s.CPUModel = fmt.Sprintf("%s ", cpuStat[0].ModelName)
	s.GPUVendor = gpuStat.GraphicsCards[0].DeviceInfo.Vendor.Name
	s.GPUModel = gpuStat.GraphicsCards[0].DeviceInfo.Product.Name

	return s, nil
}
