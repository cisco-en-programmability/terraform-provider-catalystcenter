---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_wireless_provision_access_point Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on Wireless.
  Access Point Provision and ReProvision
---

# catalystcenter_wireless_provision_access_point (Resource)

It performs create operation on Wireless.

- Access Point Provision and ReProvision



~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in DNACenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_wireless_provision_access_point" "example" {
  provider = catalystcenter
  parameters {
    payload {
      custom_ap_group_name   = "string"
      custom_flex_group_name = ["string"]
      device_name            = "string"
      rf_profile             = "string"
      site_name_hierarchy    = "string"
      type                   = "string"
    }
  }
}

output "catalystcenter_wireless_provision_access_point_example" {
  value = catalystcenter_wireless_provision_access_point.example
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

Optional:

- `payload` (Block List) Array of RequestWirelessAPProvision2 (see [below for nested schema](#nestedblock--parameters--payload))

<a id="nestedblock--parameters--payload"></a>
### Nested Schema for `parameters.payload`

Optional:

- `custom_ap_group_name` (String) Custom AP group name
- `custom_flex_group_name` (List of String) ["Custom flex group name"]
- `device_name` (String) Device name
- `rf_profile` (String) Radio frequency profile name
- `site_name_hierarchy` (String) Site name hierarchy(ex: Global/...)
- `type` (String) ApWirelessConfiguration



<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `execution_id` (String)
- `execution_status_url` (String)
- `message` (String)