package formatting_test

import (
	"testing"

	"github.com/inquizarus/cftsmt/pkg/formatting"
	"github.com/inquizarus/cftsmt/pkg/terraform"
	"github.com/stretchr/testify/assert"
)

func TestThatFormatterCreatesValidRemoveStatement(t *testing.T) {
	fp := formatting.DefaultProvider{}
	f := formatting.NewFormatter(map[string]formatting.Provider{})
	f.SetProvider(&fp)

	cases := []struct {
		expected string
		resource terraform.StateResource
	}{
		{
			expected: "terraform state rm aws_ec2_instance.this",
			resource: terraform.StateResource{
				Address:      "aws_ec2_instance.this",
				ProviderName: "registry.terraform.io/hashicorp/aws",
				Values: map[string]interface{}{
					"id": "i-abcd1234",
				},
			},
		},
		{
			expected: "terraform state rm module.db_init_lambda.aws_s3_bucket.this",
			resource: terraform.StateResource{
				Address:      "module.db_init_lambda.aws_s3_bucket.this",
				ProviderName: "registry.terraform.io/hashicorp/aws",
				Values: map[string]interface{}{
					"id": "core-prod-db-init-lambda-code",
				},
			},
		},
		{
			expected: "terraform state rm 'module.grafana_db.aws_db_instance.this[0]'",
			resource: terraform.StateResource{
				Address:      "module.grafana_db.aws_db_instance.this[0]",
				ProviderName: "registry.terraform.io/hashicorp/aws",
				Values: map[string]interface{}{
					"id": "db-1",
				},
			},
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, f.Remove(c.resource))
	}
}

func TestThatFormatterCreatesValidImportStatement(t *testing.T) {
	fp := formatting.DefaultProvider{}
	f := formatting.NewFormatter(map[string]formatting.Provider{})
	f.SetProvider(&fp)

	cases := []struct {
		expected string
		resource terraform.StateResource
	}{
		{
			expected: "terraform import aws_ec2_instance.this i-abcd1234",
			resource: terraform.StateResource{
				Address:      "aws_ec2_instance.this",
				ProviderName: "registry.terraform.io/hashicorp/aws",
				Values: map[string]interface{}{
					"id": "i-abcd1234",
				},
			},
		},
		{
			expected: "terraform import module.db_init_lambda.aws_s3_bucket.this core-prod-db-init-lambda-code",
			resource: terraform.StateResource{
				Address:      "module.db_init_lambda.aws_s3_bucket.this",
				ProviderName: "registry.terraform.io/hashicorp/aws",
				Values: map[string]interface{}{
					"id": "core-prod-db-init-lambda-code",
				},
			},
		},
		{
			expected: "terraform import 'module.grafana_db.aws_db_instance.this[0]' db-1",
			resource: terraform.StateResource{
				Address:      "module.grafana_db.aws_db_instance.this[0]",
				ProviderName: "registry.terraform.io/hashicorp/aws",
				Values: map[string]interface{}{
					"id": "db-1",
				},
			},
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, f.Import(c.resource))
	}
}
