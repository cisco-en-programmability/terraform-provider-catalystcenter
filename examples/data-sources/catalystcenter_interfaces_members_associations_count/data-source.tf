
data "catalystcenter_interfaces_members_associations_count" "example" {
  provider = catalystcenter
}

output "catalystcenter_interfaces_members_associations_count_example" {
  value = data.catalystcenter_interfaces_members_associations_count.example.item
}
