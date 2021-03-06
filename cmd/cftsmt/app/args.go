package app

var (
	cfgFile         string
	resourceAddress string
	depth           int
)

const (
	// ArgConfigFile ...
	ArgConfigFile = "config-file"
	// ArgResourceAddress ...
	ArgResourceAddress = "resource-address"
	// ArgDepth ...
	ArgDepth = "depth"
)

// ShortArgMap ...
var ShortArgMap = map[string]string{
	ArgConfigFile:      "c",
	ArgResourceAddress: "r",
	ArgDepth:           "d",
}

// ArgDescriptionMap ...
var ArgDescriptionMap = map[string]string{
	ArgConfigFile:      "Absolute path to configuration for cftsmt",
	ArgResourceAddress: "Full Terraform path to a specific resource",
	ArgDepth:           "How many levels that should be traveresed to find resources",
}

// ArgDefaultValueMap ...
var ArgDefaultValueMap = map[string]interface{}{
	ArgConfigFile:      "",
	ArgResourceAddress: "",
	ArgDepth:           0,
}
