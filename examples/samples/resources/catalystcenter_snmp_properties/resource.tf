
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
  debug = "true"
}

resource "catalystcenter_snmp_properties" "example" {
  provider = catalystcenter
  parameters {

    #id                   = "string"
    #instance_tenant_id   = "string"
    #instance_uuid        = "string"
    #int_value            = 1
  }
}

output "catalystcenter_snmp_properties_example" {
  value = catalystcenter_snmp_properties.example
}
