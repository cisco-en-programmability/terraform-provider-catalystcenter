---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_assurance_events_count Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Devices.
  API to fetch the count of assurance events that match the filter criteria. Please refer to the 'API Support
  Documentation' section to understand which fields are supported. For detailed information about the usage of the API,
  please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
  specs/blob/main/Assurance/CECatCenter_Org-AssuranceEvents-1.0.0-resolved.yaml
---

# catalystcenter_assurance_events_count (Data Source)

It performs read operation on Devices.

- API to fetch the count of assurance events that match the filter criteria. Please refer to the 'API Support
Documentation' section to understand which fields are supported. For detailed information about the usage of the API,
please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml

## Example Usage

```terraform
data "catalystcenter_assurance_events_count" "example" {
  provider            = catalystcenter
  ap_mac              = "string"
  client_mac          = "string"
  device_family       = "string"
  end_time            = "string"
  message_type        = "string"
  network_device_id   = "string"
  network_device_name = "string"
  severity            = "string"
  site_hierarchy_id   = "string"
  site_id             = "string"
  start_time          = "string"
  xca_lle_rid         = "string"
}

output "catalystcenter_assurance_events_count_example" {
  value = data.catalystcenter_assurance_events_count.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device_family` (String) deviceFamily query parameter. Device family. Please note that multiple families across network device type and client type is not allowed. For example, choosing **Routers** along with **Wireless Client** or **Unified AP** is not supported. Examples:
**deviceFamily=Switches and Hubs** (single deviceFamily requested)
**deviceFamily=Switches and Hubs&deviceFamily=Routers** (multiple deviceFamily requested)
- `xca_lle_rid` (String) X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.

### Optional

- `ap_mac` (String) apMac query parameter. MAC address of the access point. This parameter is applicable for **Unified AP** and **Wireless Client** events.
This field supports wildcard (*****) character-based search. Ex: ***50:0F*** or **50:0F*** or ***50:0F**
Examples:
**apMac=50:0F:80:0F:F7:E0** (single apMac requested)
**apMac=50:0F:80:0F:F7:E0&apMac=18:80:90:AB:7E:A0** (multiple apMac requested)
- `client_mac` (String) clientMac query parameter. MAC address of the client. This parameter is applicable for **Wired Client** and **Wireless Client** events.
This field supports wildcard (*****) character-based search. Ex: ***66:2B*** or **66:2B*** or ***66:2B**
Examples:
**clientMac=66:2B:B8:D2:01:56** (single clientMac requested)
**clientMac=66:2B:B8:D2:01:56&clientMac=DC:A6:32:F5:5A:89** (multiple clientMac requested)
- `end_time` (String) endTime query parameter. End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
If **endTime** is not provided, API will default to current time.
- `message_type` (String) messageType query parameter. Message type for the event.
Examples:
**messageType=Syslog** (single messageType requested)
**messageType=Trap&messageType=Syslog** (multiple messageType requested)
- `network_device_id` (String) networkDeviceId query parameter. The list of Network Device Uuids. (Ex. **6bef213c-19ca-4170-8375-b694e251101c**)
Examples:
**networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c** (single networkDeviceId requested)
**networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c&networkDeviceId=32219612-819e-4b5e-a96b-cf22aca13dd9&networkDeviceId=2541e9a7-b80d-4955-8aa2-79b233318ba0** (multiple networkDeviceId requested)
- `network_device_name` (String) networkDeviceName query parameter. Network device name. This parameter is applicable for network device related families. This field supports wildcard (*****) character-based search. Ex: ***Branch*** or **Branch*** or ***Branch** Examples:
**networkDeviceName=Branch-3-Gateway** (single networkDeviceName requested)
**networkDeviceName=Branch-3-Gateway&networkDeviceName=Branch-3-Switch** (multiple networkDeviceName requested)
- `severity` (String) severity query parameter. Severity of the event between 0 and 6. This is applicable only for events related to network devices (other than AP) and **Wired Client** events.
| Value | Severity    | | ----| ----------| | 0     | Emergency   | | 1     | Alert       | | 2     | Critical    | | 3     | Error       | | 4     | Warning     | | 5     | Notice      | | 6     | Info        |
Examples:
**severity=0** (single severity requested)
**severity=0&severity=1** (multiple severity requested)
- `site_hierarchy_id` (String) siteHierarchyId query parameter. The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. **globalUuid/areaUuid/buildingUuid/floorUuid**)
This field supports wildcard asterisk (*****) character search support. E.g. ***uuid*, *uuid, uuid***
Examples:
**?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid **(single siteHierarchyId requested)
**?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2** (multiple siteHierarchyId requested)
- `site_id` (String) siteId query parameter. The UUID of the site. (Ex. **flooruuid**)
Examples:
**?siteId=id1** (single siteId requested)
**?siteId=id1&siteId=id2&siteId=id3** (multiple siteId requested)
- `start_time` (String) startTime query parameter. Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
If **startTime** is not provided, API will default to current time minus 24 hours.

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `count` (Number)
