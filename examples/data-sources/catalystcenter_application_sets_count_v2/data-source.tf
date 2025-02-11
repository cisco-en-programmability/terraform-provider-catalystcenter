
data "catalystcenter_application_sets_count_v2" "example" {
    provider = catalystcenter
    scalable_group_type = "string"
}

output "catalystcenter_application_sets_count_v2_example" {
    value = data.catalystcenter_application_sets_count_v2.example.item
}
