
data "catalystcenter_event_artifact_count" "example" {
    provider = catalystcenter
}

output "catalystcenter_event_artifact_count_example" {
    value = data.catalystcenter_event_artifact_count.example.item
}
