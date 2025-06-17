
resource "catalystcenter_wireless_controllers_provision" "example" {
  provider  = catalystcenter
  device_id = "string"
  parameters = [{

    ap_authorization_list_name                = "string"
    authorize_mesh_and_non_mesh_access_points = "false"
    feature_templates_overriden_attributes = [{

      edit_feature_templates = [{

        additional_identifiers = [{

          site_uuid         = "string"
          wlan_profile_name = "string"
        }]
        attributes          = "string"
        excluded_attributes = ["string"]
        feature_template_id = "string"
      }]
    }]
    interfaces = [{

      interface_gateway          = "string"
      interface_ipaddress        = "string"
      interface_name             = "string"
      interface_netmask_in_cid_r = 1
      lag_or_port_number         = 1
      vlan_id                    = 1
    }]
    rolling_ap_upgrade = [{

      ap_reboot_percentage      = 1
      enable_rolling_ap_upgrade = "false"
    }]
    skip_ap_provision = "false"
  }]
}

output "catalystcenter_wireless_controllers_provision_example" {
  value = catalystcenter_wireless_controllers_provision.example
}