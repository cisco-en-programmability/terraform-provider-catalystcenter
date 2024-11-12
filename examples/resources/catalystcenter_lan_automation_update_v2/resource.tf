
resource "catalystcenter_lan_automation_update_v2" "example" {
  provider = meraki
  id       = "string"
  parameters {

    device_management_ipaddress = "string"
    new_host_name               = "string"
    new_loopback0_ipaddress     = "string"
  }
}

output "catalystcenter_lan_automation_update_v2_example" {
  value = catalystcenter_lan_automation_update_v2.example
}