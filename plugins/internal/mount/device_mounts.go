//go:build windows

package mount

import (
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"

	"github.com/tensorworks/directx-device-plugins/plugins/internal/discovery"
)

// Generates the device specs for the supplied list of devices
func SpecsForDevices(devices []*discovery.Device) []*pluginapi.DeviceSpec {

	specs := []*pluginapi.DeviceSpec{}

	// class:// GUID_DEVINTERFACE_DISPLAY_ADAPTER
	for range devices {
		specs = appendUniqueSpec(specs, &pluginapi.DeviceSpec{
			HostPath:      "class://5B45201D-F2F2-4F3B-85BB-30FF1F953599",
			ContainerPath: "",
			Permissions:   "",
		})
	}

	return specs
}

// Generates the runtime file mounts for the supplied list of devices
func MountsForDevices(devices []*discovery.Device) []*pluginapi.Mount {
	return []*pluginapi.Mount{}
}

// Appends a device spec to an existing list of device specs if it's not already present in the list
func appendUniqueSpec(specs []*pluginapi.DeviceSpec, newSpec *pluginapi.DeviceSpec) []*pluginapi.DeviceSpec {
	for _, existing := range specs {
		if existing.HostPath == newSpec.HostPath {
			return specs
		}
	}

	return append(specs, newSpec)
}

// Appends a mount to an existing list of mounts if it's not already present in the list
func appendUniqueMount(mounts []*pluginapi.Mount, newMount *pluginapi.Mount) []*pluginapi.Mount {
	for _, existing := range mounts {
		if existing.ContainerPath == newMount.ContainerPath {
			return mounts
		}
	}

	return append(mounts, newMount)
}
