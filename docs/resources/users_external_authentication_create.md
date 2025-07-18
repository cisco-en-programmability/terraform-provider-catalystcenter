---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_users_external_authentication_create Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on User and Roles.
  Enable or disable external authentication on Cisco Catalyst Center System.
  Please find the Administrator Guide for your particular release from the list linked below and follow the steps required
  to enable external authentication before trying to do so from this API.
  https://www.cisco.com/c/en/us/support/cloud-systems-management/dna-center/products-maintenance-guides-list.html
---

# catalystcenter_users_external_authentication_create (Resource)

It performs create operation on User and Roles.

- Enable or disable external authentication on Cisco Catalyst Center System.
Please find the Administrator Guide for your particular release from the list linked below and follow the steps required
to enable external authentication before trying to do so from this API.
https://www.cisco.com/c/en/us/support/cloud-systems-management/dna-center/products-maintenance-guides-list.html


~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in CatalystCenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_users_external_authentication_create" "example" {
  provider = catalystcenter
  parameters = [{

    enable = "false"
  }]
}

output "catalystcenter_users_external_authentication_create_example" {
  value = catalystcenter_users_external_authentication_create.example
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

- `enable` (String) Enable/disable External Authentication.


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `message` (String)
