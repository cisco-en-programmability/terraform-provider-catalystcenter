
data "catalystcenter_sda_fabric_sites_count" "example" {
  provider = catalystcenter
}

output "catalystcenter_sda_fabric_sites_count_example" {
  value = data.catalystcenter_sda_fabric_sites_count.example.item
}
