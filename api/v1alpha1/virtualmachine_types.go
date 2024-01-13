/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VirtualMachineSpec defines the desired state of VirtualMachine

type VirtualMachineSpec struct {
	XMLName        xmlName         `xml:"domain"  json:"XMLName"`
	Type           string          `xml:"type,attr"  json:"Type,omitempty"`
	XmlNS          string          `xml:"xmlns:qemu,attr,omitempty" json:"XmlNS"`
	Name           string          `xml:"name" json:"Name,omitempty"`
	UUID           string          `xml:"uuid,omitempty" json:"UUID,omitempty"`
	Memory         Memory          `xml:"memory" json:"Memory"`
	CurrentMemory  *Memory         `xml:"currentMemory,omitempty" json:"CurrentMemory,omitempty"`
	MaxMemory      *MaxMemory      `xml:"maxMemory,omitempty" json:"MaxMemory,omitempty"`
	MemoryBacking  *MemoryBacking  `xml:"memoryBacking,omitempty" json:"MemoryBacking,omitempty"`
	OS             OS              `xml:"os" json:"OS"`
	SysInfo        *SysInfo        `xml:"sysinfo,omitempty" json:"SysInfo,omitempty"`
	Devices        Devices         `xml:"devices" json:"Devices"`
	Clock          *Clock          `xml:"clock,omitempty" json:"Clock,omitempty"`
	Resource       *Resource       `xml:"resource,omitempty" json:"Resource,omitempty"`
	QEMUCmd        *Commandline    `xml:"qemu:commandline,omitempty" json:"QEMUCmd,omitempty"`
	Features       *Features       `xml:"features,omitempty" json:"Features,omitempty"`
	CPU            CPU             `xml:"cpu" json:"CPU"`
	VCPU           *VCPU           `xml:"vcpu" json:"VCPU,omitempty"`
	VCPUs          *VCPUs          `xml:"vcpus" json:"VCPUs,omitempty"`
	CPUTune        *CPUTune        `xml:"cputune" json:"CPUTune,omitempty"`
	NUMATune       *NUMATune       `xml:"numatune" json:"NUMATune,omitempty"`
	IOThreads      *IOThreads      `xml:"iothreads,omitempty" json:"IOThreads,omitempty"`
	LaunchSecurity *LaunchSecurity `xml:"launchSecurity,omitempty" json:"LaunchSecurity,omitempty"`
}

type CPUTune struct {
	VCPUPin     []CPUTuneVCPUPin     `xml:"vcpupin" json:"VCPUPin,omitempty"`
	IOThreadPin []CPUTuneIOThreadPin `xml:"iothreadpin,omitempty" json:"IOThreadPin,omitempty"`
	EmulatorPin *CPUEmulatorPin      `xml:"emulatorpin" json:"EmulatorPin,omitempty"`
}

type NUMATune struct {
	Memory   NumaTuneMemory `xml:"memory" json:"Memory"`
	MemNodes []MemNode      `xml:"memnode" json:"MemNodes,omitempty"`
}

type MemNode struct {
	CellID  uint32 `xml:"cellid,attr" json:"CellID,omitempty"`
	Mode    string `xml:"mode,attr" json:"Mode,omitempty"`
	NodeSet string `xml:"nodeset,attr" json:"NodeSet,omitempty"`
}

type NumaTuneMemory struct {
	Mode    string `xml:"mode,attr" json:"Mode,omitempty"`
	NodeSet string `xml:"nodeset,attr" json:"NodeSet,omitempty"`
}

type CPUTuneVCPUPin struct {
	VCPU   uint32 `xml:"vcpu,attr" json:"VCPU,omitempty"`
	CPUSet string `xml:"cpuset,attr" json:"CPUSet,omitempty"`
}

type CPUTuneIOThreadPin struct {
	IOThread uint32 `xml:"iothread,attr" json:"IOThread,omitempty"`
	CPUSet   string `xml:"cpuset,attr" json:"CPUSet,omitempty"`
}

type CPUEmulatorPin struct {
	CPUSet string `xml:"cpuset,attr" json:"CPUSet,omitempty"`
}

type VCPU struct {
	Placement string `xml:"placement,attr" json:"Placement,omitempty"`
	CPUs      uint32 `xml:",chardata" json:"CPUs,omitempty"`
}

type VCPUsVCPU struct {
	ID           uint32 `xml:"id,attr" json:"ID,omitempty"`
	Enabled      string `xml:"enabled,attr,omitempty" json:"Enabled,omitempty"`
	Hotpluggable string `xml:"hotpluggable,attr,omitempty" json:"Hotpluggable,omitempty"`
	Order        uint32 `xml:"order,attr,omitempty" json:"Order,omitempty"`
}

type VCPUs struct {
	VCPU []VCPUsVCPU `xml:"vcpu" json:"VCPU,omitempty"`
}

type CPU struct {
	Mode     string       `xml:"mode,attr,omitempty" json:"Mode,omitempty"`
	Model    string       `xml:"model,omitempty" json:"Model,omitempty"`
	Features []CPUFeature `xml:"feature" json:"Features,omitempty"`
	Topology *CPUTopology `xml:"topology" json:"Topology,omitempty"`
	NUMA     *NUMA        `xml:"numa,omitempty" json:"NUMA,omitempty"`
}

type NUMA struct {
	Cells []NUMACell `xml:"cell" json:"Cells,omitempty"`
}

type NUMACell struct {
	ID           string `xml:"id,attr" json:"ID,omitempty"`
	CPUs         string `xml:"cpus,attr" json:"CPUs,omitempty"`
	Memory       uint64 `xml:"memory,attr,omitempty" json:"Memory,omitempty"`
	Unit         string `xml:"unit,attr,omitempty" json:"Unit,omitempty"`
	MemoryAccess string `xml:"memAccess,attr,omitempty" json:"MemoryAccess,omitempty"`
}

