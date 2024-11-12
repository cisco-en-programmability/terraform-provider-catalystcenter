
resource "catalystcenter_pnp_device_authorize" "example" {
  provider = meraki
  parameters {

    device_id_list = ["string"]
  }
}

output "catalystcenter_pnp_device_authorize_example" {
  value = catalystcenter_pnp_device_authorize.example
}