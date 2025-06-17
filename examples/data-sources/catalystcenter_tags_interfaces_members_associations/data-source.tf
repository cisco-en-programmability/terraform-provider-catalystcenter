
data "catalystcenter_tags_interfaces_members_associations" "example" {
  provider = catalystcenter
  limit    = 1
  offset   = 1
}

output "catalystcenter_tags_interfaces_members_associations_example" {
  value = data.catalystcenter_tags_interfaces_members_associations.example.items
}
