
data "catalystcenter_applications" "example" {
    provider = catalystcenter
    limit = 1
    name = "string"
    offset = 1
}

output "catalystcenter_applications_example" {
    value = data.catalystcenter_applications.example.items
}
