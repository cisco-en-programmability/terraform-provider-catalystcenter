---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_event Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Event Management.
  Gets the list of registered Events with provided eventIds or tags as mandatory
---

# catalystcenter_event (Data Source)

It performs read operation on Event Management.

- Gets the list of registered Events with provided eventIds or tags as mandatory

## Example Usage

```terraform
data "catalystcenter_event" "example" {
  provider = catalystcenter
  event_id = "string"
  limit    = 1
  offset   = 1
  order    = "string"
  sort_by  = "string"
  tags     = "string"
}

output "catalystcenter_event_example" {
  value = data.catalystcenter_event.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `tags` (String) tags query parameter. The registered Tags should be provided

### Optional

- `event_id` (String) eventId query parameter. The registered EventId should be provided
- `limit` (Number) limit query parameter. The number of Registries to limit in the resultset whose default value 10
- `offset` (Number) offset query parameter. The number of Registries to offset in the resultset whose default value 0
- `order` (String) order query parameter.
- `sort_by` (String) sortBy query parameter. SortBy field name

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
- `event_id` (String)
- `name` (String)
- `name_space` (String)
- `severity` (Number)
- `sub_domain` (String)
- `subscription_types` (List of String)
- `tags` (List of String)
- `type` (String)
- `version` (String)