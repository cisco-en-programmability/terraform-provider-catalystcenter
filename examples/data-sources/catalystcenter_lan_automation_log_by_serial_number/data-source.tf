
data "catalystcenter_lan_automation_log_by_serial_number" "example" {
    provider = catalystcenter
    id = "string"
    log_level = "string"
    serial_number = "string"
}

output "catalystcenter_lan_automation_log_by_serial_number_example" {
    value = data.catalystcenter_lan_automation_log_by_serial_number.example.items
}
