
data "catalystcenter_sda_transit_networks" "example" {
  provider = catalystcenter
  id       = "string"
  limit    = 1
  name     = "string"
  offset   = 1
  type     = "string"
}

output "catalystcenter_sda_transit_networks_example" {
  value = data.catalystcenter_sda_transit_networks.example.items
}
