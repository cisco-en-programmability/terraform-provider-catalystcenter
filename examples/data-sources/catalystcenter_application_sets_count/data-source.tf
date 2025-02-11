
data "catalystcenter_application_sets_count" "example" {
    provider = catalystcenter
}

output "catalystcenter_application_sets_count_example" {
    value = data.catalystcenter_application_sets_count.example.item
}
