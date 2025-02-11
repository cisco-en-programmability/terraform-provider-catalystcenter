terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

data "catalystcenter_app_policy" "example" {
  provider = catalystcenter
  # policy_scope = "draft_WiredTest"
}

output "catalystcenter_app_policy_example" {
  value = data.catalystcenter_app_policy.example.items
}
