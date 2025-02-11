
data "catalystcenter_event_count" "example" {
    provider = catalystcenter
    event_id = "string"
    tags = "string"
}

output "catalystcenter_event_count_example" {
    value = data.catalystcenter_event_count.example.item
}
