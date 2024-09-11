
data "catalystcenter_event_api_status" "example" {
  provider     = catalystcenter
  execution_id = "string"
}

output "catalystcenter_event_api_status_example" {
  value = data.catalystcenter_event_api_status.example.item
}
