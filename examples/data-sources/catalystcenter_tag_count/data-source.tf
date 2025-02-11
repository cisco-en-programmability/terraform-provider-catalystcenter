
data "catalystcenter_tag_count" "example" {
    provider = catalystcenter
    attribute_name = "string"
    name = "string"
    name_space = "string"
    size = "string"
    system_tag = "string"
}

output "catalystcenter_tag_count_example" {
    value = data.catalystcenter_tag_count.example.item
}
