
data "catalystcenter_sda_pending_fabric_events" "example" {
    provider = catalystcenter
    fabric_id = "string"
    limit = 1
    offset = 1
}

output "catalystcenter_sda_pending_fabric_events_example" {
    value = data.catalystcenter_sda_pending_fabric_events.example.items
}
