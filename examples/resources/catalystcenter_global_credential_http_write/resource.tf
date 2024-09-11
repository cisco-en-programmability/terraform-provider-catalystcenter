provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_global_credential_http_write" "example" {
  provider = catalystcenter
  parameters {
    secure             = "false"
    username           = "string"
    password           = "string"
    port               = 1
    description        = "string"
    comments           = "string"
    credential_type    = "string"
    instance_tenant_id = "string"
    instance_uuid      = "string"
    id                 = "string"
  }
}

output "catalystcenter_global_credential_http_write_example" {
  value     = catalystcenter_global_credential_http_write.example
  sensitive = true
}