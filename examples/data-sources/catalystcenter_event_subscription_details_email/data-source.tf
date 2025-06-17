
data "catalystcenter_event_subscription_details_email" "example" {
  provider    = catalystcenter
  instance_id = "string"
  limit       = 1
  name        = "string"
  offset      = 1
  order       = "string"
  sort_by     = "string"
}

output "catalystcenter_event_subscription_details_email_example" {
  value = data.catalystcenter_event_subscription_details_email.example.items
}
