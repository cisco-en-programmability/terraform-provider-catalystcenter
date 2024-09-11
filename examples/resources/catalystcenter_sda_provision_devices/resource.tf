
resource "catalystcenter_sda_provision_devices" "example" {
  provider = catalystcenter

  parameters {

    id                = "string"
    network_device_id = "string"
    site_id           = "string"
  }
}

output "catalystcenter_sda_provision_devices_example" {
  value = catalystcenter_sda_provision_devices.example
}