type CPUFeature struct {
	Name   string `xml:"name,attr" json:"Name,omitempty"`
	Policy string `xml:"policy,attr,omitempty" json:"Policy,omitempty"`
}

type CPUTopology struct {
	Sockets uint32 `xml:"sockets,attr,omitempty" json:"Sockets,omitempty"`
	Cores   uint32 `xml:"cores,attr,omitempty" json:"Cores,omitempty"`
	Threads uint32 `xml:"threads,attr,omitempty" json:"Threads,omitempty"`
}

type Features struct {
	ACPI       *FeatureEnabled    `xml:"acpi,omitempty" json:"ACPI,omitempty"`
	APIC       *FeatureEnabled    `xml:"apic,omitempty" json:"APIC,omitempty"`
	Hyperv     *FeatureHyperv     `xml:"hyperv,omitempty" json:"Hyperv,omitempty"`
	SMM        *FeatureEnabled    `xml:"smm,omitempty" json:"SMM,omitempty"`
	KVM        *FeatureKVM        `xml:"kvm,omitempty" json:"KVM,omitempty"`
	PVSpinlock *FeaturePVSpinlock `xml:"pvspinlock,omitempty" json:"PVSpinlock,omitempty"`
	PMU        *FeatureState      `xml:"pmu,omitempty" json:"PMU,omitempty"`
}

type FeatureHyperv struct {
	Relaxed         *FeatureState     `xml:"relaxed,omitempty" json:"Relaxed,omitempty"`
	VAPIC           *FeatureState     `xml:"vapic,omitempty" json:"VAPIC,omitempty"`
	Spinlocks       *FeatureSpinlocks `xml:"spinlocks,omitempty" json:"Spinlocks,omitempty"`
	VPIndex         *FeatureState     `xml:"vpindex,omitempty" json:"VPIndex,omitempty"`
	Runtime         *FeatureState     `xml:"runtime,omitempty" json:"Runtime,omitempty"`
	SyNIC           *FeatureState     `xml:"synic,omitempty" json:"SyNIC,omitempty"`
	SyNICTimer      *SyNICTimer       `xml:"stimer,omitempty" json:"SyNICTimer,omitempty"`
	Reset           *FeatureState     `xml:"reset,omitempty" json:"Reset,omitempty"`
	VendorID        *FeatureVendorID  `xml:"vendor_id,omitempty" json:"VendorID,omitempty"`
	Frequencies     *FeatureState     `xml:"frequencies,omitempty" json:"Frequencies,omitempty"`
	Reenlightenment *FeatureState     `xml:"reenlightenment,omitempty" json:"Reenlightenment,omitempty"`
	TLBFlush        *FeatureState     `xml:"tlbflush,omitempty" json:"TLBFlush,omitempty"`
	IPI             *FeatureState     `xml:"ipi,omitempty" json:"IPI,omitempty"`
	EVMCS           *FeatureState     `xml:"evmcs,omitempty" json:"EVMCS,omitempty"`
}

type FeatureSpinlocks struct {
	State   string  `xml:"state,attr,omitempty" json:"State,omitempty"`
	Retries *uint32 `xml:"retries,attr,omitempty" json:"Retries,omitempty"`
}

type SyNICTimer struct {
	Direct *FeatureState `xml:"direct,omitempty" json:"Direct,omitempty"`
	State  string        `xml:"state,attr,omitempty" json:"State,omitempty"`
}

type FeaturePVSpinlock struct {
	State string `xml:"state,attr,omitempty" json:"State,omitempty"`
}

type FeatureVendorID struct {
	State string `xml:"state,attr,omitempty" json:"State,omitempty"`
	Value string `xml:"value,attr,omitempty" json:"Value,omitempty"`
}

type FeatureEnabled struct {
}

type Shareable struct{}

type FeatureState struct {
	State string `xml:"state,attr,omitempty" json:"State,omitempty"`
}

type FeatureKVM struct {
	Hidden        *FeatureState `xml:"hidden,omitempty" json:"Hidden,omitempty"`
	HintDedicated *FeatureState `xml:"hint-dedicated,omitempty" json:"HintDedicated,omitempty"`
}

type AccessCredentialMetadata struct {
	Succeeded bool   `xml:"succeeded,omitempty"`
	Message   string `xml:"message,omitempty"`
}

type MemoryDumpMetadata struct {
	FileName       string       `xml:"fileName,omitempty"`
	StartTimestamp *metav1.Time `xml:"startTimestamp,omitempty"`
	EndTimestamp   *metav1.Time `xml:"endTimestamp,omitempty"`
	Completed      bool         `xml:"completed,omitempty"`
	Failed         bool         `xml:"failed,omitempty"`
	FailureReason  string       `xml:"failureReason,omitempty"`
}

type GracePeriodMetadata struct {
	DeletionGracePeriodSeconds int64        `xml:"deletionGracePeriodSeconds"`
	DeletionTimestamp          *metav1.Time `xml:"deletionTimestamp,omitempty"`
	MarkedForGracefulShutdown  *bool        `xml:"markedForGracefulShutdown,omitempty"`
}

type Commandline struct {
	QEMUEnv []Env `xml:"qemu:env,omitempty" json:"QEMUEnv,omitempty"`
	QEMUArg []Arg `xml:"qemu:arg,omitempty" json:"QEMUArg,omitempty"`
}

type Env struct {
	Name  string `xml:"name,attr" json:"Name,omitempty"`
	Value string `xml:"value,attr" json:"Value,omitempty"`
}

type Arg struct {
	Value string `xml:"value,attr" json:"Value,omitempty"`
}

type Resource struct {
	Partition string `xml:"partition" json:"Partition,omitempty"`
}

type Memory struct {
	Value uint64 `xml:",chardata" json:"Value,omitempty"`
	Unit  string `xml:"unit,attr" json:"Unit,omitempty"`
}

