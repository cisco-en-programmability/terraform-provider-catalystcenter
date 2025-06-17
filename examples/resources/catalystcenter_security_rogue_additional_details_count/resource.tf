
resource "catalystcenter_security_rogue_additional_details_count" "example" {
  provider = catalystcenter
  parameters = [{

    end_time     = 1609459200
    site_id      = ["string"]
    start_time   = 1609459200
    threat_level = ["string"]
    threat_type  = ["string"]
  }]
}

output "catalystcenter_security_rogue_additional_details_count_example" {
  value = catalystcenter_security_rogue_additional_details_count.example
}