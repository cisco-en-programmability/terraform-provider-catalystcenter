
data "catalystcenter_tag_member_count" "example" {
    provider = catalystcenter
    id = "string"
    member_association_type = "string"
    member_type = "string"
}

output "catalystcenter_tag_member_count_example" {
    value = data.catalystcenter_tag_member_count.example.item
}
