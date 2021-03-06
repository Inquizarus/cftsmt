package terraform

import (
	"bufio"
	"encoding/json"
	"io"
)

// StateValue ...
type StateValue struct {
	Sensitive bool `json:"sensitive"`
	Value     interface{}
}

// StateResource ...
type StateResource struct {
	Address      string                 `json:"address"`
	Name         string                 `json:"name"`
	Mode         string                 `json:"mode"`
	Type         string                 `json:"type"`
	ProviderName string                 `json:"provider_name"`
	Values       map[string]interface{} `json:"values"`
}

// StateModule ...
type StateModule struct {
	Address      string          `json:"address"`
	Resources    []StateResource `json:"resources"`
	ChildModules []StateModule   `json:"child_modules"`
}

// StateValues ...
type StateValues struct {
	Outputs    map[string]StateValue `json:"outputs"`
	RootModule StateModule           `json:"root_module"`
}

// State ...
type State struct {
	Values StateValues `json:"values"`
}

// FromReader ...
func FromReader(r *bufio.Reader) (State, error) {
	var err error
	tfstate := State{}

	var output []byte
	for {
		input, err := r.ReadByte()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	if nil == err {
		err = json.Unmarshal(output, &tfstate)
	}

	return tfstate, err
}
