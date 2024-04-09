package build

var CurrentCommit string

const BuildVersion = "0.4.5"

const UBITaskImageIntelCpu = "filswan/ubi-worker-cpu-intel:v2.0"
const UBITaskImageIntelGpu = "filswan/ubi-worker-gpu-intel:v2.0"
const UBITaskImageAmdCpu = "filswan/ubi-worker-cpu-amd:v2.0"
const UBITaskImageAmdGpu = "filswan/ubi-worker-gpu-amd:v2.0"

// aleo image
const UBITaskAleoProofImageAmdGpu = "martian123/swan-aleo-proof:0.0.5"

func UserVersion() string {
	return BuildVersion + CurrentCommit
}
