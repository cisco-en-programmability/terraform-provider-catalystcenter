
data "catalystcenter_sda_multicast_virtual_networks_count" "example" {
  provider  = catalystcenter
  fabric_id = "string"
}

output "catalystcenter_sda_multicast_virtual_networks_count_example" {
  value = data.catalystcenter_sda_multicast_virtual_networks_count.example.item
}
