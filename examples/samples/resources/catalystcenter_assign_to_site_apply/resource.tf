terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.1-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

resource "catalystcenter_assign_to_site_apply" "example" {
  provider = catalystcenter
  parameters {

    device_ids = ["string"]
    site_id    = "string"
  }
}

output "catalystcenter_assign_to_site_apply_example" {
  value = catalystcenter_assign_to_site_apply.example
}
