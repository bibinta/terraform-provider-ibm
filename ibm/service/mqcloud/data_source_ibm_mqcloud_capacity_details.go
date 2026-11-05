// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package mqcloud

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/mqcloud-go-sdk/mqcloudv1"
)

func DataSourceIbmMqcloudCapacityDetails() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmMqcloudCapacityDetailsRead,

		Schema: map[string]*schema.Schema{
			"service_instance_guid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The GUID that uniquely identifies the MQ SaaS Reserved Capacity service instance.",
			},
			"capacity_units": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of capacity units currently deployed. One capacity unit is 5 VPC.",
			},
			"vpc_entitlement": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "The VPC entitlement of the service instance, expressed as a raw VPC count (capacity_units x 5).",
			},
			"vpc_usage": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "The number of VPCs currently consumed by the queue managers deployed in the service instance.",
			},
		},
	}
}

func dataSourceIbmMqcloudCapacityDetailsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mqcloudClient, err := meta.(conns.ClientSession).MqcloudV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_mqcloud_capacity_details", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getUsageDetailsOptions := &mqcloudv1.GetUsageDetailsOptions{}

	getUsageDetailsOptions.SetServiceInstanceGuid(d.Get("service_instance_guid").(string))

	usage, _, err := mqcloudClient.GetUsageDetailsWithContext(context, getUsageDetailsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetUsageDetailsWithContext failed: %s", err.Error()), "(Data) ibm_mqcloud_capacity_details", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmMqcloudCapacityDetailsID(d))

	if !core.IsNil(usage.CapacityUnits) {
		if err = d.Set("capacity_units", flex.IntValue(usage.CapacityUnits)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting capacity_units: %s", err), "(Data) ibm_mqcloud_capacity_details", "read", "set-capacity_units").GetDiag()
		}
	}

	if !core.IsNil(usage.VpcEntitlement) {
		if err = d.Set("vpc_entitlement", float64(*usage.VpcEntitlement)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vpc_entitlement: %s", err), "(Data) ibm_mqcloud_capacity_details", "read", "set-vpc_entitlement").GetDiag()
		}
	}

	if !core.IsNil(usage.VpcUsage) {
		if err = d.Set("vpc_usage", float64(*usage.VpcUsage)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vpc_usage: %s", err), "(Data) ibm_mqcloud_capacity_details", "read", "set-vpc_usage").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmMqcloudCapacityDetailsID returns a reasonable ID for the data source.
func dataSourceIbmMqcloudCapacityDetailsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
