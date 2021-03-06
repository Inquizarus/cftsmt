package terraform

// ResourceFilter ...
type ResourceFilter struct {
	Type   string
	Mode   string
	Values map[string]interface{}
}

func (rf *ResourceFilter) typeMatches(r StateResource) bool {
	if rf.Type != "" {
		return rf.Type == r.Type
	}
	return true
}

func (rf *ResourceFilter) modeMatches(r StateResource) bool {
	if rf.Mode != "" {
		return rf.Mode == r.Mode
	}
	return true
}

func (rf *ResourceFilter) valueMatches(x interface{}, y interface{}) bool {
	switch t := x.(type) {
	case string:
		return t == y
	case bool:
		return t == y
	case map[string]interface{}:
		for k, v := range x.(map[string]interface{}) {
			if nil == y.(map[string]interface{})[k] {
				return false
			}
			if false == rf.valueMatches(v, y.(map[string]interface{})[k]) {
				return false
			}
		}
	}
	return true
}

func (rf *ResourceFilter) valuesMatches(r StateResource) bool {
	for rfk, rfv := range rf.Values {
		if rv := r.Values[rfk]; nil == rv {
			return false
		}
		if false == rf.valueMatches(rfv, r.Values[rfk]) {
			return false
		}
	}
	return true
}

// Matches returns true if the passed resource matches all the configured criterias
func (rf *ResourceFilter) Matches(r StateResource) bool {
	if false == rf.typeMatches(r) {
		return false
	}
	if false == rf.modeMatches(r) {
		return false
	}
	if false == rf.valuesMatches(r) {
		return false
	}
	return true
}
