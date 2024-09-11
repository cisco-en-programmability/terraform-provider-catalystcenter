
data "catalystcenter_device_interface" "example" {
  provider = catalystcenter
  limit    = 1
  offset   = 1
}

output "catalystcenter_device_interface_example" {
  value = data.catalystcenter_device_interface.example.items
}

data "catalystcenter_device_interface" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_device_interface_example" {
  value = data.catalystcenter_device_interface.example.item
}
