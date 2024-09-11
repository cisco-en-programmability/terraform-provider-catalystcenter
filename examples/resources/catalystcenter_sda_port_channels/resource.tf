
resource "catalystcenter_sda_port_channels" "example" {
  provider = catalystcenter

  parameters {

    connected_device_type = "string"
    description           = "string"
    fabric_id             = "string"
    id                    = "string"
    interface_names       = ["string"]
    network_device_id     = "string"
    port_channel_name     = "string"
    protocol              = "string"
  }
}

output "catalystcenter_sda_port_channels_example" {
  value = catalystcenter_sda_port_channels.example
}