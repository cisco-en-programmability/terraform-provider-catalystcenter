terraform {
  required_providers {
    catalystcenter = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

resource "catalystcenter_sda_virtual_network" "example" {
  provider = catalystcenter
  parameters {
    site_name_hierarchy  = "Global/New Jersey/MurrayHill/test/TestFloor"
    virtual_network_name = "ANSIBLE80"
  }
}

# resource "catalystcenter_sda_virtual_network" "example2" {
#   provider = catalystcenter
#   depends_on = [
#     catalystcenter_sda_virtual_network.example
#   ]
#   parameters {
#     payload {
#       site_name_hierarchy  = "Global/New Jersey/MurrayHill/test/TestFloor"
#       virtual_network_name = " TEST_VNs"
#     }
#   }
# }

output "catalystcenter_sda_virtual_network_example" {
  value = catalystcenter_sda_virtual_network.example
}
