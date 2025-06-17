
data "catalystcenter_tags_interfaces_members_associations_count" "example" {
  provider = catalystcenter
}

output "catalystcenter_tags_interfaces_members_associations_count_example" {
  value = data.catalystcenter_tags_interfaces_members_associations_count.example.item
}
