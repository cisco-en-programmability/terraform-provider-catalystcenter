terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.1-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
}
/*
data "catalystcenter_reports" "example" {
  provider      = catalystcenter
  view_group_id = "string"
  view_id       = "string"
}

output "catalystcenter_reports_example" {
  value = data.catalystcenter_reports.example.items
}
*/
data "catalystcenter_reports" "example" {
  provider  = catalystcenter
  report_id = "string"
}

output "catalystcenter_reports_example" {
  value = data.catalystcenter_reports.example.item
}
