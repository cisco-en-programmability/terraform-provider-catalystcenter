---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_file_import Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on File.
  Uploads a new file within a specific nameSpace
---

# catalystcenter_file_import (Resource)

It performs create operation on File.

- Uploads a new file within a specific nameSpace



~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in DNACenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_file_import" "example" {
  provider   = catalystcenter
  file_name  = "string"
  file_path  = "string"
  name_space = "string"
  parameters {

  }
}

output "catalystcenter_file_import_example" {
  value = catalystcenter_file_import.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `parameters` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of String)
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Required:

- `file_name` (String) File name.
- `file_path` (String) File absolute path.
- `name_space` (String) nameSpace path parameter.