terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
}

data "catalystcenter_app_policy_queuing_profile" "example" {
  provider = catalystcenter
  name     = "CVD_QUEUING_PROFILE"
}

output "catalystcenter_app_policy_queuing_profile_example" {
  value = data.catalystcenter_app_policy_queuing_profile.example
}
