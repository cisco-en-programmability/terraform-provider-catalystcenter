
data "catalystcenter_event_series_audit_logs_parent_records" "example" {
    provider = catalystcenter
    category = "string"
    context = "string"
    description = "string"
    device_id = "string"
    domain = "string"
    end_time = 1609459200
    event_hierarchy = "string"
    event_id = "string"
    instance_id = "string"
    is_system_events = "false"
    limit = 1
    name = "string"
    offset = 1
    order = "string"
    severity = "string"
    site_id = "string"
    sort_by = "string"
    source = "string"
    start_time = 1609459200
    sub_domain = "string"
    user_id = "string"
}

output "catalystcenter_event_series_audit_logs_parent_records_example" {
    value = data.catalystcenter_event_series_audit_logs_parent_records.example.items
}
