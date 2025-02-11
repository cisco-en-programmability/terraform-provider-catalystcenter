
data "catalystcenter_event_series" "example" {
    provider = catalystcenter
    category = "string"
    domain = "string"
    end_time = 1609459200
    event_ids = "string"
    limit = 1
    namespace = "string"
    offset = 1
    order = "string"
    severity = "string"
    site_id = "string"
    sort_by = "string"
    source = "string"
    start_time = 1609459200
    sub_domain = "string"
    tags = "string"
    type = "string"
}

output "catalystcenter_event_series_example" {
    value = data.catalystcenter_event_series.example.items
}
