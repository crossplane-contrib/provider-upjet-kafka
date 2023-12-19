//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACL) DeepCopyInto(out *ACL) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACL.
func (in *ACL) DeepCopy() *ACL {
	if in == nil {
		return nil
	}
	out := new(ACL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ACL) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACLInitParameters) DeepCopyInto(out *ACLInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACLInitParameters.
func (in *ACLInitParameters) DeepCopy() *ACLInitParameters {
	if in == nil {
		return nil
	}
	out := new(ACLInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACLList) DeepCopyInto(out *ACLList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ACL, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACLList.
func (in *ACLList) DeepCopy() *ACLList {
	if in == nil {
		return nil
	}
	out := new(ACLList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ACLList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACLObservation) DeepCopyInto(out *ACLObservation) {
	*out = *in
	if in.ACLHost != nil {
		in, out := &in.ACLHost, &out.ACLHost
		*out = new(string)
		**out = **in
	}
	if in.ACLOperation != nil {
		in, out := &in.ACLOperation, &out.ACLOperation
		*out = new(string)
		**out = **in
	}
	if in.ACLPermissionType != nil {
		in, out := &in.ACLPermissionType, &out.ACLPermissionType
		*out = new(string)
		**out = **in
	}
	if in.ACLPrincipal != nil {
		in, out := &in.ACLPrincipal, &out.ACLPrincipal
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ResourceName != nil {
		in, out := &in.ResourceName, &out.ResourceName
		*out = new(string)
		**out = **in
	}
	if in.ResourcePatternTypeFilter != nil {
		in, out := &in.ResourcePatternTypeFilter, &out.ResourcePatternTypeFilter
		*out = new(string)
		**out = **in
	}
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACLObservation.
func (in *ACLObservation) DeepCopy() *ACLObservation {
	if in == nil {
		return nil
	}
	out := new(ACLObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACLParameters) DeepCopyInto(out *ACLParameters) {
	*out = *in
	if in.ACLHost != nil {
		in, out := &in.ACLHost, &out.ACLHost
		*out = new(string)
		**out = **in
	}
	if in.ACLOperation != nil {
		in, out := &in.ACLOperation, &out.ACLOperation
		*out = new(string)
		**out = **in
	}
	if in.ACLPermissionType != nil {
		in, out := &in.ACLPermissionType, &out.ACLPermissionType
		*out = new(string)
		**out = **in
	}
	if in.ACLPrincipal != nil {
		in, out := &in.ACLPrincipal, &out.ACLPrincipal
		*out = new(string)
		**out = **in
	}
	if in.ResourceName != nil {
		in, out := &in.ResourceName, &out.ResourceName
		*out = new(string)
		**out = **in
	}
	if in.ResourcePatternTypeFilter != nil {
		in, out := &in.ResourcePatternTypeFilter, &out.ResourcePatternTypeFilter
		*out = new(string)
		**out = **in
	}
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACLParameters.
func (in *ACLParameters) DeepCopy() *ACLParameters {
	if in == nil {
		return nil
	}
	out := new(ACLParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACLSpec) DeepCopyInto(out *ACLSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	out.InitProvider = in.InitProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACLSpec.
func (in *ACLSpec) DeepCopy() *ACLSpec {
	if in == nil {
		return nil
	}
	out := new(ACLSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACLStatus) DeepCopyInto(out *ACLStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACLStatus.
func (in *ACLStatus) DeepCopy() *ACLStatus {
	if in == nil {
		return nil
	}
	out := new(ACLStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Quota) DeepCopyInto(out *Quota) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Quota.
func (in *Quota) DeepCopy() *Quota {
	if in == nil {
		return nil
	}
	out := new(Quota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Quota) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaInitParameters) DeepCopyInto(out *QuotaInitParameters) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]float64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EntityName != nil {
		in, out := &in.EntityName, &out.EntityName
		*out = new(string)
		**out = **in
	}
	if in.EntityType != nil {
		in, out := &in.EntityType, &out.EntityType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaInitParameters.
func (in *QuotaInitParameters) DeepCopy() *QuotaInitParameters {
	if in == nil {
		return nil
	}
	out := new(QuotaInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaList) DeepCopyInto(out *QuotaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Quota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaList.
func (in *QuotaList) DeepCopy() *QuotaList {
	if in == nil {
		return nil
	}
	out := new(QuotaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QuotaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaObservation) DeepCopyInto(out *QuotaObservation) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]float64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EntityName != nil {
		in, out := &in.EntityName, &out.EntityName
		*out = new(string)
		**out = **in
	}
	if in.EntityType != nil {
		in, out := &in.EntityType, &out.EntityType
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaObservation.
func (in *QuotaObservation) DeepCopy() *QuotaObservation {
	if in == nil {
		return nil
	}
	out := new(QuotaObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaParameters) DeepCopyInto(out *QuotaParameters) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]float64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EntityName != nil {
		in, out := &in.EntityName, &out.EntityName
		*out = new(string)
		**out = **in
	}
	if in.EntityType != nil {
		in, out := &in.EntityType, &out.EntityType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaParameters.
func (in *QuotaParameters) DeepCopy() *QuotaParameters {
	if in == nil {
		return nil
	}
	out := new(QuotaParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaSpec) DeepCopyInto(out *QuotaSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaSpec.
func (in *QuotaSpec) DeepCopy() *QuotaSpec {
	if in == nil {
		return nil
	}
	out := new(QuotaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaStatus) DeepCopyInto(out *QuotaStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaStatus.
func (in *QuotaStatus) DeepCopy() *QuotaStatus {
	if in == nil {
		return nil
	}
	out := new(QuotaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Topic) DeepCopyInto(out *Topic) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Topic.
func (in *Topic) DeepCopy() *Topic {
	if in == nil {
		return nil
	}
	out := new(Topic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Topic) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopicInitParameters) DeepCopyInto(out *TopicInitParameters) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Partitions != nil {
		in, out := &in.Partitions, &out.Partitions
		*out = new(int64)
		**out = **in
	}
	if in.ReplicationFactor != nil {
		in, out := &in.ReplicationFactor, &out.ReplicationFactor
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopicInitParameters.
func (in *TopicInitParameters) DeepCopy() *TopicInitParameters {
	if in == nil {
		return nil
	}
	out := new(TopicInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopicList) DeepCopyInto(out *TopicList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Topic, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopicList.
func (in *TopicList) DeepCopy() *TopicList {
	if in == nil {
		return nil
	}
	out := new(TopicList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TopicList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopicObservation) DeepCopyInto(out *TopicObservation) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Partitions != nil {
		in, out := &in.Partitions, &out.Partitions
		*out = new(int64)
		**out = **in
	}
	if in.ReplicationFactor != nil {
		in, out := &in.ReplicationFactor, &out.ReplicationFactor
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopicObservation.
func (in *TopicObservation) DeepCopy() *TopicObservation {
	if in == nil {
		return nil
	}
	out := new(TopicObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopicParameters) DeepCopyInto(out *TopicParameters) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Partitions != nil {
		in, out := &in.Partitions, &out.Partitions
		*out = new(int64)
		**out = **in
	}
	if in.ReplicationFactor != nil {
		in, out := &in.ReplicationFactor, &out.ReplicationFactor
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopicParameters.
func (in *TopicParameters) DeepCopy() *TopicParameters {
	if in == nil {
		return nil
	}
	out := new(TopicParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopicSpec) DeepCopyInto(out *TopicSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopicSpec.
func (in *TopicSpec) DeepCopy() *TopicSpec {
	if in == nil {
		return nil
	}
	out := new(TopicSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopicStatus) DeepCopyInto(out *TopicStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopicStatus.
func (in *TopicStatus) DeepCopy() *TopicStatus {
	if in == nil {
		return nil
	}
	out := new(TopicStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserScramCredential) DeepCopyInto(out *UserScramCredential) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserScramCredential.
func (in *UserScramCredential) DeepCopy() *UserScramCredential {
	if in == nil {
		return nil
	}
	out := new(UserScramCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserScramCredential) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserScramCredentialInitParameters) DeepCopyInto(out *UserScramCredentialInitParameters) {
	*out = *in
	if in.ScramIterations != nil {
		in, out := &in.ScramIterations, &out.ScramIterations
		*out = new(int64)
		**out = **in
	}
	if in.ScramMechanism != nil {
		in, out := &in.ScramMechanism, &out.ScramMechanism
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserScramCredentialInitParameters.
func (in *UserScramCredentialInitParameters) DeepCopy() *UserScramCredentialInitParameters {
	if in == nil {
		return nil
	}
	out := new(UserScramCredentialInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserScramCredentialList) DeepCopyInto(out *UserScramCredentialList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UserScramCredential, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserScramCredentialList.
func (in *UserScramCredentialList) DeepCopy() *UserScramCredentialList {
	if in == nil {
		return nil
	}
	out := new(UserScramCredentialList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserScramCredentialList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserScramCredentialObservation) DeepCopyInto(out *UserScramCredentialObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ScramIterations != nil {
		in, out := &in.ScramIterations, &out.ScramIterations
		*out = new(int64)
		**out = **in
	}
	if in.ScramMechanism != nil {
		in, out := &in.ScramMechanism, &out.ScramMechanism
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserScramCredentialObservation.
func (in *UserScramCredentialObservation) DeepCopy() *UserScramCredentialObservation {
	if in == nil {
		return nil
	}
	out := new(UserScramCredentialObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserScramCredentialParameters) DeepCopyInto(out *UserScramCredentialParameters) {
	*out = *in
	out.PasswordSecretRef = in.PasswordSecretRef
	if in.ScramIterations != nil {
		in, out := &in.ScramIterations, &out.ScramIterations
		*out = new(int64)
		**out = **in
	}
	if in.ScramMechanism != nil {
		in, out := &in.ScramMechanism, &out.ScramMechanism
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserScramCredentialParameters.
func (in *UserScramCredentialParameters) DeepCopy() *UserScramCredentialParameters {
	if in == nil {
		return nil
	}
	out := new(UserScramCredentialParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserScramCredentialSpec) DeepCopyInto(out *UserScramCredentialSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserScramCredentialSpec.
func (in *UserScramCredentialSpec) DeepCopy() *UserScramCredentialSpec {
	if in == nil {
		return nil
	}
	out := new(UserScramCredentialSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserScramCredentialStatus) DeepCopyInto(out *UserScramCredentialStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserScramCredentialStatus.
func (in *UserScramCredentialStatus) DeepCopy() *UserScramCredentialStatus {
	if in == nil {
		return nil
	}
	out := new(UserScramCredentialStatus)
	in.DeepCopyInto(out)
	return out
}
