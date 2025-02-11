
data "catalystcenter_tag_member_type" "example" {
    provider = catalystcenter
}

output "catalystcenter_tag_member_type_example" {
    value = data.catalystcenter_tag_member_type.example.items
}
