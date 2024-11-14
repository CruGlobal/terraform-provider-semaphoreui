package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"strconv"
	"terraform-provider-semaphoreui/semaphoreui/client/project"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccProjectEnvironmentExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		if rs.Primary.Attributes["id"] == "" {
			return fmt.Errorf("no ID is set")
		}
		if rs.Primary.Attributes["project_id"] == "" {
			return fmt.Errorf("no ProjectID is set")
		}

		id, _ := strconv.ParseInt(rs.Primary.Attributes["id"], 10, 64)
		projectId, _ := strconv.ParseInt(rs.Primary.Attributes["project_id"], 10, 64)

		response, err := testClient().Project.GetProjectProjectIDEnvironmentEnvironmentID(&project.GetProjectProjectIDEnvironmentEnvironmentIDParams{
			ProjectID:     projectId,
			EnvironmentID: id,
		}, nil)
		if err != nil {
			return fmt.Errorf("error reading project environment: %s", err.Error())
		}

		if rs.Primary.Attributes["name"] != response.Payload.Name {
			return fmt.Errorf("environment name mismatch: %s != %s", rs.Primary.Attributes["name"], response.Payload.Name)
		}

		return nil
	}
}

func testAccProjectEnvironmentEmptyConfig(nameSuffix string) string {
	return fmt.Sprintf(`
resource "semaphoreui_project" "test" {
  name = "test-%[1]s"
}
`, nameSuffix)
}

type testAccProjectEnvironmentSecret struct {
	Name  string
	Value string
	Type  string
}

func testAccProjectEnvironmentConfig(
	nameSuffix string,
	variables *map[string]string,
	environment *map[string]string,
	secrets *[]testAccProjectEnvironmentSecret,
) string {
	var vars, envs, secs string
	if variables != nil {
		for k, v := range *variables {
			vars += fmt.Sprintf("%s = \"%s\"\n", k, v)
		}
		vars = fmt.Sprintf(`variables = {
      %s
  }`, vars)
	}

	if environment != nil {
		for k, v := range *environment {
			envs += fmt.Sprintf("%s = \"%s\"\n", k, v)
		}
		envs = fmt.Sprintf(`environment = {
	  %s
  }`, envs)
	}

	if secrets != nil {
		for _, s := range *secrets {
			secs += fmt.Sprintf(`{
	  name = "%s"
	  value = "%s"
      type = "%s"
	},`, s.Name, s.Value, s.Type)
		}
		secs = fmt.Sprintf(`secrets = [
	%s
  ]`, secs)
	}

	return fmt.Sprintf(`
%[1]s
resource "semaphoreui_project_environment" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Test %[2]s"
  %[3]s
  %[4]s
  %[5]s
}`, testAccProjectEnvironmentEmptyConfig(nameSuffix), nameSuffix, vars, envs, secs)
}

func testAccProjectEnvironmentImportID(n string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return "", fmt.Errorf("not found: %s", n)
		}

		return fmt.Sprintf("project/%[1]s/environment/%[2]s", rs.Primary.Attributes["project_id"], rs.Primary.Attributes["id"]), nil
	}
}

func TestAcc_ProjectEnvironmentResource_basic(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectEnvironmentConfig(nameSuffix, nil, nil, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectEnvironmentExists("semaphoreui_project_environment.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "name", fmt.Sprintf("Test %s", nameSuffix)),

					resource.TestCheckNoResourceAttr("semaphoreui_project_environment.test", "variables"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_environment.test", "environment"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_environment.test", "secrets"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "project_id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project_environment.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectEnvironmentImportID("semaphoreui_project_environment.test"),
			},
			// Update and Read testing
			{
				Config: testAccProjectEnvironmentConfig(nameSuffix, &map[string]string{}, &map[string]string{}, &[]testAccProjectEnvironmentSecret{}),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectEnvironmentExists("semaphoreui_project_environment.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "name", fmt.Sprintf("Test %s", nameSuffix)),

					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "variables.%", "0"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "environment.%", "0"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "secrets.#", "0"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "project_id"),
				),
			},
			// Delete testing
			{
				Config: testAccProjectEnvironmentEmptyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_environment.test"),
				),
			},
		},
	})
}

