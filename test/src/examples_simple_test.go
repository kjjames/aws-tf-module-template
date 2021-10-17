package src

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	testStructure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)


func TestTerraformSimpleExample(t *testing.T) {
	t.Parallel()

	fixtureFolder := "../../examples/simple"

	// Use Terratest to deploy the infrastructure
	testStructure.RunTestStage(t, "setup", func() {
		terraformOptions := &terraform.Options{
			// Indicate the directory that contains the Terraform configuration to deploy
			TerraformDir: fixtureFolder,
		}

		// Save options for later test stages
		testStructure.SaveTerraformOptions(t, fixtureFolder, terraformOptions)

		// Triggers the terraform init and terraform apply command
		terraform.InitAndApply(t, terraformOptions)
	})

	testStructure.RunTestStage(t, "validate", func() {
		// run validation checks here
		terraformOptions := testStructure.LoadTerraformOptions(t, fixtureFolder)
		id := terraform.Output(t, terraformOptions, "id")
		assert.Equal(t, "", id)
	})

	// When the test is completed, teardown the infrastructure by calling terraform destroy
	testStructure.RunTestStage(t, "teardown", func() {
		terraformOptions := testStructure.LoadTerraformOptions(t, fixtureFolder)
		terraform.Destroy(t, terraformOptions)
		testStructure.CleanupTestDataFolder(t, fixtureFolder)
	})
}
