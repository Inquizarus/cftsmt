package app

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/inquizarus/cftsmt/pkg/terraform"
	"github.com/rodaine/table"
	"github.com/spf13/viper"
)

func addModulesModulesRows(mlist []terraform.StateModule, t table.Table, level int) {
	for _, m := range mlist {
		addModuleRows(m, t, level)
	}
}

func addModuleRows(m terraform.StateModule, t table.Table, level int) {
	address := "root"
	if "" != m.Address {
		address = m.Address
	}
	t.AddRow(address, len(m.ChildModules), len(m.Resources), level)
	addModulesModulesRows(m.ChildModules, t, level+1)
}

func addModulesResourcesRows(mlist []terraform.StateModule, t table.Table, level int) {
	for _, m := range mlist {
		addModuleResourcesRows(m, t, level)
	}
}

func addModuleResourcesRows(m terraform.StateModule, t table.Table, level int) {
	for _, r := range m.Resources {
		t.AddRow(r.Type, r.Name, r.Address, r.Mode, level)
	}
	addModulesResourcesRows(m.ChildModules, t, level+1)
}

func addResourceRows(rl []terraform.StateResource, t table.Table) {
	for _, r := range rl {
		t.AddRow(r.Type, r.Name, r.Address, r.Mode, strings.Count(r.Address, "module"))
	}
}

func createResourceFilter(v *viper.Viper) *terraform.ResourceFilter {
	values := map[string]interface{}{}
	valuesFilter := v.GetString(ArgValuesFilter)
	if "" != valuesFilter {
		if err := json.Unmarshal([]byte(valuesFilter), &values); nil != err {
			color.Red("values filter must be valid JSON, '%s' is invalid", valuesFilter)
			os.Exit(1)
		}
	}
	return &terraform.ResourceFilter{
		Type:   viper.GetString(ArgTypeFilter),
		Mode:   viper.GetString(ArgModeFilter),
		Values: values,
	}
}
