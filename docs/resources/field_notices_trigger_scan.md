---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_field_notices_trigger_scan Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on Compliance.
  Triggers a field notices scan for the supported network devices. The supported devices are switches, routers and
  wireless controllers. If a device is not supported, the FieldNoticeNetworkDevice scanStatus will be Failed with
  appropriate comments. The consent to connect agreement must have been accepted in the UI for this to succeed. Please
  refer to the user guide at
  for more details on consent to connect.
---

# catalystcenter_field_notices_trigger_scan (Resource)

It performs create operation on Compliance.

- Triggers a field notices scan for the supported network devices. The supported devices are switches, routers and
wireless controllers. If a device is not supported, the FieldNoticeNetworkDevice scanStatus will be Failed with
appropriate comments. The consent to connect agreement must have been accepted in the UI for this to succeed. Please
refer to the user guide at
 for more details on consent to connect.


~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in CatalystCenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_field_notices_trigger_scan" "example" {
  provider            = catalystcenter
  failed_devices_only = "false"
}

output "catalystcenter_field_notices_trigger_scan_example" {
  value = catalystcenter_field_notices_trigger_scan.example
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

- `failed_devices_only` (Boolean) failedDevicesOnly query parameter. Used to specify if the scan should run only for the network devices that failed during the previous scan. If not specified, this parameter defaults to false.


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `task_id` (String)
- `url` (String)
