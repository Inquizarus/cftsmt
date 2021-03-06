package terraform

import "strings"

// FindResourceInModule attempts to traverse through a module and its sub-modules
// to find a specific resource by its address
func FindResourceInModule(a string, m StateModule) *StateResource {
	parts := strings.Split(a, ".")
	if 1 > len(parts) {
		return nil
	}
	if "module" != parts[0] {
		return FindResourceInModuleByTypeAndName(parts[0], parts[1], m)
	}

	// TODO: Add support for the case of module.module_name.module.module_name (and so on)

	sm := FindModuleInModule(strings.Join(parts[:len(parts)-2], "."), m)

	if nil != sm {
		return FindResourceInModuleByTypeAndName(parts[len(parts)-2], parts[len(parts)-1], *sm)
	}

	return nil
}

// FindResourcesInModule traverses passed module and finds resources by digging
// into child modules to a depth of passed value. Depth of 0 only return the root
// modules resources
func FindResourcesInModule(m StateModule, depth int) []StateResource {
	resources := []StateResource{}
	if depth < 0 {
		return resources
	}
	resources = append(resources, m.Resources...)
	for _, sm := range m.ChildModules {
		resources = append(resources, FindResourcesInModule(sm, depth-1)...)
	}
	return resources
}

// FindResourceInModuleByTypeAndName iterates over a modules resources and attempts to find a specific
// resource by type and name
func FindResourceInModuleByTypeAndName(t string, n string, m StateModule) *StateResource {
	for _, r := range m.Resources {
		if t == r.Type && n == r.Name {
			return &r
		}
	}
	return nil
}

// FindModuleInModule iterates over a modules sub-mobules and attempts to find a specific
// module by address
func FindModuleInModule(a string, m StateModule) *StateModule {
	for _, v := range m.ChildModules {
		if v.Address == a {
			return &v
		}
	}
	return nil
}
