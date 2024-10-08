---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_system_health_count Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Health and Performance.
  This data source gives the count of the latest system events
---

# catalystcenter_system_health_count (Data Source)

It performs read operation on Health and Performance.

- This data source gives the count of the latest system events

## Example Usage

```terraform
data "catalystcenter_system_health_count" "example" {
  provider  = catalystcenter
  domain    = "string"
  subdomain = "string"
}

output "catalystcenter_system_health_count_example" {
  value = data.catalystcenter_system_health_count.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `domain` (String) domain query parameter. Fetch system events with this domain. Possible values of domain are listed here : /dna/platform/app/consumer-portal/developer-toolkit/events
- `subdomain` (String) subdomain query parameter. Fetch system events with this subdomain. Possible values of subdomain are listed here : /dna/platform/app/consumer-portal/developer-toolkit/events

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `count` (Number)
