---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_analytics_dictionaries Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on AI Endpoint Analytics.
  Fetches the list of attribute dictionaries.
---

# catalystcenter_analytics_dictionaries (Data Source)

It performs read operation on AI Endpoint Analytics.

- Fetches the list of attribute dictionaries.

## Example Usage

```terraform
data "catalystcenter_analytics_dictionaries" "example" {
  provider           = catalystcenter
  include_attributes = "false"
}

output "catalystcenter_analytics_dictionaries_example" {
  value = data.catalystcenter_analytics_dictionaries.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `include_attributes` (Boolean) includeAttributes query parameter. Flag to indicate whether attribute list for each dictionary should be included in response.

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `attributes` (List of Object) (see [below for nested schema](#nestedobjatt--items--attributes))
- `description` (String)
- `name` (String)

<a id="nestedobjatt--items--attributes"></a>
### Nested Schema for `items.attributes`

Read-Only:

- `description` (String)
- `name` (String)
