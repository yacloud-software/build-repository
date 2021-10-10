// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/datastore/admin/v1beta1/datastore_admin.proto

package admin

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Operation types.
type OperationType int32

const (
	// Unspecified.
	OperationType_OPERATION_TYPE_UNSPECIFIED OperationType = 0
	// ExportEntities.
	OperationType_EXPORT_ENTITIES OperationType = 1
	// ImportEntities.
	OperationType_IMPORT_ENTITIES OperationType = 2
)

var OperationType_name = map[int32]string{
	0: "OPERATION_TYPE_UNSPECIFIED",
	1: "EXPORT_ENTITIES",
	2: "IMPORT_ENTITIES",
}

var OperationType_value = map[string]int32{
	"OPERATION_TYPE_UNSPECIFIED": 0,
	"EXPORT_ENTITIES":            1,
	"IMPORT_ENTITIES":            2,
}

func (x OperationType) String() string {
	return proto.EnumName(OperationType_name, int32(x))
}

func (OperationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b63216f38706ce20, []int{0}
}

// The various possible states for an ongoing Operation.
type CommonMetadata_State int32

const (
	// Unspecified.
	CommonMetadata_STATE_UNSPECIFIED CommonMetadata_State = 0
	// Request is being prepared for processing.
	CommonMetadata_INITIALIZING CommonMetadata_State = 1
	// Request is actively being processed.
	CommonMetadata_PROCESSING CommonMetadata_State = 2
	// Request is in the process of being cancelled after user called
	// google.longrunning.Operations.CancelOperation on the operation.
	CommonMetadata_CANCELLING CommonMetadata_State = 3
	// Request has been processed and is in its finalization stage.
	CommonMetadata_FINALIZING CommonMetadata_State = 4
	// Request has completed successfully.
	CommonMetadata_SUCCESSFUL CommonMetadata_State = 5
	// Request has finished being processed, but encountered an error.
	CommonMetadata_FAILED CommonMetadata_State = 6
	// Request has finished being cancelled after user called
	// google.longrunning.Operations.CancelOperation.
	CommonMetadata_CANCELLED CommonMetadata_State = 7
)

var CommonMetadata_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "INITIALIZING",
	2: "PROCESSING",
	3: "CANCELLING",
	4: "FINALIZING",
	5: "SUCCESSFUL",
	6: "FAILED",
	7: "CANCELLED",
}

var CommonMetadata_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"INITIALIZING":      1,
	"PROCESSING":        2,
	"CANCELLING":        3,
	"FINALIZING":        4,
	"SUCCESSFUL":        5,
	"FAILED":            6,
	"CANCELLED":         7,
}

func (x CommonMetadata_State) String() string {
	return proto.EnumName(CommonMetadata_State_name, int32(x))
}

func (CommonMetadata_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b63216f38706ce20, []int{0, 0}
}

