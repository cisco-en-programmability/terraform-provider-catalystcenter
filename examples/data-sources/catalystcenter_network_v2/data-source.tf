
data "catalystcenter_network_v2" "example" {
  provider = catalystcenter
  site_id  = "string"
}

output "catalystcenter_network_v2_example" {
  value = data.catalystcenter_network_v2.example.items
}
