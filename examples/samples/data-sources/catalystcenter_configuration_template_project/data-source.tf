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


data "catalystcenter_configuration_template_project" "example" {
  provider   = catalystcenter
  name       = "Onboarding Configuration"
  sort_order = "asc"
}

output "catalystcenter_configuration_template_project_example" {
  value = data.catalystcenter_configuration_template_project.example.items
}
/*
data "catalystcenter_configuration_template_project" "example" {
  provider   = catalystcenter
  project_id = "2128b364-8751-45f5-95cb-f5ecaa2e9085"
}

output "catalystcenter_configuration_template_project_example" {
  value = data.catalystcenter_configuration_template_project.example.item
}
*/
