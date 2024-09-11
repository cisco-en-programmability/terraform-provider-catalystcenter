
resource "catalystcenter_network_device" "example" {
  provider = catalystcenter

  parameters {

    id = "string"
  }
}

output "catalystcenter_network_device_example" {
  value = catalystcenter_network_device.example
}