
data "catalystcenter_lan_automation_status" "example" {
  provider = catalystcenter
  limit    = 1
  offset   = 1
}

output "catalystcenter_lan_automation_status_example" {
  value = data.catalystcenter_lan_automation_status.example.items
}

data "catalystcenter_lan_automation_status" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_lan_automation_status_example" {
  value = data.catalystcenter_lan_automation_status.example.item
}
