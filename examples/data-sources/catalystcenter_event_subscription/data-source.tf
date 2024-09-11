
data "catalystcenter_event_subscription" "example" {
  provider  = catalystcenter
  event_ids = "string"
  limit     = 1
  offset    = 1
  order     = "string"
  sort_by   = "string"
}

output "catalystcenter_event_subscription_example" {
  value = data.catalystcenter_event_subscription.example.items
}
