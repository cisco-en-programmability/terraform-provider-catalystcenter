
data "catalystcenter_lan_automation_log" "example" {
    provider = catalystcenter
    limit = 1
    offset = 1
}

output "catalystcenter_lan_automation_log_example" {
    value = data.catalystcenter_lan_automation_log.example.items
}

data "catalystcenter_lan_automation_log" "example" {
    provider = catalystcenter
    id = "string"
}

output "catalystcenter_lan_automation_log_example" {
    value = data.catalystcenter_lan_automation_log.example.item
}
