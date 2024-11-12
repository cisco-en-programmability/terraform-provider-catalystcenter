
resource "catalystcenter_pnp_device_unclaim" "example" {
  provider = meraki
  parameters {

    device_id_list = ["string"]
  }
}

output "catalystcenter_pnp_device_unclaim_example" {
  value = catalystcenter_pnp_device_unclaim.example
}