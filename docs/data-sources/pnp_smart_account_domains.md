---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_pnp_smart_account_domains Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Device Onboarding (PnP).
  Returns the list of Smart Account domains
---

# catalystcenter_pnp_smart_account_domains (Data Source)

It performs read operation on Device Onboarding (PnP).

- Returns the list of Smart Account domains

## Example Usage

```terraform
data "catalystcenter_pnp_smart_account_domains" "example" {
  provider = catalystcenter
}

output "catalystcenter_pnp_smart_account_domains_example" {
  value = data.catalystcenter_pnp_smart_account_domains.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of String)
