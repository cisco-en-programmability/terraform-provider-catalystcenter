
resource "catalystcenter_itsm_integration_events_retry" "example" {
  provider   = meraki
  parameters = ["string"]
}

output "catalystcenter_itsm_integration_events_retry_example" {
  value = catalystcenter_itsm_integration_events_retry.example
}