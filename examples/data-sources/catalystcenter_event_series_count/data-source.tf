
data "catalystcenter_event_series_count" "example" {
    provider = catalystcenter
    category = "string"
    domain = "string"
    end_time = 1609459200
    event_ids = "string"
    severity = "string"
    source = "string"
    start_time = 1609459200
    sub_domain = "string"
    type = "string"
}

output "catalystcenter_event_series_count_example" {
    value = data.catalystcenter_event_series_count.example.item
}
