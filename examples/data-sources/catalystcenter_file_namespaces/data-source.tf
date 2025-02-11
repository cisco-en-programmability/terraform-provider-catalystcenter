
data "catalystcenter_file_namespaces" "example" {
    provider = catalystcenter
}

output "catalystcenter_file_namespaces_example" {
    value = data.catalystcenter_file_namespaces.example.items
}
