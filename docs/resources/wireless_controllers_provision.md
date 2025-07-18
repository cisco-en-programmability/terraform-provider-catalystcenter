---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_wireless_controllers_provision Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on Wireless.
  This data source action is used to provision wireless controller
---

# catalystcenter_wireless_controllers_provision (Resource)

It performs create operation on Wireless.

- This data source action is used to provision wireless controller


~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in CatalystCenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `parameters` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Required:

- `device_id` (String) deviceId path parameter. Network Device ID. This value can be obtained by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}

Optional:

- `ap_authorization_list_name` (String) AP Authorization List name. 'Obtain the AP Authorization List names by using the API call GET: /intent/api/v1/wirelessSettings/apAuthorizationLists. During re-provision, obtain the AP Authorization List configured for the given provisioned network device Id using the API call GET: /intent/api/v1/wireless/apAuthorizationLists/{networkDeviceId}'
- `authorize_mesh_and_non_mesh_access_points` (String) True if AP Authorization List should  authorize against All Mesh/Non-Mesh APs, else false if AP Authorization List should only authorize against Mesh APs (Applicable only when Mesh is enabled on sites)
- `feature_templates_overriden_attributes` (Block List) (see [below for nested schema](#nestedblock--parameters--feature_templates_overriden_attributes))
- `interfaces` (Block List) (see [below for nested schema](#nestedblock--parameters--interfaces))
- `rolling_ap_upgrade` (Block List) (see [below for nested schema](#nestedblock--parameters--rolling_ap_upgrade))
- `skip_ap_provision` (String) True if Skip AP Provision is enabled, else False

<a id="nestedblock--parameters--feature_templates_overriden_attributes"></a>
### Nested Schema for `parameters.feature_templates_overriden_attributes`

Optional:

- `edit_feature_templates` (Block List) (see [below for nested schema](#nestedblock--parameters--feature_templates_overriden_attributes--edit_feature_templates))

<a id="nestedblock--parameters--feature_templates_overriden_attributes--edit_feature_templates"></a>
### Nested Schema for `parameters.feature_templates_overriden_attributes.edit_feature_templates`

Optional:

- `additional_identifiers` (Block List) (see [below for nested schema](#nestedblock--parameters--feature_templates_overriden_attributes--edit_feature_templates--additional_identifiers))
- `attributes` (String) This dynamic map should contain attribute name and overridden value of respective Feature Template whose **featureTemplateId**. List of attributes applicable to given **featureTemplateId** can be retrieved from its GET API call /dna/intent/api/v1/featureTemplates/wireless/<featureTemplateName>/featureTemplateId
- `excluded_attributes` (List of String) List of attributes which will NOT be provisioned.
- `feature_template_id` (String) Feature Template ID

<a id="nestedblock--parameters--feature_templates_overriden_attributes--edit_feature_templates--additional_identifiers"></a>
### Nested Schema for `parameters.feature_templates_overriden_attributes.edit_feature_templates.additional_identifiers`

Optional:

- `site_uuid` (String) Site UUID. This must be provided if **featureTemplateId** belongs to **Flex Configuration** feature template.
- `wlan_profile_name` (String) WLAN Profile Name. This must be passed if **featureTemplateId** belongs to **Advanced SSID Configuration** Feature Template.




<a id="nestedblock--parameters--interfaces"></a>
### Nested Schema for `parameters.interfaces`

Optional:

- `interface_gateway` (String) Interface Gateway
- `interface_ipaddress` (String) Interface IP Address
- `interface_name` (String) Interface Name
- `interface_netmask_in_cid_r` (Number) Interface Netmask In CIDR, range is 1-30
- `lag_or_port_number` (Number) Lag Or Port Number
- `vlan_id` (Number) VLAN ID range is 1 - 4094


<a id="nestedblock--parameters--rolling_ap_upgrade"></a>
### Nested Schema for `parameters.rolling_ap_upgrade`

Optional:

- `ap_reboot_percentage` (Number) AP Reboot Percentage. Permissible values - 5, 15, 25
- `enable_rolling_ap_upgrade` (String) True if Rolling AP Upgrade is enabled, else False



<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `task_id` (String)
- `url` (String)
