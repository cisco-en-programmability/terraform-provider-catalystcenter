---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_license_device_license_details Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Licenses.
  Get detailed license information of a device.
---

# catalystcenter_license_device_license_details (Data Source)

It performs read operation on Licenses.

- Get detailed license information of a device.

## Example Usage

```terraform
data "catalystcenter_license_device_license_details" "example" {
  provider    = catalystcenter
  device_uuid = "string"
}

output "catalystcenter_license_device_license_details_example" {
  value = data.catalystcenter_license_device_license_details.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device_uuid` (String) device_uuid path parameter. Id of device

### Read-Only

- `id` (String) The ID of this resource.
- `item` (String)
