terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

data "catalystcenter_event_series_audit_logs_parent_records" "example" {
  provider = catalystcenter
  /*category         = "string"
  context          = "string"
  description      = "string"
  device_id        = "string"
  domain           = "string"
  end_time         = "hh:mm"
  event_hierarchy  = "string"
  event_id         = "string"
  instance_id      = "string"
  is_system_events = "false"
  limit            = "#"
  name             = "string"
  offset           = "#"
  order            = "string"
  severity         = "string"
  site_id          = "string"
  sort_by          = "string"
  source           = "string"
  start_time       = "hh:mm"
  sub_domain       = "string"
  user_id          = "string"*/
}

output "catalystcenter_event_series_audit_logs_parent_records_example" {
  value = data.catalystcenter_event_series_audit_logs_parent_records.example.items
}
