
data "catalystcenter_event_subscription_syslog" "example" {
  provider   = catalystcenter
  category   = "string"
  domain     = "string"
  event_ids  = "string"
  limit      = 1
  name       = "string"
  offset     = 1
  order      = "string"
  sort_by    = "string"
  sub_domain = "string"
  type       = "string"
}

output "catalystcenter_event_subscription_syslog_example" {
  value = data.catalystcenter_event_subscription_syslog.example.items
}
