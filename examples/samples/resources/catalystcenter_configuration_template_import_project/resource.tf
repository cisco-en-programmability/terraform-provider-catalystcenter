terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

resource "catalystcenter_configuration_template_import_project" "example" {
  provider = catalystcenter

  parameters {
    do_version = "false"
  }
}
