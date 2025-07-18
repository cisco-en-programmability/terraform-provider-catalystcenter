---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_site_health_summaries_trend_analytics Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Sites.
  Retrieves the time series information of health and issue data for sites specified by query parameters, or all sites.
  The data will be grouped based on the specified trend time interval. For detailed information about the usage of the
  API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-
  api-specs/blob/main/Assurance/CECatCenter_Org-siteHealthSummaries-2.0.0-resolved.yaml
---

# catalystcenter_site_health_summaries_trend_analytics (Data Source)

It performs read operation on Sites.

- Retrieves the time series information of health and issue data for sites specified by query parameters, or all sites.
The data will be grouped based on the specified trend time interval. For detailed information about the usage of the
API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-
api-specs/blob/main/Assurance/CE_Cat_Center_Org-siteHealthSummaries-2.0.0-resolved.yaml

## Example Usage

```terraform
data "catalystcenter_site_health_summaries_trend_analytics" "example" {
  provider          = catalystcenter
  attribute         = "string"
  end_time          = 1609459200
  id                = "string"
  limit             = 1
  offset            = 1
  site_hierarchy    = "string"
  site_hierarchy_id = "string"
  site_type         = "string"
  start_time        = 1609459200
  task_id           = "string"
  time_sort_order   = "string"
  trend_interval    = "string"
  xca_lle_rid       = "string"
}

output "catalystcenter_site_health_summaries_trend_analytics_example" {
  value = data.catalystcenter_site_health_summaries_trend_analytics.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `xca_lle_rid` (String) X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.

### Optional

- `attribute` (String) attribute query parameter. Supported Analytics Attributes:
[networkDeviceCount, networkDeviceGoodHealthCount,wirelessDeviceCount, wirelessDeviceGoodHealthCount, accessDeviceCount, accessDeviceGoodHealthCount, coreDeviceCount, coreDeviceGoodHealthCount, distributionDeviceCount, distributionDeviceGoodHealthCount, routerDeviceCount, routerDeviceGoodHealthCount, apDeviceCount, apDeviceGoodHealthCount, wlcDeviceCount, wlcDeviceGoodHealthCount, switchDeviceCount, switchDeviceGoodHealthCount, networkDeviceGoodHealthPercentage, accessDeviceGoodHealthPercentage, coreDeviceGoodHealthPercentage, distributionDeviceGoodHealthPercentage, routerDeviceGoodHealthPercentage, apDeviceGoodHealthPercentage, wlcDeviceGoodHealthPercentage, switchDeviceGoodHealthPercentage, wirelessDeviceGoodHealthPercentage, clientCount, clientGoodHealthCount, wiredClientCount, wirelessClientCount, wiredClientGoodHealthCount, wirelessClientGoodHealthCount, clientGoodHealthPercentage, wiredClientGoodHealthPercentage, wirelessClientGoodHealthPercentage, clientDataUsage, p1IssueCount, p2IssueCount, p3IssueCount, p4IssueCount, issueCount]
attribute=networkDeviceCount (single attribute requested)
attribute=networkDeviceCount&attribute=clientCount (multiple attributes requested)
- `end_time` (Number) endTime query parameter. End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
- `id` (String) id query parameter. The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
- `limit` (Number) limit query parameter. Maximum number of records to return
- `offset` (Number) offset query parameter. Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
- `site_hierarchy` (String) siteHierarchy query parameter. The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. **Global/AreaName/BuildingName/FloorName**)
This field supports wildcard asterisk (*****) character search support. E.g. ***/San*, */San, /San***
Examples:
**?siteHierarchy=Global/AreaName/BuildingName/FloorName** (single siteHierarchy requested)
**?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2** (multiple siteHierarchies requested)
- `site_hierarchy_id` (String) siteHierarchyId query parameter. The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. **globalUuid/areaUuid/buildingUuid/floorUuid**)
This field supports wildcard asterisk (*****) character search support. E.g. ***uuid*, *uuid, uuid***
Examples:
**?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid **(single siteHierarchyId requested)
**?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2** (multiple siteHierarchyIds requested)
- `site_type` (String) siteType query parameter. The type of the site. A site can be an area, building, or floor.
Default when not provided will be **[floor,building,area]**
Examples:
**?siteType=area** (single siteType requested)
**?siteType=area&siteType=building&siteType=floor** (multiple siteTypes requested)
- `start_time` (Number) startTime query parameter. Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
- `task_id` (String) taskId query parameter. used to retrieve asynchronously processed & stored data. When this parameter is used, the rest of the request params will be ignored.
- `time_sort_order` (String) timeSortOrder query parameter. The sort order of a time sorted API response.
- `trend_interval` (String) trendInterval query parameter. The time window to aggregate the metrics. Interval can be 5 minutes or 10 minutes or 1 hour or 1 day or 7 days

### Read-Only

- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `attributes` (List of Object) (see [below for nested schema](#nestedobjatt--items--attributes))
- `timestamp` (Number)

<a id="nestedobjatt--items--attributes"></a>
### Nested Schema for `items.attributes`

Read-Only:

- `name` (String)
- `value` (Number)
