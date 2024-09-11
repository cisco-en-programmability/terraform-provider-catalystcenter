---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_assurance_issues_top_n_analytics Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on Issues.
  Gets the Top N analytics data related to issues based on given filters and group by field. This data can be used to
  find top sites which has most issues or top device types with most issue etc,. https://github.com/cisco-en-
  programmability/catalyst-center-api-specs/blob/main/Assurance/CECatCenter_Org-IssuesList-1.0.0-resolved.yaml
---

# catalystcenter_assurance_issues_top_n_analytics (Resource)

It performs create operation on Issues.

- Gets the Top N analytics data related to issues based on given filters and group by field. This data can be used to
find top sites which has most issues or top device types with most issue etc,. https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesList-1.0.0-resolved.yaml

## Example Usage

```terraform
resource "catalystcenter_assurance_issues_top_n_analytics" "example" {
  provider        = catalystcenter
  accept_language = "string"
  xca_lle_rid     = "string"
  parameters {

    aggregate_attributes {

      function = "string"
      name     = "string"
    }
    attributes = ["string"]
    end_time   = 1
    filters {

      filters {

        key      = "string"
        operator = "string"
        value    = "string"
      }
      key              = "string"
      logical_operator = "string"
      operator         = "string"
      value            = "string"
    }
    group_by = ["string"]
    page {

      limit  = 1
      offset = 1
      sort_by {

        name  = "string"
        order = "string"
      }
    }
    start_time = 1
    top_n      = 1
  }
}

output "catalystcenter_assurance_issues_top_n_analytics_example" {
  value = catalystcenter_assurance_issues_top_n_analytics.example
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

- `accept_language` (String) Accept-Language header parameter. This header parameter can be used to specify the language in which issue display name need to be returned. Available options are 'en' (English), 'ja' (Japanese), 'ko' (Korean), 'zh' (Chinese). If this parameter is not present the issue display name is returned in English language.
- `xca_lle_rid` (String) X-CALLER-ID header parameter. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc

Optional:

- `aggregate_attributes` (Block List) (see [below for nested schema](#nestedblock--parameters--aggregate_attributes))
- `attributes` (List of String) Attributes
- `end_time` (Number) End Time
- `filters` (Block List) (see [below for nested schema](#nestedblock--parameters--filters))
- `group_by` (List of String) Group By
- `page` (Block List) (see [below for nested schema](#nestedblock--parameters--page))
- `start_time` (Number) Start Time
- `top_n` (Number) Top N

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



<a id="nestedblock--parameters--page"></a>
### Nested Schema for `parameters.page`

Optional:

- `limit` (Number) Limit
- `offset` (Number) Offset
- `sort_by` (Block List) (see [below for nested schema](#nestedblock--parameters--page--sort_by))

<a id="nestedblock--parameters--page--sort_by"></a>
### Nested Schema for `parameters.page.sort_by`

Optional:

- `name` (String) Name
- `order` (String) Order



<a id="nestedatt--parameters--items"></a>
### Nested Schema for `parameters.items`

Read-Only:

- `aggregate_attributes` (List of Object) (see [below for nested schema](#nestedobjatt--parameters--items--aggregate_attributes))
- `attributes` (List of Object) (see [below for nested schema](#nestedobjatt--parameters--items--attributes))
- `id` (String)

<a id="nestedobjatt--parameters--items--aggregate_attributes"></a>
### Nested Schema for `parameters.items.aggregate_attributes`

Read-Only:

- `function` (String)
- `name` (String)
- `value` (Number)


<a id="nestedobjatt--parameters--items--attributes"></a>
### Nested Schema for `parameters.items.attributes`

Read-Only:

- `name` (String)
- `value` (String)