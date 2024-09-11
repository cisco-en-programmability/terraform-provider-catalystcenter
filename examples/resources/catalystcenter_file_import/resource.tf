
resource "catalystcenter_file_import" "example" {
  provider   = catalystcenter
  file_name  = "string"
  file_path  = "string"
  name_space = "string"
  parameters {

  }
}

output "catalystcenter_file_import_example" {
  value = catalystcenter_file_import.example
}