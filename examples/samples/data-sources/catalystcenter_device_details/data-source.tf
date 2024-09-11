
data "catalystcenter_device_details" "example" {
  provider   = catalystcenter
  identifier = "string"
  search_by  = "string"
  timestamp  = "string"
}

output "catalystcenter_device_details_example" {
  value = data.catalystcenter_device_details.example.item
}
