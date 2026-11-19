---
layout: "ibm"
page_title: "IBM : ibm_mqcloud_capacity_details"
description: |-
  Get information about mqcloud_capacity_details
subcategory: "MQ SaaS"
---

# ibm_mqcloud_capacity_details

Provides a read-only data source to retrieve information about mqcloud_capacity_details. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_mqcloud_capacity_details" "mqcloud_capacity_details" {
	service_instance_guid = "a2b4d4bc-dadb-4637-bcec-9b7d1e723af8"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `service_instance_guid` - (Required, Forces new resource, String) The GUID that uniquely identifies the MQ SaaS service instance.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the mqcloud_capacity_details.
* `capacity_status` - (String) Capacity lifecycle status of a Reserved Capacity service instance, for example `Available`, `Scaling_In_Progress` or `Scaling_Failed`.
* `capacity_units` - (Integer) Current deployed capacity units for a Reserved Capacity service instance.
  * Constraints: The maximum value is `20`. The minimum value is `0`.
* `vpc_entitlement` - (Float) VPC entitlement.
* `vpc_usage` - (Float) VPC usage.

