
data "catalystcenter_event_artifact" "example" {
    provider = catalystcenter
    event_ids = "string"
    limit = 1
    offset = 1
    order = "string"
    search = "string"
    sort_by = "string"
    tags = "string"
}

output "catalystcenter_event_artifact_example" {
    value = data.catalystcenter_event_artifact.example.items
}
