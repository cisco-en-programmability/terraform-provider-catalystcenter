---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_site_assign_credential Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on Sites.
  Assign Device Credential to a site.
---

# catalystcenter_site_assign_credential (Resource)

It performs create operation on Sites.

- Assign Device Credential to a site.


~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in CatalystCenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_site_assign_credential" "example" {
  provider = catalystcenter
  site_id  = "string"
  parameters = [{

    cli_id           = "string"
    http_read        = "string"
    http_write       = "string"
    snmp_v2_read_id  = "string"
    snmp_v2_write_id = "string"
    snmp_v3_id       = "string"
  }]
}

output "catalystcenter_site_assign_credential_example" {
  value = catalystcenter_site_assign_credential.example
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

- `site_id` (String) siteId path parameter. site id to assign credential.

Optional:

- `cli_id` (String) Cli Id
- `http_read` (String) Http Read
- `http_write` (String) Http Write
- `snmp_v2_read_id` (String) Snmp V2 Read Id
- `snmp_v2_write_id` (String) Snmp V2 Write Id
- `snmp_v3_id` (String) Snmp V3 Id


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `execution_id` (String)
- `execution_status_url` (String)
- `message` (String)
