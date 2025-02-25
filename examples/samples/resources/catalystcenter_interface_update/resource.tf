
terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.1-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

resource "catalystcenter_interface_update" "example" {
  provider = catalystcenter
  parameters {

    //admin_status = "string"
    description    = "test"
    interface_uuid = "c6820b57-ecde-4b6d-98db-06ba10486809"
    vlan_id        = 2
    //voice_vlan_id = 1
  }
}

output "catalystcenter_interface_update_example" {
  value = catalystcenter_interface_update.example
}
