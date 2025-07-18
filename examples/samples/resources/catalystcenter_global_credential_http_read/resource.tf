
terraform {
  required_providers {
    catalystcenter = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_global_credential_http_read" "example" {
  provider = catalystcenter
  parameters {
    secure      = true
    username    = "userTF3"
    password    = "123"
    port        = 23
    description = "This is a test for tf"
    #comments= null
    credential_type = "APP"
    #instance_tenant_id= "6168b750e7a2701a37d64526"
    #instance_uuid= "aed1c6d9-e32d-47b5-a7c4-9e8cb15060a1"
    #id= "aed1c6d9-e32d-47b5-a7c4-9e8cb15060a1"
  }
}

output "catalystcenter_global_credential_http_read_example" {
  value     = catalystcenter_global_credential_http_read.example
  sensitive = true
}
