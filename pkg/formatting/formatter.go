package formatting

import "github.com/inquizarus/cftsmt/pkg/terraform"

// Provider is a common interface for ...
type Provider interface {
	Name() string
	Import(r terraform.StateResource) string
	Remove(r terraform.StateResource) string
}

// Formatter handles formatting of specific things ...
type Formatter struct {
	providers map[string]Provider
}

func (f *Formatter) getProvider(s string) Provider {
	if p := f.providers[s]; p != nil {
		return p
	}
	if p := f.providers["default"]; p != nil {
		return p
	}
	return nil
}

// SetProvider sets or updates the provider entry
func (f *Formatter) SetProvider(p Provider) {
	f.providers[p.Name()] = p
}

// RemoveProvider remnoves the provider entry with name n
func (f *Formatter) RemoveProvider(n string) {
	delete(f.providers, n)
}

// Import generates an import statement for the passed resource
func (f *Formatter) Import(r terraform.StateResource) string {
	p := f.getProvider(r.ProviderName)
	if nil == p {
		return ""
	}
	return p.Import(r)
}

// Remove generates a remove statement for the passed resource
func (f *Formatter) Remove(r terraform.StateResource) string {
	p := f.getProvider(r.ProviderName)
	if nil == p {
		return ""
	}
	return p.Remove(r)
}

// NewFormatter returns a Formatter with the passed providers set
func NewFormatter(providers map[string]Provider) Formatter {
	return Formatter{
		providers: providers,
	}
}
