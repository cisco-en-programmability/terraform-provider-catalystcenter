---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_wireless_access_points_factory_reset_request_provision Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on Wireless.
  This data source action is used to factory reset Access Points. It is supported for maximum 100 Access Points per
  request. Factory reset clears all configurations from the Access Points. After factory reset the Access Point may become
  unreachable from the currently associated Wireless Controller and may or may not join back the same controller.
---

# catalystcenter_wireless_access_points_factory_reset_request_provision (Resource)

It performs create operation on Wireless.

- This data source action is used to factory reset Access Points. It is supported for maximum 100 Access Points per
request. Factory reset clears all configurations from the Access Points. After factory reset the Access Point may become
unreachable from the currently associated Wireless Controller and may or may not join back the same controller.


~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in CatalystCenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_wireless_access_points_factory_reset_request_provision" "example" {
  provider = catalystcenter
  parameters = [{

    ap_mac_addresses     = ["string"]
    keep_static_ipconfig = "false"
  }]
}

output "catalystcenter_wireless_access_points_factory_reset_request_provision_example" {
  value = catalystcenter_wireless_access_points_factory_reset_request_provision.example
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

- `ap_mac_addresses` (List of String) List of Access Point's Ethernet MAC addresses, set maximum 100 ethernet MAC addresses per request.
- `keep_static_ipconfig` (String) Set the value of keepStaticIPConfig to false, to clear all configurations from Access Points and set the value of keepStaticIPConfig to true, to clear all configurations from Access Points without clearing static IP configuration.


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `task_id` (String)
- `url` (String)
