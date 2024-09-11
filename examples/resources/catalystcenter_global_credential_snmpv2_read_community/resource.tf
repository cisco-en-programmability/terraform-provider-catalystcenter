provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_global_credential_snmpv2_read_community" "example" {
  provider = catalystcenter
  parameters {
    description     = "string"
    comments        = "string"
    credential_type = "string"
    read_community  = "string"
  }
}

output "catalystcenter_global_credential_snmpv2_read_community_example" {
  value = catalystcenter_global_credential_snmpv2_read_community.example
}