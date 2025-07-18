---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_sites_aaa_settings Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It manages read and update operations on Network Settings.
  Set AAA settings for a site; null values indicate that the settings will be inherited from the parent site; empty
  objects ({}) indicate that the settings is unset.
---

# catalystcenter_sites_aaa_settings (Resource)

It manages read and update operations on Network Settings.

- Set AAA settings for a site; **null** values indicate that the settings will be inherited from the parent site; empty
objects (**{}**) indicate that the settings is unset.

## Example Usage

```terraform
resource "catalystcenter_sites_aaa_settings" "example" {
  provider = catalystcenter
 
  parameters {

    aaa_client {

      pan                 = "string"
      primary_server_ip   = "string"
      protocol            = "string"
      secondary_server_ip = "string"
      server_type         = "string"
      shared_secret       = "string"
    }
    aaa_network {

      pan                 = "string"
      primary_server_ip   = "string"
      protocol            = "string"
      secondary_server_ip = "string"
      server_type         = "string"
      shared_secret       = "string"
    }
    id = "string"
  }
}

output "catalystcenter_sites_aaa_settings_example" {
  value = catalystcenter_sites_aaa_settings.example
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

- `id` (String) id path parameter. Site Id

Optional:

- `aaa_client` (Block List) (see [below for nested schema](#nestedblock--parameters--aaa_client))
- `aaa_network` (Block List) (see [below for nested schema](#nestedblock--parameters--aaa_network))

<a id="nestedblock--parameters--aaa_client"></a>
### Nested Schema for `parameters.aaa_client`

Optional:

- `pan` (String) Administration Node.  Required for ISE.
- `primary_server_ip` (String) The server to use as a primary.
- `protocol` (String) Protocol
- `secondary_server_ip` (String) The server to use as a secondary.
- `server_type` (String) Server Type
- `shared_secret` (String) Shared Secret


<a id="nestedblock--parameters--aaa_network"></a>
### Nested Schema for `parameters.aaa_network`

Optional:

- `pan` (String) Administration Node. Required for ISE.
- `primary_server_ip` (String) The server to use as a primary.
- `protocol` (String) Protocol
- `secondary_server_ip` (String) The server to use as a secondary.
- `server_type` (String) Server Type
- `shared_secret` (String) Shared Secret



<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `aaa_client` (List of Object) (see [below for nested schema](#nestedobjatt--item--aaa_client))
- `aaa_network` (List of Object) (see [below for nested schema](#nestedobjatt--item--aaa_network))

<a id="nestedobjatt--item--aaa_client"></a>
### Nested Schema for `item.aaa_client`

Read-Only:

- `inherited_site_id` (String)
- `inherited_site_name` (String)
- `pan` (String)
- `primary_server_ip` (String)
- `protocol` (String)
- `secondary_server_ip` (String)
- `server_type` (String)
- `shared_secret` (String)


<a id="nestedobjatt--item--aaa_network"></a>
### Nested Schema for `item.aaa_network`

Read-Only:

- `inherited_site_id` (String)
- `inherited_site_name` (String)
- `pan` (String)
- `primary_server_ip` (String)
- `protocol` (String)
- `secondary_server_ip` (String)
- `server_type` (String)
- `shared_secret` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import catalystcenter_sites_aaa_settings.example "id:=string"
```
