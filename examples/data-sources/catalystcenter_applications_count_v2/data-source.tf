
data "catalystcenter_applications_count_v2" "example" {
    provider = catalystcenter
    scalable_group_type = "string"
}

output "catalystcenter_applications_count_v2_example" {
    value = data.catalystcenter_applications_count_v2.example.item
}
