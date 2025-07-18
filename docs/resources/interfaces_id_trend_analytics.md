---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_interfaces_id_trend_analytics Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on Devices.
  The Trend analytcis data for the interface, identified by its instanceUuid, in the specified time range. The data is
  grouped based on the trend time Interval, other input parameters like attributes and aggregate attributes. The default
  time interval range is 3 hours when start and endTime is not provided.
  The field trendIntervalInMinutes is requiered and either the attributes or the aggregateAttributes fields is also
  required.
  For detailed information about the usage of the API, please refer to the Open API specification document
  https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CECatCenter_Org-
  interfaces-2.0.0-resolved.yaml
---

# catalystcenter_interfaces_id_trend_analytics (Resource)

It performs create operation on Devices.

- The Trend analytcis data for the interface, identified by its instanceUuid, in the specified time range. The data is
grouped based on the trend time Interval, other input parameters like attributes and aggregate attributes. The default
time interval range is 3 hours when start and endTime is not provided.
The field trendIntervalInMinutes is requiered and either the attributes or the aggregateAttributes fields is also
required.
For detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
interfaces-2.0.0-resolved.yaml


~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in CatalystCenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_interfaces_id_trend_analytics" "example" {
  provider = catalystcenter
  id       = "string"
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
      value    = "string"
    }]
    start_time                = 1
    timestamp_order           = "string"
    trend_interval_in_minutes = 1
  }]
}

output "catalystcenter_interfaces_id_trend_analytics_example" {
  value = catalystcenter_interfaces_id_trend_analytics.example
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

- `id` (String) id path parameter. The interface instance Uuid

Optional:

- `aggregate_attributes` (Block List) (see [below for nested schema](#nestedblock--parameters--aggregate_attributes))
- `attributes` (List of String) Attributes
- `end_time` (Number) End Time
- `filters` (Block List) (see [below for nested schema](#nestedblock--parameters--filters))
- `start_time` (Number) Start Time
- `timestamp_order` (String) Timestamp Order
- `trend_interval_in_minutes` (Number) Trend Interval In Minutes

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
- `value` (String) Value


<a id="nestedatt--parameters--items"></a>
### Nested Schema for `parameters.items`

Read-Only:

- `aggregate_attributes` (List of Object) (see [below for nested schema](#nestedobjatt--parameters--items--aggregate_attributes))
- `attributes` (List of Object) (see [below for nested schema](#nestedobjatt--parameters--items--attributes))
- `timestamp` (Number)

<a id="nestedobjatt--parameters--items--aggregate_attributes"></a>
### Nested Schema for `parameters.items.aggregate_attributes`

Read-Only:

- `name` (String)


<a id="nestedobjatt--parameters--items--attributes"></a>
### Nested Schema for `parameters.items.attributes`

Read-Only:

- `name` (String)
- `value` (String)
