
resource "catalystcenter_tags_network_devices_members_associations_query" "example" {
  provider = catalystcenter
  parameters = [{

    ids = ["string"]
  }]
}

output "catalystcenter_tags_network_devices_members_associations_query_example" {
  value = catalystcenter_tags_network_devices_members_associations_query.example
}