
terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_application_sets" "example" {
  provider = catalystcenter
  parameters {
    payload {
      name = "TestAppTerraform222Update"
    }
  }
}

output "catalystcenter_application_sets_example" {
  value = catalystcenter_application_sets.example
}
