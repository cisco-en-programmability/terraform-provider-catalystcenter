---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_sda_virtual_network_v2 Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It manages create, read, update and delete operations on SDA.
  Add virtual network with scalable groups at global levelDelete virtual network with scalable groupsUpdate virtual network with scalable groups
---

# catalystcenter_sda_virtual_network_v2 (Resource)

It manages create, read, update and delete operations on SDA.

- Add virtual network with scalable groups at global level

- Delete virtual network with scalable groups

- Update virtual network with scalable groups

## Example Usage

```terraform
resource "catalystcenter_sda_virtual_network_v2" "example" {
  provider = catalystcenter

  parameters {

    is_guest_virtual_network = "false"
    scalable_group_names     = ["string"]
    v_manage_vpn_id          = "string"
    virtual_network_name     = "string"
  }
}

output "catalystcenter_sda_virtual_network_v2_example" {
  value = catalystcenter_sda_virtual_network_v2.example
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

Optional:

- `is_guest_virtual_network` (String) Guest Virtual Network enablement flag, default value is False.
- `scalable_group_names` (List of String) Scalable Group to be associated to virtual network
- `v_manage_vpn_id` (String) vManage vpn id for SD-WAN
- `virtual_network_name` (String) Virtual Network Name to be assigned at global level


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `description` (String)
- `execution_id` (String)
- `is_guest_virtual_network` (String)
- `scalable_group_names` (List of String)
- `status` (String)
- `v_manage_vpn_id` (String)
- `virtual_network_context_id` (String)
- `virtual_network_name` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import catalystcenter_sda_virtual_network_v2.example "id:=string"
```
