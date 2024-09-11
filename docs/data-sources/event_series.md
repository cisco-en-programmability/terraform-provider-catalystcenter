---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_event_series Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Event Management.
  Get the list of Published Notifications
---

# catalystcenter_event_series (Data Source)

It performs read operation on Event Management.

- Get the list of Published Notifications

## Example Usage

```terraform
data "catalystcenter_event_series" "example" {
  provider   = catalystcenter
  category   = "string"
  domain     = "string"
  end_time   = 1609459200
  event_ids  = "string"
  limit      = 1
  namespace  = "string"
  offset     = 1
  order      = "string"
  severity   = "string"
  site_id    = "string"
  sort_by    = "string"
  source     = "string"
  start_time = 1609459200
  sub_domain = "string"
  tags       = "string"
  type       = "string"
}

output "catalystcenter_event_series_example" {
  value = data.catalystcenter_event_series.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `category` (String) category query parameter.
- `domain` (String) domain query parameter.
- `end_time` (Number) endTime query parameter. End Time in milliseconds
- `event_ids` (String) eventIds query parameter. The registered EventId should be provided
- `limit` (Number) limit query parameter. # of records
- `namespace` (String) namespace query parameter.
- `offset` (Number) offset query parameter. Start Offset
- `order` (String) order query parameter. Ascending/Descending order [asc/desc]
- `severity` (String) severity query parameter.
- `site_id` (String) siteId query parameter. Site Id
- `sort_by` (String) sortBy query parameter. Sort By column
- `source` (String) source query parameter.
- `start_time` (Number) startTime query parameter. Start Time in milliseconds
- `sub_domain` (String) subDomain query parameter. Sub Domain
- `tags` (String) tags query parameter.
- `type` (String) type query parameter.

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `category` (String)
- `description` (String)
- `details` (String)
- `domain` (String)
- `event_hierarchy` (String)
- `event_id` (String)
- `instance_id` (String)
- `name` (String)
- `namespace` (String)
- `network` (List of Object) (see [below for nested schema](#nestedobjatt--items--network))
- `severity` (String)
- `source` (String)
- `sub_domain` (String)
- `timestamp` (String)
- `type` (String)
- `version` (String)

<a id="nestedobjatt--items--network"></a>
### Nested Schema for `items.network`

Read-Only:

- `device_id` (String)
- `site_id` (String)