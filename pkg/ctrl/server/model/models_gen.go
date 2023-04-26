// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"k8s.io/api/core/v1"
)

type Cgroups struct {
	Raw    string         `json:"raw"`
	CPU    *CgroupsCPU    `json:"cpu"`
	Memory *CgroupsMemory `json:"memory"`
}

type CgroupsCPU struct {
	Quota  int `json:"quota"`
	Period int `json:"period"`
}

type CgroupsMemory struct {
	Limit int64 `json:"limit"`
}

type Fd struct {
	Fd     string `json:"fd"`
	Target string `json:"target"`
}

type KillProcessResult struct {
	Pid     string `json:"pid"`
	Command string `json:"command"`
}

type MutablePod struct {
	Pod           *v1.Pod              `json:"pod"`
	KillProcesses []*KillProcessResult `json:"killProcesses"`
	CleanTcs      []string             `json:"cleanTcs"`
	CleanIptables []string             `json:"cleanIptables"`
}

type Namespace struct {
	Ns              string                      `json:"ns"`
	Component       []*v1.Pod                   `json:"component"`
	Pod             []*v1.Pod                   `json:"pod"`
	Stresschaos     []*v1alpha1.StressChaos     `json:"stresschaos"`
	Iochaos         []*v1alpha1.IOChaos         `json:"iochaos"`
	Podiochaos      []*v1alpha1.PodIOChaos      `json:"podiochaos"`
	Httpchaos       []*v1alpha1.HTTPChaos       `json:"httpchaos"`
	Podhttpchaos    []*v1alpha1.PodHttpChaos    `json:"podhttpchaos"`
	Networkchaos    []*v1alpha1.NetworkChaos    `json:"networkchaos"`
	Podnetworkchaos []*v1alpha1.PodNetworkChaos `json:"podnetworkchaos"`
}

type PodSelectorInput struct {
	Namespaces          []string               `json:"namespaces"`
	Nodes               []string               `json:"nodes"`
	Pods                map[string]interface{} `json:"pods"`
	NodeSelectors       map[string]interface{} `json:"nodeSelectors"`
	FieldSelectors      map[string]interface{} `json:"fieldSelectors"`
	LabelSelectors      map[string]interface{} `json:"labelSelectors"`
	AnnotationSelectors map[string]interface{} `json:"annotationSelectors"`
	PodPhaseSelectors   []string               `json:"podPhaseSelectors"`
}

type PodStressChaos struct {
	StressChaos   *v1alpha1.StressChaos `json:"stressChaos"`
	Pod           *v1.Pod               `json:"pod"`
	Cgroups       *Cgroups              `json:"cgroups"`
	ProcessStress []*ProcessStress      `json:"processStress"`
}

type Process struct {
	Pod     *v1.Pod `json:"pod"`
	Pid     string  `json:"pid"`
	Command string  `json:"command"`
	Fds     []*Fd   `json:"fds"`
}

type ProcessStress struct {
	Process *Process `json:"process"`
	Cgroup  string   `json:"cgroup"`
}

type Component string

const (
	ComponentManager   Component = "MANAGER"
	ComponentDaemon    Component = "DAEMON"
	ComponentDashboard Component = "DASHBOARD"
	ComponentDNSServer Component = "DNSSERVER"
)

var AllComponent = []Component{
	ComponentManager,
	ComponentDaemon,
	ComponentDashboard,
	ComponentDNSServer,
}

func (e Component) IsValid() bool {
	switch e {
	case ComponentManager, ComponentDaemon, ComponentDashboard, ComponentDNSServer:
		return true
	}
	return false
}

func (e Component) String() string {
	return string(e)
}

func (e *Component) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Component(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Component", str)
	}
	return nil
}

func (e Component) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
