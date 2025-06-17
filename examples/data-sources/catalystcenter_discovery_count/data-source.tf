
data "catalystcenter_discovery_count" "example" {
  provider = catalystcenter
}

output "catalystcenter_discovery_count_example" {
  value = data.catalystcenter_discovery_count.example.item
}
