
data "catalystcenter_interface" "example" {
  provider       = catalystcenter
  interface_uuid = "string"
}

output "catalystcenter_interface_example" {
  value = data.catalystcenter_interface.example.item
}
