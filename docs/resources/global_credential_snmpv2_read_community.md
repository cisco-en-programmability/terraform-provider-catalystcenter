---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_global_credential_snmpv2_read_community Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It manages create, read and update operations on Discovery.
  Updates global SNMP read communityAdds global SNMP read communityDeletes global credential for the given ID
---

# catalystcenter_global_credential_snmpv2_read_community (Resource)

It manages create, read and update operations on Discovery.

- Updates global SNMP read community

- Adds global SNMP read community

- Deletes global credential for the given ID

## Example Usage

```terraform
provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_global_credential_snmpv2_read_community" "example" {
  provider = catalystcenter
  parameters {
    description     = "string"
    comments        = "string"
    credential_type = "string"
    read_community  = "string"
  }
}

output "catalystcenter_global_credential_snmpv2_read_community_example" {
  value = catalystcenter_global_credential_snmpv2_read_community.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `parameters` (Block List, Min: 1, Max: 1) Array of RequestDiscoveryCreateSNMPReadCommunity (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `comments` (String) Comments to identify the credential
- `credential_type` (String) Credential type to identify the application that uses the credential
- `description` (String) Name/Description of the credential
- `instance_uuid` (String)
- `read_community` (String) SNMP read community

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
- `port` (Number)
- `read_community` (String)
- `secure` (String)
- `username` (String)
- `write_community` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import catalystcenter_global_credential_snmpv2_read_community.example "id:=string"
```