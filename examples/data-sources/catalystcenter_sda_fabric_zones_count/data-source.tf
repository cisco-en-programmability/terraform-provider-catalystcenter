
data "catalystcenter_sda_fabric_zones_count" "example" {
  provider = catalystcenter
}

output "catalystcenter_sda_fabric_zones_count_example" {
  value = data.catalystcenter_sda_fabric_zones_count.example.item
}