func TestAcc_ProjectEnvironmentResource_basicVariables(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectEnvironmentConfig(nameSuffix, &map[string]string{"lorem": "ipsum", "dolor": "sit"}, nil, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectEnvironmentExists("semaphoreui_project_environment.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "name", fmt.Sprintf("Test %s", nameSuffix)),

					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "variables.%", "2"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "variables.lorem", "ipsum"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "variables.dolor", "sit"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_environment.test", "environment"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_environment.test", "secrets"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "project_id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project_environment.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectEnvironmentImportID("semaphoreui_project_environment.test"),
			},
			// Update and Read testing
			{
				Config: testAccProjectEnvironmentConfig(nameSuffix, &map[string]string{"dolor": "sit", "amet": "tempor"}, nil, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectEnvironmentExists("semaphoreui_project_environment.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "name", fmt.Sprintf("Test %s", nameSuffix)),

					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "variables.%", "2"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "variables.amet", "tempor"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "variables.dolor", "sit"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_environment.test", "environment"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_environment.test", "secrets"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "project_id"),
				),
			},
			// Delete testing
			{
				Config: testAccProjectEnvironmentEmptyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_environment.test"),
				),
			},
		},
	})
}

func TestAcc_ProjectEnvironmentResource_basicEnvironment(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectEnvironmentConfig(nameSuffix, nil, &map[string]string{"FOO": "BAR", "BAZ": "QUX"}, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectEnvironmentExists("semaphoreui_project_environment.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "name", fmt.Sprintf("Test %s", nameSuffix)),

					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "environment.%", "2"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "environment.FOO", "BAR"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "environment.BAZ", "QUX"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_environment.test", "variables"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_environment.test", "secrets"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "project_id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project_environment.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectEnvironmentImportID("semaphoreui_project_environment.test"),
			},
			// Update and Read testing
			{
				Config: testAccProjectEnvironmentConfig(nameSuffix, nil, &map[string]string{"BAZ": "QUX", "LOREM": "IPSUM"}, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectEnvironmentExists("semaphoreui_project_environment.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "name", fmt.Sprintf("Test %s", nameSuffix)),

					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "environment.%", "2"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "environment.LOREM", "IPSUM"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "environment.BAZ", "QUX"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_environment.test", "variables"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_environment.test", "secrets"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "project_id"),
				),
			},
			// Delete testing
			{
				Config: testAccProjectEnvironmentEmptyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_environment.test"),
				),
			},
		},
	})
}

func TestAcc_ProjectEnvironmentResource_basicSecrets(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectEnvironmentConfig(nameSuffix, nil, nil, &[]testAccProjectEnvironmentSecret{
					{Name: "FOO", Value: "BAR", Type: "var"},
					{Name: "BAZ", Value: "QUX", Type: "env"},
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectEnvironmentExists("semaphoreui_project_environment.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "name", fmt.Sprintf("Test %s", nameSuffix)),

					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "secrets.#", "2"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "secrets.0.id"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "secrets.0.name", "FOO"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "secrets.0.value", "BAR"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "secrets.0.type", "var"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "secrets.1.id"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "secrets.1.name", "BAZ"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "secrets.1.value", "QUX"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "secrets.1.type", "env"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_environment.test", "variables"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_environment.test", "environment"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "project_id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project_environment.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectEnvironmentImportID("semaphoreui_project_environment.test"),
				// Secret values can't be imported and are set to empty strings
				ImportStateVerifyIgnore: []string{"secrets.0.value", "secrets.1.value"},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "secrets.0.value", ""),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "secrets.1.value", ""),
				),
			},
			// Update and Read testing
			{
				// Update names and values
				Config: testAccProjectEnvironmentConfig(nameSuffix, nil, nil, &[]testAccProjectEnvironmentSecret{
					{Name: "NAME", Value: "BAR", Type: "var"},
					{Name: "BAZ", Value: "VALUE", Type: "env"},
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectEnvironmentExists("semaphoreui_project_environment.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "name", fmt.Sprintf("Test %s", nameSuffix)),

					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "secrets.#", "2"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "secrets.0.id"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "secrets.0.name", "NAME"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "secrets.0.value", "BAR"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "secrets.0.type", "var"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "secrets.1.id"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "secrets.1.name", "BAZ"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "secrets.1.value", "VALUE"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "secrets.1.type", "env"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_environment.test", "variables"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_environment.test", "environment"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "project_id"),
				),
			},
			// Update and Read testing
			{
				// Change type and delete
				Config: testAccProjectEnvironmentConfig(nameSuffix, nil, nil, &[]testAccProjectEnvironmentSecret{
					{Name: "NAME", Value: "BAR", Type: "env"},
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectEnvironmentExists("semaphoreui_project_environment.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "name", fmt.Sprintf("Test %s", nameSuffix)),

					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "secrets.#", "1"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "secrets.0.id"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "secrets.0.name", "NAME"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "secrets.0.value", "BAR"),
					resource.TestCheckResourceAttr("semaphoreui_project_environment.test", "secrets.0.type", "env"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_environment.test", "variables"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_environment.test", "environment"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_environment.test", "project_id"),
				),
			},
			// Delete testing
			{
				Config: testAccProjectEnvironmentEmptyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_environment.test"),
				),
			},
		},
	})
}
