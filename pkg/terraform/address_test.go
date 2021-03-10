package terraform_test

import (
	"testing"

	"github.com/inquizarus/cftsmt/pkg/terraform"
	"github.com/stretchr/testify/assert"
)

func TestThatIsModuleAddressReturnCorrectValues(t *testing.T) {
	assert.True(t, terraform.IsModuleAddress("module.test.this"))
	assert.False(t, terraform.IsModuleAddress("resource.this"))
	assert.False(t, terraform.IsModuleAddress("data.resource.this"))
	assert.False(t, terraform.IsModuleAddress(""))
}

func TestThatIsDataAddressReturnCorrectValues(t *testing.T) {
	assert.True(t, terraform.IsDataAddress("data.test_data.this"))
	assert.False(t, terraform.IsDataAddress("resource.this"))
	assert.False(t, terraform.IsDataAddress("module.test.this"))
	assert.False(t, terraform.IsDataAddress(""))
}

func TestThatIsManagedAddressReturnCorrectValues(t *testing.T) {
	assert.True(t, terraform.IsManagedAddress("test_resource.this"))
	assert.False(t, terraform.IsManagedAddress("data.test_data.this"))
	assert.False(t, terraform.IsManagedAddress(""))
}