type MaxMemory struct {
	Value uint64 `xml:",chardata" json:"Value,omitempty"`
	Unit  string `xml:"unit,attr" json:"Unit,omitempty"`
	Slots uint64 `xml:"slots,attr" json:"Slots,omitempty"`
}

// MemoryBacking mirroring libvirt XML under https://libvirt.org/formatdomain.html#elementsMemoryBacking
type MemoryBacking struct {
	HugePages    *HugePages           `xml:"hugepages,omitempty" json:"HugePages,omitempty"`
	Source       *MemoryBackingSource `xml:"source,omitempty" json:"Source,omitempty"`
	Access       *MemoryBackingAccess `xml:"access,omitempty" json:"Access,omitempty"`
	Allocation   *MemoryAllocation    `xml:"allocation,omitempty" json:"Allocation,omitempty"`
	NoSharePages *NoSharePages        `xml:"nosharepages,omitempty" json:"NoSharePages,omitempty"`
}

type MemoryAllocationMode string

const (
	MemoryAllocationModeImmediate MemoryAllocationMode = "immediate"
)

type MemoryAllocation struct {
	Mode MemoryAllocationMode `xml:"mode,attr" json:"Mode,omitempty"`
}

type MemoryBackingSource struct {
	Type string `xml:"type,attr" json:"Type,omitempty"`
}

// HugePages mirroring libvirt XML under memoryBacking
type HugePages struct {
	HugePage []HugePage `xml:"page,omitempty" json:"HugePage,omitempty"`
}

// HugePage mirroring libvirt XML under hugepages
type HugePage struct {
	Size    string `xml:"size,attr" json:"Size,omitempty"`
	Unit    string `xml:"unit,attr" json:"Unit,omitempty"`
	NodeSet string `xml:"nodeset,attr" json:"NodeSet,omitempty"`
}

type MemoryBackingAccess struct {
	Mode string `xml:"mode,attr" json:"Mode,omitempty"`
}

type NoSharePages struct {
}

type MemoryTarget struct {
	Size      Memory `xml:"size" json:"Size" :"Size"`
	Requested Memory `xml:"requested" json:"Requested" :"Requested"`
	Current   Memory `xml:"current" json:"Current" :"Current"`
	Node      string `xml:"node" json:"Node,omitempty" :"Node"`
	Block     Memory `xml:"block" json:"Block" :"Block"`
}

type MemoryDevice struct {
	XMLName xmlName       `xml:"memory" json:"XMLName"`
	Model   string        `xml:"model,attr" json:"Model,omitempty"`
	Target  *MemoryTarget `xml:"target" json:"Target,omitempty"`
	Alias   *Alias        `xml:"alias,omitempty" json:"Alias,omitempty"`
	Address *Address      `xml:"address,omitempty" json:"Address,omitempty"`
}

type Devices struct {
	Emulator    string             `xml:"emulator,omitempty" json:"Emulator,omitempty"`
	Interfaces  []Interface        `xml:"interface" json:"Interfaces,omitempty"`
	Channels    []Channel          `xml:"channel" json:"Channels,omitempty"`
	HostDevices []HostDevice       `xml:"hostdev,omitempty" json:"HostDevices,omitempty"`
	Controllers []Controller       `xml:"controller,omitempty" json:"Controllers,omitempty"`
	Video       []Video            `xml:"video" json:"Video,omitempty"`
	Graphics    []Graphics         `xml:"graphics" json:"Graphics,omitempty"`
	Ballooning  *MemBalloon        `xml:"memballoon,omitempty" json:"Ballooning,omitempty"`
	Disks       []Disk             `xml:"disk" json:"Disks,omitempty"`
	Serials     []Serial           `xml:"serial" json:"Serials,omitempty"`
	Consoles    []Console          `xml:"console" json:"Consoles,omitempty"`
	Watchdogs   []Watchdog         `xml:"watchdog,omitempty" json:"Watchdogs,omitempty"`
	Rng         *Rng               `xml:"rng,omitempty" json:"Rng,omitempty"`
	Filesystems []FilesystemDevice `xml:"filesystem,omitempty" json:"Filesystems,omitempty"`
	Redirs      []RedirectedDevice `xml:"redirdev,omitempty" json:"Redirs,omitempty"`
	SoundCards  []SoundCard        `xml:"sound,omitempty" json:"SoundCards,omitempty"`
	TPMs        []TPM              `xml:"tpm,omitempty" json:"TPMs,omitempty"`
	VSOCK       *VSOCK             `xml:"vsock,omitempty" json:"VSOCK,omitempty"`
	Memory      *MemoryDevice      `xml:"memory,omitempty" json:"Memory,omitempty"`
}

type TPM struct {
	Model   string     `xml:"model,attr" json:"Model,omitempty"`
	Backend TPMBackend `xml:"backend" json:"Backend"`
}

type TPMBackend struct {
	Type            string `xml:"type,attr" json:"Type,omitempty"`
	Version         string `xml:"version,attr" json:"Version,omitempty"`
	PersistentState string `xml:"persistent_state,attr,omitempty" json:"PersistentState,omitempty"`
}

// RedirectedDevice describes a device to be redirected
// See: https://libvirt.org/formatdomain.html#redirected-devices
type RedirectedDevice struct {
	Type   string                 `xml:"type,attr" json:"Type,omitempty"`
	Bus    string                 `xml:"bus,attr" json:"Bus,omitempty"`
	Source RedirectedDeviceSource `xml:"source" json:"Source"`
}

type RedirectedDeviceSource struct {
	Mode string `xml:"mode,attr" json:"Mode,omitempty"`
	Path string `xml:"path,attr" json:"Path,omitempty"`
}

