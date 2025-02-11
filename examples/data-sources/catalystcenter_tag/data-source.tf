
data "catalystcenter_tag" "example" {
    provider = catalystcenter
    additional_info_attributes = "string"
    additional_info_name_space = "string"
    field = "string"
    level = "string"
    limit = 1
    name = "string"
    offset = 1
    order = "string"
    size = "string"
    sort_by = "string"
    system_tag = "string"
}

output "catalystcenter_tag_example" {
    value = data.catalystcenter_tag.example.items
}

data "catalystcenter_tag" "example" {
    provider = catalystcenter
    id = "string"
}

output "catalystcenter_tag_example" {
    value = data.catalystcenter_tag.example.item
}
