package formatting

import (
	"fmt"
	"strings"

	"github.com/inquizarus/cftsmt/pkg/terraform"
)

// DefaultProvider is used to generate statements for AWS specific resources
type DefaultProvider struct{}

// Name ...
func (f *DefaultProvider) Name() string {
	return "default"
}

func (f *DefaultProvider) getImportFormat(r terraform.StateResource) string {
	defaultFormat := "terraform import %s %s"
	quotedFormat := "terraform import '%s' %s"
	formats := map[string]string{}

	if format := formats[r.Type]; "" != format {
		return format
	}

	if true == strings.HasSuffix(r.Address, "]") {
		return quotedFormat
	}

	return defaultFormat
}

func (f *DefaultProvider) getRemoveFormat(r terraform.StateResource) string {
	defaultFormat := "terraform state rm %s"
	quotedFormat := "terraform state rm '%s'"
	formats := map[string]string{}

	if format := formats[r.Type]; "" != format {
		return format
	}

	if true == strings.HasSuffix(r.Address, "]") {
		return quotedFormat
	}

	return defaultFormat
}

func (f *DefaultProvider) getKey(r terraform.StateResource) string {
	defaultKey := "id"
	keys := map[string]string{}
	if key := keys[r.Type]; "" != key {
		return key
	}
	return defaultKey
}

// Import generates an import statement for the passed resource
func (f *DefaultProvider) Import(r terraform.StateResource) string {
	// Refactor to determine if address should be wrapped in '' and which value that is used to import it
	return fmt.Sprintf(f.getImportFormat(r), r.Address, r.Values[f.getKey(r)])
}

// Remove generates a remove statement for the passed resource
func (f *DefaultProvider) Remove(r terraform.StateResource) string {
	return fmt.Sprintf(f.getRemoveFormat(r), r.Address)
}
