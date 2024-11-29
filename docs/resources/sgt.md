---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_sgt Resource - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This resource can manage a SGT.
---

# fmc_sgt (Resource)

This resource can manage a SGT.

## Example Usage

```terraform
resource "fmc_sgt" "example" {
  name        = "SGT1"
  description = "My SGT object"
  tag         = "11"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the SGT object.
- `tag` (String) Security Group Tag.

### Optional

- `description` (String) Description
- `domain` (String) The name of the FMC domain
- `type` (String) Type of the object; this value is always 'SecurityGroupTag'.
  - Default value: `SecurityGroupTag`

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

```shell
terraform import fmc_sgt.example "<id>"
```