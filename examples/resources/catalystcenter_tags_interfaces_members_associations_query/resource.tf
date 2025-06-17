
resource "catalystcenter_tags_interfaces_members_associations_query" "example" {
  provider = catalystcenter
  parameters = [{

    ids = ["string"]
  }]
}

output "catalystcenter_tags_interfaces_members_associations_query_example" {
  value = catalystcenter_tags_interfaces_members_associations_query.example
}