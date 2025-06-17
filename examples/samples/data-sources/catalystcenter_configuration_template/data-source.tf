terraform {
  required_providers {
    catalystcenter = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
}

/*
data "catalystcenter_configuration_template" "example" {
  provider                     = catalystcenter
  filter_conflicting_templates = "false"
  product_family               = "string"
  product_series               = "string"
  product_type                 = "string"
  project_id                   = "string"
  project_names                = ["string"]
  software_type                = "string"
  software_version             = "string"
  sort_order                   = "asc"
  tags                         = ["string"]
  un_committed                 = "false"
}

output "catalystcenter_configuration_template_example" {
  value = data.catalystcenter_configuration_template.example.items
}
*/

data "catalystcenter_configuration_template" "example" {
  provider       = catalystcenter
  latest_version = "true"
  template_id    = "2121f1d7-d9ea-4eea-8f83-45af2941f5f9"
}

output "catalystcenter_configuration_template_example" {
  value = data.catalystcenter_configuration_template.example.item
}