type FilesystemDevice struct {
	Type       string            `xml:"type,attr" json:"Type,omitempty"`
	AccessMode string            `xml:"accessMode,attr" json:"AccessMode,omitempty"`
	Source     *FilesystemSource `xml:"source,omitempty" json:"Source,omitempty"`
	Target     *FilesystemTarget `xml:"target,omitempty" json:"Target,omitempty"`
	Driver     *FilesystemDriver `xml:"driver,omitempty" json:"Driver,omitempty"`
	Binary     *FilesystemBinary `xml:"binary,omitempty" json:"Binary,omitempty"`
}

type FilesystemTarget struct {
	Dir string `xml:"dir,attr,omitempty" json:"Dir,omitempty"`
}

type FilesystemSource struct {
	Dir    string `xml:"dir,attr" json:"Dir,omitempty"`
	Socket string `xml:"socket,attr,omitempty" json:"Socket,omitempty"`
}

type FilesystemDriver struct {
	Type  string `xml:"type,attr" json:"Type,omitempty"`
	Queue string `xml:"queue,attr,omitempty" json:"Queue,omitempty"`
}

type FilesystemBinary struct {
	Path  string                 `xml:"path,attr,omitempty" json:"Path,omitempty"`
	Xattr string                 `xml:"xattr,attr,omitempty" json:"Xattr,omitempty"`
	Cache *FilesystemBinaryCache `xml:"cache,omitempty" json:"Cache,omitempty"`
	Lock  *FilesystemBinaryLock  `xml:"lock,omitempty" json:"Lock,omitempty"`
}

type FilesystemBinaryCache struct {
	Mode string `xml:"mode,attr,omitempty" json:"Mode,omitempty"`
}

type FilesystemBinaryLock struct {
	Posix string `xml:"posix,attr,omitempty" json:"Posix,omitempty"`
	Flock string `xml:"flock,attr,omitempty" json:"Flock,omitempty"`
}

// BEGIN HostDevice -----------------------------
type HostDevice struct {
	XMLName   xmlName          `xml:"hostdev" json:"XMLName"`
	Source    HostDeviceSource `xml:"source" json:"Source"`
	Type      string           `xml:"type,attr" json:"Type,omitempty"`
	BootOrder *BootOrder       `xml:"boot,omitempty" json:"BootOrder,omitempty"`
	Managed   string           `xml:"managed,attr,omitempty" json:"Managed,omitempty"`
	Mode      string           `xml:"mode,attr,omitempty" json:"Mode,omitempty"`
	Model     string           `xml:"model,attr,omitempty" json:"Model,omitempty"`
	Address   *Address         `xml:"address,omitempty" json:"Address,omitempty"`
	Alias     *Alias           `xml:"alias,omitempty" json:"Alias,omitempty"`
	Display   string           `xml:"display,attr,omitempty" json:"Display,omitempty"`
	RamFB     string           `xml:"ramfb,attr,omitempty" json:"RamFB,omitempty"`
}

type HostDeviceSource struct {
	Address *Address `xml:"address,omitempty" json:"Address,omitempty"`
}

// END HostDevice -----------------------------

// BEGIN Controller -----------------------------

// Controller represens libvirt controller element https://libvirt.org/formatdomain.html#elementsControllers
type Controller struct {
	Type    string            `xml:"type,attr" json:"Type,omitempty"`
	Index   string            `xml:"index,attr" json:"Index,omitempty"`
	Model   string            `xml:"model,attr,omitempty" json:"Model,omitempty"`
	Driver  *ControllerDriver `xml:"driver,omitempty" json:"Driver,omitempty"`
	Alias   *Alias            `xml:"alias,omitempty" json:"Alias,omitempty"`
	Address *Address          `xml:"address,omitempty" json:"Address,omitempty"`
}

// END Controller -----------------------------

// BEGIN ControllerDriver
type ControllerDriver struct {
	IOThread *uint  `xml:"iothread,attr,omitempty" json:"IOThread,omitempty"`
	Queues   *uint  `xml:"queues,attr,omitempty" json:"Queues,omitempty"`
	IOMMU    string `xml:"iommu,attr,omitempty" json:"IOMMU,omitempty"`
}

// END ControllerDriver

// BEGIN Disk -----------------------------

type Disk struct {
	Device             string        `xml:"device,attr" json:"Device,omitempty"`
	Snapshot           string        `xml:"snapshot,attr,omitempty" json:"Snapshot,omitempty"`
	Type               string        `xml:"type,attr" json:"Type,omitempty"`
	Source             DiskSource    `xml:"source" json:"Source"`
	Target             DiskTarget    `xml:"target" json:"Target"`
	Serial             string        `xml:"serial,omitempty" json:"Serial,omitempty"`
	Driver             *DiskDriver   `xml:"driver,omitempty" json:"Driver,omitempty"`
	ReadOnly           *ReadOnly     `xml:"readonly,omitempty" json:"ReadOnly,omitempty"`
	Auth               *DiskAuth     `xml:"auth,omitempty" json:"Auth,omitempty"`
	Alias              *Alias        `xml:"alias,omitempty" json:"Alias,omitempty"`
	BackingStore       *BackingStore `xml:"backingStore,omitempty" json:"BackingStore,omitempty"`
	BootOrder          *BootOrder    `xml:"boot,omitempty" json:"BootOrder,omitempty"`
	Address            *Address      `xml:"address,omitempty" json:"Address,omitempty"`
	Model              string        `xml:"model,attr,omitempty" json:"Model,omitempty"`
	BlockIO            *BlockIO      `xml:"blockio,omitempty" json:"BlockIO,omitempty"`
	Capacity           *int64        `xml:"capacity,omitempty" json:"Capacity,omitempty"`
	ExpandDisksEnabled bool          `xml:"expandDisksEnabled,omitempty" json:"ExpandDisksEnabled,omitempty"`
	Shareable          *Shareable    `xml:"shareable,omitempty" json:"Shareable,omitempty"`
}

type DiskAuth struct {
	Username string      `xml:"username,attr" json:"Username,omitempty"`
	Secret   *DiskSecret `xml:"secret,omitempty" json:"Secret,omitempty"`
}

