
data "catalystcenter_network_device_custom_prompt" "example" {
  provider = catalystcenter
}

output "catalystcenter_network_device_custom_prompt_example" {
  value = data.catalystcenter_network_device_custom_prompt.example.item
}
