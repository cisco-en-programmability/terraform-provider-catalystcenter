terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}
resource "catalystcenter_pnp_device_authorize" "example" {
  provider = catalystcenter
  parameters {

    device_id_list = ["618c2b7ff6b3b66f71acbaf8"]
  }
}

output "catalystcenter_pnp_device_authorize_example" {
  value = catalystcenter_pnp_device_authorize.example
}