type DiskSecret struct {
	Type  string `xml:"type,attr" json:"Type,omitempty"`
	Usage string `xml:"usage,attr,omitempty" json:"Usage,omitempty"`
	UUID  string `xml:"uuid,attr,omitempty" json:"UUID,omitempty"`
}

type ReadOnly struct{}

type DiskSource struct {
	Dev           string          `xml:"dev,attr,omitempty" json:"Dev,omitempty"`
	File          string          `xml:"file,attr,omitempty" json:"File,omitempty"`
	StartupPolicy string          `xml:"startupPolicy,attr,omitempty" json:"StartupPolicy,omitempty"`
	Protocol      string          `xml:"protocol,attr,omitempty" json:"Protocol,omitempty"`
	Name          string          `xml:"name,attr,omitempty" json:"Name,omitempty"`
	Host          *DiskSourceHost `xml:"host,omitempty" json:"Host,omitempty"`
	Reservations  *Reservations   `xml:"reservations,omitempty" json:"Reservations,omitempty"`
}

type DiskTarget struct {
	Device string `xml:"dev,attr,omitempty" json:"Device,omitempty"`
	Tray   string `xml:"tray,attr,omitempty" json:"Tray,omitempty"`
}

type DiskDriver struct {
	Cache    string `xml:"cache,attr,omitempty" json:"Cache,omitempty"`
	Name     string `xml:"name,attr" json:"Name,omitempty"`
	Type     string `xml:"type,attr" json:"Type,omitempty"`
	IOThread *uint  `xml:"iothread,attr,omitempty" json:"IOThread,omitempty"`
	Queues   *uint  `xml:"queues,attr,omitempty" json:"Queues,omitempty"`
	Discard  string `xml:"discard,attr,omitempty" json:"Discard,omitempty"`
	IOMMU    string `xml:"iommu,attr,omitempty" json:"IOMMU,omitempty"`
}

type DiskSourceHost struct {
	Name string `xml:"name,attr" json:"Name,omitempty"`
	Port string `xml:"port,attr,omitempty" json:"Port,omitempty"`
}

type BackingStore struct {
	Type   string              `xml:"type,attr,omitempty" json:"Type,omitempty"`
	Format *BackingStoreFormat `xml:"format,omitempty" json:"Format,omitempty"`
	Source *DiskSource         `xml:"source,omitempty" json:"Source,omitempty"`
}

type BackingStoreFormat struct {
	Type string `xml:"type,attr" json:"Type,omitempty"`
}

type BlockIO struct {
	LogicalBlockSize  uint `xml:"logical_block_size,attr,omitempty" json:"LogicalBlockSize,omitempty"`
	PhysicalBlockSize uint `xml:"physical_block_size,attr,omitempty" json:"PhysicalBlockSize,omitempty"`
}

type Reservations struct {
	Managed            string              `xml:"managed,attr,omitempty" json:"Managed,omitempty"`
	SourceReservations *SourceReservations `xml:"source,omitempty" json:"SourceReservations,omitempty"`
}

type SourceReservations struct {
	Type string `xml:"type,attr" json:"Type,omitempty"`
	Path string `xml:"path,attr,omitempty" json:"Path,omitempty"`
	Mode string `xml:"mode,attr,omitempty" json:"Mode,omitempty"`
}

// END Disk -----------------------------

// BEGIN Serial -----------------------------

type Serial struct {
	Type   string        `xml:"type,attr" json:"Type,omitempty"`
	Target *SerialTarget `xml:"target,omitempty" json:"Target,omitempty"`
	Source *SerialSource `xml:"source,omitempty" json:"Source,omitempty"`
	Alias  *Alias        `xml:"alias,omitempty" json:"Alias,omitempty"`
	Log    *SerialLog    `xml:"log,omitempty" json:"Log,omitempty"`
}

type SerialTarget struct {
	Port *uint `xml:"port,attr,omitempty" json:"Port,omitempty"`
}

type SerialSource struct {
	Mode string `xml:"mode,attr,omitempty" json:"Mode,omitempty"`
	Path string `xml:"path,attr,omitempty" json:"Path,omitempty"`
}

type SerialLog struct {
	File   string `xml:"file,attr,omitempty" json:"File,omitempty"`
	Append string `xml:"append,attr,omitempty" json:"Append,omitempty"`
}

// END Serial -----------------------------

// BEGIN Console -----------------------------

type Console struct {
	Type   string         `xml:"type,attr" json:"Type,omitempty"`
	Target *ConsoleTarget `xml:"target,omitempty" json:"Target,omitempty"`
	Source *ConsoleSource `xml:"source,omitempty" json:"Source,omitempty"`
	Alias  *Alias         `xml:"alias,omitempty" json:"Alias,omitempty"`
}

type ConsoleTarget struct {
	Type *string `xml:"type,attr,omitempty" json:"Type,omitempty"`
	Port *uint   `xml:"port,attr,omitempty" json:"Port,omitempty"`
}

type ConsoleSource struct {
	Mode string `xml:"mode,attr,omitempty" json:"Mode,omitempty"`
	Path string `xml:"path,attr,omitempty" json:"Path,omitempty"`
}

// END Serial -----------------------------

// BEGIN Inteface -----------------------------

