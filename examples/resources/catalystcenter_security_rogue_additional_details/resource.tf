
resource "catalystcenter_security_rogue_additional_details" "example" {
  provider = catalystcenter
  parameters = [{

    end_time     = 1609459200
    limit        = 1
    offset       = 1
    site_id      = ["string"]
    start_time   = 1609459200
    threat_level = ["string"]
    threat_type  = ["string"]
  }]
}

output "catalystcenter_security_rogue_additional_details_example" {
  value = catalystcenter_security_rogue_additional_details.example
}