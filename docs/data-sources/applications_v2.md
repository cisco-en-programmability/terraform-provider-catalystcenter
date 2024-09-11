---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_applications_v2 Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Application Policy.
  Get application/s by offset/limit or by name
---

# catalystcenter_applications_v2 (Data Source)

It performs read operation on Application Policy.

- Get application/s by offset/limit or by name

## Example Usage

```terraform
data "catalystcenter_applications_v2" "example" {
  provider   = catalystcenter
  attributes = "string"
  limit      = 1
  name       = "string"
  offset     = 1
}

output "catalystcenter_applications_v2_example" {
  value = data.catalystcenter_applications_v2.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `attributes` (String) attributes query parameter. Attributes to retrieve, valid value application
- `limit` (Number) limit query parameter. The limit which is the maximum number of items to include in a single page of results, max value 500
- `offset` (Number) offset query parameter. The starting point or index from where the paginated results should begin.

### Optional

- `name` (String) name query parameter. The application name

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `display_name` (String)
- `id` (String)
- `identity_source` (List of Object) (see [below for nested schema](#nestedobjatt--items--identity_source))
- `indicative_network_identity` (List of Object) (see [below for nested schema](#nestedobjatt--items--indicative_network_identity))
- `instance_id` (Number)
- `instance_version` (Number)
- `name` (String)
- `namespace` (String)
- `network_applications` (List of Object) (see [below for nested schema](#nestedobjatt--items--network_applications))
- `network_identity` (List of Object) (see [below for nested schema](#nestedobjatt--items--network_identity))
- `parent_scalable_group` (List of Object) (see [below for nested schema](#nestedobjatt--items--parent_scalable_group))
- `qualifier` (String)
- `scalable_group_external_handle` (String)
- `scalable_group_type` (String)
- `type` (String)

<a id="nestedobjatt--items--identity_source"></a>
### Nested Schema for `items.identity_source`

Read-Only:

- `id` (String)
- `type` (String)


<a id="nestedobjatt--items--indicative_network_identity"></a>
### Nested Schema for `items.indicative_network_identity`

Read-Only:

- `display_name` (String)
- `id` (String)
- `lower_port` (Number)
- `ports` (String)
- `protocol` (String)
- `upper_port` (Number)


<a id="nestedobjatt--items--network_applications"></a>
### Nested Schema for `items.network_applications`

Read-Only:

- `app_protocol` (String)
- `application_sub_type` (String)
- `application_type` (String)
- `category_id` (String)
- `display_name` (String)
- `dscp` (String)
- `engine_id` (String)
- `help_string` (String)
- `id` (String)
- `long_description` (String)
- `name` (String)
- `popularity` (Number)
- `rank` (Number)
- `selector_id` (String)
- `server_name` (String)
- `traffic_class` (String)
- `url` (String)


<a id="nestedobjatt--items--network_identity"></a>
### Nested Schema for `items.network_identity`

Read-Only:

- `display_name` (String)
- `id` (String)
- `ipv4_subnet` (List of String)
- `ipv6_subnet` (List of String)
- `lower_port` (Number)
- `ports` (String)
- `protocol` (String)
- `upper_port` (Number)


<a id="nestedobjatt--items--parent_scalable_group"></a>
### Nested Schema for `items.parent_scalable_group`

Read-Only:

- `id` (String)
- `id_ref` (String)