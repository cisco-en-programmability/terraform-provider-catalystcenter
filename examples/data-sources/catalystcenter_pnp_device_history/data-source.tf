
data "catalystcenter_pnp_device_history" "example" {
  provider      = catalystcenter
  serial_number = "string"
  sort          = ["string"]
  sort_order    = "string"
}

output "catalystcenter_pnp_device_history_example" {
  value = data.catalystcenter_pnp_device_history.example.items
}
