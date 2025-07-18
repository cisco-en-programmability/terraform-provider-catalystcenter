---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_wireless_controllers_ssid_details Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Wireless.
  Retrieves all details of SSIDs associated with the specific Wireless Controller.
---

# catalystcenter_wireless_controllers_ssid_details (Data Source)

It performs read operation on Wireless.

- Retrieves all details of SSIDs associated with the specific Wireless Controller.

## Example Usage

```terraform
data "catalystcenter_wireless_controllers_ssid_details" "example" {
  provider          = catalystcenter
  admin_status      = "false"
  limit             = 1
  managed           = "false"
  network_device_id = "string"
  offset            = 1
  ssid_name         = "string"
}

output "catalystcenter_wireless_controllers_ssid_details_example" {
  value = data.catalystcenter_wireless_controllers_ssid_details.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_device_id` (String) networkDeviceId path parameter. Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.

### Optional

- `admin_status` (Boolean) adminStatus query parameter. Utilize this query parameter to obtain the administrative status. A 'true' value signifies that the admin status of the SSID is enabled, while a 'false' value indicates that the admin status of the SSID is disabled.
- `limit` (Number) limit query parameter. The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
- `managed` (Boolean) managed query parameter. If value is 'true' means SSIDs are configured through design.If the value is 'false' means out of band configuration from the Wireless Controller.
- `offset` (Number) offset query parameter. The first record to show for this page; the first record is numbered 1.
- `ssid_name` (String) ssidName query parameter. Employ this query parameter to obtain the details of the SSID corresponding to the provided SSID name.

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `admin_status` (String)
- `l2_security` (String)
- `l3_security` (String)
- `managed` (String)
- `radio_policy` (String)
- `ssid_name` (String)
- `wlan_id` (Number)
- `wlan_profile_name` (String)
