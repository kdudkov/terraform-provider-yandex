package yandex

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1/instancegroup"
)

func dataSourceYandexComputeInstanceGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceYandexComputeInstanceGroupRead,

		SchemaVersion: 0,

		Schema: map[string]*schema.Schema{
			"instance_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"folder_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"service_account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"instance_template": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resources": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"memory": {
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"cores": {
										Type:     schema.TypeInt,
										Computed: true,
									},

									"core_fraction": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"boot_disk": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mode": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"initialize_params": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"description": {
													Type:     schema.TypeString,
													Computed: true,
												},

												"size": {
													Type:     schema.TypeInt,
													Computed: true,
												},

												"type": {
													Type:     schema.TypeString,
													Computed: true,
												},

												"image_id": {
													Type:     schema.TypeString,
													Computed: true,
												},

												"snapshot_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"device_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"platform_id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"metadata": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},

						"labels": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},

						"network_interface": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"network_id": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"subnet_ids": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"nat": {
										Type:     schema.TypeBool,
										Computed: true,
									},

									"ipv6": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},

						"secondary_disk": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mode": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"initialize_params": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"description": {
													Type:     schema.TypeString,
													Computed: true,
												},

												"size": {
													Type:     schema.TypeInt,
													Computed: true,
												},

												"type": {
													Type:     schema.TypeString,
													Computed: true,
												},

												"image_id": {
													Type:     schema.TypeString,
													Computed: true,
												},

												"snapshot_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"device_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"scheduling_policy": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"preemptible": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},

						"service_account_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"scale_policy": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fixed_scale": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"size": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},

			"deploy_policy": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_unavailable": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"max_expansion": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"max_deleting": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"max_creating": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"startup_duration": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},

			"allocation_policy": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zones": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"health_check": {
				Type:     schema.TypeList,
				MinItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interval": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"timeout": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"healthy_threshold": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"unhealthy_threshold": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"tcp_options": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"http_options": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port": {
										Type:     schema.TypeInt,
										Computed: true,
									},

									"path": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},

			"load_balancer": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target_group_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_group_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_group_labels": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						"target_group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status_message": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"instances": {
				Computed: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fqdn": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status_message": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"zone_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_interface": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"index": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"mac_address": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_address": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ipv6": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"ipv6_address": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"subnet_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"nat": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"nat_ip_address": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"nat_ip_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"load_balancer_state": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target_group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status_message": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceYandexComputeInstanceGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	ctx := context.Background()

	instanceGroupID := d.Get("instance_group_id").(string)

	if instanceGroupID == "" {
		return fmt.Errorf("instance_group_id should be provided")
	}

	instanceGroup, err := config.sdk.InstanceGroup().InstanceGroup().Get(ctx, &instancegroup.GetInstanceGroupRequest{
		InstanceGroupId: instanceGroupID,
		View:            instancegroup.InstanceGroupView_FULL,
	})

	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("Instance group %q", d.Get("name").(string)))
	}

	instances, err := config.sdk.InstanceGroup().InstanceGroup().ListInstances(ctx, &instancegroup.ListInstanceGroupInstancesRequest{
		InstanceGroupId: instanceGroupID,
	})

	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("Can't read instances for instance group with ID %q", instanceGroupID))
	}

	return flattenInstanceGroupDataSource(d, instanceGroup, instances.GetInstances())
}

func flattenInstanceGroupDataSource(d *schema.ResourceData, instanceGroup *instancegroup.InstanceGroup, instances []*instancegroup.ManagedInstance) error {
	err := flattenInstanceGroup(d, instanceGroup)

	if err != nil {
		return err
	}

	loadBalancerState, err := flattenInstanceGroupLoadBalancerState(instanceGroup)
	if err != nil {
		return err
	}
	if err := d.Set("load_balancer_state", loadBalancerState); err != nil {
		return err
	}

	if instances == nil {
		return nil
	}

	res := make([]map[string]interface{}, len(instances))

	for i, instance := range instances {
		instDict := make(map[string]interface{})
		instDict["status"] = instance.Status.String()
		instDict["instance_id"] = instance.InstanceId
		instDict["fqdn"] = instance.Fqdn
		instDict["name"] = instance.Name
		instDict["status_message"] = instance.StatusMessage
		instDict["zone_id"] = instance.ZoneId

		networkInterfaces, _, _, err := flattenInstanceGroupManagedInstanceNetworkInterfaces(instance)
		if err != nil {
			return err
		}

		instDict["network_interface"] = networkInterfaces
		res[i] = instDict
	}

	err = d.Set("instances", res)
	if err != nil {
		return err
	}

	d.SetId(instanceGroup.Id)

	return nil
}