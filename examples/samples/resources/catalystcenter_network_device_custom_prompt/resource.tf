terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

resource "catalystcenter_network_device_custom_prompt" "example" {
  provider = catalystcenter
  parameters {

    password_prompt = "******"
    username_prompt = "string2"
  }
}

output "catalystcenter_network_device_custom_prompt_example" {
  value     = catalystcenter_network_device_custom_prompt.example
  sensitive = true
}
