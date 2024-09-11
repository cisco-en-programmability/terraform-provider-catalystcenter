
data "catalystcenter_file" "example" {
  provider = catalystcenter
  dirpath  = "string"
  file_id  = "string"
}

output "catalystcenter_file_example" {
  value = data.catalystcenter_file.example.item
}
