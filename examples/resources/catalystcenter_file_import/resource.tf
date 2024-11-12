
resource "catalystcenter_file_import" "example" {
  provider   = meraki
  name_space = "string"
  parameters {

  }
}

output "catalystcenter_file_import_example" {
  value = catalystcenter_file_import.example
}