
data "catalystcenter_itsm_integration_events_failed" "example" {
  provider    = catalystcenter
  instance_id = "string"
}

output "catalystcenter_itsm_integration_events_failed_example" {
  value = data.catalystcenter_itsm_integration_events_failed.example.items
}
