package terraform

import "strings"

// FindResourceInModule attempts to traverse through a module and its sub-modules
// to find a specific resource by its address
func FindResourceInModule(a string, m StateModule) *StateResource {
	parts := strings.Split(a, ".")
	if 1 > len(parts) || "" == parts[0] {
		return nil
	}

	if 3 > len(parts) { // This is not a module or data prefixed resource
		return FindResourceInModuleByTypeAndName(parts[0], parts[1], m)
	}

	if "data" == parts[0] {
		return FindResourceInModuleByTypeAndName(parts[1], parts[2], m)
	}

	trimCount := 2

	if strings.Contains(a, "data") {
		trimCount = 3
	}

	// TODO: Add support for the case of module.module_name.module.module_name (and so on)

	sm := FindModuleInModule(strings.Join(parts[:len(parts)-trimCount], "."), m)

	if nil != sm {
		return FindResourceInModuleByTypeAndName(parts[len(parts)-2], parts[len(parts)-1], *sm)
	}

	return nil
}

// FindResourcesInModule traverses passed module and finds resources by digging
// into child modules to a depth of passed value. Depth of 0 only return the root
// modules resources
func FindResourcesInModule(m StateModule, depth int, filter *ResourceFilter) []StateResource {
	resources := []StateResource{}
	if depth < 0 {
		return resources
	}
	for _, r := range m.Resources {
		if true == filter.Matches(r) {
			resources = append(resources, r)
		}
	}
	for _, sm := range m.ChildModules {
		resources = append(resources, FindResourcesInModule(sm, depth-1, filter)...)
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
