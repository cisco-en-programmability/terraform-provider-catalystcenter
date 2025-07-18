---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_assurance_issues_query_count Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on Issues.
  Returns the total number issues for given set of filters. If there is no start and/or end time, then end time will be
  defaulted to current time and start time will be defaulted to 24-hours ago from end time. For detailed information about
  the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
  programmability/catalyst-center-api-specs/blob/main/Assurance/CECatCenter_Org-IssuesList-1.0.0-resolved.yaml
---

# catalystcenter_assurance_issues_query_count (Resource)

It performs create operation on Issues.

- Returns the total number issues for given set of filters. If there is no start and/or end time, then end time will be
defaulted to current time and start time will be defaulted to 24-hours ago from end time. For detailed information about
the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesList-1.0.0-resolved.yaml


~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in CatalystCenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_assurance_issues_query_count" "example" {
  provider    = catalystcenter
  xca_lle_rid = "string"
  parameters = [{

    end_time = 1
    filters = [{

      filters = [{

        key      = "string"
        operator = "string"
        value    = "string"
      }]
      key              = "string"
      logical_operator = "string"
      operator         = "string"
      value            = "string"
    }]
    start_time = 1
  }]
}

output "catalystcenter_assurance_issues_query_count_example" {
  value = catalystcenter_assurance_issues_query_count.example
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

- `xca_lle_rid` (String) X-CALLER-ID header parameter. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc

Optional:

- `end_time` (Number) End Time
- `filters` (Block List) (see [below for nested schema](#nestedblock--parameters--filters))
- `start_time` (Number) Start Time

Read-Only:

- `items` (List of Object) (see [below for nested schema](#nestedatt--parameters--items))

<a id="nestedblock--parameters--filters"></a>
### Nested Schema for `parameters.filters`

Optional:

- `filters` (Block List) (see [below for nested schema](#nestedblock--parameters--filters--filters))
- `key` (String) Key
- `logical_operator` (String) Logical Operator
- `operator` (String) Operator
- `value` (String) Value

<a id="nestedblock--parameters--filters--filters"></a>
### Nested Schema for `parameters.filters.filters`

Optional:

- `key` (String) Key
- `operator` (String) Operator
- `value` (String) Value



<a id="nestedatt--parameters--items"></a>
### Nested Schema for `parameters.items`

Read-Only:

- `key` (String)
- `operator` (String)
- `value` (String)
