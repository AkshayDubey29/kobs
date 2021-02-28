// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: clusters.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// GetClustersRequest is the request to get all loaded Kubernetes clusters via the GetClusters method.
type GetClustersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetClustersRequest) Reset() {
	*x = GetClustersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clusters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClustersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClustersRequest) ProtoMessage() {}

func (x *GetClustersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clusters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClustersRequest.ProtoReflect.Descriptor instead.
func (*GetClustersRequest) Descriptor() ([]byte, []int) {
	return file_clusters_proto_rawDescGZIP(), []int{0}
}

// GetClustersResponse is the response for a GetClusters request. It contains a clusers field, which contains the names
// of all clusters.
type GetClustersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clusters []string `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *GetClustersResponse) Reset() {
	*x = GetClustersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clusters_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClustersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClustersResponse) ProtoMessage() {}

func (x *GetClustersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clusters_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClustersResponse.ProtoReflect.Descriptor instead.
func (*GetClustersResponse) Descriptor() ([]byte, []int) {
	return file_clusters_proto_rawDescGZIP(), []int{1}
}

func (x *GetClustersResponse) GetClusters() []string {
	if x != nil {
		return x.Clusters
	}
	return nil
}

// GetNamespacesRequest is the request to get all namespaces via the GetNamespaces method. It must contain a list of
// clusters, for which the namespaces should be returned.
type GetNamespacesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clusters []string `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *GetNamespacesRequest) Reset() {
	*x = GetNamespacesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clusters_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNamespacesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNamespacesRequest) ProtoMessage() {}

func (x *GetNamespacesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clusters_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNamespacesRequest.ProtoReflect.Descriptor instead.
func (*GetNamespacesRequest) Descriptor() ([]byte, []int) {
	return file_clusters_proto_rawDescGZIP(), []int{2}
}

func (x *GetNamespacesRequest) GetClusters() []string {
	if x != nil {
		return x.Clusters
	}
	return nil
}

// GetNamespacesResponse is the response for a GetNamespaces request, which contains a list of all namespaces for the
// given clusters.
type GetNamespacesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespaces []string `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
}

func (x *GetNamespacesResponse) Reset() {
	*x = GetNamespacesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clusters_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNamespacesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNamespacesResponse) ProtoMessage() {}

func (x *GetNamespacesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clusters_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNamespacesResponse.ProtoReflect.Descriptor instead.
func (*GetNamespacesResponse) Descriptor() ([]byte, []int) {
	return file_clusters_proto_rawDescGZIP(), []int{3}
}

func (x *GetNamespacesResponse) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

// GetResourcesRequest is the request to get a specific resource for multiple clusters and namespaces. It contains the
// Kubernetes API endpoint for the resource and a list of clusters and namespaces. It is also possible to specify a
// parameter, which can be used to set a labelSelector or fieldSelector for the request.
type GetResourcesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clusters   []string `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	Namespaces []string `protobuf:"bytes,2,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	Path       string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Resource   string   `protobuf:"bytes,4,opt,name=resource,proto3" json:"resource,omitempty"`
	ParamName  string   `protobuf:"bytes,5,opt,name=paramName,proto3" json:"paramName,omitempty"`
	Param      string   `protobuf:"bytes,6,opt,name=param,proto3" json:"param,omitempty"`
}

func (x *GetResourcesRequest) Reset() {
	*x = GetResourcesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clusters_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourcesRequest) ProtoMessage() {}

func (x *GetResourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clusters_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourcesRequest.ProtoReflect.Descriptor instead.
func (*GetResourcesRequest) Descriptor() ([]byte, []int) {
	return file_clusters_proto_rawDescGZIP(), []int{4}
}

func (x *GetResourcesRequest) GetClusters() []string {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *GetResourcesRequest) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

func (x *GetResourcesRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *GetResourcesRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *GetResourcesRequest) GetParamName() string {
	if x != nil {
		return x.ParamName
	}
	return ""
}

func (x *GetResourcesRequest) GetParam() string {
	if x != nil {
		return x.Param
	}
	return ""
}

// GetResourcesResponse is the response for a GetResources request, which contains a list of resources. The resources
// are returned by cluster and namespace.
type GetResourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources []*Resources `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *GetResourcesResponse) Reset() {
	*x = GetResourcesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clusters_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourcesResponse) ProtoMessage() {}

