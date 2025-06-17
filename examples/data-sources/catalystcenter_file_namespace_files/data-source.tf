
data "catalystcenter_file_namespace_files" "example" {
  provider   = catalystcenter
  name_space = "string"
}

output "catalystcenter_file_namespace_files_example" {
  value = data.catalystcenter_file_namespace_files.example.items
}
