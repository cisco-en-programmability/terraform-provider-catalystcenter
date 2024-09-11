terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

resource "catalystcenter_assign_device_to_site" "example" {
  provider = catalystcenter
  parameters {

    device {

      ip = "10.122.1.1"
    }
    site_id = "1ba55132-0fb8-4986-a31e-674d30e8b8ee"
  }
}

output "catalystcenter_assign_device_to_site_example" {
  value = catalystcenter_assign_device_to_site.example
}