func (x *GetResourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clusters_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourcesResponse.ProtoReflect.Descriptor instead.
func (*GetResourcesResponse) Descriptor() ([]byte, []int) {
	return file_clusters_proto_rawDescGZIP(), []int{5}
}

func (x *GetResourcesResponse) GetResources() []*Resources {
	if x != nil {
		return x.Resources
	}
	return nil
}

// Resources is the structure for a list of resources for a cluster and a namespaces. The resourceList string contains
// the JSON string for the list type of a resource (e.g. PodList).
type Resources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster      string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Namespace    string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ResourceList string `protobuf:"bytes,3,opt,name=resourceList,proto3" json:"resourceList,omitempty"`
}

func (x *Resources) Reset() {
	*x = Resources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clusters_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resources) ProtoMessage() {}

func (x *Resources) ProtoReflect() protoreflect.Message {
	mi := &file_clusters_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resources.ProtoReflect.Descriptor instead.
func (*Resources) Descriptor() ([]byte, []int) {
	return file_clusters_proto_rawDescGZIP(), []int{6}
}

func (x *Resources) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *Resources) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Resources) GetResourceList() string {
	if x != nil {
		return x.ResourceList
	}
	return ""
}

// GetApplicationsRequest is the message formate to get a list of applications. To get a list of applications the
// clusters and namespaces for which the applications should be retrieved must be specified.
type GetApplicationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clusters   []string `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	Namespaces []string `protobuf:"bytes,2,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
}

func (x *GetApplicationsRequest) Reset() {
	*x = GetApplicationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clusters_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApplicationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApplicationsRequest) ProtoMessage() {}

func (x *GetApplicationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clusters_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApplicationsRequest.ProtoReflect.Descriptor instead.
func (*GetApplicationsRequest) Descriptor() ([]byte, []int) {
	return file_clusters_proto_rawDescGZIP(), []int{7}
}

func (x *GetApplicationsRequest) GetClusters() []string {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *GetApplicationsRequest) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

// GetApplicationsResponse is the response for a GetApplications request, which returns a list of applications.
type GetApplicationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Applications []*Application `protobuf:"bytes,1,rep,name=applications,proto3" json:"applications,omitempty"`
}

func (x *GetApplicationsResponse) Reset() {
	*x = GetApplicationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clusters_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApplicationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApplicationsResponse) ProtoMessage() {}

func (x *GetApplicationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clusters_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApplicationsResponse.ProtoReflect.Descriptor instead.
func (*GetApplicationsResponse) Descriptor() ([]byte, []int) {
	return file_clusters_proto_rawDescGZIP(), []int{8}
}

func (x *GetApplicationsResponse) GetApplications() []*Application {
	if x != nil {
		return x.Applications
	}
	return nil
}

// GetApplicationRequest is the format to get a single application. Each application can be identified by the cluster,
// namespace and name of the application.
type GetApplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster   string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetApplicationRequest) Reset() {
	*x = GetApplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clusters_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApplicationRequest) ProtoMessage() {}

