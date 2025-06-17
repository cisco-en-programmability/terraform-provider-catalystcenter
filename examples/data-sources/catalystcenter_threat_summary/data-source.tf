
data "catalystcenter_threat_summary" "example" {
  provider = catalystcenter
  end_time = 1
  site_id      = ["string"]
  start_time   = 1
  threat_level = ["string"]
  threat_type  = ["string"]
}