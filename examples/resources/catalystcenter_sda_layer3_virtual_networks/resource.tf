
resource "catalystcenter_sda_layer3_virtual_networks" "example" {
  provider = catalystcenter

  parameters {

    anchored_site_id     = "string"
    fabric_ids           = ["string"]
    id                   = "string"
    virtual_network_name = "string"
  }
}

output "catalystcenter_sda_layer3_virtual_networks_example" {
  value = catalystcenter_sda_layer3_virtual_networks.example
}