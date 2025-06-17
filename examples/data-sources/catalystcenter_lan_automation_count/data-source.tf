
data "catalystcenter_lan_automation_count" "example" {
  provider = catalystcenter
}

output "catalystcenter_lan_automation_count_example" {
  value = data.catalystcenter_lan_automation_count.example.item
}
