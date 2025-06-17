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
}

data "catalystcenter_app_policy_queuing_profile_count" "example" {
  provider = catalystcenter
}

output "catalystcenter_app_policy_queuing_profile_count_example" {
  value = data.catalystcenter_app_policy_queuing_profile_count.example.item
}
