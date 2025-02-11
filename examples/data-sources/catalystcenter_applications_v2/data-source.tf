
data "catalystcenter_applications_v2" "example" {
    provider = catalystcenter
    attributes = "string"
    limit = 1
    name = "string"
    offset = 1
}

output "catalystcenter_applications_v2_example" {
    value = data.catalystcenter_applications_v2.example.items
}
