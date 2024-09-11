
data "catalystcenter_network_devices_members_associations_count" "example" {
  provider = catalystcenter
}

output "catalystcenter_network_devices_members_associations_count_example" {
  value = data.catalystcenter_network_devices_members_associations_count.example.item
}
