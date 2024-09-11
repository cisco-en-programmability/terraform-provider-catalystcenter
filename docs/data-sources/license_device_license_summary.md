---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_license_device_license_summary Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Licenses.
  Show license summary of device(s).
---

# catalystcenter_license_device_license_summary (Data Source)

It performs read operation on Licenses.

- Show license summary of device(s).

## Example Usage

```terraform
data "catalystcenter_license_device_license_summary" "example" {
  provider             = catalystcenter
  device_type          = "string"
  device_uuid          = "string"
  dna_level            = "string"
  limit                = 1
  order                = "string"
  page_number          = 1.0
  registration_status  = "string"
  smart_account_id     = "string"
  sort_by              = "string"
  virtual_account_name = "string"
}

output "catalystcenter_license_device_license_summary_example" {
  value = data.catalystcenter_license_device_license_summary.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `limit` (Number) limit query parameter.
- `order` (String) order query parameter. Sorting order
- `page_number` (Number) page_number query parameter. Page number of response

### Optional

- `device_type` (String) device_type query parameter. Type of device. The valid values are Routers, Switches and Hubs, Wireless Controller
- `device_uuid` (String) device_uuid query parameter. Id of device
- `dna_level` (String) dna_level query parameter. Device Cisco DNA license level. The valid values are Advantage, Essentials
- `registration_status` (String) registration_status query parameter. Smart license registration status of device. The valid values are Unknown, NA, Unregistered, Registered, Registration_expired, Reservation_in_progress, Registered_slr, Registered_plr, Registered_satellite
- `smart_account_id` (String) smart_account_id query parameter. Id of smart account
- `sort_by` (String) sort_by query parameter. Sort result by field
- `virtual_account_name` (String) virtual_account_name query parameter. Name of virtual account

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `authorization_status` (String)
- `customer_tag1` (String)
- `customer_tag2` (String)
- `customer_tag3` (String)
- `customer_tag4` (String)
- `device_name` (String)
- `device_type` (String)
- `device_uuid` (String)
- `dna_level` (String)
- `evaluation_license_expiry` (String)
- `hsec_status` (String)
- `ip_address` (String)
- `is_license_expired` (String)
- `is_performance_allowed` (String)
- `is_wireless` (String)
- `is_wireless_capable` (String)
- `last_successful_rum_usage_upload_time` (String)
- `last_updated_time` (String)
- `license_mode` (String)
- `mac_address` (String)
- `model` (String)
- `network_license` (String)
- `performance_license` (String)
- `registration_status` (String)
- `reservation_status` (String)
- `site` (String)
- `sle_auth_code` (String)
- `sle_state` (String)
- `smart_account_name` (String)
- `software_version` (String)
- `throughput_level` (String)
- `total_access_point_count` (Number)
- `virtual_account_name` (String)
- `wireless_capable_dna_license` (String)
- `wireless_capable_network_license` (String)