
data "catalystcenter_applications_count" "example" {
    provider = catalystcenter
}

output "catalystcenter_applications_count_example" {
    value = data.catalystcenter_applications_count.example.item
}
