terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
}

data "catalystcenter_pnp_workflow_count" "example" {
  provider = catalystcenter
  name     = ["string"]
}

output "catalystcenter_pnp_workflow_count_example" {
  value = data.catalystcenter_pnp_workflow_count.example.item
}
