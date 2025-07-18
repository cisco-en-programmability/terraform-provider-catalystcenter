---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_pnp_device_unclaim Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on Device Onboarding (PnP).
  Un-Claims one of more devices with specified workflow (Deprecated).
---

# catalystcenter_pnp_device_unclaim (Resource)

It performs create operation on Device Onboarding (PnP).

- Un-Claims one of more devices with specified workflow (Deprecated).


~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in CatalystCenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_pnp_device_unclaim" "example" {
  provider = catalystcenter
  parameters = [{

    device_id_list = ["string"]
  }]
}

output "catalystcenter_pnp_device_unclaim_example" {
  value = catalystcenter_pnp_device_unclaim.example
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

- `device_id_list` (List of String)


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `json_array_response` (List of String)
- `json_response` (String)
- `message` (String)
- `status_code` (Number)
