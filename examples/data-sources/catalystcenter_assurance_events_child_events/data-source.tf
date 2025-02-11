
data "catalystcenter_assurance_events_child_events" "example" {
    provider = catalystcenter
    id = "string"
    xca_lle_rid = "string"
}

output "catalystcenter_assurance_events_child_events_example" {
    value = data.catalystcenter_assurance_events_child_events.example.items
}
