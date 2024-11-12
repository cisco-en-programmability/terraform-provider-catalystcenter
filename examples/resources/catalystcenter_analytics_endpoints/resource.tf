
resource "catalystcenter_analytics_endpoints" "example" {
  provider = catalystcenter

  parameters {

    device_type           = "string"
    ep_id                 = "string"
    hardware_manufacturer = "string"
    hardware_model        = "string"
    mac_address           = "string"
  }
}

output "catalystcenter_analytics_endpoints_example" {
  value = catalystcenter_analytics_endpoints.example
}