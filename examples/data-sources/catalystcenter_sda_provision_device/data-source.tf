
data "catalystcenter_sda_provision_device" "example" {
  provider                     = catalystcenter
  device_management_ip_address = "string"
}

output "catalystcenter_sda_provision_device_example" {
  value = data.catalystcenter_sda_provision_device.example.item
}
