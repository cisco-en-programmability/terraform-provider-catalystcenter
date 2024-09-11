provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_global_credential_cli" "example" {
  provider = catalystcenter
  parameters {
    comments           = "string"
    credential_type    = "string"
    description        = "string"
    enable_password    = "string"
    id                 = "string"
    instance_tenant_id = "string"
    instance_uuid      = "string"
    password           = "string"
    username           = "string"
  }
}

output "catalystcenter_global_credential_cli_example" {
  value     = catalystcenter_global_credential_cli.example
  sensitive = true
}