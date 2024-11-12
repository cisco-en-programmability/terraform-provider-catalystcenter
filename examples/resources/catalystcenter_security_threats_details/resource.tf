
resource "catalystcenter_security_threats_details" "example" {
  provider = meraki
  parameters {

    end_time      = 1
    is_new_threat = "false"
    limit         = 1
    offset        = 1
    site_id       = ["string"]
    start_time    = 1
    threat_level  = ["string"]
    threat_type   = ["string"]
  }
}

output "catalystcenter_security_threats_details_example" {
  value = catalystcenter_security_threats_details.example
}