
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

resource "catalystcenter_business_sda_wireless_controller" "example" {
  provider = catalystcenter
  parameters {
    device_name         = "C9K-Branch-SFO.dcloud.cisco.com"
    site_name_hierarchy = "Global/San Francisco"
  }
}

output "catalystcenter_business_sda_wireless_controller_example" {
  value = catalystcenter_business_sda_wireless_controller.example
}

data "catalystcenter_dnacaap_management_execution_status" "example" {
  depends_on   = [catalystcenter_business_sda_wireless_controller.example]
  provider     = catalystcenter
  execution_id = catalystcenter_business_sda_wireless_controller.example.item.0.execution_id
}

output "catalystcenter_dnacaap_management_execution_status_example" {
  value = data.catalystcenter_dnacaap_management_execution_status.example.item
}
