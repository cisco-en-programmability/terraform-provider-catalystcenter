---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_pnp_virtual_accounts Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Device Onboarding (PnP).
  Returns list of virtual accounts associated with the specified smart account
---

# catalystcenter_pnp_virtual_accounts (Data Source)

It performs read operation on Device Onboarding (PnP).

- Returns list of virtual accounts associated with the specified smart account

## Example Usage

```terraform
data "catalystcenter_pnp_virtual_accounts" "example" {
  provider = catalystcenter
  domain   = "string"
}

output "catalystcenter_pnp_virtual_accounts_example" {
  value = data.catalystcenter_pnp_virtual_accounts.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `domain` (String) domain path parameter. Smart Account Domain

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of String)
