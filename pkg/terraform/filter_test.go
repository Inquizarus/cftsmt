package terraform_test

import (
	"testing"

	"github.com/inquizarus/cftsmt/pkg/terraform"
	"github.com/stretchr/testify/assert"
)

func TestThatFilterMatchReturnsTrueWhenResourceMatches(t *testing.T) {
	f := terraform.ResourceFilter{
		Type: "test_resource",
		Mode: "managed",
		Values: map[string]interface{}{
			"id": "abcd",
			"life_cycle": map[string]interface{}{
				"prevent_destroy": true,
			},
		},
	}
	r := terraform.StateResource{
		Type: "test_resource",
		Mode: "managed",
		Values: map[string]interface{}{
			"id": "abcd",
			"life_cycle": map[string]interface{}{
				"prevent_destroy": true,
			},
		},
	}

	assert.True(t, f.Matches(r))
}

func TestThatFilterMatchReturnsFalseWhenResourceDontMatchType(t *testing.T) {
	f := terraform.ResourceFilter{
		Type: "test_resource",
	}

	r := terraform.StateResource{
		Type: "test_other",
	}

	assert.False(t, f.Matches(r))
}

func TestThatFilterMatchReturnsFalseWhenResourceDontMatchMode(t *testing.T) {
	f := terraform.ResourceFilter{
		Mode: "managed",
	}
	r := terraform.StateResource{
		Mode: "data",
	}

	assert.False(t, f.Matches(r))
}

func TestThatFilterMatchReturnsFalseWhenResourceDontMatchValuesValue(t *testing.T) {
	f := terraform.ResourceFilter{
		Values: map[string]interface{}{
			"id": "abcd",
		},
	}
	r := terraform.StateResource{
		Values: map[string]interface{}{
			"id": "xyz",
		},
	}
	assert.False(t, f.Matches(r))
}

func TestThatFilterMatchReturnsFalseWhenResourceDontMatchValues(t *testing.T) {
	f := terraform.ResourceFilter{
		Values: map[string]interface{}{
			"id": "abcd",
		},
	}
	r := terraform.StateResource{
		Values: map[string]interface{}{},
	}
	assert.False(t, f.Matches(r))
}

func TestThatFilterMatchReturnsFalseWhenResourceDontMatchNestedValue(t *testing.T) {
	f := terraform.ResourceFilter{
		Type: "test_resource",
		Mode: "managed",
		Values: map[string]interface{}{
			"id": "abcd",
			"life_cycle": map[string]interface{}{
				"prevent_destroy": true,
			},
		},
	}
	r := terraform.StateResource{
		Type: "test_resource",
		Mode: "managed",
		Values: map[string]interface{}{
			"id": "abcd",
			"life_cycle": map[string]interface{}{
				"prevent_destroy": false,
			},
		},
	}
	assert.False(t, f.Matches(r))
}

func TestThatFilterMatchReturnsFalseWhenResourceDontMatchNestedValueKey(t *testing.T) {
	f := terraform.ResourceFilter{
		Type: "test_resource",
		Mode: "managed",
		Values: map[string]interface{}{
			"id": "abcd",
			"life_cycle": map[string]interface{}{
				"prevent_destroy": true,
			},
		},
	}
	r := terraform.StateResource{
		Type: "test_resource",
		Mode: "managed",
		Values: map[string]interface{}{
			"id":         "abcd",
			"life_cycle": map[string]interface{}{},
		},
	}
	assert.False(t, f.Matches(r))
}
