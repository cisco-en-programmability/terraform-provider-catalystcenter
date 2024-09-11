---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_configuration_template_version Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Configuration Templates.
  Get all the versions of template by its id
---

# catalystcenter_configuration_template_version (Data Source)

It performs read operation on Configuration Templates.

- Get all the versions of template by its id

## Example Usage

```terraform
data "catalystcenter_configuration_template_version" "example" {
  provider    = catalystcenter
  template_id = "string"
}

output "catalystcenter_configuration_template_version_example" {
  value = data.catalystcenter_configuration_template_version.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `template_id` (String) templateId path parameter. templateId(UUID) to get list of versioned templates

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `composite` (String)
- `name` (String)
- `project_id` (String)
- `project_name` (String)
- `template_id` (String)
- `versions_info` (List of Object) (see [below for nested schema](#nestedobjatt--items--versions_info))

<a id="nestedobjatt--items--versions_info"></a>
### Nested Schema for `items.versions_info`

Read-Only:

- `author` (String)
- `description` (String)
- `id` (String)
- `version` (String)
- `version_comment` (String)
- `version_time` (Number)