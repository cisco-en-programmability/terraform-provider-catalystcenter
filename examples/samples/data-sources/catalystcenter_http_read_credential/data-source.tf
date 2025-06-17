terraform {
  required_providers {
    catalystcenter = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
}

data "dnaccatalystcenter_http_read_credential" "example" {
  provider = catalystcenter
  payload {

    comments           = "string"
    credential_type    = "string"
    description        = "string"
    id                 = "string"
    instance_tenant_id = "string"
    instance_uuid      = "string"
    password           = "******"
    port               = 1
    secure             = "false"
    username           = "string"
  }
}
