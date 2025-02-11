
data "catalystcenter_event_syslog_config" "example" {
    provider = catalystcenter
    config_id = "string"
    limit = 1
    name = "string"
    offset = 1
    order = "string"
    protocol = "string"
    sort_by = "string"
}

output "catalystcenter_event_syslog_config_example" {
    value = data.catalystcenter_event_syslog_config.example.item
}
