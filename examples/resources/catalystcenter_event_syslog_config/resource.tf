
resource "catalystcenter_event_syslog_config" "example" {
  provider = catalystcenter

  parameters {

    config_id   = "string"
    description = "string"
    host        = "string"
    name        = "string"
    port        = 1
    protocol    = "string"
  }
}

output "catalystcenter_event_syslog_config_example" {
  value = catalystcenter_event_syslog_config.example
}