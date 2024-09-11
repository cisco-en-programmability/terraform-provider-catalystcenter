provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_global_credential_snmpv2_write_community" "example" {
  provider = catalystcenter
  parameters {
    description     = "string"
    comments        = "string"
    credential_type = "string"
    write_community = "string"
  }
}

output "catalystcenter_global_credential_snmpv2_write_community_example" {
  value = catalystcenter_global_credential_snmpv2_read_community.example
}