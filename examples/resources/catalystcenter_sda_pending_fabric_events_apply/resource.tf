
resource "catalystcenter_sda_pending_fabric_events_apply" "example" {
  provider = catalystcenter
  parameters = [{

    fabric_id = "string"
    id        = "string"
  }]
}

output "catalystcenter_sda_pending_fabric_events_apply_example" {
  value = catalystcenter_sda_pending_fabric_events_apply.example
}