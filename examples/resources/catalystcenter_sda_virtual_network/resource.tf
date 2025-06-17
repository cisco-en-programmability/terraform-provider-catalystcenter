
resource "catalystcenter_sda_virtual_network" "example" {
  provider = catalystcenter
 
  parameters {

    site_name_hierarchy  = "string"
    virtual_network_name = "string"
  }
}

output "catalystcenter_sda_virtual_network_example" {
  value = catalystcenter_sda_virtual_network.example
}