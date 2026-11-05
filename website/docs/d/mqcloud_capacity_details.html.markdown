---
layout: "ibm"
page_title: "IBM : ibm_mqcloud_capacity_details"
description: |-
  Get information about mqcloud_capacity_details
subcategory: "MQ SaaS"
---

# ibm_mqcloud_capacity_details

Provides a read-only data source to retrieve the deployed capacity of an MQ SaaS Reserved Capacity service instance. The values reflect the capacity that is actually deployed, which can differ from the last requested `capacity_units` parameter on the `ibm_resource_instance` resource after a failed scaling operation. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_mqcloud_capacity_details" "mqcloud_capacity_details" {
	service_instance_guid = "a2b4d4bc-dadb-4637-bcec-9b7d1e723af8"
}
```

To scale a Reserved Capacity instance, change `capacity_units` in the `parameters_json` of the `ibm_resource_instance` resource. Scale-up is supported only, and the maximum is `20` capacity units. Because scaling can take up to 60 minutes, override the default update timeout.

```hcl
resource "ibm_resource_instance" "rc_env" {
  name              = "my-rc-environment"
  service           = "mqcloud"
  plan              = "reserved-capacity"
  location          = "us-east"
  resource_group_id = var.resource_group_id

  parameters_json = jsonencode({
    capacity_units = var.capacity_units
  })

  timeouts {
    create = "120m"
    update = "90m"
    delete = "20m"
  }
}

data "ibm_mqcloud_capacity_details" "current" {
  service_instance_guid = ibm_resource_instance.rc_env.guid
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `service_instance_guid` - (Required, Forces new resource, String) The GUID that uniquely identifies the MQ SaaS Reserved Capacity service instance.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the mqcloud_capacity_details.
* `capacity_units` - (Integer) The number of capacity units currently deployed. One capacity unit is 5 VPC.
* `vpc_entitlement` - (Float) The VPC entitlement of the service instance, expressed as a raw VPC count (`capacity_units` x 5).
* `vpc_usage` - (Float) The number of VPCs currently consumed by the queue managers deployed in the service instance.
