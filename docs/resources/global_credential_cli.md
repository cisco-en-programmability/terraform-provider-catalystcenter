---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_global_credential_cli Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It manages create, read and update operations on Discovery.
  Updates global CLI credentialsAdds global CLI credentialDeletes global credential for the given ID
---

# catalystcenter_global_credential_cli (Resource)

It manages create, read and update operations on Discovery.

- Updates global CLI credentials

- Adds global CLI credential

- Deletes global credential for the given ID



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `parameters` (Block List, Min: 1, Max: 1) Array of RequestDiscoveryCreateCliCredentials (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `comments` (String)
- `credential_type` (String)
- `description` (String)
- `enable_password` (String)
- `instance_tenant_id` (String)
- `instance_uuid` (String)
- `password` (String, Sensitive)
- `username` (String)

Read-Only:

- `id` (String) The ID of this resource.


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `comments` (String)
- `credential_type` (String)
- `description` (String)
- `id` (String)
- `instance_tenant_id` (String)
- `instance_uuid` (String)
- `netconf_port` (String)
- `password` (String)
- `port` (String)
- `read_community` (String)
- `secure` (String)
- `username` (String)
- `write_community` (String)
