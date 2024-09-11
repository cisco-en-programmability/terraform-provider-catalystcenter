
resource "catalystcenter_network_device_custom_prompt" "example" {
  provider = catalystcenter

  parameters {

    password_prompt = "******"
    username_prompt = "string"
  }
}

output "catalystcenter_network_device_custom_prompt_example" {
  value = catalystcenter_network_device_custom_prompt.example
}