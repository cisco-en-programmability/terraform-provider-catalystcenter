
data "catalystcenter_sda_transit_networks_count" "example" {
  provider = catalystcenter
  type     = "string"
}

output "catalystcenter_sda_transit_networks_count_example" {
  value = data.catalystcenter_sda_transit_networks_count.example.item
}
