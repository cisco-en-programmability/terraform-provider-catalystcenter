---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_fabric_site_health_summaries_count Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on SDA.
  Get a count of Fabric sites. Use available query parameters to get the count of a subset of fabric sites.
  This data source provides the latest health data until the given endTime. If data is not ready for the provided
  endTime, the request will fail with error code 400 Bad Request, and the error message will indicate the recommended
  endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we
  are not a real time system. When endTime is not provided, the API returns the latest data.
  For detailed information about the usage of the API, please refer to the Open API specification document
  https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CECatCenter_Org-
  fabricSiteHealthSummaries-1.0.1-resolved.yaml
---

# catalystcenter_fabric_site_health_summaries_count (Data Source)

It performs read operation on SDA.

- Get a count of Fabric sites. Use available query parameters to get the count of a subset of fabric sites.
This data source provides the latest health data until the given **endTime**. If data is not ready for the provided
endTime, the request will fail with error code **400 Bad Request**, and the error message will indicate the recommended
endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we
are not a real time system. When **endTime** is not provided, the API returns the latest data.
For detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
fabricSiteHealthSummaries-1.0.1-resolved.yaml

## Example Usage

```terraform
data "catalystcenter_fabric_site_health_summaries_count" "example" {
  provider          = catalystcenter
  end_time          = 1609459200
  id                = "string"
  site_hierarchy    = "string"
  site_hierarchy_id = "string"
  start_time        = 1609459200
  xca_lle_rid       = "string"
}

output "catalystcenter_fabric_site_health_summaries_count_example" {
  value = data.catalystcenter_fabric_site_health_summaries_count.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `xca_lle_rid` (String) X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.

### Optional

- `end_time` (Number) endTime query parameter. End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
- `id` (String) id query parameter. The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
- `site_hierarchy` (String) siteHierarchy query parameter. The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. **Global/AreaName/BuildingName/FloorName**)          This field supports wildcard asterisk (*****) character search support. E.g. ***/San*, */San, /San***          Examples:          **?siteHierarchy=Global/AreaName/BuildingName/FloorName** (single siteHierarchy requested)          **?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2** (multiple siteHierarchies requested)
- `site_hierarchy_id` (String) siteHierarchyId query parameter. The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. **globalUuid/areaUuid/buildingUuid/floorUuid**)          This field supports wildcard asterisk (*****) character search support. E.g. ***uuid*, *uuid, uuid***          Examples:          **?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid **(single siteHierarchyId requested)          **?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2** (multiple siteHierarchyIds requested)
- `start_time` (Number) startTime query parameter. Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `count` (Number)
