---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_assurance_tasks Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Task.
  returns all existing tasks in a paginated list
  default sorting of list is startTime, asc
  valid field to sort by are [startTime,endTime,updateTime,status] For detailed information about the
  usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
  programmability/catalyst-center-api-specs/blob/main/Assurance/CECatCenter_Org-AssuranceTasks-1.0.0-resolved.yaml
---

# catalystcenter_assurance_tasks (Data Source)

It performs read operation on Task.

- returns all existing tasks in a paginated list
default sorting of list is **startTime**, **asc**
valid field to sort by are [**startTime**,**endTime**,**updateTime**,**status**] For detailed information about the
usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceTasks-1.0.0-resolved.yaml

## Example Usage

```terraform
data "catalystcenter_assurance_tasks" "example" {
  provider    = catalystcenter
  limit       = 1
  offset      = 1
  order       = "string"
  sort_by     = "string"
  status      = "string"
  xca_lle_rid = "string"
}

output "catalystcenter_assurance_tasks_example" {
  value = data.catalystcenter_assurance_tasks.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `xca_lle_rid` (String) X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.

### Optional

- `limit` (Number) limit query parameter. Maximum number of records to return
- `offset` (Number) offset query parameter. Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
- `order` (String) order query parameter. The sort order of the field ascending or descending.
- `sort_by` (String) sortBy query parameter. A field within the response to sort by.
- `status` (String) status query parameter. used to get a subset of tasks by their status

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `data` (String)
- `end_time` (Number)
- `error_code` (String)
- `failure_reason` (String)
- `id` (String)
- `progress` (String)
- `request_type` (String)
- `result_url` (String)
- `start_time` (Number)
- `status` (String)
- `update_time` (Number)
