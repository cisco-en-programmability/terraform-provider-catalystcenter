terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

data "catalystcenter_network_device_custom_prompt" "example" {
  provider = catalystcenter
}

output "catalystcenter_network_device_custom_prompt_example" {
  value = data.catalystcenter_network_device_custom_prompt.example.item
}
