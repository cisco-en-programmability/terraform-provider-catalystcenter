---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_projects_project_id Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Configuration Templates.
  Get a template project by the project's ID.
---

# catalystcenter_projects_project_id (Data Source)

It performs read operation on Configuration Templates.

- Get a template project by the project's ID.

## Example Usage

```terraform
data "catalystcenter_projects_project_id" "example" {
  provider   = catalystcenter
  project_id = "string"
}

output "catalystcenter_projects_project_id_example" {
  value = data.catalystcenter_projects_project_id.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `project_id` (String) projectId path parameter. The id of the project to get, retrieveable from **GET /dna/intent/api/v1/projects**

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `description` (String)
- `last_update_time` (Number)
- `name` (String)
- `project_id` (String)
