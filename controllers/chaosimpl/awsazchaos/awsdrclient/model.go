package awsdrclient

type AutoScalingGroupState struct {
	AutoScalingGroupName string
	AvailabilityZones    []string
	DesiredCapacity      int32
	MaxSize              int32
	MinSize              int32
}
