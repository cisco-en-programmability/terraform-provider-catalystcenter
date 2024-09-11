
terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.32-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_business_sda_hostonboarding_ssid_ippool" "example" {
  provider = catalystcenter
  parameters {

    # scalable_group_name = "string"
    site_name_hierarchy = "siteNameHierarchy 1"
    ssid_names          = ["lab"]
    vlan_name           = "vlanName 1"
  }
}

output "catalystcenter_business_sda_hostonboarding_ssid_ippool_example" {
  value = catalystcenter_business_sda_hostonboarding_ssid_ippool.example
}
