---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_compliance Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on Compliance.
  Run compliance check for device(s).
---

# catalystcenter_compliance (Resource)

It performs create operation on Compliance.

- Run compliance check for device(s).


~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in CatalystCenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_compliance" "example" {
  provider = catalystcenter
  parameters = [{

    categories   = ["string"]
    device_uuids = ["string"]
    trigger_full = "false"
  }]
}

output "catalystcenter_compliance_example" {
  value = catalystcenter_compliance.example
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

- `categories` (List of String) Category can have any value among 'INTENT'(mapped to compliance types: NETWORK_SETTINGS,NETWORK_PROFILE,WORKFLOW,FABRIC,APPLICATION_VISIBILITY), 'RUNNING_CONFIG' , 'IMAGE' , 'PSIRT' , 'EOX'
- `device_uuids` (List of String) UUID of the device.
- `trigger_full` (String) if it is true then compliance will be triggered for all categories. If it is false then compliance will be triggered for categories mentioned in categories section .


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `task_id` (String)
- `url` (String)
