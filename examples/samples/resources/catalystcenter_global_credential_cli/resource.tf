
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

resource "catalystcenter_global_credential_cli" "example" {
  provider = catalystcenter
  parameters {
    comments        = "test4_comment"
    credential_type = "GLOBAL"
    description     = "string"
    enable_password = "string"
    #id= "string"
    #instance_tenant_id= "string"
    #instance_uuid= "string"
    password = "string"
    username = "TEST4"
  }
}

output "catalystcenter_global_credential_cli_example" {
  value     = catalystcenter_global_credential_cli.example
  sensitive = true
}
