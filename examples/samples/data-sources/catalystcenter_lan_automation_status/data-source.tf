terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

data "catalystcenter_lan_automation_status" "example" {
  provider = catalystcenter
  //limit = 1
  //offset = 1
}

output "catalystcenter_lan_automation_status_example" {
  value = data.catalystcenter_lan_automation_status.example.items
}

/*data "catalystcenter_lan_automation_status" "example" {
    provider = catalystcenter
    id = "string"
}

output "catalystcenter_lan_automation_status_example" {
    value = data.catalystcenter_lan_automation_status.example.item
}
*/