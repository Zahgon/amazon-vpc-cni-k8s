package utils

import (
	"github.com/aws/amazon-vpc-cni-k8s/test/framework/resources/aws/services"
)

func WaitTillAddonIsDeleted(eks services.EKS, addonName string, clusterName string) error {
	_ = "STUB: not implemented"
	return nil
}

func WaitTillAddonIsActive(eks services.EKS, addonName string, clusterName string) error {
	_ = "STUB: not implemented"
	return nil
}