func (x *GetApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clusters_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApplicationRequest.ProtoReflect.Descriptor instead.
func (*GetApplicationRequest) Descriptor() ([]byte, []int) {
	return file_clusters_proto_rawDescGZIP(), []int{9}
}

func (x *GetApplicationRequest) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *GetApplicationRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *GetApplicationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// GetApplicationResponse is the response for a GetApplication request, which returns a single application.
type GetApplicationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Application *Application `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
}

func (x *GetApplicationResponse) Reset() {
	*x = GetApplicationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clusters_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApplicationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApplicationResponse) ProtoMessage() {}

func (x *GetApplicationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clusters_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApplicationResponse.ProtoReflect.Descriptor instead.
func (*GetApplicationResponse) Descriptor() ([]byte, []int) {
	return file_clusters_proto_rawDescGZIP(), []int{10}
}

func (x *GetApplicationResponse) GetApplication() *Application {
	if x != nil {
		return x.Application
	}
	return nil
}

var File_clusters_proto protoreflect.FileDescriptor

var file_clusters_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x11, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x32, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x37, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x22, 0xb5, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x49, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x09, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x67, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x54, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x22, 0x57, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3c, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x63,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x54, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a,
	0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xae, 0x03, 0x0a, 0x08, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x4c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x2e, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x62, 0x73, 0x69, 0x6f, 0x2f,
	0x6b, 0x6f, 0x62, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_clusters_proto_rawDescOnce sync.Once
	file_clusters_proto_rawDescData = file_clusters_proto_rawDesc
)

func file_clusters_proto_rawDescGZIP() []byte {
	file_clusters_proto_rawDescOnce.Do(func() {
		file_clusters_proto_rawDescData = protoimpl.X.CompressGZIP(file_clusters_proto_rawDescData)
	})
	return file_clusters_proto_rawDescData
}

var file_clusters_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_clusters_proto_goTypes = []interface{}{
	(*GetClustersRequest)(nil),      // 0: clusters.GetClustersRequest
	(*GetClustersResponse)(nil),     // 1: clusters.GetClustersResponse
	(*GetNamespacesRequest)(nil),    // 2: clusters.GetNamespacesRequest
	(*GetNamespacesResponse)(nil),   // 3: clusters.GetNamespacesResponse
	(*GetResourcesRequest)(nil),     // 4: clusters.GetResourcesRequest
	(*GetResourcesResponse)(nil),    // 5: clusters.GetResourcesResponse
	(*Resources)(nil),               // 6: clusters.Resources
	(*GetApplicationsRequest)(nil),  // 7: clusters.GetApplicationsRequest
	(*GetApplicationsResponse)(nil), // 8: clusters.GetApplicationsResponse
	(*GetApplicationRequest)(nil),   // 9: clusters.GetApplicationRequest
	(*GetApplicationResponse)(nil),  // 10: clusters.GetApplicationResponse
	(*Application)(nil),             // 11: application.Application
}
var file_clusters_proto_depIdxs = []int32{
	6,  // 0: clusters.GetResourcesResponse.resources:type_name -> clusters.Resources
	11, // 1: clusters.GetApplicationsResponse.applications:type_name -> application.Application
	11, // 2: clusters.GetApplicationResponse.application:type_name -> application.Application
	0,  // 3: clusters.Clusters.GetClusters:input_type -> clusters.GetClustersRequest
	2,  // 4: clusters.Clusters.GetNamespaces:input_type -> clusters.GetNamespacesRequest
	4,  // 5: clusters.Clusters.GetResources:input_type -> clusters.GetResourcesRequest
	7,  // 6: clusters.Clusters.GetApplications:input_type -> clusters.GetApplicationsRequest
	9,  // 7: clusters.Clusters.GetApplication:input_type -> clusters.GetApplicationRequest
	1,  // 8: clusters.Clusters.GetClusters:output_type -> clusters.GetClustersResponse
	3,  // 9: clusters.Clusters.GetNamespaces:output_type -> clusters.GetNamespacesResponse
	5,  // 10: clusters.Clusters.GetResources:output_type -> clusters.GetResourcesResponse
	8,  // 11: clusters.Clusters.GetApplications:output_type -> clusters.GetApplicationsResponse
	10, // 12: clusters.Clusters.GetApplication:output_type -> clusters.GetApplicationResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_clusters_proto_init() }
func file_clusters_proto_init() {
	if File_clusters_proto != nil {
		return
	}
	file_application_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_clusters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClustersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_clusters_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClustersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_clusters_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNamespacesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_clusters_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNamespacesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_clusters_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourcesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_clusters_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourcesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_clusters_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resources); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_clusters_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApplicationsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_clusters_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApplicationsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_clusters_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApplicationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_clusters_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApplicationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_clusters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_clusters_proto_goTypes,
		DependencyIndexes: file_clusters_proto_depIdxs,
		MessageInfos:      file_clusters_proto_msgTypes,
	}.Build()
	File_clusters_proto = out.File
	file_clusters_proto_rawDesc = nil
	file_clusters_proto_goTypes = nil
	file_clusters_proto_depIdxs = nil
}
