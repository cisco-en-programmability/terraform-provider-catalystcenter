
terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}


resource "catalystcenter_wireless_profile" "example" {
  provider = catalystcenter

  parameters {
    profile_details {
      name  = "Test22"
      sites = ["Global/CR"]



      ssid_details {
        enable_fabric = "true"
        flex_connect {
          enable_flex_connect = "false"
          local_to_vlan       = 0
        }
        interface_name = "management"
        name           = "BTest22"
        type           = "string"
      }
      ssid_details {
        enable_fabric = "true"
        flex_connect {
          enable_flex_connect = "false"
          local_to_vlan       = 0
        }
        interface_name = "management"
        name           = "ATest222"
        type           = "eduroam"
      }
      ssid_details {
        enable_fabric = "true"
        flex_connect {
          enable_flex_connect = "false"
          local_to_vlan       = 0
        }
        interface_name = "management"
        name           = "new"
        type           = "string"
      }
    }
  }
}


# output "catalystcenter_wireless_profile_example" {
#   value = catalystcenter_wireless_profile.example
# }
