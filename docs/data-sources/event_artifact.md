---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_event_artifact Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Event Management.
  Gets the list of artifacts based on provided offset and limit
---

# catalystcenter_event_artifact (Data Source)

It performs read operation on Event Management.

- Gets the list of artifacts based on provided offset and limit

## Example Usage

```terraform
data "catalystcenter_event_artifact" "example" {
  provider  = catalystcenter
  event_ids = "string"
  limit     = 1
  offset    = 1
  order     = "string"
  search    = "string"
  sort_by   = "string"
  tags      = "string"
}

output "catalystcenter_event_artifact_example" {
  value = data.catalystcenter_event_artifact.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `event_ids` (String) eventIds query parameter. List of eventIds
- `limit` (Number) limit query parameter. # of records to return in result set
- `offset` (Number) offset query parameter. Record start offset
- `order` (String) order query parameter. sorting order (asc/desc)
- `search` (String) search query parameter. findd matches in name, description, eventId, type, category
- `sort_by` (String) sortBy query parameter. Sort by field
- `tags` (String) tags query parameter. Tags defined

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `artifact_id` (String)
- `cisco_dna_event_link` (String)
- `configs` (List of Object) (see [below for nested schema](#nestedobjatt--items--configs))
- `deprecated` (String)
- `deprecation_message` (String)
- `description` (String)
- `domain` (String)
- `event_payload` (List of Object) (see [below for nested schema](#nestedobjatt--items--event_payload))
- `event_templates` (List of String)
- `is_private` (String)
- `is_template_enabled` (String)
- `is_tenant_aware` (String)
- `name` (String)
- `namespace` (String)
- `note` (String)
- `sub_domain` (String)
- `supported_connector_types` (List of String)
- `tags` (List of String)
- `tenant_id` (String)
- `version` (String)

<a id="nestedobjatt--items--configs"></a>
### Nested Schema for `items.configs`

Read-Only:

- `is_acknowledgeable` (String)
- `is_alert` (String)


<a id="nestedobjatt--items--event_payload"></a>
### Nested Schema for `items.event_payload`

Read-Only:

- `additional_details` (String)
- `category` (String)
- `details` (List of Object) (see [below for nested schema](#nestedobjatt--items--event_payload--details))
- `event_id` (String)
- `severity` (String)
- `source` (String)
- `type` (String)
- `version` (String)

<a id="nestedobjatt--items--event_payload--details"></a>
### Nested Schema for `items.event_payload.details`

Read-Only:

- `device_ip` (String)
- `message` (String)
