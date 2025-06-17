
data "catalystcenter_network" "example" {
  provider = catalystcenter
  site_id  = "string"
}

output "catalystcenter_network_example" {
  value = data.catalystcenter_network.example.items
}
