
resource "catalystcenter_security_threats_summary" "example" {
  provider = catalystcenter
  parameters = [{

    end_time     = 1
    site_id      = ["string"]
    start_time   = 1
    threat_level = ["string"]
    threat_type  = ["string"]
  }]
}

output "catalystcenter_security_threats_summary_example" {
  value = catalystcenter_security_threats_summary.example
}