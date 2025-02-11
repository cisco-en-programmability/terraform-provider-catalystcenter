
data "catalystcenter_sda_layer3_virtual_networks" "example" {
    provider = catalystcenter
    anchored_site_id = "string"
    fabric_id = "string"
    limit = 1
    offset = 1
    virtual_network_name = "string"
}

output "catalystcenter_sda_layer3_virtual_networks_example" {
    value = data.catalystcenter_sda_layer3_virtual_networks.example.items
}
