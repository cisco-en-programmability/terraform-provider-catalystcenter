
data "catalystcenter_sda_port_channels" "example" {
    provider = catalystcenter
    connected_device_type = "string"
    fabric_id = "string"
    limit = 1
    network_device_id = "string"
    offset = 1
    port_channel_name = "string"
}

output "catalystcenter_sda_port_channels_example" {
    value = data.catalystcenter_sda_port_channels.example.items
}
