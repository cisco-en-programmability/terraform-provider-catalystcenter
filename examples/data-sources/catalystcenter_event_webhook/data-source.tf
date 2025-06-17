
data "catalystcenter_event_webhook" "example" {
  provider    = catalystcenter
  limit       = 1
  offset      = 1
  order       = "string"
  sort_by     = "string"
  webhook_ids = "string"
}

output "catalystcenter_event_webhook_example" {
  value = data.catalystcenter_event_webhook.example.item
}
