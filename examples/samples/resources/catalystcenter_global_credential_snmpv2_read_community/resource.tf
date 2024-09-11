
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

resource "catalystcenter_global_credential_snmpv2_read_community" "example" {
  provider = catalystcenter
  parameters {
    description     = "Description 4"
    comments        = "New Comments"
    credential_type = "APP"
    read_community  = "Test4"
  }
}

output "catalystcenter_global_credential_snmpv2_read_community_example" {
  value = catalystcenter_global_credential_snmpv2_read_community.example
}
