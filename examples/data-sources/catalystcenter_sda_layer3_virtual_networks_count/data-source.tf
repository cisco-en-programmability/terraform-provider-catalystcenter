
data "catalystcenter_sda_layer3_virtual_networks_count" "example" {
  provider         = catalystcenter
  anchored_site_id = "string"
  fabric_id        = "string"
}

output "catalystcenter_sda_layer3_virtual_networks_count_example" {
  value = data.catalystcenter_sda_layer3_virtual_networks_count.example.item
}
