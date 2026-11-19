// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.117.3-3e91ab11-20261028-182643
 */

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
			"service_instance_guid": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The GUID that uniquely identifies the MQ SaaS service instance.",
			},
			"vpc_entitlement": &schema.Schema{
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "VPC entitlement.",
			},
			"vpc_usage": &schema.Schema{
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "VPC usage.",
			},
			"capacity_units": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Current deployed capacity units for a Reserved Capacity service instance.",
			},
			"capacity_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Capacity lifecycle status of a Reserved Capacity service instance, for example `Available`, `Scaling_In_Progress` or `Scaling_Failed`.",
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

	if !core.IsNil(usage.VpcEntitlement) {
		if err = d.Set("vpc_entitlement", flex.Float64Value(usage.VpcEntitlement)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting vpc_entitlement: %s", err), "(Data) ibm_mqcloud_capacity_details", "read", "set-vpc_entitlement").GetDiag()
		}
	}

	if !core.IsNil(usage.VpcUsage) {
		if err = d.Set("vpc_usage", flex.Float64Value(usage.VpcUsage)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting vpc_usage: %s", err), "(Data) ibm_mqcloud_capacity_details", "read", "set-vpc_usage").GetDiag()
		}
	}

	if !core.IsNil(usage.CapacityUnits) {
		if err = d.Set("capacity_units", flex.IntValue(usage.CapacityUnits)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting capacity_units: %s", err), "(Data) ibm_mqcloud_capacity_details", "read", "set-capacity_units").GetDiag()
		}
	}

	if !core.IsNil(usage.CapacityStatus) {
		if err = d.Set("capacity_status", usage.CapacityStatus); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting capacity_status: %s", err), "(Data) ibm_mqcloud_capacity_details", "read", "set-capacity_status").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmMqcloudCapacityDetailsID returns a reasonable ID for the list.
func dataSourceIbmMqcloudCapacityDetailsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
