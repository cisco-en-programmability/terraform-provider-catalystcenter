
data "catalystcenter_network_device_meraki_organization" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_network_device_meraki_organization_example" {
  value = data.catalystcenter_network_device_meraki_organization.example.items
}
