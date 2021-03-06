package app

var (
	cfgFile         string
	resourceAddress string
	depth           int
	// filter argument containers
	typeFilter   string
	modeFilter   string
	valuesFilter string
)

const (
	// ArgConfigFile ...
	ArgConfigFile = "config-file"
	// ArgResourceAddress ...
	ArgResourceAddress = "resource-address"
	// ArgDepth ...
	ArgDepth = "depth"
	// ArgTypeFilter ...
	ArgTypeFilter = "type-filter"
	// ArgModeFilter ...
	ArgModeFilter = "mode-filter"
	// ArgValuesFilter ...
	ArgValuesFilter = "values-filter"
)

// ShortArgMap ...
var ShortArgMap = map[string]string{
	ArgConfigFile:      "c",
	ArgResourceAddress: "r",
	ArgDepth:           "d",
	ArgTypeFilter:      "t",
	ArgModeFilter:      "m",
	ArgValuesFilter:    "v",
}

// ArgDescriptionMap ...
var ArgDescriptionMap = map[string]string{
	ArgConfigFile:      "Absolute path to configuration for cftsmt",
	ArgResourceAddress: "Full Terraform path to a specific resource",
	ArgDepth:           "How many levels that should be traveresed to find resources",
	ArgTypeFilter:      "Only show resources with this type.",
	ArgModeFilter:      "Only show resources with this mode.",
	ArgValuesFilter:    "Only show resources where these values matches.",
}

// ArgDefaultValueMap ...
var ArgDefaultValueMap = map[string]interface{}{
	ArgConfigFile:      "",
	ArgResourceAddress: "",
	ArgTypeFilter:      "",
	ArgModeFilter:      "",
	ArgValuesFilter:    "",
	ArgDepth:           0,
}
