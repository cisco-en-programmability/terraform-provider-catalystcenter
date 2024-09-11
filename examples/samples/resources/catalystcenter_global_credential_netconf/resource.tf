
terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_global_credential_netconf" "example" {
  provider = catalystcenter
  parameters {
    description     = "Description"
    comments        = "New comments"
    credential_type = "APP"
    netconf_port    = "1"
    #instance_tenant_id= "6168b750e7a2701a37d64526"
    #instance_uuid= "aed1c6d9-e32d-47b5-a7c4-9e8cb15060a1"
    #id= "aed1c6d9-e32d-47b5-a7c4-9e8cb15060a1"
  }
}

output "catalystcenter_global_credential_netconf_example" {
  value = catalystcenter_global_credential_netconf.example
}
