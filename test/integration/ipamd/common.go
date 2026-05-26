package ipamd

import (
	"github.com/aws/amazon-vpc-cni-k8s/test/framework"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

var primaryInstance types.Instance
var f *framework.Framework
var err error

func ceil(x, y int) int { _ = "STUB: not implemented"; return 0 }

func Max(x, y int) int { _ = "STUB: not implemented"; return 0 }

// MinIgnoreZero returns smaller of two number, if any number is zero returns the other number
func MinIgnoreZero(x, y int) int { _ = "STUB: not implemented"; return 0 }
