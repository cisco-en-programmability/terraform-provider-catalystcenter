
resource "catalystcenter_site_assign_credential" "example" {
  provider = catalystcenter
  site_id  = "string"
  parameters {

    cli_id           = "string"
    http_read        = "string"
    http_write       = "string"
    snmp_v2_read_id  = "string"
    snmp_v2_write_id = "string"
    snmp_v3_id       = "string"
  }
}

output "catalystcenter_site_assign_credential_example" {
  value = catalystcenter_site_assign_credential.example
}