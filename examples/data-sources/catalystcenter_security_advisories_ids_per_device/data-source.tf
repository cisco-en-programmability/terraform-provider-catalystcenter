
data "catalystcenter_security_advisories_ids_per_device" "example" {
  provider  = catalystcenter
  device_id = "string"
}

output "catalystcenter_security_advisories_ids_per_device_example" {
  value = data.catalystcenter_security_advisories_ids_per_device.example.item
}
