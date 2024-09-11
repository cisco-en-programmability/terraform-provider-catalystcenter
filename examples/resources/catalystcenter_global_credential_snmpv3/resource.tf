provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_global_credential_snmpv3" "example" {
  provider = catalystcenter
  parameters {
    auth_password    = "string"
    auth_type        = "string"
    comments         = "string"
    credential_type  = "string"
    description      = "string"
    id               = "string"
    instanceTenantId = "string"
    instanceUuid     = "string"
    privacy_password = "string"
    privacy_type     = "string"
    snmp_mode        = "string"
    username         = "string"
  }
}

output "catalystcenter_global_credential_snmpv3_example" {
  value = catalystcenter_global_credential_snmpv3.example
}