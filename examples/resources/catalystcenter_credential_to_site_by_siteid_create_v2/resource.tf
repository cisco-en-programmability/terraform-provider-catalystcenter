
resource "catalystcenter_credential_to_site_by_siteid_create_v2" "example" {
    provider = meraki
    site_id = "string"
    parameters = [{
      
      cli_id = "string"
      http_read = "string"
      http_write = "string"
      snmp_v2_read_id = "string"
      snmp_v2_write_id = "string"
      snmp_v3_id = "string"
    }]
}

output "catalystcenter_credential_to_site_by_siteid_create_v2_example" {
    value = catalystcenter_credential_to_site_by_siteid_create_v2.example
}