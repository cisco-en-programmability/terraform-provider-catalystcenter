terraform {
  required_providers {
    catalystcenter = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}
resource "catalystcenter_lan_automation_delete" "example" {
  provider = catalystcenter
  parameters {

    id = "string"
  }
}

output "catalystcenter_lan_automation_delete_example" {
  value = catalystcenter_lan_automation_delete.example
}
