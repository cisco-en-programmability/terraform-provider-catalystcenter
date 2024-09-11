---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_sda_fabrics_vlan_to_ssids_fabric_id_count Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Fabric Wireless.
  Returns the count of VLANs mapped to SSIDs in a Fabric Site. The 'fabricId' represents the Fabric ID of a particular
  Fabric Site.
---

# catalystcenter_sda_fabrics_vlan_to_ssids_fabric_id_count (Data Source)

It performs read operation on Fabric Wireless.

- Returns the count of VLANs mapped to SSIDs in a Fabric Site. The 'fabricId' represents the Fabric ID of a particular
Fabric Site.

## Example Usage

```terraform
data "catalystcenter_sda_fabrics_vlan_to_ssids_fabric_id_count" "example" {
  provider  = catalystcenter
  fabric_id = "string"
}

output "catalystcenter_sda_fabrics_vlan_to_ssids_fabric_id_count_example" {
  value = data.catalystcenter_sda_fabrics_vlan_to_ssids_fabric_id_count.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `fabric_id` (String) fabricId path parameter. The 'fabricId' represents the Fabric ID of a particular Fabric Site

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `count` (Number)