//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkDeployment) DeepCopyInto(out *FlinkDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkDeployment.
func (in *FlinkDeployment) DeepCopy() *FlinkDeployment {
	if in == nil {
		return nil
	}
	out := new(FlinkDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlinkDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkDeploymentList) DeepCopyInto(out *FlinkDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlinkDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkDeploymentList.
func (in *FlinkDeploymentList) DeepCopy() *FlinkDeploymentList {
	if in == nil {
		return nil
	}
	out := new(FlinkDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlinkDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkDeploymentSpec) DeepCopyInto(out *FlinkDeploymentSpec) {
	*out = *in
	if in.FlinkConfiguration != nil {
		in, out := &in.FlinkConfiguration, &out.FlinkConfiguration
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.LogConfiguration != nil {
		in, out := &in.LogConfiguration, &out.LogConfiguration
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.JobManager.DeepCopyInto(&out.JobManager)
	in.TaskManager.DeepCopyInto(&out.TaskManager)
	in.JobSpec.DeepCopyInto(&out.JobSpec)
	in.PodTemplateSpec.DeepCopyInto(&out.PodTemplateSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkDeploymentSpec.
func (in *FlinkDeploymentSpec) DeepCopy() *FlinkDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(FlinkDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkDeploymentStatus) DeepCopyInto(out *FlinkDeploymentStatus) {
	*out = *in
	if in.ClusterInfo != nil {
		in, out := &in.ClusterInfo, &out.ClusterInfo
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.TaskManager = in.TaskManager
	out.JobStatus = in.JobStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkDeploymentStatus.
func (in *FlinkDeploymentStatus) DeepCopy() *FlinkDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(FlinkDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkJobStatusSpec) DeepCopyInto(out *FlinkJobStatusSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkJobStatusSpec.
func (in *FlinkJobStatusSpec) DeepCopy() *FlinkJobStatusSpec {
	if in == nil {
		return nil
	}
	out := new(FlinkJobStatusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkSessionJob) DeepCopyInto(out *FlinkSessionJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkSessionJob.
func (in *FlinkSessionJob) DeepCopy() *FlinkSessionJob {
	if in == nil {
		return nil
	}
	out := new(FlinkSessionJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlinkSessionJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkSessionJobList) DeepCopyInto(out *FlinkSessionJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlinkSessionJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkSessionJobList.
func (in *FlinkSessionJobList) DeepCopy() *FlinkSessionJobList {
	if in == nil {
		return nil
	}
	out := new(FlinkSessionJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlinkSessionJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkSessionJobSpec) DeepCopyInto(out *FlinkSessionJobSpec) {
	*out = *in
	in.Job.DeepCopyInto(&out.Job)
	if in.FlinkConfiguration != nil {
		in, out := &in.FlinkConfiguration, &out.FlinkConfiguration
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkSessionJobSpec.
func (in *FlinkSessionJobSpec) DeepCopy() *FlinkSessionJobSpec {
	if in == nil {
		return nil
	}
	out := new(FlinkSessionJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkSessionJobStatus) DeepCopyInto(out *FlinkSessionJobStatus) {
	*out = *in
	out.FlinkJobStatus = in.FlinkJobStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkSessionJobStatus.
func (in *FlinkSessionJobStatus) DeepCopy() *FlinkSessionJobStatus {
	if in == nil {
		return nil
	}
	out := new(FlinkSessionJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobManagerSpec) DeepCopyInto(out *JobManagerSpec) {
	*out = *in
	out.Resource = in.Resource
	in.PodTemplate.DeepCopyInto(&out.PodTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobManagerSpec.
func (in *JobManagerSpec) DeepCopy() *JobManagerSpec {
	if in == nil {
		return nil
	}
	out := new(JobManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobSpec) DeepCopyInto(out *JobSpec) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobSpec.
func (in *JobSpec) DeepCopy() *JobSpec {
	if in == nil {
		return nil
	}
	out := new(JobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobStatusInfo) DeepCopyInto(out *JobStatusInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobStatusInfo.
func (in *JobStatusInfo) DeepCopy() *JobStatusInfo {
	if in == nil {
		return nil
	}
	out := new(JobStatusInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSpec) DeepCopyInto(out *ResourceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSpec.
func (in *ResourceSpec) DeepCopy() *ResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskManagerInfo) DeepCopyInto(out *TaskManagerInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskManagerInfo.
func (in *TaskManagerInfo) DeepCopy() *TaskManagerInfo {
	if in == nil {
		return nil
	}
	out := new(TaskManagerInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskManagerSpec) DeepCopyInto(out *TaskManagerSpec) {
	*out = *in
	out.Resource = in.Resource
	in.PodTemplate.DeepCopyInto(&out.PodTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskManagerSpec.
func (in *TaskManagerSpec) DeepCopy() *TaskManagerSpec {
	if in == nil {
		return nil
	}
	out := new(TaskManagerSpec)
	in.DeepCopyInto(out)
	return out
}