type Interface struct {
	Address             *Address               `xml:"address,omitempty" json:"Address,omitempty"`
	Type                string                 `xml:"type,attr" json:"Type,omitempty"`
	TrustGuestRxFilters string                 `xml:"trustGuestRxFilters,attr,omitempty" json:"TrustGuestRxFilters,omitempty"`
	Source              InterfaceSource        `xml:"source" json:"Source"`
	Target              *InterfaceTarget       `xml:"target,omitempty" json:"Target,omitempty"`
	Model               *Model                 `xml:"model,omitempty" json:"Model,omitempty"`
	MAC                 *MAC                   `xml:"mac,omitempty" json:"MAC,omitempty"`
	MTU                 *MTU                   `xml:"mtu,omitempty" json:"MTU,omitempty"`
	BandWidth           *BandWidth             `xml:"bandwidth,omitempty" json:"BandWidth,omitempty"`
	BootOrder           *BootOrder             `xml:"boot,omitempty" json:"BootOrder,omitempty"`
	LinkState           *LinkState             `xml:"link,omitempty" json:"LinkState,omitempty"`
	FilterRef           *FilterRef             `xml:"filterref,omitempty" json:"FilterRef,omitempty"`
	Alias               *Alias                 `xml:"alias,omitempty" json:"Alias,omitempty"`
	Driver              *InterfaceDriver       `xml:"driver,omitempty" json:"Driver,omitempty"`
	Rom                 *Rom                   `xml:"rom,omitempty" json:"Rom,omitempty"`
	ACPI                *ACPI                  `xml:"acpi,omitempty" json:"ACPI,omitempty"`
	Backend             *InterfaceBackend      `xml:"backend,omitempty" json:"Backend,omitempty"`
	PortForward         []InterfacePortForward `xml:"portForward,omitempty" json:"PortForward,omitempty"`
}

type InterfacePortForward struct {
	Proto   string                      `xml:"proto,attr" json:"Proto,omitempty"`
	Address string                      `xml:"address,attr,omitempty" json:"Address,omitempty"`
	Dev     string                      `xml:"dev,attr,omitempty" json:"Dev,omitempty"`
	Ranges  []InterfacePortForwardRange `xml:"range,omitempty" json:"Ranges,omitempty"`
}

type InterfacePortForwardRange struct {
	Start   uint   `xml:"start,attr" json:"Start,omitempty"`
	End     uint   `xml:"end,attr,omitempty" json:"End,omitempty"`
	To      uint   `xml:"to,attr,omitempty" json:"To,omitempty"`
	Exclude string `xml:"exclude,attr,omitempty" json:"Exclude,omitempty"`
}

type InterfaceBackend struct {
	Type    string `xml:"type,attr,omitempty" json:"Type,omitempty"`
	LogFile string `xml:"logFile,attr,omitempty" json:"LogFile,omitempty"`
}

type ACPI struct {
	Index uint `xml:"index,attr" json:"Index,omitempty"`
}

type InterfaceDriver struct {
	Name   string `xml:"name,attr" json:"Name,omitempty"`
	Queues *uint  `xml:"queues,attr,omitempty" json:"Queues,omitempty"`
	IOMMU  string `xml:"iommu,attr,omitempty" json:"IOMMU,omitempty"`
}

type LinkState struct {
	State string `xml:"state,attr" json:"State,omitempty"`
}

type BandWidth struct {
}

type BootOrder struct {
	Order uint `xml:"order,attr" json:"Order,omitempty"`
}

type MAC struct {
	MAC string `xml:"address,attr" json:"MAC,omitempty"`
}

type MTU struct {
	Size string `xml:"size,attr" json:"Size,omitempty"`
}

type FilterRef struct {
	Filter string `xml:"filter,attr" json:"Filter,omitempty"`
}

type xmlName struct {
	Space string `xml:"space,attr,omitempty" json:"Space,omitempty"`
	Local string `xml:"local,attr,omitempty" json:"Local,omitempty"`
}

type InterfaceSource struct {
	Network string   `xml:"network,attr,omitempty" json:"Network,omitempty"`
	Device  string   `xml:"dev,attr,omitempty" json:"Device,omitempty"`
	Bridge  string   `xml:"bridge,attr,omitempty" json:"Bridge,omitempty"`
	Mode    string   `xml:"mode,attr,omitempty" json:"Mode,omitempty"`
	Address *Address `xml:"address,omitempty" json:"Address,omitempty"`
}

type Model struct {
	Type string `xml:"type,attr" json:"Type,omitempty"`
}

type InterfaceTarget struct {
	Device  string `xml:"dev,attr" json:"Device,omitempty"`
	Managed string `xml:"managed,attr,omitempty" json:"Managed,omitempty"`
}

type Alias struct {
	Name        string `xml:"name,attr" json:"Name,omitempty"`
	userDefined bool   `json:"UserDefined,omitempty"`
}

// Package private, responsible to interact with xml and json marshal/unmarshal
type userAliasMarshal struct {
	Name        string `xml:"name,attr"`
	UserDefined bool   `xml:"-"`
}

type Rom struct {
	Enabled string `xml:"enabled,attr" json:"Enabled,omitempty"`
}

func NewUserDefinedAlias(aliasName string) *Alias {
	return &Alias{Name: aliasName, userDefined: true}
}

func (alias Alias) GetName() string {
	return alias.Name
}

func (alias Alias) IsUserDefined() bool {
	return alias.userDefined
}

type OS struct {
	Type       OSType    `xml:"type" json:"Type"`
	ACPI       *OSACPI   `xml:"acpi,omitempty" json:"ACPI,omitempty"`
	SMBios     *SMBios   `xml:"smbios,omitempty" json:"SMBios,omitempty"`
	BootOrder  []Boot    `xml:"boot" json:"BootOrder,omitempty"`
	BootMenu   *BootMenu `xml:"bootmenu,omitempty" json:"BootMenu,omitempty"`
	BIOS       *BIOS     `xml:"bios,omitempty" json:"BIOS,omitempty"`
	BootLoader *Loader   `xml:"loader,omitempty" json:"BootLoader,omitempty"`
	NVRam      *NVRam    `xml:"nvram,omitempty" json:"NVRam,omitempty"`
	Kernel     string    `xml:"kernel,omitempty" json:"Kernel,omitempty"`
	Initrd     string    `xml:"initrd,omitempty" json:"Initrd,omitempty"`
	KernelArgs string    `xml:"cmdline,omitempty" json:"KernelArgs,omitempty"`
}

