---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_assign_to_site_apply Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on Site Design.
  Assign unprovisioned network devices to a site. Along with that it can also be used to assign unprovisioned network
  devices to a different site. If device controllability is enabled, it will be triggered once device assigned to site
  successfully. Device Controllability can be enabled/disabled using
  '/dna/intent/api/v1/networkDevices/deviceControllability/settings'.
---

# catalystcenter_assign_to_site_apply (Resource)

It performs create operation on Site Design.

- Assign unprovisioned network devices to a site. Along with that it can also be used to assign unprovisioned network
devices to a different site. If device controllability is enabled, it will be triggered once device assigned to site
successfully. Device Controllability can be enabled/disabled using
'/dna/intent/api/v1/networkDevices/deviceControllability/settings'.



~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in DNACenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_assign_to_site_apply" "example" {
  provider = catalystcenter
  parameters {

    device_ids = ["string"]
    site_id    = "string"
  }
}

output "catalystcenter_assign_to_site_apply_example" {
  value = catalystcenter_assign_to_site_apply.example
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

- `device_ids` (List of String) Unassigned network devices.
- `site_id` (String) This must be building Id or floor Id. Access points, Sensors are assigned to floor. Remaining network devices are assigned to building. Site Id can be retrieved using '/intent/api/v1/sites'.


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `task_id` (String)
- `url` (String)