terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

data "catalystcenter_app_policy_default" "example" {
  provider = catalystcenter
}

output "catalystcenter_app_policy_default_example" {
  value = data.catalystcenter_app_policy_default.example.items
}
