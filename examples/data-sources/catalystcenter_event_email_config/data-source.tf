
data "catalystcenter_event_email_config" "example" {
  provider = catalystcenter
}

output "catalystcenter_event_email_config_example" {
  value = data.catalystcenter_event_email_config.example.items
}