type OSType struct {
	OS      string `xml:",chardata" json:"OS,omitempty"`
	Arch    string `xml:"arch,attr,omitempty" json:"Arch,omitempty"`
	Machine string `xml:"machine,attr,omitempty" json:"Machine,omitempty"`
}

type OSACPI struct {
	Table ACPITable `xml:"table,omitempty" json:"Table"`
}

type ACPITable struct {
	Path string `xml:",chardata" json:"Path,omitempty"`
	Type string `xml:"type,attr,omitempty" json:"Type,omitempty"`
}

type SMBios struct {
	Mode string `xml:"mode,attr" json:"Mode,omitempty"`
}

type NVRam struct {
	Template string `xml:"template,attr,omitempty" json:"Template,omitempty"`
	NVRam    string `xml:",chardata" json:"NVRam,omitempty"`
}

type Boot struct {
	Dev string `xml:"dev,attr" json:"Dev,omitempty"`
}

type BootMenu struct {
	Enable  string `xml:"enable,attr" json:"Enable,omitempty"`
	Timeout *uint  `xml:"timeout,attr,omitempty" json:"Timeout,omitempty"`
}

type Loader struct {
	ReadOnly string `xml:"readonly,attr,omitempty" json:"ReadOnly,omitempty"`
	Secure   string `xml:"secure,attr,omitempty" json:"Secure,omitempty"`
	Type     string `xml:"type,attr,omitempty" json:"Type,omitempty"`
	Path     string `xml:",chardata" json:"Path,omitempty"`
}

// TODO <bios rebootTimeout='0'/>
type BIOS struct {
	UseSerial string `xml:"useserial,attr,omitempty" json:"UseSerial,omitempty"`
}

type SysInfo struct {
	Type      string  `xml:"type,attr" json:"Type,omitempty"`
	System    []Entry `xml:"system>entry" json:"System,omitempty"`
	BIOS      []Entry `xml:"bios>entry" json:"BIOS,omitempty"`
	BaseBoard []Entry `xml:"baseBoard>entry" json:"BaseBoard,omitempty"`
	Chassis   []Entry `xml:"chassis>entry" json:"Chassis,omitempty"`
}

type Entry struct {
	Name  string `xml:"name,attr" json:"Name,omitempty"`
	Value string `xml:",chardata" json:"Value,omitempty"`
}

//END OS --------------------
//BEGIN LaunchSecurity --------------------

type LaunchSecurity struct {
	Type            string `xml:"type,attr" json:"Type,omitempty"`
	Cbitpos         string `xml:"cbitpos,omitempty" json:"Cbitpos,omitempty"`
	ReducedPhysBits string `xml:"reducedPhysBits,omitempty" json:"ReducedPhysBits,omitempty"`
	Policy          string `xml:"policy,omitempty" json:"Policy,omitempty"`
	DHCert          string `xml:"dhCert,omitempty" json:"DHCert,omitempty"`
	Session         string `xml:"session,omitempty" json:"Session,omitempty"`
}

//END LaunchSecurity --------------------
//BEGIN Clock --------------------

type Clock struct {
	Offset     string  `xml:"offset,attr,omitempty" json:"Offset,omitempty"`
	Timezone   string  `xml:"timezone,attr,omitempty" json:"Timezone,omitempty"`
	Adjustment string  `xml:"adjustment,attr,omitempty" json:"Adjustment,omitempty"`
	Timer      []Timer `xml:"timer,omitempty" json:"Timer,omitempty"`
}

type Timer struct {
	Name       string `xml:"name,attr" json:"Name,omitempty"`
	TickPolicy string `xml:"tickpolicy,attr,omitempty" json:"TickPolicy,omitempty"`
	Present    string `xml:"present,attr,omitempty" json:"Present,omitempty"`
	Track      string `xml:"track,attr,omitempty" json:"Track,omitempty"`
	Frequency  string `xml:"frequency,attr,omitempty" json:"Frequency,omitempty"`
}

//END Clock --------------------

//BEGIN Channel --------------------

type Channel struct {
	Type   string         `xml:"type,attr" json:"Type,omitempty"`
	Source *ChannelSource `xml:"source,omitempty" json:"Source,omitempty"`
	Target *ChannelTarget `xml:"target,omitempty" json:"Target,omitempty"`
}

type ChannelTarget struct {
	Name    string `xml:"name,attr,omitempty" json:"Name,omitempty"`
	Type    string `xml:"type,attr" json:"Type,omitempty"`
	Address string `xml:"address,attr,omitempty" json:"Address,omitempty"`
	Port    uint   `xml:"port,attr,omitempty" json:"Port,omitempty"`
	State   string `xml:"state,attr,omitempty" json:"State,omitempty"`
}

type ChannelSource struct {
	Mode string `xml:"mode,attr" json:"Mode,omitempty"`
	Path string `xml:"path,attr" json:"Path,omitempty"`
}

//END Channel --------------------

//BEGIN Sound -------------------

type SoundCard struct {
	Alias *Alias `xml:"alias,omitempty" json:"Alias,omitempty"`
	Model string `xml:"model,attr" json:"Model,omitempty"`
}

//END Sound -------------------

//BEGIN Video -------------------

type Video struct {
	Model VideoModel `xml:"model" json:"Model"`
}

type VideoModel struct {
	Type   string `xml:"type,attr" json:"Type,omitempty"`
	Heads  *uint  `xml:"heads,attr,omitempty" json:"Heads,omitempty"`
	Ram    *uint  `xml:"ram,attr,omitempty" json:"Ram,omitempty"`
	VRam   *uint  `xml:"vram,attr,omitempty" json:"VRam,omitempty"`
	VGAMem *uint  `xml:"vgamem,attr,omitempty" json:"VGAMem,omitempty"`
}

