package healthcheck

const (
	// SourceConfig the check is managed by the configuration file
	SourceConfig string = ""
	// SourceAPI the check is managed by the API
	SourceAPI string = "api"
	// SourceKubernetesPod the check was created from a Kubernetes pod
	SourceKubernetesPod string = "kubernetes-pod"
	// SourceKubernetesService the check was created from a service pod
	SourceKubernetesService string = "kubernetes-service"
)

// Base shared fields between healthchecks
type Base struct {
	Name string `json:"name"`
	// +kubebuilder:validation:Optional
	Description string `json:"description"`
	// +kubebuilder:validation:Type=string
	Interval Duration `json:"interval"`
	// +kubebuilder:validation:Optional
	OneOff bool `json:"one-off"`
	// +kubebuilder:validation:Optional
	Source string `json:"source"`
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`
}

// SourceChecksNames returns all checks managed by the given source
func (c *Component) SourceChecksNames(source string) map[string]bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	checks := make(map[string]bool)
	for i := range c.Healthchecks {
		wrapper := c.Healthchecks[i]
		if wrapper.healthcheck.Base().Source == source {
			checks[wrapper.healthcheck.Base().Name] = true
		}
	}
	return checks
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Base) DeepCopyInto(out *Base) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Base.
func (in *Base) DeepCopy() *Base {
	if in == nil {
		return nil
	}
	out := new(Base)
	in.DeepCopyInto(out)
	return out
}
