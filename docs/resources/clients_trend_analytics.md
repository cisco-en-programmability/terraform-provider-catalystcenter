---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_clients_trend_analytics Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on Clients.
  Retrieves the trend analytics of client data for the specified time range. The data will be grouped based on the given
  trend time interval. This data source action facilitates obtaining consolidated insights into the performance and status
  of the clients over the specified start and end time. For detailed information about the usage of the API, please refer
  to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
  specs/blob/main/Assurance/CECatCenter_Org-clients1-1.0.0-resolved.yaml
---

# catalystcenter_clients_trend_analytics (Resource)

It performs create operation on Clients.

- Retrieves the trend analytics of client data for the specified time range. The data will be grouped based on the given
trend time interval. This data source action facilitates obtaining consolidated insights into the performance and status
of the clients over the specified start and end time. For detailed information about the usage of the API, please refer
to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml


~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in CatalystCenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_clients_trend_analytics" "example" {
  provider    = catalystcenter
  xca_lle_rid = "string"
  parameters = [{

    aggregate_attributes = [{

      function = "string"
      name     = "string"
    }]
    attributes = ["string"]
    end_time   = 1
    filters = [{

      key      = "string"
      operator = "string"
      value    = 1
    }]
    group_by = ["string"]
    page = [{

      cursor          = "string"
      limit           = 1
      time_sort_order = "string"
    }]
    start_time     = 1
    trend_interval = "string"
  }]
}

output "catalystcenter_clients_trend_analytics_example" {
  value = catalystcenter_clients_trend_analytics.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `parameters` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Required:

- `xca_lle_rid` (String) X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.

Optional:

- `aggregate_attributes` (Block List) (see [below for nested schema](#nestedblock--parameters--aggregate_attributes))
- `attributes` (List of String) Attributes
- `end_time` (Number) End Time
- `filters` (Block List) (see [below for nested schema](#nestedblock--parameters--filters))
- `group_by` (List of String) Group By
- `page` (Block List) (see [below for nested schema](#nestedblock--parameters--page))
- `start_time` (Number) Start Time
- `trend_interval` (String) Trend Interval

Read-Only:

- `items` (List of Object) (see [below for nested schema](#nestedatt--parameters--items))

<a id="nestedblock--parameters--aggregate_attributes"></a>
### Nested Schema for `parameters.aggregate_attributes`

Optional:

- `function` (String) Function
- `name` (String) Name


<a id="nestedblock--parameters--filters"></a>
### Nested Schema for `parameters.filters`

Optional:

- `key` (String) Key
- `operator` (String) Operator
- `value` (Number) Value


<a id="nestedblock--parameters--page"></a>
### Nested Schema for `parameters.page`

Optional:

- `cursor` (String) Cursor
- `limit` (Number) Limit
- `time_sort_order` (String) Time Sort Order


<a id="nestedatt--parameters--items"></a>
### Nested Schema for `parameters.items`

Read-Only:

- `groups` (List of Object) (see [below for nested schema](#nestedobjatt--parameters--items--groups))
- `timestamp` (Number)

<a id="nestedobjatt--parameters--items--groups"></a>
### Nested Schema for `parameters.items.groups`

Read-Only:

- `aggregate_attributes` (List of Object) (see [below for nested schema](#nestedobjatt--parameters--items--groups--aggregate_attributes))
- `attributes` (List of Object) (see [below for nested schema](#nestedobjatt--parameters--items--groups--attributes))
- `id` (String)

<a id="nestedobjatt--parameters--items--groups--aggregate_attributes"></a>
### Nested Schema for `parameters.items.groups.id`

Read-Only:

- `function` (String)
- `name` (String)
- `value` (Number)


<a id="nestedobjatt--parameters--items--groups--attributes"></a>
### Nested Schema for `parameters.items.groups.id`

Read-Only:

- `name` (String)
- `value` (Number)
