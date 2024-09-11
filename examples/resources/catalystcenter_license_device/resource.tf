
resource "catalystcenter_license_device" "example" {
  provider = catalystcenter

  parameters {

    device_uuids         = ["string"]
    smart_account_id     = "string"
    virtual_account_name = "string"
  }
}

output "catalystcenter_license_device_example" {
  value = catalystcenter_license_device.example
}