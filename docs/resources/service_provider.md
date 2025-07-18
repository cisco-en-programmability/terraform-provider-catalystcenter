---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_service_provider Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It manages create, read, update and delete operations on Network Settings.
  API to create Service Provider Profile(QOS).API to update Service Provider Profile (QoS).API to delete Service Provider Profile (QoS).
---

# catalystcenter_service_provider (Resource)

It manages create, read, update and delete operations on Network Settings.

- API to create Service Provider Profile(QOS).

- API to update Service Provider Profile (QoS).

- API to delete Service Provider Profile (QoS).

## Example Usage

```terraform
resource "catalystcenter_service_provider" "example" {
  provider = catalystcenter
 
  parameters {

    settings {

      qos {

        model            = "string"
        old_profile_name = "string"
        profile_name     = "string"
        wan_provider     = "string"
      }
    }
    sp_profile_name = "string"
  }
}

output "catalystcenter_service_provider_example" {
  value = catalystcenter_service_provider.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `parameters` (Block List) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Required:

- `sp_profile_name` (String) spProfileName path parameter. sp profile name

Optional:

- `settings` (Block List) (see [below for nested schema](#nestedblock--parameters--settings))

<a id="nestedblock--parameters--settings"></a>
### Nested Schema for `parameters.settings`

Optional:

- `qos` (Block List) (see [below for nested schema](#nestedblock--parameters--settings--qos))

<a id="nestedblock--parameters--settings--qos"></a>
### Nested Schema for `parameters.settings.qos`

Optional:

- `model` (String) Model
- `old_profile_name` (String) Old Profile Name
- `profile_name` (String) Profile Name
- `wan_provider` (String) Wan Provider




<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `group_uuid` (String)
- `inherited_group_name` (String)
- `inherited_group_uuid` (String)
- `instance_type` (String)
- `instance_uuid` (String)
- `key` (String)
- `namespace` (String)
- `type` (String)
- `value` (List of Object) (see [below for nested schema](#nestedobjatt--item--value))
- `version` (Number)

<a id="nestedobjatt--item--value"></a>
### Nested Schema for `item.value`

Read-Only:

- `sla_profile_name` (String)
- `sp_profile_name` (String)
- `wan_provider` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import catalystcenter_service_provider.example "id:=string"
```
