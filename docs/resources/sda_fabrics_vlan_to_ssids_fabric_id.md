---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_sda_fabrics_vlan_to_ssids_fabric_id Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It manages read and update operations on Fabric Wireless.
  Add, update, or remove SSID mappings to a VLAN. If the payload doesn't contain a 'vlanName' which has SSIDs mapping
  done earlier then all the mapped SSIDs of the 'vlanName' is cleared. The request must include all SSIDs currently mapped
  to a VLAN, as determined by the response from the GET operation for the same fabricId used in the request. If an
  already-mapped SSID is not included in the payload, its mapping will be removed by this API. Conversely, if a new SSID
  is provided, it will be added to the Mapping. Ensure that any new SSID added is a Fabric SSID. This resource can also be
  used to add a VLAN and associate the relevant SSIDs with it. The 'vlanName' must be 'Fabric Wireless Enabled' and should
  be part of the Fabric Site representing 'Fabric ID' specified in the API request.
---

# catalystcenter_sda_fabrics_vlan_to_ssids_fabric_id (Resource)

It manages read and update operations on Fabric Wireless.

- Add, update, or remove SSID mappings to a VLAN. If the payload doesn't contain a 'vlanName' which has SSIDs mapping
done earlier then all the mapped SSIDs of the 'vlanName' is cleared. The request must include all SSIDs currently mapped
to a VLAN, as determined by the response from the GET operation for the same fabricId used in the request. If an
already-mapped SSID is not included in the payload, its mapping will be removed by this API. Conversely, if a new SSID
is provided, it will be added to the Mapping. Ensure that any new SSID added is a Fabric SSID. This resource can also be
used to add a VLAN and associate the relevant SSIDs with it. The 'vlanName' must be 'Fabric Wireless Enabled' and should
be part of the Fabric Site representing 'Fabric ID' specified in the API request.

## Example Usage

```terraform
resource "catalystcenter_sda_fabrics_vlan_to_ssids_fabric_id" "example" {
  provider = catalystcenter
 
  parameters {

    fabric_id = "string"
    ssid_details {

      name               = "string"
      security_group_tag = "string"
    }
    vlan_name = "string"
  }
}

output "catalystcenter_sda_fabrics_vlan_to_ssids_fabric_id_example" {
  value = catalystcenter_sda_fabrics_vlan_to_ssids_fabric_id.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `parameters` (Block List) Array of RequestFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLAN (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `payload` (Block List) Array of RequestApplicationPolicyCreateApplication (see [below for nested schema](#nestedblock--parameters--payload))

<a id="nestedblock--parameters--payload"></a>
### Nested Schema for `parameters.payload`

Required:

- `fabric_id` (String) fabricId path parameter. The 'fabricId' represents the Fabric ID of a particular Fabric Site

Optional:

- `ssid_details` (Block List) (see [below for nested schema](#nestedblock--parameters--payload--ssid_details))
- `vlan_name` (String) Vlan Name

<a id="nestedblock--parameters--payload--ssid_details"></a>
### Nested Schema for `parameters.payload.ssid_details`

Optional:

- `name` (String) Name of the SSID
- `security_group_tag` (String) Represents the name of the Security Group. Example: Auditors, BYOD, Developers, etc.




<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `ssid_details` (List of Object) (see [below for nested schema](#nestedobjatt--item--ssid_details))
- `vlan_name` (String)

<a id="nestedobjatt--item--ssid_details"></a>
### Nested Schema for `item.ssid_details`

Read-Only:

- `name` (String)
- `security_group_tag` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import catalystcenter_sda_fabrics_vlan_to_ssids_fabric_id.example "fabric_id:=string"
```
