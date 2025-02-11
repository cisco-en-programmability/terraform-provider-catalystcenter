
data "catalystcenter_assurance_events_count" "example" {
    provider = catalystcenter
    ap_mac = "string"
    client_mac = "string"
    device_family = "string"
    end_time = "string"
    message_type = "string"
    network_device_id = "string"
    network_device_name = "string"
    severity = "string"
    site_hierarchy_id = "string"
    site_id = "string"
    start_time = "string"
    xca_lle_rid = "string"
}

output "catalystcenter_assurance_events_count_example" {
    value = data.catalystcenter_assurance_events_count.example.item
}
