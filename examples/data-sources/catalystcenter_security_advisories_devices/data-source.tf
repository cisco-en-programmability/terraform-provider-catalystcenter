
data "catalystcenter_security_advisories_devices" "example" {
  provider    = catalystcenter
  advisory_id = "string"
}

output "catalystcenter_security_advisories_devices_example" {
  value = data.catalystcenter_security_advisories_devices.example.items
}
