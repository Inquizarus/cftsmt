package terraform_test

import (
	"testing"

	"github.com/inquizarus/cftsmt/pkg/terraform"
	"github.com/stretchr/testify/assert"
)

func TestThatNilIsReturnedWhenTryingToFindEmptyResource(t *testing.T) {
	r := terraform.StateResource{
		Type:    "test_resource",
		Name:    "this",
		Address: "",
	}
	m := terraform.StateModule{
		Address: "module.test",
		Resources: []terraform.StateResource{
			r,
		},
		ChildModules: []terraform.StateModule{},
	}
	o := terraform.FindResourceInModule(r.Address, m)

	assert.Nil(t, o)
}

func TestThatNilIsReturnedWhenResourceIsNotFound(t *testing.T) {
	r := terraform.StateResource{
		Type:    "test_resource",
		Name:    "this",
		Address: "test_resource.this",
	}
	m := terraform.StateModule{
		Address:      "module.test",
		Resources:    []terraform.StateResource{},
		ChildModules: []terraform.StateModule{},
	}
	o := terraform.FindResourceInModule(r.Address, m)

	assert.Nil(t, o)
}

func TestThatNilIsReturnedWhenSubModuleeIsNotFound(t *testing.T) {
	r := terraform.StateResource{
		Type:    "test_resource",
		Name:    "this",
		Address: "module.sub.test_resource.this",
	}
	m := terraform.StateModule{
		Address:      "module.test",
		Resources:    []terraform.StateResource{},
		ChildModules: []terraform.StateModule{},
	}
	o := terraform.FindResourceInModule(r.Address, m)

	assert.Nil(t, o)
}

func TestThatResourceCanBeFoundInModuleByAddress(t *testing.T) {
	r := terraform.StateResource{
		Type:    "test_resource",
		Name:    "this",
		Address: "test_resource.this",
	}
	m := terraform.StateModule{
		Address: "module.test",
		Resources: []terraform.StateResource{
			r,
		},
		ChildModules: []terraform.StateModule{},
	}
	o := terraform.FindResourceInModule(r.Address, m)

	assert.Equal(t, r, *o)
}

func TestThatDataResourceCanBeFoundInModuleByAddress(t *testing.T) {
	r := terraform.StateResource{
		Type:    "test_resource",
		Name:    "this",
		Mode:    "data",
		Address: "data.test_resource.this",
	}
	m := terraform.StateModule{
		Address: "module.test",
		Resources: []terraform.StateResource{
			r,
		},
		ChildModules: []terraform.StateModule{},
	}
	o := terraform.FindResourceInModule(r.Address, m)

	assert.Equal(t, r, *o)
}

func TestThatResourceInSubModuleCanBeFoundInModuleByAddress(t *testing.T) {
	r := terraform.StateResource{
		Type:    "test_resource",
		Name:    "this",
		Address: "module.test.module.sub_test.test_resource.this",
	}
	m := terraform.StateModule{
		Address: "module.test",
		ChildModules: []terraform.StateModule{
			{
				Address: "module.test.module.sub_test",
				Resources: []terraform.StateResource{
					r,
				},
			},
		},
	}
	o := terraform.FindResourceInModule(r.Address, m)

	assert.Equal(t, r, *o)
}

func TestThatDataResourceInSubModuleCanBeFoundInModuleByAddress(t *testing.T) {
	r := terraform.StateResource{
		Type:    "test_resource",
		Name:    "this",
		Mode:    "data",
		Address: "module.test.module.sub_test.data.test_resource.this",
	}
	m := terraform.StateModule{
		Address: "module.test",
		ChildModules: []terraform.StateModule{
			{
				Address: "module.test.module.sub_test",
				Resources: []terraform.StateResource{
					r,
				},
			},
		},
	}
	o := terraform.FindResourceInModule(r.Address, m)

	assert.Equal(t, r, *o)
}

func TestThatResourceCanBeFoundInModuleByTypeAndName(t *testing.T) {
	r := terraform.StateResource{
		Type:    "test_resource",
		Name:    "this",
		Address: "test_resource.this",
	}
	m := terraform.StateModule{
		Address: "module.test",
		Resources: []terraform.StateResource{
			r,
		},
		ChildModules: []terraform.StateModule{},
	}
	o := terraform.FindResourceInModuleByTypeAndName(r.Type, r.Name, m)

	assert.Equal(t, r, *o)
}

func TestThatModuleCanBeFoundInModuleByAddress(t *testing.T) {
	r := terraform.StateModule{
		Address: "module.test.module.test_sub",
	}
	m := terraform.StateModule{
		Address: "module.test",
		ChildModules: []terraform.StateModule{
			r,
		},
	}
	o := terraform.FindModuleInModule(r.Address, m)

	assert.Equal(t, r, *o)
}

func TestThatFindResourcesInModuleWorksAsIntended(t *testing.T) {
	m := terraform.StateModule{
		Address: "root",
		Resources: []terraform.StateResource{
			{
				Address: "test_resource.one",
				Type:    "test_resource",
				Name:    "one",
			},
			{
				Address: "test_resource.two",
				Type:    "test_resource",
				Name:    "two",
			},
		},
		ChildModules: []terraform.StateModule{
			{
				Address: "module.test_module_d1",
				Resources: []terraform.StateResource{
					{
						Address: "module.test_module_d1.test_resource.one",
						Type:    "test_resource",
						Name:    "one",
					},
					{
						Address: "module.test_module_d1.test_resource.two",
						Type:    "test_resource",
						Name:    "two",
					},
				},
				ChildModules: []terraform.StateModule{
					{
						Address: "module.test_module_d1.module.test_module_d2",
						Resources: []terraform.StateResource{
							{
								Address: "module.test_module_d1.module.test_module_d2.test_resource.one",
								Type:    "test_resource",
								Name:    "one",
							},
							{
								Address: "module.test_module_d1.module.test_module_d2.test_resource.two",
								Type:    "test_resource",
								Name:    "two",
							},
						},
					},
				},
			},
		},
	}
	cases := []struct {
		m terraform.StateModule
		d int
		e int
	}{
		{
			m: m,
			d: 0,
			e: 2,
		},
		{
			m: m,
			d: 1,
			e: 4,
		},
		{
			m: m,
			d: 2,
			e: 6,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.e, len(terraform.FindResourcesInModule(c.m, c.d, &terraform.ResourceFilter{})))
	}
}
