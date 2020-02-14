package log

type Severity int32

const (
	Severity_Unknown Severity = 0
	Severity_Error   Severity = 1
	Severity_Warning Severity = 2
	Severity_Info    Severity = 3
	Severity_Debug   Severity = 4
)

var Severity_name = map[int32]string{
	0: "Unknown",
	1: "Error",
	2: "Warning",
	3: "Info",
	4: "Debug",
}

var Severity_value = map[string]int32{
	"Unknown": 0,
	"Error":   1,
	"Warning": 2,
	"Info":    3,
	"Debug":   4,
}
