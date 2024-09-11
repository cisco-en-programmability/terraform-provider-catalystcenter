---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_network_profiles_for_sites_site_assignments_bulk_create Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on Site Design.
  Assign a network profile for sites to a list of sites. Also assigns the profile to child sites.
---

# catalystcenter_network_profiles_for_sites_site_assignments_bulk_create (Resource)

It performs create operation on Site Design.

- Assign a network profile for sites to a list of sites. Also assigns the profile to child sites.



~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in DNACenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_network_profiles_for_sites_site_assignments_bulk_create" "example" {
  provider   = catalystcenter
  profile_id = "string"
  parameters {

    items = []
  }
}

output "catalystcenter_network_profiles_for_sites_site_assignments_bulk_create_example" {
  value = catalystcenter_network_profiles_for_sites_site_assignments_bulk_create.example
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

- `profile_id` (String) profileId path parameter. The 'id' of the network profile, retrievable from 'GET /intent/api/v1/networkProfilesForSites'

Optional:

- `items` (List of String) Items


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `task_id` (String)
- `url` (String)