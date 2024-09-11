
terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}
resource "catalystcenter_interface_operation_create" "example" {
  provider = catalystcenter
  parameters {

    interface_uuid = "c6820b57-ecde-4b6d-98db-06ba10486809"
    operation      = "string"
    payload        = "string"
  }
}

output "catalystcenter_interface_operation_create_example" {
  value = catalystcenter_interface_operation_create.example
}
