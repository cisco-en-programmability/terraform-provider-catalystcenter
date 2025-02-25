
terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.1-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_global_credential_snmpv3" "example" {
  provider = catalystcenter
  parameters {
    auth_password   = "12345678"
    auth_type       = "SHA"
    comments        = "Test"
    credential_type = "APP"
    description     = "Description 3"
    #id= "string"
    #instanceTenantId= "string"
    #instanceUuid= "string"
    privacy_password = "privacy_password_test"
    privacy_type     = "AES128"
    snmp_mode        = "AUTHPRIV"
    username         = "Global_credential_test3"
  }
}

output "catalystcenter_global_credential_snmpv3_example" {
  value = catalystcenter_global_credential_snmpv3.example
}
