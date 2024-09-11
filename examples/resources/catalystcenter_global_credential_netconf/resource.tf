provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_global_credential_netconf" "example" {
  provider = catalystcenter
  parameters {
    description        = "string"
    comments           = "string"
    credential_type    = "string"
    netconf_port       = "string"
    instance_tenant_id = "string"
    instance_uuid      = "string"
    id                 = "string"
  }
}

output "catalystcenter_global_credential_netconf_example" {
  value = catalystcenter_global_credential_netconf.example
}