// Metadata common to all Datastore Admin operations.
type CommonMetadata struct {
	// The time that work began on the operation.
	StartTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The time the operation ended, either successfully or otherwise.
	EndTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// The type of the operation. Can be used as a filter in
	// ListOperationsRequest.
	OperationType OperationType `protobuf:"varint,3,opt,name=operation_type,json=operationType,proto3,enum=google.datastore.admin.v1beta1.OperationType" json:"operation_type,omitempty"`
	// The client-assigned labels which were provided when the operation was
	// created. May also include additional labels.
	Labels map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The current state of the Operation.
	State                CommonMetadata_State `protobuf:"varint,5,opt,name=state,proto3,enum=google.datastore.admin.v1beta1.CommonMetadata_State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CommonMetadata) Reset()         { *m = CommonMetadata{} }
func (m *CommonMetadata) String() string { return proto.CompactTextString(m) }
func (*CommonMetadata) ProtoMessage()    {}
func (*CommonMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_b63216f38706ce20, []int{0}
}

func (m *CommonMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonMetadata.Unmarshal(m, b)
}
func (m *CommonMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonMetadata.Marshal(b, m, deterministic)
}
func (m *CommonMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonMetadata.Merge(m, src)
}
func (m *CommonMetadata) XXX_Size() int {
	return xxx_messageInfo_CommonMetadata.Size(m)
}
func (m *CommonMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CommonMetadata proto.InternalMessageInfo

func (m *CommonMetadata) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *CommonMetadata) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *CommonMetadata) GetOperationType() OperationType {
	if m != nil {
		return m.OperationType
	}
	return OperationType_OPERATION_TYPE_UNSPECIFIED
}

func (m *CommonMetadata) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *CommonMetadata) GetState() CommonMetadata_State {
	if m != nil {
		return m.State
	}
	return CommonMetadata_STATE_UNSPECIFIED
}

// Measures the progress of a particular metric.
type Progress struct {
	// The amount of work that has been completed. Note that this may be greater
	// than work_estimated.
	WorkCompleted int64 `protobuf:"varint,1,opt,name=work_completed,json=workCompleted,proto3" json:"work_completed,omitempty"`
	// An estimate of how much work needs to be performed. May be zero if the
	// work estimate is unavailable.
	WorkEstimated        int64    `protobuf:"varint,2,opt,name=work_estimated,json=workEstimated,proto3" json:"work_estimated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Progress) Reset()         { *m = Progress{} }
func (m *Progress) String() string { return proto.CompactTextString(m) }
func (*Progress) ProtoMessage()    {}
func (*Progress) Descriptor() ([]byte, []int) {
	return fileDescriptor_b63216f38706ce20, []int{1}
}

func (m *Progress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Progress.Unmarshal(m, b)
}
func (m *Progress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Progress.Marshal(b, m, deterministic)
}
func (m *Progress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Progress.Merge(m, src)
}
func (m *Progress) XXX_Size() int {
	return xxx_messageInfo_Progress.Size(m)
}
func (m *Progress) XXX_DiscardUnknown() {
	xxx_messageInfo_Progress.DiscardUnknown(m)
}

var xxx_messageInfo_Progress proto.InternalMessageInfo

func (m *Progress) GetWorkCompleted() int64 {
	if m != nil {
		return m.WorkCompleted
	}
	return 0
}

func (m *Progress) GetWorkEstimated() int64 {
	if m != nil {
		return m.WorkEstimated
	}
	return 0
}

// The request for
// [google.datastore.admin.v1beta1.DatastoreAdmin.ExportEntities][google.datastore.admin.v1beta1.DatastoreAdmin.ExportEntities].
type ExportEntitiesRequest struct {
	// Project ID against which to make the request.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Client-assigned labels.
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Description of what data from the project is included in the export.
	EntityFilter *EntityFilter `protobuf:"bytes,3,opt,name=entity_filter,json=entityFilter,proto3" json:"entity_filter,omitempty"`
	// Location for the export metadata and data files.
	//
	// The full resource URL of the external storage location. Currently, only
	// Google Cloud Storage is supported. So output_url_prefix should be of the
	// form: `gs://BUCKET_NAME[/NAMESPACE_PATH]`, where `BUCKET_NAME` is the
	// name of the Cloud Storage bucket and `NAMESPACE_PATH` is an optional Cloud
	// Storage namespace path (this is not a Cloud Datastore namespace). For more
	// information about Cloud Storage namespace paths, see
	// [Object name
	// considerations](https://cloud.google.com/storage/docs/naming#object-considerations).
	//
	// The resulting files will be nested deeper than the specified URL prefix.
	// The final output URL will be provided in the
	// [google.datastore.admin.v1beta1.ExportEntitiesResponse.output_url][google.datastore.admin.v1beta1.ExportEntitiesResponse.output_url]
	// field. That value should be used for subsequent ImportEntities operations.
	//
	// By nesting the data files deeper, the same Cloud Storage bucket can be used
	// in multiple ExportEntities operations without conflict.
	OutputUrlPrefix      string   `protobuf:"bytes,4,opt,name=output_url_prefix,json=outputUrlPrefix,proto3" json:"output_url_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportEntitiesRequest) Reset()         { *m = ExportEntitiesRequest{} }
func (m *ExportEntitiesRequest) String() string { return proto.CompactTextString(m) }
func (*ExportEntitiesRequest) ProtoMessage()    {}
func (*ExportEntitiesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b63216f38706ce20, []int{2}
}

func (m *ExportEntitiesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportEntitiesRequest.Unmarshal(m, b)
}
func (m *ExportEntitiesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportEntitiesRequest.Marshal(b, m, deterministic)
}
func (m *ExportEntitiesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportEntitiesRequest.Merge(m, src)
}
func (m *ExportEntitiesRequest) XXX_Size() int {
	return xxx_messageInfo_ExportEntitiesRequest.Size(m)
}
func (m *ExportEntitiesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportEntitiesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportEntitiesRequest proto.InternalMessageInfo

func (m *ExportEntitiesRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *ExportEntitiesRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ExportEntitiesRequest) GetEntityFilter() *EntityFilter {
	if m != nil {
		return m.EntityFilter
	}
	return nil
}

func (m *ExportEntitiesRequest) GetOutputUrlPrefix() string {
	if m != nil {
		return m.OutputUrlPrefix
	}
	return ""
}

// The request for
// [google.datastore.admin.v1beta1.DatastoreAdmin.ImportEntities][google.datastore.admin.v1beta1.DatastoreAdmin.ImportEntities].
type ImportEntitiesRequest struct {
	// Project ID against which to make the request.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Client-assigned labels.
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The full resource URL of the external storage location. Currently, only
	// Google Cloud Storage is supported. So input_url should be of the form:
	// `gs://BUCKET_NAME[/NAMESPACE_PATH]/OVERALL_EXPORT_METADATA_FILE`, where
	// `BUCKET_NAME` is the name of the Cloud Storage bucket, `NAMESPACE_PATH` is
	// an optional Cloud Storage namespace path (this is not a Cloud Datastore
	// namespace), and `OVERALL_EXPORT_METADATA_FILE` is the metadata file written
	// by the ExportEntities operation. For more information about Cloud Storage
	// namespace paths, see
	// [Object name
	// considerations](https://cloud.google.com/storage/docs/naming#object-considerations).
	//
	// For more information, see
	// [google.datastore.admin.v1beta1.ExportEntitiesResponse.output_url][google.datastore.admin.v1beta1.ExportEntitiesResponse.output_url].
	InputUrl string `protobuf:"bytes,3,opt,name=input_url,json=inputUrl,proto3" json:"input_url,omitempty"`
	// Optionally specify which kinds/namespaces are to be imported. If provided,
	// the list must be a subset of the EntityFilter used in creating the export,
	// otherwise a FAILED_PRECONDITION error will be returned. If no filter is
	// specified then all entities from the export are imported.
	EntityFilter         *EntityFilter `protobuf:"bytes,4,opt,name=entity_filter,json=entityFilter,proto3" json:"entity_filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ImportEntitiesRequest) Reset()         { *m = ImportEntitiesRequest{} }
func (m *ImportEntitiesRequest) String() string { return proto.CompactTextString(m) }
func (*ImportEntitiesRequest) ProtoMessage()    {}
func (*ImportEntitiesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b63216f38706ce20, []int{3}
}

func (m *ImportEntitiesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportEntitiesRequest.Unmarshal(m, b)
}
func (m *ImportEntitiesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportEntitiesRequest.Marshal(b, m, deterministic)
}
func (m *ImportEntitiesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportEntitiesRequest.Merge(m, src)
}
func (m *ImportEntitiesRequest) XXX_Size() int {
	return xxx_messageInfo_ImportEntitiesRequest.Size(m)
}
func (m *ImportEntitiesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportEntitiesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ImportEntitiesRequest proto.InternalMessageInfo

func (m *ImportEntitiesRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *ImportEntitiesRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ImportEntitiesRequest) GetInputUrl() string {
	if m != nil {
		return m.InputUrl
	}
	return ""
}

func (m *ImportEntitiesRequest) GetEntityFilter() *EntityFilter {
	if m != nil {
		return m.EntityFilter
	}
	return nil
}

// The response for
// [google.datastore.admin.v1beta1.DatastoreAdmin.ExportEntities][google.datastore.admin.v1beta1.DatastoreAdmin.ExportEntities].
type ExportEntitiesResponse struct {
	// Location of the output metadata file. This can be used to begin an import
	// into Cloud Datastore (this project or another project). See
	// [google.datastore.admin.v1beta1.ImportEntitiesRequest.input_url][google.datastore.admin.v1beta1.ImportEntitiesRequest.input_url].
	// Only present if the operation completed successfully.
	OutputUrl            string   `protobuf:"bytes,1,opt,name=output_url,json=outputUrl,proto3" json:"output_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportEntitiesResponse) Reset()         { *m = ExportEntitiesResponse{} }
func (m *ExportEntitiesResponse) String() string { return proto.CompactTextString(m) }
func (*ExportEntitiesResponse) ProtoMessage()    {}
func (*ExportEntitiesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b63216f38706ce20, []int{4}
}

func (m *ExportEntitiesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportEntitiesResponse.Unmarshal(m, b)
}
func (m *ExportEntitiesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportEntitiesResponse.Marshal(b, m, deterministic)
}
func (m *ExportEntitiesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportEntitiesResponse.Merge(m, src)
}
func (m *ExportEntitiesResponse) XXX_Size() int {
	return xxx_messageInfo_ExportEntitiesResponse.Size(m)
}
func (m *ExportEntitiesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportEntitiesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportEntitiesResponse proto.InternalMessageInfo

func (m *ExportEntitiesResponse) GetOutputUrl() string {
	if m != nil {
		return m.OutputUrl
	}
	return ""
}

// Metadata for ExportEntities operations.
type ExportEntitiesMetadata struct {
	// Metadata common to all Datastore Admin operations.
	Common *CommonMetadata `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	// An estimate of the number of entities processed.
	ProgressEntities *Progress `protobuf:"bytes,2,opt,name=progress_entities,json=progressEntities,proto3" json:"progress_entities,omitempty"`
	// An estimate of the number of bytes processed.
	ProgressBytes *Progress `protobuf:"bytes,3,opt,name=progress_bytes,json=progressBytes,proto3" json:"progress_bytes,omitempty"`
	// Description of which entities are being exported.
	EntityFilter *EntityFilter `protobuf:"bytes,4,opt,name=entity_filter,json=entityFilter,proto3" json:"entity_filter,omitempty"`
	// Location for the export metadata and data files. This will be the same
	// value as the
	// [google.datastore.admin.v1beta1.ExportEntitiesRequest.output_url_prefix][google.datastore.admin.v1beta1.ExportEntitiesRequest.output_url_prefix]
	// field. The final output location is provided in
	// [google.datastore.admin.v1beta1.ExportEntitiesResponse.output_url][google.datastore.admin.v1beta1.ExportEntitiesResponse.output_url].
	OutputUrlPrefix      string   `protobuf:"bytes,5,opt,name=output_url_prefix,json=outputUrlPrefix,proto3" json:"output_url_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportEntitiesMetadata) Reset()         { *m = ExportEntitiesMetadata{} }
func (m *ExportEntitiesMetadata) String() string { return proto.CompactTextString(m) }
func (*ExportEntitiesMetadata) ProtoMessage()    {}
func (*ExportEntitiesMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_b63216f38706ce20, []int{5}
}

func (m *ExportEntitiesMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportEntitiesMetadata.Unmarshal(m, b)
}
func (m *ExportEntitiesMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportEntitiesMetadata.Marshal(b, m, deterministic)
}
func (m *ExportEntitiesMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportEntitiesMetadata.Merge(m, src)
}
func (m *ExportEntitiesMetadata) XXX_Size() int {
	return xxx_messageInfo_ExportEntitiesMetadata.Size(m)
}
func (m *ExportEntitiesMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportEntitiesMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ExportEntitiesMetadata proto.InternalMessageInfo

func (m *ExportEntitiesMetadata) GetCommon() *CommonMetadata {
	if m != nil {
		return m.Common
	}
	return nil
}

func (m *ExportEntitiesMetadata) GetProgressEntities() *Progress {
	if m != nil {
		return m.ProgressEntities
	}
	return nil
}

func (m *ExportEntitiesMetadata) GetProgressBytes() *Progress {
	if m != nil {
		return m.ProgressBytes
	}
	return nil
}

func (m *ExportEntitiesMetadata) GetEntityFilter() *EntityFilter {
	if m != nil {
		return m.EntityFilter
	}
	return nil
}

func (m *ExportEntitiesMetadata) GetOutputUrlPrefix() string {
	if m != nil {
		return m.OutputUrlPrefix
	}
	return ""
}

// Metadata for ImportEntities operations.
type ImportEntitiesMetadata struct {
	// Metadata common to all Datastore Admin operations.
	Common *CommonMetadata `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	// An estimate of the number of entities processed.
	ProgressEntities *Progress `protobuf:"bytes,2,opt,name=progress_entities,json=progressEntities,proto3" json:"progress_entities,omitempty"`
	// An estimate of the number of bytes processed.
	ProgressBytes *Progress `protobuf:"bytes,3,opt,name=progress_bytes,json=progressBytes,proto3" json:"progress_bytes,omitempty"`
	// Description of which entities are being imported.
	EntityFilter *EntityFilter `protobuf:"bytes,4,opt,name=entity_filter,json=entityFilter,proto3" json:"entity_filter,omitempty"`
	// The location of the import metadata file. This will be the same value as
	// the
	// [google.datastore.admin.v1beta1.ExportEntitiesResponse.output_url][google.datastore.admin.v1beta1.ExportEntitiesResponse.output_url]
	// field.
	InputUrl             string   `protobuf:"bytes,5,opt,name=input_url,json=inputUrl,proto3" json:"input_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImportEntitiesMetadata) Reset()         { *m = ImportEntitiesMetadata{} }
func (m *ImportEntitiesMetadata) String() string { return proto.CompactTextString(m) }
func (*ImportEntitiesMetadata) ProtoMessage()    {}
func (*ImportEntitiesMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_b63216f38706ce20, []int{6}
}

func (m *ImportEntitiesMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportEntitiesMetadata.Unmarshal(m, b)
}
func (m *ImportEntitiesMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportEntitiesMetadata.Marshal(b, m, deterministic)
}
func (m *ImportEntitiesMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportEntitiesMetadata.Merge(m, src)
}
func (m *ImportEntitiesMetadata) XXX_Size() int {
	return xxx_messageInfo_ImportEntitiesMetadata.Size(m)
}
func (m *ImportEntitiesMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportEntitiesMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ImportEntitiesMetadata proto.InternalMessageInfo

func (m *ImportEntitiesMetadata) GetCommon() *CommonMetadata {
	if m != nil {
		return m.Common
	}
	return nil
}

func (m *ImportEntitiesMetadata) GetProgressEntities() *Progress {
	if m != nil {
		return m.ProgressEntities
	}
	return nil
}

func (m *ImportEntitiesMetadata) GetProgressBytes() *Progress {
	if m != nil {
		return m.ProgressBytes
	}
	return nil
}

func (m *ImportEntitiesMetadata) GetEntityFilter() *EntityFilter {
	if m != nil {
		return m.EntityFilter
	}
	return nil
}

func (m *ImportEntitiesMetadata) GetInputUrl() string {
	if m != nil {
		return m.InputUrl
	}
	return ""
}

// Identifies a subset of entities in a project. This is specified as
// combinations of kinds and namespaces (either or both of which may be all, as
// described in the following examples).
// Example usage:
//
// Entire project:
//   kinds=[], namespace_ids=[]
//
// Kinds Foo and Bar in all namespaces:
//   kinds=['Foo', 'Bar'], namespace_ids=[]
//
// Kinds Foo and Bar only in the default namespace:
//   kinds=['Foo', 'Bar'], namespace_ids=['']
//
// Kinds Foo and Bar in both the default and Baz namespaces:
//   kinds=['Foo', 'Bar'], namespace_ids=['', 'Baz']
//
// The entire Baz namespace:
//   kinds=[], namespace_ids=['Baz']
type EntityFilter struct {
	// If empty, then this represents all kinds.
	Kinds []string `protobuf:"bytes,1,rep,name=kinds,proto3" json:"kinds,omitempty"`
	// An empty list represents all namespaces. This is the preferred
	// usage for projects that don't use namespaces.
	//
	// An empty string element represents the default namespace. This should be
	// used if the project has data in non-default namespaces, but doesn't want to
	// include them.
	// Each namespace in this list must be unique.
	NamespaceIds         []string `protobuf:"bytes,2,rep,name=namespace_ids,json=namespaceIds,proto3" json:"namespace_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityFilter) Reset()         { *m = EntityFilter{} }
func (m *EntityFilter) String() string { return proto.CompactTextString(m) }
func (*EntityFilter) ProtoMessage()    {}
func (*EntityFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_b63216f38706ce20, []int{7}
}

func (m *EntityFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityFilter.Unmarshal(m, b)
}
func (m *EntityFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityFilter.Marshal(b, m, deterministic)
}
func (m *EntityFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityFilter.Merge(m, src)
}
func (m *EntityFilter) XXX_Size() int {
	return xxx_messageInfo_EntityFilter.Size(m)
}
func (m *EntityFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityFilter.DiscardUnknown(m)
}

var xxx_messageInfo_EntityFilter proto.InternalMessageInfo

func (m *EntityFilter) GetKinds() []string {
	if m != nil {
		return m.Kinds
	}
	return nil
}

func (m *EntityFilter) GetNamespaceIds() []string {
	if m != nil {
		return m.NamespaceIds
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.datastore.admin.v1beta1.OperationType", OperationType_name, OperationType_value)
	proto.RegisterEnum("google.datastore.admin.v1beta1.CommonMetadata_State", CommonMetadata_State_name, CommonMetadata_State_value)
	proto.RegisterType((*CommonMetadata)(nil), "google.datastore.admin.v1beta1.CommonMetadata")
	proto.RegisterMapType((map[string]string)(nil), "google.datastore.admin.v1beta1.CommonMetadata.LabelsEntry")
	proto.RegisterType((*Progress)(nil), "google.datastore.admin.v1beta1.Progress")
	proto.RegisterType((*ExportEntitiesRequest)(nil), "google.datastore.admin.v1beta1.ExportEntitiesRequest")
	proto.RegisterMapType((map[string]string)(nil), "google.datastore.admin.v1beta1.ExportEntitiesRequest.LabelsEntry")
	proto.RegisterType((*ImportEntitiesRequest)(nil), "google.datastore.admin.v1beta1.ImportEntitiesRequest")
	proto.RegisterMapType((map[string]string)(nil), "google.datastore.admin.v1beta1.ImportEntitiesRequest.LabelsEntry")
	proto.RegisterType((*ExportEntitiesResponse)(nil), "google.datastore.admin.v1beta1.ExportEntitiesResponse")
	proto.RegisterType((*ExportEntitiesMetadata)(nil), "google.datastore.admin.v1beta1.ExportEntitiesMetadata")
	proto.RegisterType((*ImportEntitiesMetadata)(nil), "google.datastore.admin.v1beta1.ImportEntitiesMetadata")
	proto.RegisterType((*EntityFilter)(nil), "google.datastore.admin.v1beta1.EntityFilter")
}

func init() {
	proto.RegisterFile("google/datastore/admin/v1beta1/datastore_admin.proto", fileDescriptor_b63216f38706ce20)
}

var fileDescriptor_b63216f38706ce20 = []byte{
	// 996 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0x41, 0x8f, 0xdb, 0x44,
	0x14, 0xc6, 0xce, 0x26, 0x6d, 0xde, 0x6e, 0xd2, 0xec, 0x94, 0xad, 0xa2, 0x40, 0xcb, 0xca, 0xa5,
	0xd2, 0x6a, 0x05, 0x0e, 0x1b, 0x5a, 0x41, 0x97, 0x53, 0x36, 0xeb, 0x54, 0x46, 0x69, 0x12, 0x1c,
	0x07, 0x75, 0x7b, 0xb1, 0x9c, 0x78, 0x36, 0x32, 0x6b, 0x7b, 0x8c, 0x3d, 0x29, 0x8d, 0x10, 0x17,
	0x2e, 0x1c, 0x38, 0x72, 0xe1, 0x1f, 0x20, 0xf1, 0x1b, 0xb8, 0x70, 0xe1, 0xc2, 0x91, 0xbf, 0xc0,
	0x8f, 0xe0, 0x88, 0x66, 0x3c, 0x76, 0xe2, 0x25, 0x10, 0xca, 0x16, 0x4e, 0xdc, 0xfc, 0xde, 0xbc,
	0xef, 0x9b, 0x37, 0xdf, 0x9b, 0xf7, 0x3c, 0x70, 0x7f, 0x46, 0xc8, 0xcc, 0xc3, 0x4d, 0xc7, 0xa6,
	0x76, 0x4c, 0x49, 0x84, 0x9b, 0xb6, 0xe3, 0xbb, 0x41, 0xf3, 0xd9, 0xd1, 0x04, 0x53, 0xfb, 0x68,
	0xe9, 0xb7, 0xb8, 0x5f, 0x0d, 0x23, 0x42, 0x09, 0xba, 0x93, 0xa0, 0xd4, 0x6c, 0x55, 0x4d, 0x56,
	0x05, 0xaa, 0xf1, 0xba, 0x60, 0xb5, 0x43, 0xb7, 0x69, 0x07, 0x01, 0xa1, 0x36, 0x75, 0x49, 0x10,
	0x27, 0xe8, 0xc6, 0x5d, 0xb1, 0xea, 0x91, 0x60, 0x16, 0xcd, 0x83, 0xc0, 0x0d, 0x66, 0x4d, 0x12,
	0xe2, 0x28, 0x17, 0xf4, 0x86, 0x08, 0xe2, 0xd6, 0x64, 0x7e, 0xde, 0xa4, 0xae, 0x8f, 0x63, 0x6a,
	0xfb, 0x61, 0x12, 0xa0, 0xfc, 0xb8, 0x05, 0xd5, 0x0e, 0xf1, 0x7d, 0x12, 0x3c, 0xc6, 0xd4, 0x66,
	0x99, 0xa0, 0x87, 0x00, 0x31, 0xb5, 0x23, 0x6a, 0xb1, 0xd8, 0xba, 0xb4, 0x2f, 0x1d, 0x6c, 0xb7,
	0x1a, 0xaa, 0xc8, 0x35, 0x25, 0x52, 0xcd, 0x94, 0xc8, 0x28, 0xf3, 0x68, 0x66, 0xa3, 0x07, 0x70,
	0x1d, 0x07, 0x4e, 0x02, 0x94, 0x37, 0x02, 0xaf, 0xe1, 0xc0, 0xe1, 0x30, 0x13, 0xaa, 0x59, 0xe6,
	0x16, 0x5d, 0x84, 0xb8, 0x5e, 0xd8, 0x97, 0x0e, 0xaa, 0xad, 0xb7, 0xd5, 0xbf, 0x56, 0x48, 0x1d,
	0xa4, 0x28, 0x73, 0x11, 0x62, 0xa3, 0x42, 0x56, 0x4d, 0x64, 0x40, 0xc9, 0xb3, 0x27, 0xd8, 0x8b,
	0xeb, 0x5b, 0xfb, 0x85, 0x83, 0xed, 0xd6, 0xf1, 0x26, 0xb6, 0xbc, 0x0e, 0x6a, 0x8f, 0x83, 0xb5,
	0x80, 0x46, 0x0b, 0x43, 0x30, 0xa1, 0x0f, 0xa1, 0x18, 0x53, 0x9b, 0xe2, 0x7a, 0x91, 0x27, 0x78,
	0xff, 0x05, 0x29, 0x47, 0x0c, 0x6b, 0x24, 0x14, 0x8d, 0x87, 0xb0, 0xbd, 0xb2, 0x05, 0xaa, 0x41,
	0xe1, 0x02, 0x2f, 0xb8, 0xde, 0x65, 0x83, 0x7d, 0xa2, 0x57, 0xa1, 0xf8, 0xcc, 0xf6, 0xe6, 0x89,
	0x94, 0x65, 0x23, 0x31, 0x8e, 0xe5, 0xf7, 0x25, 0xe5, 0x6b, 0x09, 0x8a, 0x9c, 0x0b, 0xed, 0xc1,
	0xee, 0xc8, 0x6c, 0x9b, 0x9a, 0x35, 0xee, 0x8f, 0x86, 0x5a, 0x47, 0xef, 0xea, 0xda, 0x69, 0xed,
	0x15, 0x54, 0x83, 0x1d, 0xbd, 0xaf, 0x9b, 0x7a, 0xbb, 0xa7, 0x3f, 0xd5, 0xfb, 0x8f, 0x6a, 0x12,
	0xaa, 0x02, 0x0c, 0x8d, 0x41, 0x47, 0x1b, 0x8d, 0x98, 0x2d, 0x33, 0xbb, 0xd3, 0xee, 0x77, 0xb4,
	0x5e, 0x8f, 0xd9, 0x05, 0x66, 0x77, 0xf5, 0x7e, 0x1a, 0xbf, 0xc5, 0xec, 0xd1, 0xb8, 0xc3, 0xe2,
	0xbb, 0xe3, 0x5e, 0xad, 0x88, 0x00, 0x4a, 0xdd, 0xb6, 0xde, 0xd3, 0x4e, 0x6b, 0x25, 0x54, 0x81,
	0xb2, 0xc0, 0x6a, 0xa7, 0xb5, 0x6b, 0xca, 0x13, 0xb8, 0x3e, 0x8c, 0xc8, 0x2c, 0xc2, 0x71, 0x8c,
	0xee, 0x41, 0xf5, 0x33, 0x12, 0x5d, 0x58, 0x53, 0xe2, 0x87, 0x1e, 0xa6, 0xd8, 0xe1, 0x07, 0x2a,
	0x18, 0x15, 0xe6, 0xed, 0xa4, 0xce, 0x2c, 0x0c, 0xc7, 0xd4, 0xf5, 0x6d, 0x16, 0x26, 0x2f, 0xc3,
	0xb4, 0xd4, 0xa9, 0xfc, 0x2c, 0xc3, 0x9e, 0xf6, 0x3c, 0x24, 0x11, 0xd5, 0x02, 0xea, 0x52, 0x17,
	0xc7, 0x06, 0xfe, 0x74, 0x8e, 0x63, 0x8a, 0x6e, 0x03, 0x84, 0x11, 0xf9, 0x04, 0x4f, 0xa9, 0xe5,
	0x3a, 0x42, 0xb4, 0xb2, 0xf0, 0xe8, 0x0e, 0x3a, 0xcb, 0x6a, 0x2f, 0xf3, 0xda, 0xb7, 0x37, 0x15,
	0x6a, 0xed, 0x2e, 0x6b, 0xaf, 0xc0, 0x47, 0x50, 0xc1, 0x2c, 0x6c, 0x61, 0x9d, 0xbb, 0x1e, 0xc5,
	0x11, 0xbf, 0xab, 0xdb, 0xad, 0xb7, 0x36, 0xee, 0xc0, 0x41, 0x5d, 0x8e, 0x31, 0x76, 0xf0, 0x8a,
	0x85, 0x0e, 0x61, 0x97, 0xcc, 0x69, 0x38, 0xa7, 0xd6, 0x3c, 0xf2, 0xac, 0x30, 0xc2, 0xe7, 0xee,
	0xf3, 0xfa, 0x16, 0x3f, 0xd3, 0x8d, 0x64, 0x61, 0x1c, 0x79, 0x43, 0xee, 0xbe, 0xca, 0xad, 0xf9,
	0x41, 0x86, 0x3d, 0xdd, 0xff, 0x2f, 0xd4, 0x5c, 0xbb, 0xcb, 0x5a, 0x35, 0x5f, 0x83, 0xb2, 0x1b,
	0x88, 0x93, 0x73, 0x25, 0xcb, 0xc6, 0x75, 0xee, 0x18, 0x47, 0xde, 0x1f, 0xa5, 0xde, 0xba, 0xaa,
	0xd4, 0x57, 0x91, 0xef, 0x3d, 0xb8, 0x75, 0xf9, 0x96, 0xc4, 0x21, 0x09, 0x62, 0xcc, 0xe4, 0x5b,
	0xd6, 0x2f, 0x95, 0x2f, 0x2b, 0x9c, 0xf2, 0x55, 0xe1, 0x32, 0x32, 0x9b, 0xb5, 0x5d, 0x28, 0x4d,
	0xf9, 0x88, 0x10, 0x73, 0x56, 0x7d, 0xb1, 0x81, 0x62, 0x08, 0x34, 0x1a, 0xc3, 0x6e, 0x28, 0x5a,
	0xd0, 0xc2, 0x62, 0x13, 0x31, 0x81, 0x0f, 0x36, 0x51, 0xa6, 0xbd, 0x6b, 0xd4, 0x52, 0x8a, 0x34,
	0x4d, 0x34, 0x80, 0x6a, 0x46, 0x3b, 0x59, 0x50, 0x1c, 0x8b, 0xcb, 0xfe, 0xf7, 0x39, 0x2b, 0x29,
	0xfe, 0x84, 0xc1, 0xff, 0x85, 0x8a, 0xae, 0x6f, 0x9e, 0xe2, 0xda, 0xe6, 0x51, 0x7e, 0x93, 0xe1,
	0x56, 0xfe, 0x6e, 0xfe, 0x5f, 0x89, 0x97, 0x57, 0x89, 0x5c, 0x2f, 0x17, 0xf3, 0xbd, 0xac, 0xe8,
	0xb0, 0xb3, 0x0a, 0x65, 0x7d, 0x76, 0xe1, 0x06, 0x4e, 0x5c, 0x97, 0xf6, 0x0b, 0xac, 0xcf, 0xb8,
	0x81, 0xee, 0x42, 0x25, 0xb0, 0x7d, 0x1c, 0x87, 0xf6, 0x14, 0x5b, 0xae, 0x93, 0x0c, 0x9c, 0xb2,
	0xb1, 0x93, 0x39, 0x75, 0x27, 0x3e, 0x3c, 0x83, 0x4a, 0xee, 0xc7, 0x8f, 0xee, 0x40, 0x63, 0x30,
	0xd4, 0x8c, 0xb6, 0xa9, 0x0f, 0xfa, 0x96, 0x79, 0x36, 0xbc, 0xfc, 0x37, 0xbc, 0x09, 0x37, 0xb4,
	0x27, 0xc3, 0x81, 0x61, 0x5a, 0x5a, 0xdf, 0xd4, 0x4d, 0x5d, 0x1b, 0xd5, 0x24, 0xe6, 0xd4, 0x1f,
	0xe7, 0x9d, 0x72, 0xeb, 0x27, 0x19, 0xaa, 0xa7, 0xe9, 0xc9, 0xdb, 0xec, 0xe0, 0xe8, 0x5b, 0x09,
	0xaa, 0xf9, 0xee, 0x45, 0x0f, 0xfe, 0xd1, 0xdf, 0xa4, 0x71, 0x3b, 0x85, 0xad, 0x3c, 0xd9, 0x96,
	0x4f, 0x18, 0xe5, 0x9d, 0x2f, 0x7f, 0xf9, 0xf5, 0x1b, 0xf9, 0x50, 0xb9, 0x97, 0x3d, 0x1b, 0xc5,
	0x04, 0x8e, 0x9b, 0x9f, 0x2f, 0xa7, 0xf3, 0x17, 0xc7, 0x98, 0x93, 0x1f, 0x4b, 0x87, 0x3c, 0xb5,
	0xfc, 0x75, 0xde, 0x9c, 0xda, 0xda, 0xd1, 0xfc, 0xb2, 0x52, 0x73, 0x7d, 0x91, 0xda, 0xc9, 0x77,
	0x12, 0x28, 0x53, 0xe2, 0x6f, 0xc8, 0xe6, 0xe4, 0x66, 0x5e, 0xec, 0x21, 0x7b, 0x24, 0x0e, 0xa5,
	0xa7, 0x1d, 0x01, 0x9b, 0x11, 0xcf, 0x0e, 0x66, 0x2a, 0x89, 0x66, 0xcd, 0x19, 0x0e, 0xf8, 0x13,
	0xb2, 0x99, 0x2c, 0xd9, 0xa1, 0x1b, 0xff, 0xd9, 0x73, 0xfb, 0x03, 0x6e, 0x7d, 0x2f, 0xbf, 0xf9,
	0x28, 0x61, 0xe9, 0x78, 0x64, 0xee, 0xa8, 0xd9, 0x4e, 0x2a, 0xdf, 0x4a, 0xfd, 0xf8, 0xe8, 0x84,
	0x05, 0x4f, 0x4a, 0x9c, 0xf6, 0xdd, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x77, 0x71, 0x2d, 0x88,
	0xc4, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DatastoreAdminClient is the client API for DatastoreAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatastoreAdminClient interface {
	// Exports a copy of all or a subset of entities from Google Cloud Datastore
	// to another storage system, such as Google Cloud Storage. Recent updates to
	// entities may not be reflected in the export. The export occurs in the
	// background and its progress can be monitored and managed via the
	// Operation resource that is created. The output of an export may only be
	// used once the associated operation is done. If an export operation is
	// cancelled before completion it may leave partial data behind in Google
	// Cloud Storage.
	ExportEntities(ctx context.Context, in *ExportEntitiesRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Imports entities into Google Cloud Datastore. Existing entities with the
	// same key are overwritten. The import occurs in the background and its
	// progress can be monitored and managed via the Operation resource that is
	// created. If an ImportEntities operation is cancelled, it is possible
	// that a subset of the data has already been imported to Cloud Datastore.
	ImportEntities(ctx context.Context, in *ImportEntitiesRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
}

type datastoreAdminClient struct {
	cc *grpc.ClientConn
}

func NewDatastoreAdminClient(cc *grpc.ClientConn) DatastoreAdminClient {
	return &datastoreAdminClient{cc}
}

func (c *datastoreAdminClient) ExportEntities(ctx context.Context, in *ExportEntitiesRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.datastore.admin.v1beta1.DatastoreAdmin/ExportEntities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datastoreAdminClient) ImportEntities(ctx context.Context, in *ImportEntitiesRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.datastore.admin.v1beta1.DatastoreAdmin/ImportEntities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatastoreAdminServer is the server API for DatastoreAdmin service.
type DatastoreAdminServer interface {
	// Exports a copy of all or a subset of entities from Google Cloud Datastore
	// to another storage system, such as Google Cloud Storage. Recent updates to
	// entities may not be reflected in the export. The export occurs in the
	// background and its progress can be monitored and managed via the
	// Operation resource that is created. The output of an export may only be
	// used once the associated operation is done. If an export operation is
	// cancelled before completion it may leave partial data behind in Google
	// Cloud Storage.
	ExportEntities(context.Context, *ExportEntitiesRequest) (*longrunning.Operation, error)
	// Imports entities into Google Cloud Datastore. Existing entities with the
	// same key are overwritten. The import occurs in the background and its
	// progress can be monitored and managed via the Operation resource that is
	// created. If an ImportEntities operation is cancelled, it is possible
	// that a subset of the data has already been imported to Cloud Datastore.
	ImportEntities(context.Context, *ImportEntitiesRequest) (*longrunning.Operation, error)
}

// UnimplementedDatastoreAdminServer can be embedded to have forward compatible implementations.
type UnimplementedDatastoreAdminServer struct {
}

func (*UnimplementedDatastoreAdminServer) ExportEntities(ctx context.Context, req *ExportEntitiesRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportEntities not implemented")
}
func (*UnimplementedDatastoreAdminServer) ImportEntities(ctx context.Context, req *ImportEntitiesRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportEntities not implemented")
}

func RegisterDatastoreAdminServer(s *grpc.Server, srv DatastoreAdminServer) {
	s.RegisterService(&_DatastoreAdmin_serviceDesc, srv)
}

func _DatastoreAdmin_ExportEntities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportEntitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatastoreAdminServer).ExportEntities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.datastore.admin.v1beta1.DatastoreAdmin/ExportEntities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatastoreAdminServer).ExportEntities(ctx, req.(*ExportEntitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatastoreAdmin_ImportEntities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportEntitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatastoreAdminServer).ImportEntities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.datastore.admin.v1beta1.DatastoreAdmin/ImportEntities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatastoreAdminServer).ImportEntities(ctx, req.(*ImportEntitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DatastoreAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.datastore.admin.v1beta1.DatastoreAdmin",
	HandlerType: (*DatastoreAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExportEntities",
			Handler:    _DatastoreAdmin_ExportEntities_Handler,
		},
		{
			MethodName: "ImportEntities",
			Handler:    _DatastoreAdmin_ImportEntities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/datastore/admin/v1beta1/datastore_admin.proto",
}
