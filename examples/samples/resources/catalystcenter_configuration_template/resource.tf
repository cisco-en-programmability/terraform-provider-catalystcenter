
terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

# data "catalystcenter_configuration_template" "example_source" {
#   provider       = catalystcenter
#   latest_version = "true"
#   template_id    = "2121f1d7-d9ea-4eea-8f83-45af2941f5f9"
# }

# output "catalystcenter_configuration_template_example_source" {
#   value = data.catalystcenter_configuration_template.example_source.item
# }

resource "catalystcenter_configuration_template" "example" {
  provider = catalystcenter
  parameters {
    project_id       = "e9df8f19-0cf2-4514-9703-cc049ca902e6"
    template_content = "if a > b \n hola22"
    language         = "JINJA"
    name             = "Saludo Terraform"
    software_type    = "IOS-XE"
    author           = "altus"
    description      = "Created by Terraform automation"
    device_types {
      product_family = "Switches and Hubs"
      product_series = "Cisco Catalyst 9300 Series Switches"
    }
  }

}

output "catalystcenter_configuration_template_example" {
  value = catalystcenter_configuration_template.example
}
