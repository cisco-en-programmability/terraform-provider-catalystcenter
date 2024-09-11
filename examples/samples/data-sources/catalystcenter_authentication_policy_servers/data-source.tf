terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}


data "catalystcenter_authentication_policy_servers" "example" {
  provider = catalystcenter
  # is_ise_enabled = "false"
  # role           = "string"
  # state          = "string"
}

output "catalystcenter_authentication_policy_servers_example" {
  value = data.catalystcenter_authentication_policy_servers.example.items
}
