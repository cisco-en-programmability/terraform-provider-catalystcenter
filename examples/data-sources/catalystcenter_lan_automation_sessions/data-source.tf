
data "catalystcenter_lan_automation_sessions" "example" {
  provider = catalystcenter
}

output "catalystcenter_lan_automation_sessions_example" {
  value = data.catalystcenter_lan_automation_sessions.example.item
}
