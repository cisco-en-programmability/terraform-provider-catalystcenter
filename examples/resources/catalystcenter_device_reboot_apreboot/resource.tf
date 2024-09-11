
resource "catalystcenter_device_reboot_apreboot" "example" {
  provider = catalystcenter

  parameters {

    ap_mac_addresses = ["string"]
  }
}

output "catalystcenter_device_reboot_apreboot_example" {
  value = catalystcenter_device_reboot_apreboot.example
}