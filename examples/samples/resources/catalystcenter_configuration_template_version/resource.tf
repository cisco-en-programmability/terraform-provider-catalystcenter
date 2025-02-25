terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.1-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_configuration_template_version" "example" {
  provider = catalystcenter
  parameters {
    # comments = "string"
    template_id = "524d6bfd-45df-4399-9bbf-3e6f4006b009"
  }
}

output "catalystcenter_configuration_template_version_example" {
  value = catalystcenter_configuration_template_version.example
}
