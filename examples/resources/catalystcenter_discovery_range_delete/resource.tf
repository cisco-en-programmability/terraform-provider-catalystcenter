
resource "catalystcenter_discovery_range_delete" "example" {
  provider          = meraki
  records_to_delete = 1
  start_index       = 1
  parameters {

  }
}

output "catalystcenter_discovery_range_delete_example" {
  value = catalystcenter_discovery_range_delete.example
}