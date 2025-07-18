---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_global_credential_delete Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs delete operation on Discovery.
  Deletes global credential for the given ID
---

# catalystcenter_global_credential_delete (Resource)

It performs delete operation on Discovery.

- Deletes global credential for the given ID


~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in CatalystCenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_global_credential_delete" "example" {
  provider             = catalystcenter
  global_credential_id = "string"
  parameters = [{

  }]
}

output "catalystcenter_global_credential_delete_example" {
  value = catalystcenter_global_credential_delete.example
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

- `global_credential_id` (String) globalCredentialId path parameter. ID of global-credential


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `task_id` (String)
- `url` (String)