type Graphics struct {
	AutoPort      string          `xml:"autoport,attr,omitempty" json:"AutoPort,omitempty"`
	DefaultMode   string          `xml:"defaultMode,attr,omitempty" json:"DefaultMode,omitempty"`
	Listen        *GraphicsListen `xml:"listen,omitempty" json:"Listen,omitempty"`
	PasswdValidTo string          `xml:"passwdValidTo,attr,omitempty" json:"PasswdValidTo,omitempty"`
	Port          int32           `xml:"port,attr,omitempty" json:"Port,omitempty"`
	TLSPort       int             `xml:"tlsPort,attr,omitempty" json:"TLSPort,omitempty"`
	Type          string          `xml:"type,attr" json:"Type,omitempty"`
}

type GraphicsListen struct {
	Type    string `xml:"type,attr" json:"Type,omitempty"`
	Address string `xml:"address,attr,omitempty" json:"Address,omitempty"`
	Network string `xml:"newtork,attr,omitempty" json:"Network,omitempty"`
	Socket  string `xml:"socket,attr,omitempty" json:"Socket,omitempty"`
}

type Address struct {
	Type       string `xml:"type,attr" json:"Type,omitempty"`
	Domain     string `xml:"domain,attr,omitempty" json:"Domain,omitempty"`
	Bus        string `xml:"bus,attr" json:"Bus,omitempty"`
	Slot       string `xml:"slot,attr,omitempty" json:"Slot,omitempty"`
	Function   string `xml:"function,attr,omitempty" json:"Function,omitempty"`
	Controller string `xml:"controller,attr,omitempty" json:"Controller,omitempty"`
	Target     string `xml:"target,attr,omitempty" json:"Target,omitempty"`
	Unit       string `xml:"unit,attr,omitempty" json:"Unit,omitempty"`
	UUID       string `xml:"uuid,attr,omitempty" json:"UUID,omitempty"`
	Device     string `xml:"device,attr,omitempty" json:"Device,omitempty"`
}

//END Video -------------------

//BEGIN VSOCK -------------------

type VSOCK struct {
	Model string `xml:"model,attr,omitempty" json:"Model,omitempty"`
	CID   CID    `xml:"cid" json:"CID"`
}

type CID struct {
	Auto    string `xml:"auto,attr" json:"Auto,omitempty"`
	Address uint32 `xml:"address,attr,omitempty" json:"Address,omitempty"`
}

//END VSOCK -------------------

type Stats struct {
	Period uint `xml:"period,attr" json:"Period,omitempty"`
}

type MemBalloon struct {
	Model             string            `xml:"model,attr" json:"Model,omitempty"`
	Stats             *Stats            `xml:"stats,omitempty" json:"Stats,omitempty"`
	Address           *Address          `xml:"address,omitempty" json:"Address,omitempty"`
	Driver            *MemBalloonDriver `xml:"driver,omitempty" json:"Driver,omitempty"`
	FreePageReporting string            `xml:"freePageReporting,attr,omitempty" json:"FreePageReporting,omitempty"`
}

type MemBalloonDriver struct {
	IOMMU string `xml:"iommu,attr,omitempty" json:"IOMMU,omitempty"`
}

type Watchdog struct {
	Model   string   `xml:"model,attr" json:"Model,omitempty"`
	Action  string   `xml:"action,attr" json:"Action,omitempty"`
	Alias   *Alias   `xml:"alias,omitempty" json:"Alias,omitempty"`
	Address *Address `xml:"address,omitempty" json:"Address,omitempty"`
}

// Rng represents the source of entropy from host to VM
type Rng struct {
	// Model attribute specifies what type of RNG device is provided
	Model string `xml:"model,attr" json:"Model,omitempty"`
	// Backend specifies the source of entropy to be used
	Backend *RngBackend `xml:"backend,omitempty" json:"Backend,omitempty"`
	Address *Address    `xml:"address,omitempty" json:"Address,omitempty"`
	Driver  *RngDriver  `xml:"driver,omitempty" json:"Driver,omitempty"`
}

type RngDriver struct {
	IOMMU string `xml:"iommu,attr,omitempty" json:"IOMMU,omitempty"`
}

// RngRate sets the limiting factor how to read from entropy source
type RngRate struct {
	// Period define how long is the read period
	Period uint32 `xml:"period,attr" json:"Period,omitempty"`
	// Bytes define how many bytes can guest read from entropy source
	Bytes uint32 `xml:"bytes,attr" json:"Bytes,omitempty"`
}

// RngBackend is the backend device used
type RngBackend struct {
	// Model is source model
	Model string `xml:"model,attr" json:"Model,omitempty"`
	// specifies the source of entropy to be used
	Source string `xml:",chardata" json:"Source,omitempty"`
}

type IOThreads struct {
	IOThreads uint `xml:",chardata" json:"IOThreads,omitempty"`
}

// TODO ballooning, rng, cpu ...

type SecretUsage struct {
	Type   string `xml:"type,attr"`
	Target string `xml:"target,omitempty"`
}

type SecretSpec struct {
	XMLName     xmlName     `xml:"secret"`
	Ephemeral   string      `xml:"ephemeral,attr"`
	Private     string      `xml:"private,attr"`
	Description string      `xml:"description,omitempty"`
	Usage       SecretUsage `xml:"usage,omitempty"`
}

// VirtualMachineStatus defines the observed state of VirtualMachine
type VirtualMachineStatus struct {
	Phase     string    `json:"phase,omitempty"`
	VMPodName string    `json:"vmPodName,omitempty"`
	VMPodUID  types.UID `json:"vmPodUID,omitempty"`
	NodeName  string    `json:"nodeName,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// VirtualMachine is the Schema for the virtualmachines API
type VirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineSpec   `json:"spec,omitempty"`
	Status VirtualMachineStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VirtualMachineList contains a list of VirtualMachine
type VirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualMachine{}, &VirtualMachineList{})
}
