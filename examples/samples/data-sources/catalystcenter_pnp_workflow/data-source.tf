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
data "catalystcenter_pnp_workflow" "example" {
  provider   = catalystcenter
  limit      = 1
  name       = ["string"]
  offset     = 1
  sort       = ["string"]
  sort_order = "string"
  type       = ["string"]
}

output "catalystcenter_pnp_workflow_example" {
  value = data.catalystcenter_pnp_workflow.example.items
}
*/
data "catalystcenter_pnp_workflow" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_pnp_workflow_example" {
  value = data.catalystcenter_pnp_workflow.example.item
}
