---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_event_snmp_config_create Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on Event Management.
  Create SNMP Destination
---

# catalystcenter_event_snmp_config_create (Resource)

It performs create operation on Event Management.

- Create SNMP Destination


~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in CatalystCenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_event_snmp_config_create" "example" {
  provider = catalystcenter
  parameters = [{

    auth_password     = "string"
    community         = "string"
    description       = "string"
    ip_address        = "string"
    name              = "string"
    port              = "string"
    privacy_password  = "string"
    snmp_auth_type    = "string"
    snmp_mode         = "string"
    snmp_privacy_type = "string"
    snmp_version      = "string"
    user_name         = "string"
  }]
}

output "catalystcenter_event_snmp_config_create_example" {
  value = catalystcenter_event_snmp_config_create.example
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

- `auth_password` (String) Auth Password
- `community` (String) Required only if snmpVersion is V2C
- `description` (String) Description
- `ip_address` (String) Ip Address
- `name` (String) Name
- `port` (String) Port
- `privacy_password` (String) Privacy Password
- `snmp_auth_type` (String) Snmp Auth Type
- `snmp_mode` (String) If snmpVersion is V3 it is required and cannot be NONE
- `snmp_privacy_type` (String) Snmp Privacy Type
- `snmp_version` (String) Snmp Version
- `user_name` (String) Required only if snmpVersion is V3


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `api_status` (String)
- `error_message` (List of Object) (see [below for nested schema](#nestedobjatt--item--error_message))
- `status_message` (String)

<a id="nestedobjatt--item--error_message"></a>
### Nested Schema for `item.error_message`

Read-Only:

- `errors` (List of String)
