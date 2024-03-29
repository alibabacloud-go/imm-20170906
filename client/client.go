// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CompareImageFacesRequest struct {
	FaceIdA   *string `json:"FaceIdA,omitempty" xml:"FaceIdA,omitempty"`
	FaceIdB   *string `json:"FaceIdB,omitempty" xml:"FaceIdB,omitempty"`
	ImageUriA *string `json:"ImageUriA,omitempty" xml:"ImageUriA,omitempty"`
	ImageUriB *string `json:"ImageUriB,omitempty" xml:"ImageUriB,omitempty"`
	Project   *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId     *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s CompareImageFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesRequest) GoString() string {
	return s.String()
}

func (s *CompareImageFacesRequest) SetFaceIdA(v string) *CompareImageFacesRequest {
	s.FaceIdA = &v
	return s
}

func (s *CompareImageFacesRequest) SetFaceIdB(v string) *CompareImageFacesRequest {
	s.FaceIdB = &v
	return s
}

func (s *CompareImageFacesRequest) SetImageUriA(v string) *CompareImageFacesRequest {
	s.ImageUriA = &v
	return s
}

func (s *CompareImageFacesRequest) SetImageUriB(v string) *CompareImageFacesRequest {
	s.ImageUriB = &v
	return s
}

func (s *CompareImageFacesRequest) SetProject(v string) *CompareImageFacesRequest {
	s.Project = &v
	return s
}

func (s *CompareImageFacesRequest) SetSetId(v string) *CompareImageFacesRequest {
	s.SetId = &v
	return s
}

type CompareImageFacesResponseBody struct {
	FaceA      *CompareImageFacesResponseBodyFaceA `json:"FaceA,omitempty" xml:"FaceA,omitempty" type:"Struct"`
	FaceB      *CompareImageFacesResponseBodyFaceB `json:"FaceB,omitempty" xml:"FaceB,omitempty" type:"Struct"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId      *string                             `json:"SetId,omitempty" xml:"SetId,omitempty"`
	Similarity *float32                            `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
}

func (s CompareImageFacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesResponseBody) GoString() string {
	return s.String()
}

func (s *CompareImageFacesResponseBody) SetFaceA(v *CompareImageFacesResponseBodyFaceA) *CompareImageFacesResponseBody {
	s.FaceA = v
	return s
}

func (s *CompareImageFacesResponseBody) SetFaceB(v *CompareImageFacesResponseBodyFaceB) *CompareImageFacesResponseBody {
	s.FaceB = v
	return s
}

func (s *CompareImageFacesResponseBody) SetRequestId(v string) *CompareImageFacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompareImageFacesResponseBody) SetSetId(v string) *CompareImageFacesResponseBody {
	s.SetId = &v
	return s
}

func (s *CompareImageFacesResponseBody) SetSimilarity(v float32) *CompareImageFacesResponseBody {
	s.Similarity = &v
	return s
}

type CompareImageFacesResponseBodyFaceA struct {
	FaceAttributes *CompareImageFacesResponseBodyFaceAFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" type:"Struct"`
	FaceId         *string                                           `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
}

func (s CompareImageFacesResponseBodyFaceA) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesResponseBodyFaceA) GoString() string {
	return s.String()
}

func (s *CompareImageFacesResponseBodyFaceA) SetFaceAttributes(v *CompareImageFacesResponseBodyFaceAFaceAttributes) *CompareImageFacesResponseBodyFaceA {
	s.FaceAttributes = v
	return s
}

func (s *CompareImageFacesResponseBodyFaceA) SetFaceId(v string) *CompareImageFacesResponseBodyFaceA {
	s.FaceId = &v
	return s
}

type CompareImageFacesResponseBodyFaceAFaceAttributes struct {
	FaceBoundary *CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary `json:"FaceBoundary,omitempty" xml:"FaceBoundary,omitempty" type:"Struct"`
}

func (s CompareImageFacesResponseBodyFaceAFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesResponseBodyFaceAFaceAttributes) GoString() string {
	return s.String()
}

func (s *CompareImageFacesResponseBodyFaceAFaceAttributes) SetFaceBoundary(v *CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary) *CompareImageFacesResponseBodyFaceAFaceAttributes {
	s.FaceBoundary = v
	return s
}

type CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary) GoString() string {
	return s.String()
}

func (s *CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary) SetHeight(v int32) *CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary {
	s.Height = &v
	return s
}

func (s *CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary) SetLeft(v int32) *CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary {
	s.Left = &v
	return s
}

func (s *CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary) SetTop(v int32) *CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary {
	s.Top = &v
	return s
}

func (s *CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary) SetWidth(v int32) *CompareImageFacesResponseBodyFaceAFaceAttributesFaceBoundary {
	s.Width = &v
	return s
}

type CompareImageFacesResponseBodyFaceB struct {
	FaceAttributes *CompareImageFacesResponseBodyFaceBFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" type:"Struct"`
	FaceId         *string                                           `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
}

func (s CompareImageFacesResponseBodyFaceB) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesResponseBodyFaceB) GoString() string {
	return s.String()
}

func (s *CompareImageFacesResponseBodyFaceB) SetFaceAttributes(v *CompareImageFacesResponseBodyFaceBFaceAttributes) *CompareImageFacesResponseBodyFaceB {
	s.FaceAttributes = v
	return s
}

func (s *CompareImageFacesResponseBodyFaceB) SetFaceId(v string) *CompareImageFacesResponseBodyFaceB {
	s.FaceId = &v
	return s
}

type CompareImageFacesResponseBodyFaceBFaceAttributes struct {
	FaceBoundary *CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary `json:"FaceBoundary,omitempty" xml:"FaceBoundary,omitempty" type:"Struct"`
}

func (s CompareImageFacesResponseBodyFaceBFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesResponseBodyFaceBFaceAttributes) GoString() string {
	return s.String()
}

func (s *CompareImageFacesResponseBodyFaceBFaceAttributes) SetFaceBoundary(v *CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary) *CompareImageFacesResponseBodyFaceBFaceAttributes {
	s.FaceBoundary = v
	return s
}

type CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary) GoString() string {
	return s.String()
}

func (s *CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary) SetHeight(v int32) *CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary {
	s.Height = &v
	return s
}

func (s *CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary) SetLeft(v int32) *CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary {
	s.Left = &v
	return s
}

func (s *CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary) SetTop(v int32) *CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary {
	s.Top = &v
	return s
}

func (s *CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary) SetWidth(v int32) *CompareImageFacesResponseBodyFaceBFaceAttributesFaceBoundary {
	s.Width = &v
	return s
}

type CompareImageFacesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CompareImageFacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CompareImageFacesResponse) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesResponse) GoString() string {
	return s.String()
}

func (s *CompareImageFacesResponse) SetHeaders(v map[string]*string) *CompareImageFacesResponse {
	s.Headers = v
	return s
}

func (s *CompareImageFacesResponse) SetStatusCode(v int32) *CompareImageFacesResponse {
	s.StatusCode = &v
	return s
}

func (s *CompareImageFacesResponse) SetBody(v *CompareImageFacesResponseBody) *CompareImageFacesResponse {
	s.Body = v
	return s
}

type ConvertOfficeFormatRequest struct {
	EndPage        *int64  `json:"EndPage,omitempty" xml:"EndPage,omitempty"`
	FitToPagesTall *bool   `json:"FitToPagesTall,omitempty" xml:"FitToPagesTall,omitempty"`
	FitToPagesWide *bool   `json:"FitToPagesWide,omitempty" xml:"FitToPagesWide,omitempty"`
	Hidecomments   *bool   `json:"Hidecomments,omitempty" xml:"Hidecomments,omitempty"`
	MaxSheetCol    *int64  `json:"MaxSheetCol,omitempty" xml:"MaxSheetCol,omitempty"`
	MaxSheetCount  *int64  `json:"MaxSheetCount,omitempty" xml:"MaxSheetCount,omitempty"`
	MaxSheetRow    *int64  `json:"MaxSheetRow,omitempty" xml:"MaxSheetRow,omitempty"`
	ModelId        *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	PdfVector      *bool   `json:"PdfVector,omitempty" xml:"PdfVector,omitempty"`
	Project        *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SheetOnePage   *bool   `json:"SheetOnePage,omitempty" xml:"SheetOnePage,omitempty"`
	SrcType        *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	SrcUri         *string `json:"SrcUri,omitempty" xml:"SrcUri,omitempty"`
	StartPage      *int64  `json:"StartPage,omitempty" xml:"StartPage,omitempty"`
	TgtFilePages   *string `json:"TgtFilePages,omitempty" xml:"TgtFilePages,omitempty"`
	TgtFilePrefix  *string `json:"TgtFilePrefix,omitempty" xml:"TgtFilePrefix,omitempty"`
	TgtFileSuffix  *string `json:"TgtFileSuffix,omitempty" xml:"TgtFileSuffix,omitempty"`
	TgtType        *string `json:"TgtType,omitempty" xml:"TgtType,omitempty"`
	TgtUri         *string `json:"TgtUri,omitempty" xml:"TgtUri,omitempty"`
}

func (s ConvertOfficeFormatRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertOfficeFormatRequest) GoString() string {
	return s.String()
}

func (s *ConvertOfficeFormatRequest) SetEndPage(v int64) *ConvertOfficeFormatRequest {
	s.EndPage = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetFitToPagesTall(v bool) *ConvertOfficeFormatRequest {
	s.FitToPagesTall = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetFitToPagesWide(v bool) *ConvertOfficeFormatRequest {
	s.FitToPagesWide = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetHidecomments(v bool) *ConvertOfficeFormatRequest {
	s.Hidecomments = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetMaxSheetCol(v int64) *ConvertOfficeFormatRequest {
	s.MaxSheetCol = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetMaxSheetCount(v int64) *ConvertOfficeFormatRequest {
	s.MaxSheetCount = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetMaxSheetRow(v int64) *ConvertOfficeFormatRequest {
	s.MaxSheetRow = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetModelId(v string) *ConvertOfficeFormatRequest {
	s.ModelId = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetPassword(v string) *ConvertOfficeFormatRequest {
	s.Password = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetPdfVector(v bool) *ConvertOfficeFormatRequest {
	s.PdfVector = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetProject(v string) *ConvertOfficeFormatRequest {
	s.Project = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetSheetOnePage(v bool) *ConvertOfficeFormatRequest {
	s.SheetOnePage = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetSrcType(v string) *ConvertOfficeFormatRequest {
	s.SrcType = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetSrcUri(v string) *ConvertOfficeFormatRequest {
	s.SrcUri = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetStartPage(v int64) *ConvertOfficeFormatRequest {
	s.StartPage = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetTgtFilePages(v string) *ConvertOfficeFormatRequest {
	s.TgtFilePages = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetTgtFilePrefix(v string) *ConvertOfficeFormatRequest {
	s.TgtFilePrefix = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetTgtFileSuffix(v string) *ConvertOfficeFormatRequest {
	s.TgtFileSuffix = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetTgtType(v string) *ConvertOfficeFormatRequest {
	s.TgtType = &v
	return s
}

func (s *ConvertOfficeFormatRequest) SetTgtUri(v string) *ConvertOfficeFormatRequest {
	s.TgtUri = &v
	return s
}

type ConvertOfficeFormatResponseBody struct {
	PageCount *int32  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConvertOfficeFormatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConvertOfficeFormatResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertOfficeFormatResponseBody) SetPageCount(v int32) *ConvertOfficeFormatResponseBody {
	s.PageCount = &v
	return s
}

func (s *ConvertOfficeFormatResponseBody) SetRequestId(v string) *ConvertOfficeFormatResponseBody {
	s.RequestId = &v
	return s
}

type ConvertOfficeFormatResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConvertOfficeFormatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConvertOfficeFormatResponse) String() string {
	return tea.Prettify(s)
}

func (s ConvertOfficeFormatResponse) GoString() string {
	return s.String()
}

func (s *ConvertOfficeFormatResponse) SetHeaders(v map[string]*string) *ConvertOfficeFormatResponse {
	s.Headers = v
	return s
}

func (s *ConvertOfficeFormatResponse) SetStatusCode(v int32) *ConvertOfficeFormatResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertOfficeFormatResponse) SetBody(v *ConvertOfficeFormatResponseBody) *ConvertOfficeFormatResponse {
	s.Body = v
	return s
}

type CreateGrabFrameTaskRequest struct {
	CustomMessage   *string `json:"CustomMessage,omitempty" xml:"CustomMessage,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	TargetList      *string `json:"TargetList,omitempty" xml:"TargetList,omitempty"`
	VideoUri        *string `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
}

func (s CreateGrabFrameTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGrabFrameTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateGrabFrameTaskRequest) SetCustomMessage(v string) *CreateGrabFrameTaskRequest {
	s.CustomMessage = &v
	return s
}

func (s *CreateGrabFrameTaskRequest) SetNotifyEndpoint(v string) *CreateGrabFrameTaskRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *CreateGrabFrameTaskRequest) SetNotifyTopicName(v string) *CreateGrabFrameTaskRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateGrabFrameTaskRequest) SetProject(v string) *CreateGrabFrameTaskRequest {
	s.Project = &v
	return s
}

func (s *CreateGrabFrameTaskRequest) SetTargetList(v string) *CreateGrabFrameTaskRequest {
	s.TargetList = &v
	return s
}

func (s *CreateGrabFrameTaskRequest) SetVideoUri(v string) *CreateGrabFrameTaskRequest {
	s.VideoUri = &v
	return s
}

type CreateGrabFrameTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType  *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateGrabFrameTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGrabFrameTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGrabFrameTaskResponseBody) SetRequestId(v string) *CreateGrabFrameTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGrabFrameTaskResponseBody) SetTaskId(v string) *CreateGrabFrameTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateGrabFrameTaskResponseBody) SetTaskType(v string) *CreateGrabFrameTaskResponseBody {
	s.TaskType = &v
	return s
}

type CreateGrabFrameTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGrabFrameTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGrabFrameTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGrabFrameTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateGrabFrameTaskResponse) SetHeaders(v map[string]*string) *CreateGrabFrameTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateGrabFrameTaskResponse) SetStatusCode(v int32) *CreateGrabFrameTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGrabFrameTaskResponse) SetBody(v *CreateGrabFrameTaskResponseBody) *CreateGrabFrameTaskResponse {
	s.Body = v
	return s
}

type CreateGroupFacesJobRequest struct {
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId           *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s CreateGroupFacesJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupFacesJobRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupFacesJobRequest) SetNotifyEndpoint(v string) *CreateGroupFacesJobRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *CreateGroupFacesJobRequest) SetNotifyTopicName(v string) *CreateGroupFacesJobRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateGroupFacesJobRequest) SetProject(v string) *CreateGroupFacesJobRequest {
	s.Project = &v
	return s
}

func (s *CreateGroupFacesJobRequest) SetSetId(v string) *CreateGroupFacesJobRequest {
	s.SetId = &v
	return s
}

type CreateGroupFacesJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobType   *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId     *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s CreateGroupFacesJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupFacesJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupFacesJobResponseBody) SetJobId(v string) *CreateGroupFacesJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateGroupFacesJobResponseBody) SetJobType(v string) *CreateGroupFacesJobResponseBody {
	s.JobType = &v
	return s
}

func (s *CreateGroupFacesJobResponseBody) SetRequestId(v string) *CreateGroupFacesJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupFacesJobResponseBody) SetSetId(v string) *CreateGroupFacesJobResponseBody {
	s.SetId = &v
	return s
}

type CreateGroupFacesJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGroupFacesJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGroupFacesJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupFacesJobResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupFacesJobResponse) SetHeaders(v map[string]*string) *CreateGroupFacesJobResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupFacesJobResponse) SetStatusCode(v int32) *CreateGroupFacesJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupFacesJobResponse) SetBody(v *CreateGroupFacesJobResponseBody) *CreateGroupFacesJobResponse {
	s.Body = v
	return s
}

type CreateMergeFaceGroupsJobRequest struct {
	CustomMessage   *string `json:"CustomMessage,omitempty" xml:"CustomMessage,omitempty"`
	GroupIdFrom     *string `json:"GroupIdFrom,omitempty" xml:"GroupIdFrom,omitempty"`
	GroupIdTo       *string `json:"GroupIdTo,omitempty" xml:"GroupIdTo,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId           *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s CreateMergeFaceGroupsJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeFaceGroupsJobRequest) GoString() string {
	return s.String()
}

func (s *CreateMergeFaceGroupsJobRequest) SetCustomMessage(v string) *CreateMergeFaceGroupsJobRequest {
	s.CustomMessage = &v
	return s
}

func (s *CreateMergeFaceGroupsJobRequest) SetGroupIdFrom(v string) *CreateMergeFaceGroupsJobRequest {
	s.GroupIdFrom = &v
	return s
}

func (s *CreateMergeFaceGroupsJobRequest) SetGroupIdTo(v string) *CreateMergeFaceGroupsJobRequest {
	s.GroupIdTo = &v
	return s
}

func (s *CreateMergeFaceGroupsJobRequest) SetNotifyEndpoint(v string) *CreateMergeFaceGroupsJobRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *CreateMergeFaceGroupsJobRequest) SetNotifyTopicName(v string) *CreateMergeFaceGroupsJobRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateMergeFaceGroupsJobRequest) SetProject(v string) *CreateMergeFaceGroupsJobRequest {
	s.Project = &v
	return s
}

func (s *CreateMergeFaceGroupsJobRequest) SetSetId(v string) *CreateMergeFaceGroupsJobRequest {
	s.SetId = &v
	return s
}

type CreateMergeFaceGroupsJobResponseBody struct {
	GroupIdFrom *string `json:"GroupIdFrom,omitempty" xml:"GroupIdFrom,omitempty"`
	GroupIdTo   *string `json:"GroupIdTo,omitempty" xml:"GroupIdTo,omitempty"`
	JobId       *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobType     *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId       *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s CreateMergeFaceGroupsJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeFaceGroupsJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMergeFaceGroupsJobResponseBody) SetGroupIdFrom(v string) *CreateMergeFaceGroupsJobResponseBody {
	s.GroupIdFrom = &v
	return s
}

func (s *CreateMergeFaceGroupsJobResponseBody) SetGroupIdTo(v string) *CreateMergeFaceGroupsJobResponseBody {
	s.GroupIdTo = &v
	return s
}

func (s *CreateMergeFaceGroupsJobResponseBody) SetJobId(v string) *CreateMergeFaceGroupsJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateMergeFaceGroupsJobResponseBody) SetJobType(v string) *CreateMergeFaceGroupsJobResponseBody {
	s.JobType = &v
	return s
}

func (s *CreateMergeFaceGroupsJobResponseBody) SetRequestId(v string) *CreateMergeFaceGroupsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMergeFaceGroupsJobResponseBody) SetSetId(v string) *CreateMergeFaceGroupsJobResponseBody {
	s.SetId = &v
	return s
}

type CreateMergeFaceGroupsJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMergeFaceGroupsJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMergeFaceGroupsJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeFaceGroupsJobResponse) GoString() string {
	return s.String()
}

func (s *CreateMergeFaceGroupsJobResponse) SetHeaders(v map[string]*string) *CreateMergeFaceGroupsJobResponse {
	s.Headers = v
	return s
}

func (s *CreateMergeFaceGroupsJobResponse) SetStatusCode(v int32) *CreateMergeFaceGroupsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMergeFaceGroupsJobResponse) SetBody(v *CreateMergeFaceGroupsJobResponseBody) *CreateMergeFaceGroupsJobResponse {
	s.Body = v
	return s
}

type CreateOfficeConversionTaskRequest struct {
	DisplayDpi      *int32  `json:"DisplayDpi,omitempty" xml:"DisplayDpi,omitempty"`
	EndPage         *int64  `json:"EndPage,omitempty" xml:"EndPage,omitempty"`
	FitToPagesTall  *bool   `json:"FitToPagesTall,omitempty" xml:"FitToPagesTall,omitempty"`
	FitToPagesWide  *bool   `json:"FitToPagesWide,omitempty" xml:"FitToPagesWide,omitempty"`
	Hidecomments    *bool   `json:"Hidecomments,omitempty" xml:"Hidecomments,omitempty"`
	IdempotentToken *string `json:"IdempotentToken,omitempty" xml:"IdempotentToken,omitempty"`
	MaxSheetCol     *int64  `json:"MaxSheetCol,omitempty" xml:"MaxSheetCol,omitempty"`
	MaxSheetCount   *int64  `json:"MaxSheetCount,omitempty" xml:"MaxSheetCount,omitempty"`
	MaxSheetRow     *int64  `json:"MaxSheetRow,omitempty" xml:"MaxSheetRow,omitempty"`
	ModelId         *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	PdfVector       *bool   `json:"PdfVector,omitempty" xml:"PdfVector,omitempty"`
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SheetOnePage    *bool   `json:"SheetOnePage,omitempty" xml:"SheetOnePage,omitempty"`
	SrcType         *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	SrcUri          *string `json:"SrcUri,omitempty" xml:"SrcUri,omitempty"`
	StartPage       *int64  `json:"StartPage,omitempty" xml:"StartPage,omitempty"`
	TgtFilePages    *string `json:"TgtFilePages,omitempty" xml:"TgtFilePages,omitempty"`
	TgtFilePrefix   *string `json:"TgtFilePrefix,omitempty" xml:"TgtFilePrefix,omitempty"`
	TgtFileSuffix   *string `json:"TgtFileSuffix,omitempty" xml:"TgtFileSuffix,omitempty"`
	TgtType         *string `json:"TgtType,omitempty" xml:"TgtType,omitempty"`
	TgtUri          *string `json:"TgtUri,omitempty" xml:"TgtUri,omitempty"`
	UserData        *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateOfficeConversionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOfficeConversionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateOfficeConversionTaskRequest) SetDisplayDpi(v int32) *CreateOfficeConversionTaskRequest {
	s.DisplayDpi = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetEndPage(v int64) *CreateOfficeConversionTaskRequest {
	s.EndPage = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetFitToPagesTall(v bool) *CreateOfficeConversionTaskRequest {
	s.FitToPagesTall = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetFitToPagesWide(v bool) *CreateOfficeConversionTaskRequest {
	s.FitToPagesWide = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetHidecomments(v bool) *CreateOfficeConversionTaskRequest {
	s.Hidecomments = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetIdempotentToken(v string) *CreateOfficeConversionTaskRequest {
	s.IdempotentToken = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetMaxSheetCol(v int64) *CreateOfficeConversionTaskRequest {
	s.MaxSheetCol = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetMaxSheetCount(v int64) *CreateOfficeConversionTaskRequest {
	s.MaxSheetCount = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetMaxSheetRow(v int64) *CreateOfficeConversionTaskRequest {
	s.MaxSheetRow = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetModelId(v string) *CreateOfficeConversionTaskRequest {
	s.ModelId = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetNotifyEndpoint(v string) *CreateOfficeConversionTaskRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetNotifyTopicName(v string) *CreateOfficeConversionTaskRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetPassword(v string) *CreateOfficeConversionTaskRequest {
	s.Password = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetPdfVector(v bool) *CreateOfficeConversionTaskRequest {
	s.PdfVector = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetProject(v string) *CreateOfficeConversionTaskRequest {
	s.Project = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetSheetOnePage(v bool) *CreateOfficeConversionTaskRequest {
	s.SheetOnePage = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetSrcType(v string) *CreateOfficeConversionTaskRequest {
	s.SrcType = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetSrcUri(v string) *CreateOfficeConversionTaskRequest {
	s.SrcUri = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetStartPage(v int64) *CreateOfficeConversionTaskRequest {
	s.StartPage = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetTgtFilePages(v string) *CreateOfficeConversionTaskRequest {
	s.TgtFilePages = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetTgtFilePrefix(v string) *CreateOfficeConversionTaskRequest {
	s.TgtFilePrefix = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetTgtFileSuffix(v string) *CreateOfficeConversionTaskRequest {
	s.TgtFileSuffix = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetTgtType(v string) *CreateOfficeConversionTaskRequest {
	s.TgtType = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetTgtUri(v string) *CreateOfficeConversionTaskRequest {
	s.TgtUri = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetUserData(v string) *CreateOfficeConversionTaskRequest {
	s.UserData = &v
	return s
}

type CreateOfficeConversionTaskResponseBody struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Percent    *int32  `json:"Percent,omitempty" xml:"Percent,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TgtLoc     *string `json:"TgtLoc,omitempty" xml:"TgtLoc,omitempty"`
}

func (s CreateOfficeConversionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOfficeConversionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOfficeConversionTaskResponseBody) SetCreateTime(v string) *CreateOfficeConversionTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateOfficeConversionTaskResponseBody) SetPercent(v int32) *CreateOfficeConversionTaskResponseBody {
	s.Percent = &v
	return s
}

func (s *CreateOfficeConversionTaskResponseBody) SetRequestId(v string) *CreateOfficeConversionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOfficeConversionTaskResponseBody) SetStatus(v string) *CreateOfficeConversionTaskResponseBody {
	s.Status = &v
	return s
}

func (s *CreateOfficeConversionTaskResponseBody) SetTaskId(v string) *CreateOfficeConversionTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateOfficeConversionTaskResponseBody) SetTgtLoc(v string) *CreateOfficeConversionTaskResponseBody {
	s.TgtLoc = &v
	return s
}

type CreateOfficeConversionTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateOfficeConversionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOfficeConversionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOfficeConversionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateOfficeConversionTaskResponse) SetHeaders(v map[string]*string) *CreateOfficeConversionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateOfficeConversionTaskResponse) SetStatusCode(v int32) *CreateOfficeConversionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOfficeConversionTaskResponse) SetBody(v *CreateOfficeConversionTaskResponseBody) *CreateOfficeConversionTaskResponse {
	s.Body = v
	return s
}

type CreateSetRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId   *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	SetName *string `json:"SetName,omitempty" xml:"SetName,omitempty"`
}

func (s CreateSetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSetRequest) GoString() string {
	return s.String()
}

func (s *CreateSetRequest) SetProject(v string) *CreateSetRequest {
	s.Project = &v
	return s
}

func (s *CreateSetRequest) SetSetId(v string) *CreateSetRequest {
	s.SetId = &v
	return s
}

func (s *CreateSetRequest) SetSetName(v string) *CreateSetRequest {
	s.SetName = &v
	return s
}

type CreateSetResponseBody struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FaceCount   *int32  `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	ImageCount  *int32  `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	ModifyTime  *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId       *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	SetName     *string `json:"SetName,omitempty" xml:"SetName,omitempty"`
	VideoCount  *int32  `json:"VideoCount,omitempty" xml:"VideoCount,omitempty"`
	VideoLength *int32  `json:"VideoLength,omitempty" xml:"VideoLength,omitempty"`
}

func (s CreateSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSetResponseBody) SetCreateTime(v string) *CreateSetResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateSetResponseBody) SetFaceCount(v int32) *CreateSetResponseBody {
	s.FaceCount = &v
	return s
}

func (s *CreateSetResponseBody) SetImageCount(v int32) *CreateSetResponseBody {
	s.ImageCount = &v
	return s
}

func (s *CreateSetResponseBody) SetModifyTime(v string) *CreateSetResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *CreateSetResponseBody) SetRequestId(v string) *CreateSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSetResponseBody) SetSetId(v string) *CreateSetResponseBody {
	s.SetId = &v
	return s
}

func (s *CreateSetResponseBody) SetSetName(v string) *CreateSetResponseBody {
	s.SetName = &v
	return s
}

func (s *CreateSetResponseBody) SetVideoCount(v int32) *CreateSetResponseBody {
	s.VideoCount = &v
	return s
}

func (s *CreateSetResponseBody) SetVideoLength(v int32) *CreateSetResponseBody {
	s.VideoLength = &v
	return s
}

type CreateSetResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSetResponse) GoString() string {
	return s.String()
}

func (s *CreateSetResponse) SetHeaders(v map[string]*string) *CreateSetResponse {
	s.Headers = v
	return s
}

func (s *CreateSetResponse) SetStatusCode(v int32) *CreateSetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSetResponse) SetBody(v *CreateSetResponseBody) *CreateSetResponse {
	s.Body = v
	return s
}

type CreateVideoCompressTaskRequest struct {
	CustomMessage   *string `json:"CustomMessage,omitempty" xml:"CustomMessage,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	TargetList      *string `json:"TargetList,omitempty" xml:"TargetList,omitempty"`
	TargetSegment   *string `json:"TargetSegment,omitempty" xml:"TargetSegment,omitempty"`
	TargetSubtitle  *string `json:"TargetSubtitle,omitempty" xml:"TargetSubtitle,omitempty"`
	VideoUri        *string `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
}

func (s CreateVideoCompressTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoCompressTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoCompressTaskRequest) SetCustomMessage(v string) *CreateVideoCompressTaskRequest {
	s.CustomMessage = &v
	return s
}

func (s *CreateVideoCompressTaskRequest) SetNotifyEndpoint(v string) *CreateVideoCompressTaskRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *CreateVideoCompressTaskRequest) SetNotifyTopicName(v string) *CreateVideoCompressTaskRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateVideoCompressTaskRequest) SetProject(v string) *CreateVideoCompressTaskRequest {
	s.Project = &v
	return s
}

func (s *CreateVideoCompressTaskRequest) SetTargetList(v string) *CreateVideoCompressTaskRequest {
	s.TargetList = &v
	return s
}

func (s *CreateVideoCompressTaskRequest) SetTargetSegment(v string) *CreateVideoCompressTaskRequest {
	s.TargetSegment = &v
	return s
}

func (s *CreateVideoCompressTaskRequest) SetTargetSubtitle(v string) *CreateVideoCompressTaskRequest {
	s.TargetSubtitle = &v
	return s
}

func (s *CreateVideoCompressTaskRequest) SetVideoUri(v string) *CreateVideoCompressTaskRequest {
	s.VideoUri = &v
	return s
}

type CreateVideoCompressTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType  *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateVideoCompressTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoCompressTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoCompressTaskResponseBody) SetRequestId(v string) *CreateVideoCompressTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVideoCompressTaskResponseBody) SetTaskId(v string) *CreateVideoCompressTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateVideoCompressTaskResponseBody) SetTaskType(v string) *CreateVideoCompressTaskResponseBody {
	s.TaskType = &v
	return s
}

type CreateVideoCompressTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVideoCompressTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVideoCompressTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoCompressTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoCompressTaskResponse) SetHeaders(v map[string]*string) *CreateVideoCompressTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoCompressTaskResponse) SetStatusCode(v int32) *CreateVideoCompressTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVideoCompressTaskResponse) SetBody(v *CreateVideoCompressTaskResponseBody) *CreateVideoCompressTaskResponse {
	s.Body = v
	return s
}

type DecodeBlindWatermarkRequest struct {
	ImageQuality     *int32  `json:"ImageQuality,omitempty" xml:"ImageQuality,omitempty"`
	ImageUri         *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Model            *string `json:"Model,omitempty" xml:"Model,omitempty"`
	OriginalImageUri *string `json:"OriginalImageUri,omitempty" xml:"OriginalImageUri,omitempty"`
	Project          *string `json:"Project,omitempty" xml:"Project,omitempty"`
	TargetUri        *string `json:"TargetUri,omitempty" xml:"TargetUri,omitempty"`
}

func (s DecodeBlindWatermarkRequest) String() string {
	return tea.Prettify(s)
}

func (s DecodeBlindWatermarkRequest) GoString() string {
	return s.String()
}

func (s *DecodeBlindWatermarkRequest) SetImageQuality(v int32) *DecodeBlindWatermarkRequest {
	s.ImageQuality = &v
	return s
}

func (s *DecodeBlindWatermarkRequest) SetImageUri(v string) *DecodeBlindWatermarkRequest {
	s.ImageUri = &v
	return s
}

func (s *DecodeBlindWatermarkRequest) SetModel(v string) *DecodeBlindWatermarkRequest {
	s.Model = &v
	return s
}

func (s *DecodeBlindWatermarkRequest) SetOriginalImageUri(v string) *DecodeBlindWatermarkRequest {
	s.OriginalImageUri = &v
	return s
}

func (s *DecodeBlindWatermarkRequest) SetProject(v string) *DecodeBlindWatermarkRequest {
	s.Project = &v
	return s
}

func (s *DecodeBlindWatermarkRequest) SetTargetUri(v string) *DecodeBlindWatermarkRequest {
	s.TargetUri = &v
	return s
}

type DecodeBlindWatermarkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TargetUri *string `json:"TargetUri,omitempty" xml:"TargetUri,omitempty"`
}

func (s DecodeBlindWatermarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DecodeBlindWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *DecodeBlindWatermarkResponseBody) SetRequestId(v string) *DecodeBlindWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DecodeBlindWatermarkResponseBody) SetTargetUri(v string) *DecodeBlindWatermarkResponseBody {
	s.TargetUri = &v
	return s
}

type DecodeBlindWatermarkResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DecodeBlindWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DecodeBlindWatermarkResponse) String() string {
	return tea.Prettify(s)
}

func (s DecodeBlindWatermarkResponse) GoString() string {
	return s.String()
}

func (s *DecodeBlindWatermarkResponse) SetHeaders(v map[string]*string) *DecodeBlindWatermarkResponse {
	s.Headers = v
	return s
}

func (s *DecodeBlindWatermarkResponse) SetStatusCode(v int32) *DecodeBlindWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *DecodeBlindWatermarkResponse) SetBody(v *DecodeBlindWatermarkResponseBody) *DecodeBlindWatermarkResponse {
	s.Body = v
	return s
}

type DeleteImageRequest struct {
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId    *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s DeleteImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageRequest) SetImageUri(v string) *DeleteImageRequest {
	s.ImageUri = &v
	return s
}

func (s *DeleteImageRequest) SetProject(v string) *DeleteImageRequest {
	s.Project = &v
	return s
}

func (s *DeleteImageRequest) SetSetId(v string) *DeleteImageRequest {
	s.SetId = &v
	return s
}

type DeleteImageResponseBody struct {
	ImageUri  *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId     *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s DeleteImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageResponseBody) SetImageUri(v string) *DeleteImageResponseBody {
	s.ImageUri = &v
	return s
}

func (s *DeleteImageResponseBody) SetRequestId(v string) *DeleteImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteImageResponseBody) SetSetId(v string) *DeleteImageResponseBody {
	s.SetId = &v
	return s
}

type DeleteImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageResponse) SetHeaders(v map[string]*string) *DeleteImageResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageResponse) SetStatusCode(v int32) *DeleteImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteImageResponse) SetBody(v *DeleteImageResponseBody) *DeleteImageResponse {
	s.Body = v
	return s
}

type DeleteOfficeConversionTaskRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteOfficeConversionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteOfficeConversionTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteOfficeConversionTaskRequest) SetProject(v string) *DeleteOfficeConversionTaskRequest {
	s.Project = &v
	return s
}

func (s *DeleteOfficeConversionTaskRequest) SetTaskId(v string) *DeleteOfficeConversionTaskRequest {
	s.TaskId = &v
	return s
}

type DeleteOfficeConversionTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOfficeConversionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteOfficeConversionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOfficeConversionTaskResponseBody) SetRequestId(v string) *DeleteOfficeConversionTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteOfficeConversionTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteOfficeConversionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteOfficeConversionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOfficeConversionTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteOfficeConversionTaskResponse) SetHeaders(v map[string]*string) *DeleteOfficeConversionTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteOfficeConversionTaskResponse) SetStatusCode(v int32) *DeleteOfficeConversionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOfficeConversionTaskResponse) SetBody(v *DeleteOfficeConversionTaskResponseBody) *DeleteOfficeConversionTaskResponse {
	s.Body = v
	return s
}

type DeleteProjectRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) SetProject(v string) *DeleteProjectRequest {
	s.Project = &v
	return s
}

type DeleteProjectResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponseBody) SetRequestId(v string) *DeleteProjectResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponse) SetHeaders(v map[string]*string) *DeleteProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectResponse) SetStatusCode(v int32) *DeleteProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProjectResponse) SetBody(v *DeleteProjectResponseBody) *DeleteProjectResponse {
	s.Body = v
	return s
}

type DeleteSetRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId   *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s DeleteSetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSetRequest) GoString() string {
	return s.String()
}

func (s *DeleteSetRequest) SetProject(v string) *DeleteSetRequest {
	s.Project = &v
	return s
}

func (s *DeleteSetRequest) SetSetId(v string) *DeleteSetRequest {
	s.SetId = &v
	return s
}

type DeleteSetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId     *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s DeleteSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSetResponseBody) SetRequestId(v string) *DeleteSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSetResponseBody) SetSetId(v string) *DeleteSetResponseBody {
	s.SetId = &v
	return s
}

type DeleteSetResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSetResponse) GoString() string {
	return s.String()
}

func (s *DeleteSetResponse) SetHeaders(v map[string]*string) *DeleteSetResponse {
	s.Headers = v
	return s
}

func (s *DeleteSetResponse) SetStatusCode(v int32) *DeleteSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSetResponse) SetBody(v *DeleteSetResponseBody) *DeleteSetResponse {
	s.Body = v
	return s
}

type DeleteVideoRequest struct {
	Project   *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Resources *bool   `json:"Resources,omitempty" xml:"Resources,omitempty"`
	SetId     *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	VideoUri  *string `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
}

func (s DeleteVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoRequest) GoString() string {
	return s.String()
}

func (s *DeleteVideoRequest) SetProject(v string) *DeleteVideoRequest {
	s.Project = &v
	return s
}

func (s *DeleteVideoRequest) SetResources(v bool) *DeleteVideoRequest {
	s.Resources = &v
	return s
}

func (s *DeleteVideoRequest) SetSetId(v string) *DeleteVideoRequest {
	s.SetId = &v
	return s
}

func (s *DeleteVideoRequest) SetVideoUri(v string) *DeleteVideoRequest {
	s.VideoUri = &v
	return s
}

type DeleteVideoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId     *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	VideoUri  *string `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
}

func (s DeleteVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVideoResponseBody) SetRequestId(v string) *DeleteVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVideoResponseBody) SetSetId(v string) *DeleteVideoResponseBody {
	s.SetId = &v
	return s
}

func (s *DeleteVideoResponseBody) SetVideoUri(v string) *DeleteVideoResponseBody {
	s.VideoUri = &v
	return s
}

type DeleteVideoResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVideoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoResponse) GoString() string {
	return s.String()
}

func (s *DeleteVideoResponse) SetHeaders(v map[string]*string) *DeleteVideoResponse {
	s.Headers = v
	return s
}

func (s *DeleteVideoResponse) SetStatusCode(v int32) *DeleteVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVideoResponse) SetBody(v *DeleteVideoResponseBody) *DeleteVideoResponse {
	s.Body = v
	return s
}

type DeleteVideoTaskRequest struct {
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DeleteVideoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteVideoTaskRequest) SetProject(v string) *DeleteVideoTaskRequest {
	s.Project = &v
	return s
}

func (s *DeleteVideoTaskRequest) SetTaskId(v string) *DeleteVideoTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DeleteVideoTaskRequest) SetTaskType(v string) *DeleteVideoTaskRequest {
	s.TaskType = &v
	return s
}

type DeleteVideoTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVideoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVideoTaskResponseBody) SetRequestId(v string) *DeleteVideoTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVideoTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVideoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVideoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteVideoTaskResponse) SetHeaders(v map[string]*string) *DeleteVideoTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteVideoTaskResponse) SetStatusCode(v int32) *DeleteVideoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVideoTaskResponse) SetBody(v *DeleteVideoTaskResponseBody) *DeleteVideoTaskResponse {
	s.Body = v
	return s
}

type DetectImageBodiesRequest struct {
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s DetectImageBodiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageBodiesRequest) GoString() string {
	return s.String()
}

func (s *DetectImageBodiesRequest) SetImageUri(v string) *DetectImageBodiesRequest {
	s.ImageUri = &v
	return s
}

func (s *DetectImageBodiesRequest) SetProject(v string) *DetectImageBodiesRequest {
	s.Project = &v
	return s
}

type DetectImageBodiesResponseBody struct {
	Bodies    []*DetectImageBodiesResponseBodyBodies `json:"Bodies,omitempty" xml:"Bodies,omitempty" type:"Repeated"`
	ImageUri  *string                                `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectImageBodiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectImageBodiesResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageBodiesResponseBody) SetBodies(v []*DetectImageBodiesResponseBodyBodies) *DetectImageBodiesResponseBody {
	s.Bodies = v
	return s
}

func (s *DetectImageBodiesResponseBody) SetImageUri(v string) *DetectImageBodiesResponseBody {
	s.ImageUri = &v
	return s
}

func (s *DetectImageBodiesResponseBody) SetRequestId(v string) *DetectImageBodiesResponseBody {
	s.RequestId = &v
	return s
}

type DetectImageBodiesResponseBodyBodies struct {
	BodyBoundary   *DetectImageBodiesResponseBodyBodiesBodyBoundary `json:"BodyBoundary,omitempty" xml:"BodyBoundary,omitempty" type:"Struct"`
	BodyConfidence *float32                                         `json:"BodyConfidence,omitempty" xml:"BodyConfidence,omitempty"`
}

func (s DetectImageBodiesResponseBodyBodies) String() string {
	return tea.Prettify(s)
}

func (s DetectImageBodiesResponseBodyBodies) GoString() string {
	return s.String()
}

func (s *DetectImageBodiesResponseBodyBodies) SetBodyBoundary(v *DetectImageBodiesResponseBodyBodiesBodyBoundary) *DetectImageBodiesResponseBodyBodies {
	s.BodyBoundary = v
	return s
}

func (s *DetectImageBodiesResponseBodyBodies) SetBodyConfidence(v float32) *DetectImageBodiesResponseBodyBodies {
	s.BodyConfidence = &v
	return s
}

type DetectImageBodiesResponseBodyBodiesBodyBoundary struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectImageBodiesResponseBodyBodiesBodyBoundary) String() string {
	return tea.Prettify(s)
}

func (s DetectImageBodiesResponseBodyBodiesBodyBoundary) GoString() string {
	return s.String()
}

func (s *DetectImageBodiesResponseBodyBodiesBodyBoundary) SetHeight(v int32) *DetectImageBodiesResponseBodyBodiesBodyBoundary {
	s.Height = &v
	return s
}

func (s *DetectImageBodiesResponseBodyBodiesBodyBoundary) SetLeft(v int32) *DetectImageBodiesResponseBodyBodiesBodyBoundary {
	s.Left = &v
	return s
}

func (s *DetectImageBodiesResponseBodyBodiesBodyBoundary) SetTop(v int32) *DetectImageBodiesResponseBodyBodiesBodyBoundary {
	s.Top = &v
	return s
}

func (s *DetectImageBodiesResponseBodyBodiesBodyBoundary) SetWidth(v int32) *DetectImageBodiesResponseBodyBodiesBodyBoundary {
	s.Width = &v
	return s
}

type DetectImageBodiesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectImageBodiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectImageBodiesResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectImageBodiesResponse) GoString() string {
	return s.String()
}

func (s *DetectImageBodiesResponse) SetHeaders(v map[string]*string) *DetectImageBodiesResponse {
	s.Headers = v
	return s
}

func (s *DetectImageBodiesResponse) SetStatusCode(v int32) *DetectImageBodiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectImageBodiesResponse) SetBody(v *DetectImageBodiesResponseBody) *DetectImageBodiesResponse {
	s.Body = v
	return s
}

type DetectImageFacesRequest struct {
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s DetectImageFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageFacesRequest) GoString() string {
	return s.String()
}

func (s *DetectImageFacesRequest) SetImageUri(v string) *DetectImageFacesRequest {
	s.ImageUri = &v
	return s
}

func (s *DetectImageFacesRequest) SetProject(v string) *DetectImageFacesRequest {
	s.Project = &v
	return s
}

type DetectImageFacesResponseBody struct {
	Faces     []*DetectImageFacesResponseBodyFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	ImageUri  *string                              `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectImageFacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectImageFacesResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageFacesResponseBody) SetFaces(v []*DetectImageFacesResponseBodyFaces) *DetectImageFacesResponseBody {
	s.Faces = v
	return s
}

func (s *DetectImageFacesResponseBody) SetImageUri(v string) *DetectImageFacesResponseBody {
	s.ImageUri = &v
	return s
}

func (s *DetectImageFacesResponseBody) SetRequestId(v string) *DetectImageFacesResponseBody {
	s.RequestId = &v
	return s
}

type DetectImageFacesResponseBodyFaces struct {
	Age                  *int32                                           `json:"Age,omitempty" xml:"Age,omitempty"`
	AgeConfidence        *float32                                         `json:"AgeConfidence,omitempty" xml:"AgeConfidence,omitempty"`
	Attractive           *float32                                         `json:"Attractive,omitempty" xml:"Attractive,omitempty"`
	AttractiveConfidence *float32                                         `json:"AttractiveConfidence,omitempty" xml:"AttractiveConfidence,omitempty"`
	Emotion              *string                                          `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	EmotionConfidence    *float32                                         `json:"EmotionConfidence,omitempty" xml:"EmotionConfidence,omitempty"`
	EmotionDetails       *DetectImageFacesResponseBodyFacesEmotionDetails `json:"EmotionDetails,omitempty" xml:"EmotionDetails,omitempty" type:"Struct"`
	FaceAttributes       *DetectImageFacesResponseBodyFacesFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" type:"Struct"`
	FaceConfidence       *float32                                         `json:"FaceConfidence,omitempty" xml:"FaceConfidence,omitempty"`
	FaceId               *string                                          `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	FaceQuality          *float32                                         `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	Gender               *string                                          `json:"Gender,omitempty" xml:"Gender,omitempty"`
	GenderConfidence     *float32                                         `json:"GenderConfidence,omitempty" xml:"GenderConfidence,omitempty"`
}

func (s DetectImageFacesResponseBodyFaces) String() string {
	return tea.Prettify(s)
}

func (s DetectImageFacesResponseBodyFaces) GoString() string {
	return s.String()
}

func (s *DetectImageFacesResponseBodyFaces) SetAge(v int32) *DetectImageFacesResponseBodyFaces {
	s.Age = &v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetAgeConfidence(v float32) *DetectImageFacesResponseBodyFaces {
	s.AgeConfidence = &v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetAttractive(v float32) *DetectImageFacesResponseBodyFaces {
	s.Attractive = &v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetAttractiveConfidence(v float32) *DetectImageFacesResponseBodyFaces {
	s.AttractiveConfidence = &v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetEmotion(v string) *DetectImageFacesResponseBodyFaces {
	s.Emotion = &v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetEmotionConfidence(v float32) *DetectImageFacesResponseBodyFaces {
	s.EmotionConfidence = &v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetEmotionDetails(v *DetectImageFacesResponseBodyFacesEmotionDetails) *DetectImageFacesResponseBodyFaces {
	s.EmotionDetails = v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetFaceAttributes(v *DetectImageFacesResponseBodyFacesFaceAttributes) *DetectImageFacesResponseBodyFaces {
	s.FaceAttributes = v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetFaceConfidence(v float32) *DetectImageFacesResponseBodyFaces {
	s.FaceConfidence = &v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetFaceId(v string) *DetectImageFacesResponseBodyFaces {
	s.FaceId = &v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetFaceQuality(v float32) *DetectImageFacesResponseBodyFaces {
	s.FaceQuality = &v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetGender(v string) *DetectImageFacesResponseBodyFaces {
	s.Gender = &v
	return s
}

func (s *DetectImageFacesResponseBodyFaces) SetGenderConfidence(v float32) *DetectImageFacesResponseBodyFaces {
	s.GenderConfidence = &v
	return s
}

type DetectImageFacesResponseBodyFacesEmotionDetails struct {
	ANGRY     *float32 `json:"ANGRY,omitempty" xml:"ANGRY,omitempty"`
	CALM      *float32 `json:"CALM,omitempty" xml:"CALM,omitempty"`
	DISGUSTED *float32 `json:"DISGUSTED,omitempty" xml:"DISGUSTED,omitempty"`
	HAPPY     *float32 `json:"HAPPY,omitempty" xml:"HAPPY,omitempty"`
	SAD       *float32 `json:"SAD,omitempty" xml:"SAD,omitempty"`
	SCARED    *float32 `json:"SCARED,omitempty" xml:"SCARED,omitempty"`
	SURPRISED *float32 `json:"SURPRISED,omitempty" xml:"SURPRISED,omitempty"`
}

func (s DetectImageFacesResponseBodyFacesEmotionDetails) String() string {
	return tea.Prettify(s)
}

func (s DetectImageFacesResponseBodyFacesEmotionDetails) GoString() string {
	return s.String()
}

func (s *DetectImageFacesResponseBodyFacesEmotionDetails) SetANGRY(v float32) *DetectImageFacesResponseBodyFacesEmotionDetails {
	s.ANGRY = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesEmotionDetails) SetCALM(v float32) *DetectImageFacesResponseBodyFacesEmotionDetails {
	s.CALM = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesEmotionDetails) SetDISGUSTED(v float32) *DetectImageFacesResponseBodyFacesEmotionDetails {
	s.DISGUSTED = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesEmotionDetails) SetHAPPY(v float32) *DetectImageFacesResponseBodyFacesEmotionDetails {
	s.HAPPY = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesEmotionDetails) SetSAD(v float32) *DetectImageFacesResponseBodyFacesEmotionDetails {
	s.SAD = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesEmotionDetails) SetSCARED(v float32) *DetectImageFacesResponseBodyFacesEmotionDetails {
	s.SCARED = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesEmotionDetails) SetSURPRISED(v float32) *DetectImageFacesResponseBodyFacesEmotionDetails {
	s.SURPRISED = &v
	return s
}

type DetectImageFacesResponseBodyFacesFaceAttributes struct {
	Beard             *string                                                      `json:"Beard,omitempty" xml:"Beard,omitempty"`
	BeardConfidence   *float32                                                     `json:"BeardConfidence,omitempty" xml:"BeardConfidence,omitempty"`
	FaceBoundary      *DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary `json:"FaceBoundary,omitempty" xml:"FaceBoundary,omitempty" type:"Struct"`
	Glasses           *string                                                      `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	GlassesConfidence *float32                                                     `json:"GlassesConfidence,omitempty" xml:"GlassesConfidence,omitempty"`
	HeadPose          *DetectImageFacesResponseBodyFacesFaceAttributesHeadPose     `json:"HeadPose,omitempty" xml:"HeadPose,omitempty" type:"Struct"`
	Mask              *string                                                      `json:"Mask,omitempty" xml:"Mask,omitempty"`
	MaskConfidence    *float32                                                     `json:"MaskConfidence,omitempty" xml:"MaskConfidence,omitempty"`
}

func (s DetectImageFacesResponseBodyFacesFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s DetectImageFacesResponseBodyFacesFaceAttributes) GoString() string {
	return s.String()
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributes) SetBeard(v string) *DetectImageFacesResponseBodyFacesFaceAttributes {
	s.Beard = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributes) SetBeardConfidence(v float32) *DetectImageFacesResponseBodyFacesFaceAttributes {
	s.BeardConfidence = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributes) SetFaceBoundary(v *DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary) *DetectImageFacesResponseBodyFacesFaceAttributes {
	s.FaceBoundary = v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributes) SetGlasses(v string) *DetectImageFacesResponseBodyFacesFaceAttributes {
	s.Glasses = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributes) SetGlassesConfidence(v float32) *DetectImageFacesResponseBodyFacesFaceAttributes {
	s.GlassesConfidence = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributes) SetHeadPose(v *DetectImageFacesResponseBodyFacesFaceAttributesHeadPose) *DetectImageFacesResponseBodyFacesFaceAttributes {
	s.HeadPose = v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributes) SetMask(v string) *DetectImageFacesResponseBodyFacesFaceAttributes {
	s.Mask = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributes) SetMaskConfidence(v float32) *DetectImageFacesResponseBodyFacesFaceAttributes {
	s.MaskConfidence = &v
	return s
}

type DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary) String() string {
	return tea.Prettify(s)
}

func (s DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary) GoString() string {
	return s.String()
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary) SetHeight(v int32) *DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary {
	s.Height = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary) SetLeft(v int32) *DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary {
	s.Left = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary) SetTop(v int32) *DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary {
	s.Top = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary) SetWidth(v int32) *DetectImageFacesResponseBodyFacesFaceAttributesFaceBoundary {
	s.Width = &v
	return s
}

type DetectImageFacesResponseBodyFacesFaceAttributesHeadPose struct {
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	Roll  *float32 `json:"Roll,omitempty" xml:"Roll,omitempty"`
	Yaw   *float32 `json:"Yaw,omitempty" xml:"Yaw,omitempty"`
}

func (s DetectImageFacesResponseBodyFacesFaceAttributesHeadPose) String() string {
	return tea.Prettify(s)
}

func (s DetectImageFacesResponseBodyFacesFaceAttributesHeadPose) GoString() string {
	return s.String()
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributesHeadPose) SetPitch(v float32) *DetectImageFacesResponseBodyFacesFaceAttributesHeadPose {
	s.Pitch = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributesHeadPose) SetRoll(v float32) *DetectImageFacesResponseBodyFacesFaceAttributesHeadPose {
	s.Roll = &v
	return s
}

func (s *DetectImageFacesResponseBodyFacesFaceAttributesHeadPose) SetYaw(v float32) *DetectImageFacesResponseBodyFacesFaceAttributesHeadPose {
	s.Yaw = &v
	return s
}

type DetectImageFacesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectImageFacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectImageFacesResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectImageFacesResponse) GoString() string {
	return s.String()
}

func (s *DetectImageFacesResponse) SetHeaders(v map[string]*string) *DetectImageFacesResponse {
	s.Headers = v
	return s
}

func (s *DetectImageFacesResponse) SetStatusCode(v int32) *DetectImageFacesResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectImageFacesResponse) SetBody(v *DetectImageFacesResponseBody) *DetectImageFacesResponse {
	s.Body = v
	return s
}

type DetectImageQRCodesRequest struct {
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s DetectImageQRCodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageQRCodesRequest) GoString() string {
	return s.String()
}

func (s *DetectImageQRCodesRequest) SetImageUri(v string) *DetectImageQRCodesRequest {
	s.ImageUri = &v
	return s
}

func (s *DetectImageQRCodesRequest) SetProject(v string) *DetectImageQRCodesRequest {
	s.Project = &v
	return s
}

type DetectImageQRCodesResponseBody struct {
	ImageUri  *string                                  `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	QRCodes   []*DetectImageQRCodesResponseBodyQRCodes `json:"QRCodes,omitempty" xml:"QRCodes,omitempty" type:"Repeated"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectImageQRCodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectImageQRCodesResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageQRCodesResponseBody) SetImageUri(v string) *DetectImageQRCodesResponseBody {
	s.ImageUri = &v
	return s
}

func (s *DetectImageQRCodesResponseBody) SetQRCodes(v []*DetectImageQRCodesResponseBodyQRCodes) *DetectImageQRCodesResponseBody {
	s.QRCodes = v
	return s
}

func (s *DetectImageQRCodesResponseBody) SetRequestId(v string) *DetectImageQRCodesResponseBody {
	s.RequestId = &v
	return s
}

type DetectImageQRCodesResponseBodyQRCodes struct {
	Content        *string                                              `json:"Content,omitempty" xml:"Content,omitempty"`
	QRCodeBoundary *DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary `json:"QRCodeBoundary,omitempty" xml:"QRCodeBoundary,omitempty" type:"Struct"`
}

func (s DetectImageQRCodesResponseBodyQRCodes) String() string {
	return tea.Prettify(s)
}

func (s DetectImageQRCodesResponseBodyQRCodes) GoString() string {
	return s.String()
}

func (s *DetectImageQRCodesResponseBodyQRCodes) SetContent(v string) *DetectImageQRCodesResponseBodyQRCodes {
	s.Content = &v
	return s
}

func (s *DetectImageQRCodesResponseBodyQRCodes) SetQRCodeBoundary(v *DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary) *DetectImageQRCodesResponseBodyQRCodes {
	s.QRCodeBoundary = v
	return s
}

type DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary) String() string {
	return tea.Prettify(s)
}

func (s DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary) GoString() string {
	return s.String()
}

func (s *DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary) SetHeight(v int32) *DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary {
	s.Height = &v
	return s
}

func (s *DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary) SetLeft(v int32) *DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary {
	s.Left = &v
	return s
}

func (s *DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary) SetTop(v int32) *DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary {
	s.Top = &v
	return s
}

func (s *DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary) SetWidth(v int32) *DetectImageQRCodesResponseBodyQRCodesQRCodeBoundary {
	s.Width = &v
	return s
}

type DetectImageQRCodesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectImageQRCodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectImageQRCodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectImageQRCodesResponse) GoString() string {
	return s.String()
}

func (s *DetectImageQRCodesResponse) SetHeaders(v map[string]*string) *DetectImageQRCodesResponse {
	s.Headers = v
	return s
}

func (s *DetectImageQRCodesResponse) SetStatusCode(v int32) *DetectImageQRCodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectImageQRCodesResponse) SetBody(v *DetectImageQRCodesResponseBody) *DetectImageQRCodesResponse {
	s.Body = v
	return s
}

type DetectImageTagsRequest struct {
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s DetectImageTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageTagsRequest) GoString() string {
	return s.String()
}

func (s *DetectImageTagsRequest) SetImageUri(v string) *DetectImageTagsRequest {
	s.ImageUri = &v
	return s
}

func (s *DetectImageTagsRequest) SetProject(v string) *DetectImageTagsRequest {
	s.Project = &v
	return s
}

type DetectImageTagsResponseBody struct {
	ImageUri  *string                            `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      []*DetectImageTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DetectImageTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectImageTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageTagsResponseBody) SetImageUri(v string) *DetectImageTagsResponseBody {
	s.ImageUri = &v
	return s
}

func (s *DetectImageTagsResponseBody) SetRequestId(v string) *DetectImageTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectImageTagsResponseBody) SetTags(v []*DetectImageTagsResponseBodyTags) *DetectImageTagsResponseBody {
	s.Tags = v
	return s
}

type DetectImageTagsResponseBodyTags struct {
	CentricScore    *float32 `json:"CentricScore,omitempty" xml:"CentricScore,omitempty"`
	ParentTagEnName *string  `json:"ParentTagEnName,omitempty" xml:"ParentTagEnName,omitempty"`
	ParentTagName   *string  `json:"ParentTagName,omitempty" xml:"ParentTagName,omitempty"`
	TagConfidence   *float32 `json:"TagConfidence,omitempty" xml:"TagConfidence,omitempty"`
	TagEnName       *string  `json:"TagEnName,omitempty" xml:"TagEnName,omitempty"`
	TagLevel        *int32   `json:"TagLevel,omitempty" xml:"TagLevel,omitempty"`
	TagName         *string  `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s DetectImageTagsResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DetectImageTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DetectImageTagsResponseBodyTags) SetCentricScore(v float32) *DetectImageTagsResponseBodyTags {
	s.CentricScore = &v
	return s
}

func (s *DetectImageTagsResponseBodyTags) SetParentTagEnName(v string) *DetectImageTagsResponseBodyTags {
	s.ParentTagEnName = &v
	return s
}

func (s *DetectImageTagsResponseBodyTags) SetParentTagName(v string) *DetectImageTagsResponseBodyTags {
	s.ParentTagName = &v
	return s
}

func (s *DetectImageTagsResponseBodyTags) SetTagConfidence(v float32) *DetectImageTagsResponseBodyTags {
	s.TagConfidence = &v
	return s
}

func (s *DetectImageTagsResponseBodyTags) SetTagEnName(v string) *DetectImageTagsResponseBodyTags {
	s.TagEnName = &v
	return s
}

func (s *DetectImageTagsResponseBodyTags) SetTagLevel(v int32) *DetectImageTagsResponseBodyTags {
	s.TagLevel = &v
	return s
}

func (s *DetectImageTagsResponseBodyTags) SetTagName(v string) *DetectImageTagsResponseBodyTags {
	s.TagName = &v
	return s
}

type DetectImageTagsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectImageTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectImageTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectImageTagsResponse) GoString() string {
	return s.String()
}

func (s *DetectImageTagsResponse) SetHeaders(v map[string]*string) *DetectImageTagsResponse {
	s.Headers = v
	return s
}

func (s *DetectImageTagsResponse) SetStatusCode(v int32) *DetectImageTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectImageTagsResponse) SetBody(v *DetectImageTagsResponseBody) *DetectImageTagsResponse {
	s.Body = v
	return s
}

type DetectQRCodesRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SrcUris *string `json:"SrcUris,omitempty" xml:"SrcUris,omitempty"`
}

func (s DetectQRCodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectQRCodesRequest) GoString() string {
	return s.String()
}

func (s *DetectQRCodesRequest) SetProject(v string) *DetectQRCodesRequest {
	s.Project = &v
	return s
}

func (s *DetectQRCodesRequest) SetSrcUris(v string) *DetectQRCodesRequest {
	s.SrcUris = &v
	return s
}

type DetectQRCodesResponseBody struct {
	FailDetails    []*DetectQRCodesResponseBodyFailDetails    `json:"FailDetails,omitempty" xml:"FailDetails,omitempty" type:"Repeated"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessDetails []*DetectQRCodesResponseBodySuccessDetails `json:"SuccessDetails,omitempty" xml:"SuccessDetails,omitempty" type:"Repeated"`
}

func (s DetectQRCodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectQRCodesResponseBody) GoString() string {
	return s.String()
}

func (s *DetectQRCodesResponseBody) SetFailDetails(v []*DetectQRCodesResponseBodyFailDetails) *DetectQRCodesResponseBody {
	s.FailDetails = v
	return s
}

func (s *DetectQRCodesResponseBody) SetRequestId(v string) *DetectQRCodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectQRCodesResponseBody) SetSuccessDetails(v []*DetectQRCodesResponseBodySuccessDetails) *DetectQRCodesResponseBody {
	s.SuccessDetails = v
	return s
}

type DetectQRCodesResponseBodyFailDetails struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	SrcUri       *string `json:"SrcUri,omitempty" xml:"SrcUri,omitempty"`
}

func (s DetectQRCodesResponseBodyFailDetails) String() string {
	return tea.Prettify(s)
}

func (s DetectQRCodesResponseBodyFailDetails) GoString() string {
	return s.String()
}

func (s *DetectQRCodesResponseBodyFailDetails) SetErrorCode(v string) *DetectQRCodesResponseBodyFailDetails {
	s.ErrorCode = &v
	return s
}

func (s *DetectQRCodesResponseBodyFailDetails) SetErrorMessage(v string) *DetectQRCodesResponseBodyFailDetails {
	s.ErrorMessage = &v
	return s
}

func (s *DetectQRCodesResponseBodyFailDetails) SetSrcUri(v string) *DetectQRCodesResponseBodyFailDetails {
	s.SrcUri = &v
	return s
}

type DetectQRCodesResponseBodySuccessDetails struct {
	QRCodes []*DetectQRCodesResponseBodySuccessDetailsQRCodes `json:"QRCodes,omitempty" xml:"QRCodes,omitempty" type:"Repeated"`
	SrcUri  *string                                           `json:"SrcUri,omitempty" xml:"SrcUri,omitempty"`
}

func (s DetectQRCodesResponseBodySuccessDetails) String() string {
	return tea.Prettify(s)
}

func (s DetectQRCodesResponseBodySuccessDetails) GoString() string {
	return s.String()
}

func (s *DetectQRCodesResponseBodySuccessDetails) SetQRCodes(v []*DetectQRCodesResponseBodySuccessDetailsQRCodes) *DetectQRCodesResponseBodySuccessDetails {
	s.QRCodes = v
	return s
}

func (s *DetectQRCodesResponseBodySuccessDetails) SetSrcUri(v string) *DetectQRCodesResponseBodySuccessDetails {
	s.SrcUri = &v
	return s
}

type DetectQRCodesResponseBodySuccessDetailsQRCodes struct {
	Content          *string                                                         `json:"Content,omitempty" xml:"Content,omitempty"`
	QRCodesRectangle *DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle `json:"QRCodesRectangle,omitempty" xml:"QRCodesRectangle,omitempty" type:"Struct"`
}

func (s DetectQRCodesResponseBodySuccessDetailsQRCodes) String() string {
	return tea.Prettify(s)
}

func (s DetectQRCodesResponseBodySuccessDetailsQRCodes) GoString() string {
	return s.String()
}

func (s *DetectQRCodesResponseBodySuccessDetailsQRCodes) SetContent(v string) *DetectQRCodesResponseBodySuccessDetailsQRCodes {
	s.Content = &v
	return s
}

func (s *DetectQRCodesResponseBodySuccessDetailsQRCodes) SetQRCodesRectangle(v *DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle) *DetectQRCodesResponseBodySuccessDetailsQRCodes {
	s.QRCodesRectangle = v
	return s
}

type DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle struct {
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *string `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *string `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle) String() string {
	return tea.Prettify(s)
}

func (s DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle) GoString() string {
	return s.String()
}

func (s *DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle) SetHeight(v string) *DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle {
	s.Height = &v
	return s
}

func (s *DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle) SetLeft(v string) *DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle {
	s.Left = &v
	return s
}

func (s *DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle) SetTop(v string) *DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle {
	s.Top = &v
	return s
}

func (s *DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle) SetWidth(v string) *DetectQRCodesResponseBodySuccessDetailsQRCodesQRCodesRectangle {
	s.Width = &v
	return s
}

type DetectQRCodesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectQRCodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectQRCodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectQRCodesResponse) GoString() string {
	return s.String()
}

func (s *DetectQRCodesResponse) SetHeaders(v map[string]*string) *DetectQRCodesResponse {
	s.Headers = v
	return s
}

func (s *DetectQRCodesResponse) SetStatusCode(v int32) *DetectQRCodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectQRCodesResponse) SetBody(v *DetectQRCodesResponseBody) *DetectQRCodesResponse {
	s.Body = v
	return s
}

type EncodeBlindWatermarkRequest struct {
	Content         *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ImageQuality    *string `json:"ImageQuality,omitempty" xml:"ImageQuality,omitempty"`
	ImageUri        *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Model           *string `json:"Model,omitempty" xml:"Model,omitempty"`
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	TargetImageType *string `json:"TargetImageType,omitempty" xml:"TargetImageType,omitempty"`
	TargetUri       *string `json:"TargetUri,omitempty" xml:"TargetUri,omitempty"`
	WatermarkUri    *string `json:"WatermarkUri,omitempty" xml:"WatermarkUri,omitempty"`
}

func (s EncodeBlindWatermarkRequest) String() string {
	return tea.Prettify(s)
}

func (s EncodeBlindWatermarkRequest) GoString() string {
	return s.String()
}

func (s *EncodeBlindWatermarkRequest) SetContent(v string) *EncodeBlindWatermarkRequest {
	s.Content = &v
	return s
}

func (s *EncodeBlindWatermarkRequest) SetImageQuality(v string) *EncodeBlindWatermarkRequest {
	s.ImageQuality = &v
	return s
}

func (s *EncodeBlindWatermarkRequest) SetImageUri(v string) *EncodeBlindWatermarkRequest {
	s.ImageUri = &v
	return s
}

func (s *EncodeBlindWatermarkRequest) SetModel(v string) *EncodeBlindWatermarkRequest {
	s.Model = &v
	return s
}

func (s *EncodeBlindWatermarkRequest) SetProject(v string) *EncodeBlindWatermarkRequest {
	s.Project = &v
	return s
}

func (s *EncodeBlindWatermarkRequest) SetTargetImageType(v string) *EncodeBlindWatermarkRequest {
	s.TargetImageType = &v
	return s
}

func (s *EncodeBlindWatermarkRequest) SetTargetUri(v string) *EncodeBlindWatermarkRequest {
	s.TargetUri = &v
	return s
}

func (s *EncodeBlindWatermarkRequest) SetWatermarkUri(v string) *EncodeBlindWatermarkRequest {
	s.WatermarkUri = &v
	return s
}

type EncodeBlindWatermarkResponseBody struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TargetUri *string `json:"TargetUri,omitempty" xml:"TargetUri,omitempty"`
}

func (s EncodeBlindWatermarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EncodeBlindWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *EncodeBlindWatermarkResponseBody) SetContent(v string) *EncodeBlindWatermarkResponseBody {
	s.Content = &v
	return s
}

func (s *EncodeBlindWatermarkResponseBody) SetRequestId(v string) *EncodeBlindWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *EncodeBlindWatermarkResponseBody) SetTargetUri(v string) *EncodeBlindWatermarkResponseBody {
	s.TargetUri = &v
	return s
}

type EncodeBlindWatermarkResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EncodeBlindWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EncodeBlindWatermarkResponse) String() string {
	return tea.Prettify(s)
}

func (s EncodeBlindWatermarkResponse) GoString() string {
	return s.String()
}

func (s *EncodeBlindWatermarkResponse) SetHeaders(v map[string]*string) *EncodeBlindWatermarkResponse {
	s.Headers = v
	return s
}

func (s *EncodeBlindWatermarkResponse) SetStatusCode(v int32) *EncodeBlindWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *EncodeBlindWatermarkResponse) SetBody(v *EncodeBlindWatermarkResponseBody) *EncodeBlindWatermarkResponse {
	s.Body = v
	return s
}

type FindImagesRequest struct {
	AddressLineContentsMatch *string `json:"AddressLineContentsMatch,omitempty" xml:"AddressLineContentsMatch,omitempty"`
	AgeRange                 *string `json:"AgeRange,omitempty" xml:"AgeRange,omitempty"`
	CreateTimeRange          *string `json:"CreateTimeRange,omitempty" xml:"CreateTimeRange,omitempty"`
	Emotion                  *string `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	ExternalId               *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	FacesModifyTimeRange     *string `json:"FacesModifyTimeRange,omitempty" xml:"FacesModifyTimeRange,omitempty"`
	Gender                   *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	GroupId                  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ImageSizeRange           *string `json:"ImageSizeRange,omitempty" xml:"ImageSizeRange,omitempty"`
	ImageTimeRange           *string `json:"ImageTimeRange,omitempty" xml:"ImageTimeRange,omitempty"`
	Limit                    *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	LocationBoundary         *string `json:"LocationBoundary,omitempty" xml:"LocationBoundary,omitempty"`
	Marker                   *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	ModifyTimeRange          *string `json:"ModifyTimeRange,omitempty" xml:"ModifyTimeRange,omitempty"`
	OCRContentsMatch         *string `json:"OCRContentsMatch,omitempty" xml:"OCRContentsMatch,omitempty"`
	Order                    *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OrderBy                  *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	Project                  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RemarksAPrefix           *string `json:"RemarksAPrefix,omitempty" xml:"RemarksAPrefix,omitempty"`
	RemarksArrayAIn          *string `json:"RemarksArrayAIn,omitempty" xml:"RemarksArrayAIn,omitempty"`
	RemarksArrayBIn          *string `json:"RemarksArrayBIn,omitempty" xml:"RemarksArrayBIn,omitempty"`
	RemarksBPrefix           *string `json:"RemarksBPrefix,omitempty" xml:"RemarksBPrefix,omitempty"`
	RemarksCPrefix           *string `json:"RemarksCPrefix,omitempty" xml:"RemarksCPrefix,omitempty"`
	RemarksDPrefix           *string `json:"RemarksDPrefix,omitempty" xml:"RemarksDPrefix,omitempty"`
	SetId                    *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	SourceType               *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SourceUriPrefix          *string `json:"SourceUriPrefix,omitempty" xml:"SourceUriPrefix,omitempty"`
	TagNames                 *string `json:"TagNames,omitempty" xml:"TagNames,omitempty"`
	TagsModifyTimeRange      *string `json:"TagsModifyTimeRange,omitempty" xml:"TagsModifyTimeRange,omitempty"`
}

func (s FindImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s FindImagesRequest) GoString() string {
	return s.String()
}

func (s *FindImagesRequest) SetAddressLineContentsMatch(v string) *FindImagesRequest {
	s.AddressLineContentsMatch = &v
	return s
}

func (s *FindImagesRequest) SetAgeRange(v string) *FindImagesRequest {
	s.AgeRange = &v
	return s
}

func (s *FindImagesRequest) SetCreateTimeRange(v string) *FindImagesRequest {
	s.CreateTimeRange = &v
	return s
}

func (s *FindImagesRequest) SetEmotion(v string) *FindImagesRequest {
	s.Emotion = &v
	return s
}

func (s *FindImagesRequest) SetExternalId(v string) *FindImagesRequest {
	s.ExternalId = &v
	return s
}

func (s *FindImagesRequest) SetFacesModifyTimeRange(v string) *FindImagesRequest {
	s.FacesModifyTimeRange = &v
	return s
}

func (s *FindImagesRequest) SetGender(v string) *FindImagesRequest {
	s.Gender = &v
	return s
}

func (s *FindImagesRequest) SetGroupId(v string) *FindImagesRequest {
	s.GroupId = &v
	return s
}

func (s *FindImagesRequest) SetImageSizeRange(v string) *FindImagesRequest {
	s.ImageSizeRange = &v
	return s
}

func (s *FindImagesRequest) SetImageTimeRange(v string) *FindImagesRequest {
	s.ImageTimeRange = &v
	return s
}

func (s *FindImagesRequest) SetLimit(v int32) *FindImagesRequest {
	s.Limit = &v
	return s
}

func (s *FindImagesRequest) SetLocationBoundary(v string) *FindImagesRequest {
	s.LocationBoundary = &v
	return s
}

func (s *FindImagesRequest) SetMarker(v string) *FindImagesRequest {
	s.Marker = &v
	return s
}

func (s *FindImagesRequest) SetModifyTimeRange(v string) *FindImagesRequest {
	s.ModifyTimeRange = &v
	return s
}

func (s *FindImagesRequest) SetOCRContentsMatch(v string) *FindImagesRequest {
	s.OCRContentsMatch = &v
	return s
}

func (s *FindImagesRequest) SetOrder(v string) *FindImagesRequest {
	s.Order = &v
	return s
}

func (s *FindImagesRequest) SetOrderBy(v string) *FindImagesRequest {
	s.OrderBy = &v
	return s
}

func (s *FindImagesRequest) SetProject(v string) *FindImagesRequest {
	s.Project = &v
	return s
}

func (s *FindImagesRequest) SetRemarksAPrefix(v string) *FindImagesRequest {
	s.RemarksAPrefix = &v
	return s
}

func (s *FindImagesRequest) SetRemarksArrayAIn(v string) *FindImagesRequest {
	s.RemarksArrayAIn = &v
	return s
}

func (s *FindImagesRequest) SetRemarksArrayBIn(v string) *FindImagesRequest {
	s.RemarksArrayBIn = &v
	return s
}

func (s *FindImagesRequest) SetRemarksBPrefix(v string) *FindImagesRequest {
	s.RemarksBPrefix = &v
	return s
}

func (s *FindImagesRequest) SetRemarksCPrefix(v string) *FindImagesRequest {
	s.RemarksCPrefix = &v
	return s
}

func (s *FindImagesRequest) SetRemarksDPrefix(v string) *FindImagesRequest {
	s.RemarksDPrefix = &v
	return s
}

func (s *FindImagesRequest) SetSetId(v string) *FindImagesRequest {
	s.SetId = &v
	return s
}

func (s *FindImagesRequest) SetSourceType(v string) *FindImagesRequest {
	s.SourceType = &v
	return s
}

func (s *FindImagesRequest) SetSourceUriPrefix(v string) *FindImagesRequest {
	s.SourceUriPrefix = &v
	return s
}

func (s *FindImagesRequest) SetTagNames(v string) *FindImagesRequest {
	s.TagNames = &v
	return s
}

func (s *FindImagesRequest) SetTagsModifyTimeRange(v string) *FindImagesRequest {
	s.TagsModifyTimeRange = &v
	return s
}

type FindImagesResponseBody struct {
	Images     []*FindImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	NextMarker *string                         `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId      *string                         `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s FindImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBody) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBody) SetImages(v []*FindImagesResponseBodyImages) *FindImagesResponseBody {
	s.Images = v
	return s
}

func (s *FindImagesResponseBody) SetNextMarker(v string) *FindImagesResponseBody {
	s.NextMarker = &v
	return s
}

func (s *FindImagesResponseBody) SetRequestId(v string) *FindImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindImagesResponseBody) SetSetId(v string) *FindImagesResponseBody {
	s.SetId = &v
	return s
}

type FindImagesResponseBodyImages struct {
	Address                      *FindImagesResponseBodyImagesAddress              `json:"Address,omitempty" xml:"Address,omitempty" type:"Struct"`
	AddressFailReason            *string                                           `json:"AddressFailReason,omitempty" xml:"AddressFailReason,omitempty"`
	AddressModifyTime            *string                                           `json:"AddressModifyTime,omitempty" xml:"AddressModifyTime,omitempty"`
	AddressStatus                *string                                           `json:"AddressStatus,omitempty" xml:"AddressStatus,omitempty"`
	CreateTime                   *string                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CroppingSuggestion           []*FindImagesResponseBodyImagesCroppingSuggestion `json:"CroppingSuggestion,omitempty" xml:"CroppingSuggestion,omitempty" type:"Repeated"`
	CroppingSuggestionFailReason *string                                           `json:"CroppingSuggestionFailReason,omitempty" xml:"CroppingSuggestionFailReason,omitempty"`
	CroppingSuggestionModifyTime *string                                           `json:"CroppingSuggestionModifyTime,omitempty" xml:"CroppingSuggestionModifyTime,omitempty"`
	CroppingSuggestionStatus     *string                                           `json:"CroppingSuggestionStatus,omitempty" xml:"CroppingSuggestionStatus,omitempty"`
	Exif                         *string                                           `json:"Exif,omitempty" xml:"Exif,omitempty"`
	ExternalId                   *string                                           `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	Faces                        []*FindImagesResponseBodyImagesFaces              `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	FacesFailReason              *string                                           `json:"FacesFailReason,omitempty" xml:"FacesFailReason,omitempty"`
	FacesModifyTime              *string                                           `json:"FacesModifyTime,omitempty" xml:"FacesModifyTime,omitempty"`
	FacesStatus                  *string                                           `json:"FacesStatus,omitempty" xml:"FacesStatus,omitempty"`
	FileSize                     *int32                                            `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	ImageFormat                  *string                                           `json:"ImageFormat,omitempty" xml:"ImageFormat,omitempty"`
	ImageHeight                  *int32                                            `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	ImageQuality                 *FindImagesResponseBodyImagesImageQuality         `json:"ImageQuality,omitempty" xml:"ImageQuality,omitempty" type:"Struct"`
	ImageQualityFailReason       *string                                           `json:"ImageQualityFailReason,omitempty" xml:"ImageQualityFailReason,omitempty"`
	ImageQualityModifyTime       *string                                           `json:"ImageQualityModifyTime,omitempty" xml:"ImageQualityModifyTime,omitempty"`
	ImageQualityStatus           *string                                           `json:"ImageQualityStatus,omitempty" xml:"ImageQualityStatus,omitempty"`
	ImageTime                    *string                                           `json:"ImageTime,omitempty" xml:"ImageTime,omitempty"`
	ImageUri                     *string                                           `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	ImageWidth                   *int32                                            `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
	Location                     *string                                           `json:"Location,omitempty" xml:"Location,omitempty"`
	ModifyTime                   *string                                           `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	OCR                          []*FindImagesResponseBodyImagesOCR                `json:"OCR,omitempty" xml:"OCR,omitempty" type:"Repeated"`
	OCRFailReason                *string                                           `json:"OCRFailReason,omitempty" xml:"OCRFailReason,omitempty"`
	OCRModifyTime                *string                                           `json:"OCRModifyTime,omitempty" xml:"OCRModifyTime,omitempty"`
	OCRStatus                    *string                                           `json:"OCRStatus,omitempty" xml:"OCRStatus,omitempty"`
	Orientation                  *string                                           `json:"Orientation,omitempty" xml:"Orientation,omitempty"`
	RemarksA                     *string                                           `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	RemarksArrayA                *string                                           `json:"RemarksArrayA,omitempty" xml:"RemarksArrayA,omitempty"`
	RemarksArrayB                *string                                           `json:"RemarksArrayB,omitempty" xml:"RemarksArrayB,omitempty"`
	RemarksB                     *string                                           `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	RemarksC                     *string                                           `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD                     *string                                           `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	SourcePosition               *string                                           `json:"SourcePosition,omitempty" xml:"SourcePosition,omitempty"`
	SourceType                   *string                                           `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SourceUri                    *string                                           `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
	Tags                         []*FindImagesResponseBodyImagesTags               `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TagsFailReason               *string                                           `json:"TagsFailReason,omitempty" xml:"TagsFailReason,omitempty"`
	TagsModifyTime               *string                                           `json:"TagsModifyTime,omitempty" xml:"TagsModifyTime,omitempty"`
	TagsStatus                   *string                                           `json:"TagsStatus,omitempty" xml:"TagsStatus,omitempty"`
}

func (s FindImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImages) SetAddress(v *FindImagesResponseBodyImagesAddress) *FindImagesResponseBodyImages {
	s.Address = v
	return s
}

func (s *FindImagesResponseBodyImages) SetAddressFailReason(v string) *FindImagesResponseBodyImages {
	s.AddressFailReason = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetAddressModifyTime(v string) *FindImagesResponseBodyImages {
	s.AddressModifyTime = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetAddressStatus(v string) *FindImagesResponseBodyImages {
	s.AddressStatus = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetCreateTime(v string) *FindImagesResponseBodyImages {
	s.CreateTime = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetCroppingSuggestion(v []*FindImagesResponseBodyImagesCroppingSuggestion) *FindImagesResponseBodyImages {
	s.CroppingSuggestion = v
	return s
}

func (s *FindImagesResponseBodyImages) SetCroppingSuggestionFailReason(v string) *FindImagesResponseBodyImages {
	s.CroppingSuggestionFailReason = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetCroppingSuggestionModifyTime(v string) *FindImagesResponseBodyImages {
	s.CroppingSuggestionModifyTime = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetCroppingSuggestionStatus(v string) *FindImagesResponseBodyImages {
	s.CroppingSuggestionStatus = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetExif(v string) *FindImagesResponseBodyImages {
	s.Exif = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetExternalId(v string) *FindImagesResponseBodyImages {
	s.ExternalId = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetFaces(v []*FindImagesResponseBodyImagesFaces) *FindImagesResponseBodyImages {
	s.Faces = v
	return s
}

func (s *FindImagesResponseBodyImages) SetFacesFailReason(v string) *FindImagesResponseBodyImages {
	s.FacesFailReason = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetFacesModifyTime(v string) *FindImagesResponseBodyImages {
	s.FacesModifyTime = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetFacesStatus(v string) *FindImagesResponseBodyImages {
	s.FacesStatus = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetFileSize(v int32) *FindImagesResponseBodyImages {
	s.FileSize = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetImageFormat(v string) *FindImagesResponseBodyImages {
	s.ImageFormat = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetImageHeight(v int32) *FindImagesResponseBodyImages {
	s.ImageHeight = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetImageQuality(v *FindImagesResponseBodyImagesImageQuality) *FindImagesResponseBodyImages {
	s.ImageQuality = v
	return s
}

func (s *FindImagesResponseBodyImages) SetImageQualityFailReason(v string) *FindImagesResponseBodyImages {
	s.ImageQualityFailReason = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetImageQualityModifyTime(v string) *FindImagesResponseBodyImages {
	s.ImageQualityModifyTime = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetImageQualityStatus(v string) *FindImagesResponseBodyImages {
	s.ImageQualityStatus = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetImageTime(v string) *FindImagesResponseBodyImages {
	s.ImageTime = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetImageUri(v string) *FindImagesResponseBodyImages {
	s.ImageUri = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetImageWidth(v int32) *FindImagesResponseBodyImages {
	s.ImageWidth = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetLocation(v string) *FindImagesResponseBodyImages {
	s.Location = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetModifyTime(v string) *FindImagesResponseBodyImages {
	s.ModifyTime = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetOCR(v []*FindImagesResponseBodyImagesOCR) *FindImagesResponseBodyImages {
	s.OCR = v
	return s
}

func (s *FindImagesResponseBodyImages) SetOCRFailReason(v string) *FindImagesResponseBodyImages {
	s.OCRFailReason = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetOCRModifyTime(v string) *FindImagesResponseBodyImages {
	s.OCRModifyTime = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetOCRStatus(v string) *FindImagesResponseBodyImages {
	s.OCRStatus = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetOrientation(v string) *FindImagesResponseBodyImages {
	s.Orientation = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetRemarksA(v string) *FindImagesResponseBodyImages {
	s.RemarksA = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetRemarksArrayA(v string) *FindImagesResponseBodyImages {
	s.RemarksArrayA = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetRemarksArrayB(v string) *FindImagesResponseBodyImages {
	s.RemarksArrayB = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetRemarksB(v string) *FindImagesResponseBodyImages {
	s.RemarksB = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetRemarksC(v string) *FindImagesResponseBodyImages {
	s.RemarksC = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetRemarksD(v string) *FindImagesResponseBodyImages {
	s.RemarksD = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetSourcePosition(v string) *FindImagesResponseBodyImages {
	s.SourcePosition = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetSourceType(v string) *FindImagesResponseBodyImages {
	s.SourceType = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetSourceUri(v string) *FindImagesResponseBodyImages {
	s.SourceUri = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetTags(v []*FindImagesResponseBodyImagesTags) *FindImagesResponseBodyImages {
	s.Tags = v
	return s
}

func (s *FindImagesResponseBodyImages) SetTagsFailReason(v string) *FindImagesResponseBodyImages {
	s.TagsFailReason = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetTagsModifyTime(v string) *FindImagesResponseBodyImages {
	s.TagsModifyTime = &v
	return s
}

func (s *FindImagesResponseBodyImages) SetTagsStatus(v string) *FindImagesResponseBodyImages {
	s.TagsStatus = &v
	return s
}

type FindImagesResponseBodyImagesAddress struct {
	AddressLine *string `json:"AddressLine,omitempty" xml:"AddressLine,omitempty"`
	City        *string `json:"City,omitempty" xml:"City,omitempty"`
	Country     *string `json:"Country,omitempty" xml:"Country,omitempty"`
	District    *string `json:"District,omitempty" xml:"District,omitempty"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty"`
	Township    *string `json:"Township,omitempty" xml:"Township,omitempty"`
}

func (s FindImagesResponseBodyImagesAddress) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesAddress) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesAddress) SetAddressLine(v string) *FindImagesResponseBodyImagesAddress {
	s.AddressLine = &v
	return s
}

func (s *FindImagesResponseBodyImagesAddress) SetCity(v string) *FindImagesResponseBodyImagesAddress {
	s.City = &v
	return s
}

func (s *FindImagesResponseBodyImagesAddress) SetCountry(v string) *FindImagesResponseBodyImagesAddress {
	s.Country = &v
	return s
}

func (s *FindImagesResponseBodyImagesAddress) SetDistrict(v string) *FindImagesResponseBodyImagesAddress {
	s.District = &v
	return s
}

func (s *FindImagesResponseBodyImagesAddress) SetProvince(v string) *FindImagesResponseBodyImagesAddress {
	s.Province = &v
	return s
}

func (s *FindImagesResponseBodyImagesAddress) SetTownship(v string) *FindImagesResponseBodyImagesAddress {
	s.Township = &v
	return s
}

type FindImagesResponseBodyImagesCroppingSuggestion struct {
	AspectRatio      *string                                                         `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	CroppingBoundary *FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary `json:"CroppingBoundary,omitempty" xml:"CroppingBoundary,omitempty" type:"Struct"`
	Score            *float32                                                        `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s FindImagesResponseBodyImagesCroppingSuggestion) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesCroppingSuggestion) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesCroppingSuggestion) SetAspectRatio(v string) *FindImagesResponseBodyImagesCroppingSuggestion {
	s.AspectRatio = &v
	return s
}

func (s *FindImagesResponseBodyImagesCroppingSuggestion) SetCroppingBoundary(v *FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) *FindImagesResponseBodyImagesCroppingSuggestion {
	s.CroppingBoundary = v
	return s
}

func (s *FindImagesResponseBodyImagesCroppingSuggestion) SetScore(v float32) *FindImagesResponseBodyImagesCroppingSuggestion {
	s.Score = &v
	return s
}

type FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) SetHeight(v int32) *FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary {
	s.Height = &v
	return s
}

func (s *FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) SetLeft(v int32) *FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary {
	s.Left = &v
	return s
}

func (s *FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) SetTop(v int32) *FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary {
	s.Top = &v
	return s
}

func (s *FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) SetWidth(v int32) *FindImagesResponseBodyImagesCroppingSuggestionCroppingBoundary {
	s.Width = &v
	return s
}

type FindImagesResponseBodyImagesFaces struct {
	Age               *int32                                           `json:"Age,omitempty" xml:"Age,omitempty"`
	Attractive        *float32                                         `json:"Attractive,omitempty" xml:"Attractive,omitempty"`
	Emotion           *string                                          `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	EmotionConfidence *float32                                         `json:"EmotionConfidence,omitempty" xml:"EmotionConfidence,omitempty"`
	EmotionDetails    *FindImagesResponseBodyImagesFacesEmotionDetails `json:"EmotionDetails,omitempty" xml:"EmotionDetails,omitempty" type:"Struct"`
	FaceAttributes    *FindImagesResponseBodyImagesFacesFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" type:"Struct"`
	FaceConfidence    *float32                                         `json:"FaceConfidence,omitempty" xml:"FaceConfidence,omitempty"`
	FaceId            *string                                          `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	FaceQuality       *float32                                         `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	Gender            *string                                          `json:"Gender,omitempty" xml:"Gender,omitempty"`
	GenderConfidence  *float32                                         `json:"GenderConfidence,omitempty" xml:"GenderConfidence,omitempty"`
	GroupId           *string                                          `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s FindImagesResponseBodyImagesFaces) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesFaces) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesFaces) SetAge(v int32) *FindImagesResponseBodyImagesFaces {
	s.Age = &v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetAttractive(v float32) *FindImagesResponseBodyImagesFaces {
	s.Attractive = &v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetEmotion(v string) *FindImagesResponseBodyImagesFaces {
	s.Emotion = &v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetEmotionConfidence(v float32) *FindImagesResponseBodyImagesFaces {
	s.EmotionConfidence = &v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetEmotionDetails(v *FindImagesResponseBodyImagesFacesEmotionDetails) *FindImagesResponseBodyImagesFaces {
	s.EmotionDetails = v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetFaceAttributes(v *FindImagesResponseBodyImagesFacesFaceAttributes) *FindImagesResponseBodyImagesFaces {
	s.FaceAttributes = v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetFaceConfidence(v float32) *FindImagesResponseBodyImagesFaces {
	s.FaceConfidence = &v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetFaceId(v string) *FindImagesResponseBodyImagesFaces {
	s.FaceId = &v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetFaceQuality(v float32) *FindImagesResponseBodyImagesFaces {
	s.FaceQuality = &v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetGender(v string) *FindImagesResponseBodyImagesFaces {
	s.Gender = &v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetGenderConfidence(v float32) *FindImagesResponseBodyImagesFaces {
	s.GenderConfidence = &v
	return s
}

func (s *FindImagesResponseBodyImagesFaces) SetGroupId(v string) *FindImagesResponseBodyImagesFaces {
	s.GroupId = &v
	return s
}

type FindImagesResponseBodyImagesFacesEmotionDetails struct {
	ANGRY     *float32 `json:"ANGRY,omitempty" xml:"ANGRY,omitempty"`
	CALM      *float32 `json:"CALM,omitempty" xml:"CALM,omitempty"`
	DISGUSTED *float32 `json:"DISGUSTED,omitempty" xml:"DISGUSTED,omitempty"`
	HAPPY     *float32 `json:"HAPPY,omitempty" xml:"HAPPY,omitempty"`
	SAD       *float32 `json:"SAD,omitempty" xml:"SAD,omitempty"`
	SCARED    *float32 `json:"SCARED,omitempty" xml:"SCARED,omitempty"`
	SURPRISED *float32 `json:"SURPRISED,omitempty" xml:"SURPRISED,omitempty"`
}

func (s FindImagesResponseBodyImagesFacesEmotionDetails) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesFacesEmotionDetails) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesFacesEmotionDetails) SetANGRY(v float32) *FindImagesResponseBodyImagesFacesEmotionDetails {
	s.ANGRY = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesEmotionDetails) SetCALM(v float32) *FindImagesResponseBodyImagesFacesEmotionDetails {
	s.CALM = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesEmotionDetails) SetDISGUSTED(v float32) *FindImagesResponseBodyImagesFacesEmotionDetails {
	s.DISGUSTED = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesEmotionDetails) SetHAPPY(v float32) *FindImagesResponseBodyImagesFacesEmotionDetails {
	s.HAPPY = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesEmotionDetails) SetSAD(v float32) *FindImagesResponseBodyImagesFacesEmotionDetails {
	s.SAD = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesEmotionDetails) SetSCARED(v float32) *FindImagesResponseBodyImagesFacesEmotionDetails {
	s.SCARED = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesEmotionDetails) SetSURPRISED(v float32) *FindImagesResponseBodyImagesFacesEmotionDetails {
	s.SURPRISED = &v
	return s
}

type FindImagesResponseBodyImagesFacesFaceAttributes struct {
	Beard             *string                                                      `json:"Beard,omitempty" xml:"Beard,omitempty"`
	BeardConfidence   *float32                                                     `json:"BeardConfidence,omitempty" xml:"BeardConfidence,omitempty"`
	FaceBoundary      *FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary `json:"FaceBoundary,omitempty" xml:"FaceBoundary,omitempty" type:"Struct"`
	Glasses           *string                                                      `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	GlassesConfidence *float32                                                     `json:"GlassesConfidence,omitempty" xml:"GlassesConfidence,omitempty"`
	HeadPose          *FindImagesResponseBodyImagesFacesFaceAttributesHeadPose     `json:"HeadPose,omitempty" xml:"HeadPose,omitempty" type:"Struct"`
	Mask              *string                                                      `json:"Mask,omitempty" xml:"Mask,omitempty"`
	MaskConfidence    *float32                                                     `json:"MaskConfidence,omitempty" xml:"MaskConfidence,omitempty"`
}

func (s FindImagesResponseBodyImagesFacesFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesFacesFaceAttributes) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributes) SetBeard(v string) *FindImagesResponseBodyImagesFacesFaceAttributes {
	s.Beard = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributes) SetBeardConfidence(v float32) *FindImagesResponseBodyImagesFacesFaceAttributes {
	s.BeardConfidence = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributes) SetFaceBoundary(v *FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) *FindImagesResponseBodyImagesFacesFaceAttributes {
	s.FaceBoundary = v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributes) SetGlasses(v string) *FindImagesResponseBodyImagesFacesFaceAttributes {
	s.Glasses = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributes) SetGlassesConfidence(v float32) *FindImagesResponseBodyImagesFacesFaceAttributes {
	s.GlassesConfidence = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributes) SetHeadPose(v *FindImagesResponseBodyImagesFacesFaceAttributesHeadPose) *FindImagesResponseBodyImagesFacesFaceAttributes {
	s.HeadPose = v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributes) SetMask(v string) *FindImagesResponseBodyImagesFacesFaceAttributes {
	s.Mask = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributes) SetMaskConfidence(v float32) *FindImagesResponseBodyImagesFacesFaceAttributes {
	s.MaskConfidence = &v
	return s
}

type FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) SetHeight(v int32) *FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary {
	s.Height = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) SetLeft(v int32) *FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary {
	s.Left = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) SetTop(v int32) *FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary {
	s.Top = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) SetWidth(v int32) *FindImagesResponseBodyImagesFacesFaceAttributesFaceBoundary {
	s.Width = &v
	return s
}

type FindImagesResponseBodyImagesFacesFaceAttributesHeadPose struct {
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	Roll  *float32 `json:"Roll,omitempty" xml:"Roll,omitempty"`
	Yaw   *float32 `json:"Yaw,omitempty" xml:"Yaw,omitempty"`
}

func (s FindImagesResponseBodyImagesFacesFaceAttributesHeadPose) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesFacesFaceAttributesHeadPose) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributesHeadPose) SetPitch(v float32) *FindImagesResponseBodyImagesFacesFaceAttributesHeadPose {
	s.Pitch = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributesHeadPose) SetRoll(v float32) *FindImagesResponseBodyImagesFacesFaceAttributesHeadPose {
	s.Roll = &v
	return s
}

func (s *FindImagesResponseBodyImagesFacesFaceAttributesHeadPose) SetYaw(v float32) *FindImagesResponseBodyImagesFacesFaceAttributesHeadPose {
	s.Yaw = &v
	return s
}

type FindImagesResponseBodyImagesImageQuality struct {
	Clarity          *float32 `json:"Clarity,omitempty" xml:"Clarity,omitempty"`
	ClarityScore     *float32 `json:"ClarityScore,omitempty" xml:"ClarityScore,omitempty"`
	Color            *float32 `json:"Color,omitempty" xml:"Color,omitempty"`
	ColorScore       *float32 `json:"ColorScore,omitempty" xml:"ColorScore,omitempty"`
	CompositionScore *float32 `json:"CompositionScore,omitempty" xml:"CompositionScore,omitempty"`
	Contrast         *float32 `json:"Contrast,omitempty" xml:"Contrast,omitempty"`
	ContrastScore    *float32 `json:"ContrastScore,omitempty" xml:"ContrastScore,omitempty"`
	Exposure         *float32 `json:"Exposure,omitempty" xml:"Exposure,omitempty"`
	ExposureScore    *float32 `json:"ExposureScore,omitempty" xml:"ExposureScore,omitempty"`
	OverallScore     *float32 `json:"OverallScore,omitempty" xml:"OverallScore,omitempty"`
}

func (s FindImagesResponseBodyImagesImageQuality) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesImageQuality) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesImageQuality) SetClarity(v float32) *FindImagesResponseBodyImagesImageQuality {
	s.Clarity = &v
	return s
}

func (s *FindImagesResponseBodyImagesImageQuality) SetClarityScore(v float32) *FindImagesResponseBodyImagesImageQuality {
	s.ClarityScore = &v
	return s
}

func (s *FindImagesResponseBodyImagesImageQuality) SetColor(v float32) *FindImagesResponseBodyImagesImageQuality {
	s.Color = &v
	return s
}

func (s *FindImagesResponseBodyImagesImageQuality) SetColorScore(v float32) *FindImagesResponseBodyImagesImageQuality {
	s.ColorScore = &v
	return s
}

func (s *FindImagesResponseBodyImagesImageQuality) SetCompositionScore(v float32) *FindImagesResponseBodyImagesImageQuality {
	s.CompositionScore = &v
	return s
}

func (s *FindImagesResponseBodyImagesImageQuality) SetContrast(v float32) *FindImagesResponseBodyImagesImageQuality {
	s.Contrast = &v
	return s
}

func (s *FindImagesResponseBodyImagesImageQuality) SetContrastScore(v float32) *FindImagesResponseBodyImagesImageQuality {
	s.ContrastScore = &v
	return s
}

func (s *FindImagesResponseBodyImagesImageQuality) SetExposure(v float32) *FindImagesResponseBodyImagesImageQuality {
	s.Exposure = &v
	return s
}

func (s *FindImagesResponseBodyImagesImageQuality) SetExposureScore(v float32) *FindImagesResponseBodyImagesImageQuality {
	s.ExposureScore = &v
	return s
}

func (s *FindImagesResponseBodyImagesImageQuality) SetOverallScore(v float32) *FindImagesResponseBodyImagesImageQuality {
	s.OverallScore = &v
	return s
}

type FindImagesResponseBodyImagesOCR struct {
	OCRBoundary   *FindImagesResponseBodyImagesOCROCRBoundary `json:"OCRBoundary,omitempty" xml:"OCRBoundary,omitempty" type:"Struct"`
	OCRConfidence *float32                                    `json:"OCRConfidence,omitempty" xml:"OCRConfidence,omitempty"`
	OCRContents   *string                                     `json:"OCRContents,omitempty" xml:"OCRContents,omitempty"`
}

func (s FindImagesResponseBodyImagesOCR) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesOCR) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesOCR) SetOCRBoundary(v *FindImagesResponseBodyImagesOCROCRBoundary) *FindImagesResponseBodyImagesOCR {
	s.OCRBoundary = v
	return s
}

func (s *FindImagesResponseBodyImagesOCR) SetOCRConfidence(v float32) *FindImagesResponseBodyImagesOCR {
	s.OCRConfidence = &v
	return s
}

func (s *FindImagesResponseBodyImagesOCR) SetOCRContents(v string) *FindImagesResponseBodyImagesOCR {
	s.OCRContents = &v
	return s
}

type FindImagesResponseBodyImagesOCROCRBoundary struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s FindImagesResponseBodyImagesOCROCRBoundary) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesOCROCRBoundary) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesOCROCRBoundary) SetHeight(v int32) *FindImagesResponseBodyImagesOCROCRBoundary {
	s.Height = &v
	return s
}

func (s *FindImagesResponseBodyImagesOCROCRBoundary) SetLeft(v int32) *FindImagesResponseBodyImagesOCROCRBoundary {
	s.Left = &v
	return s
}

func (s *FindImagesResponseBodyImagesOCROCRBoundary) SetTop(v int32) *FindImagesResponseBodyImagesOCROCRBoundary {
	s.Top = &v
	return s
}

func (s *FindImagesResponseBodyImagesOCROCRBoundary) SetWidth(v int32) *FindImagesResponseBodyImagesOCROCRBoundary {
	s.Width = &v
	return s
}

type FindImagesResponseBodyImagesTags struct {
	CentricScore  *float32 `json:"CentricScore,omitempty" xml:"CentricScore,omitempty"`
	ParentTagName *string  `json:"ParentTagName,omitempty" xml:"ParentTagName,omitempty"`
	TagConfidence *float32 `json:"TagConfidence,omitempty" xml:"TagConfidence,omitempty"`
	TagLevel      *int32   `json:"TagLevel,omitempty" xml:"TagLevel,omitempty"`
	TagName       *string  `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s FindImagesResponseBodyImagesTags) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponseBodyImagesTags) GoString() string {
	return s.String()
}

func (s *FindImagesResponseBodyImagesTags) SetCentricScore(v float32) *FindImagesResponseBodyImagesTags {
	s.CentricScore = &v
	return s
}

func (s *FindImagesResponseBodyImagesTags) SetParentTagName(v string) *FindImagesResponseBodyImagesTags {
	s.ParentTagName = &v
	return s
}

func (s *FindImagesResponseBodyImagesTags) SetTagConfidence(v float32) *FindImagesResponseBodyImagesTags {
	s.TagConfidence = &v
	return s
}

func (s *FindImagesResponseBodyImagesTags) SetTagLevel(v int32) *FindImagesResponseBodyImagesTags {
	s.TagLevel = &v
	return s
}

func (s *FindImagesResponseBodyImagesTags) SetTagName(v string) *FindImagesResponseBodyImagesTags {
	s.TagName = &v
	return s
}

type FindImagesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FindImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s FindImagesResponse) GoString() string {
	return s.String()
}

func (s *FindImagesResponse) SetHeaders(v map[string]*string) *FindImagesResponse {
	s.Headers = v
	return s
}

func (s *FindImagesResponse) SetStatusCode(v int32) *FindImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *FindImagesResponse) SetBody(v *FindImagesResponseBody) *FindImagesResponse {
	s.Body = v
	return s
}

type FindSimilarFacesRequest struct {
	FaceId         *string  `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	ImageUri       *string  `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Limit          *int32   `json:"Limit,omitempty" xml:"Limit,omitempty"`
	MinSimilarity  *float32 `json:"MinSimilarity,omitempty" xml:"MinSimilarity,omitempty"`
	Project        *string  `json:"Project,omitempty" xml:"Project,omitempty"`
	ResponseFormat *string  `json:"ResponseFormat,omitempty" xml:"ResponseFormat,omitempty"`
	SetId          *string  `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s FindSimilarFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s FindSimilarFacesRequest) GoString() string {
	return s.String()
}

func (s *FindSimilarFacesRequest) SetFaceId(v string) *FindSimilarFacesRequest {
	s.FaceId = &v
	return s
}

func (s *FindSimilarFacesRequest) SetImageUri(v string) *FindSimilarFacesRequest {
	s.ImageUri = &v
	return s
}

func (s *FindSimilarFacesRequest) SetLimit(v int32) *FindSimilarFacesRequest {
	s.Limit = &v
	return s
}

func (s *FindSimilarFacesRequest) SetMinSimilarity(v float32) *FindSimilarFacesRequest {
	s.MinSimilarity = &v
	return s
}

func (s *FindSimilarFacesRequest) SetProject(v string) *FindSimilarFacesRequest {
	s.Project = &v
	return s
}

func (s *FindSimilarFacesRequest) SetResponseFormat(v string) *FindSimilarFacesRequest {
	s.ResponseFormat = &v
	return s
}

func (s *FindSimilarFacesRequest) SetSetId(v string) *FindSimilarFacesRequest {
	s.SetId = &v
	return s
}

type FindSimilarFacesResponseBody struct {
	Faces     []*FindSimilarFacesResponseBodyFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FindSimilarFacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindSimilarFacesResponseBody) GoString() string {
	return s.String()
}

func (s *FindSimilarFacesResponseBody) SetFaces(v []*FindSimilarFacesResponseBodyFaces) *FindSimilarFacesResponseBody {
	s.Faces = v
	return s
}

func (s *FindSimilarFacesResponseBody) SetRequestId(v string) *FindSimilarFacesResponseBody {
	s.RequestId = &v
	return s
}

type FindSimilarFacesResponseBodyFaces struct {
	ExternalId     *string                                          `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	FaceAttributes *FindSimilarFacesResponseBodyFacesFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" type:"Struct"`
	FaceId         *string                                          `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	ImageUri       *string                                          `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	SimilarFaces   []*FindSimilarFacesResponseBodyFacesSimilarFaces `json:"SimilarFaces,omitempty" xml:"SimilarFaces,omitempty" type:"Repeated"`
	Similarity     *float32                                         `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
}

func (s FindSimilarFacesResponseBodyFaces) String() string {
	return tea.Prettify(s)
}

func (s FindSimilarFacesResponseBodyFaces) GoString() string {
	return s.String()
}

func (s *FindSimilarFacesResponseBodyFaces) SetExternalId(v string) *FindSimilarFacesResponseBodyFaces {
	s.ExternalId = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFaces) SetFaceAttributes(v *FindSimilarFacesResponseBodyFacesFaceAttributes) *FindSimilarFacesResponseBodyFaces {
	s.FaceAttributes = v
	return s
}

func (s *FindSimilarFacesResponseBodyFaces) SetFaceId(v string) *FindSimilarFacesResponseBodyFaces {
	s.FaceId = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFaces) SetImageUri(v string) *FindSimilarFacesResponseBodyFaces {
	s.ImageUri = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFaces) SetSimilarFaces(v []*FindSimilarFacesResponseBodyFacesSimilarFaces) *FindSimilarFacesResponseBodyFaces {
	s.SimilarFaces = v
	return s
}

func (s *FindSimilarFacesResponseBodyFaces) SetSimilarity(v float32) *FindSimilarFacesResponseBodyFaces {
	s.Similarity = &v
	return s
}

type FindSimilarFacesResponseBodyFacesFaceAttributes struct {
	FaceBoundary *FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary `json:"FaceBoundary,omitempty" xml:"FaceBoundary,omitempty" type:"Struct"`
}

func (s FindSimilarFacesResponseBodyFacesFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s FindSimilarFacesResponseBodyFacesFaceAttributes) GoString() string {
	return s.String()
}

func (s *FindSimilarFacesResponseBodyFacesFaceAttributes) SetFaceBoundary(v *FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary) *FindSimilarFacesResponseBodyFacesFaceAttributes {
	s.FaceBoundary = v
	return s
}

type FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary) String() string {
	return tea.Prettify(s)
}

func (s FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary) GoString() string {
	return s.String()
}

func (s *FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary) SetHeight(v int32) *FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary {
	s.Height = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary) SetLeft(v int32) *FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary {
	s.Left = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary) SetTop(v int32) *FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary {
	s.Top = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary) SetWidth(v int32) *FindSimilarFacesResponseBodyFacesFaceAttributesFaceBoundary {
	s.Width = &v
	return s
}

type FindSimilarFacesResponseBodyFacesSimilarFaces struct {
	ExternalId     *string                                                      `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	FaceAttributes *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" type:"Struct"`
	FaceId         *string                                                      `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	ImageUri       *string                                                      `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Similarity     *float32                                                     `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
}

func (s FindSimilarFacesResponseBodyFacesSimilarFaces) String() string {
	return tea.Prettify(s)
}

func (s FindSimilarFacesResponseBodyFacesSimilarFaces) GoString() string {
	return s.String()
}

func (s *FindSimilarFacesResponseBodyFacesSimilarFaces) SetExternalId(v string) *FindSimilarFacesResponseBodyFacesSimilarFaces {
	s.ExternalId = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFacesSimilarFaces) SetFaceAttributes(v *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributes) *FindSimilarFacesResponseBodyFacesSimilarFaces {
	s.FaceAttributes = v
	return s
}

func (s *FindSimilarFacesResponseBodyFacesSimilarFaces) SetFaceId(v string) *FindSimilarFacesResponseBodyFacesSimilarFaces {
	s.FaceId = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFacesSimilarFaces) SetImageUri(v string) *FindSimilarFacesResponseBodyFacesSimilarFaces {
	s.ImageUri = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFacesSimilarFaces) SetSimilarity(v float32) *FindSimilarFacesResponseBodyFacesSimilarFaces {
	s.Similarity = &v
	return s
}

type FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributes struct {
	FaceBoundary *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary `json:"FaceBoundary,omitempty" xml:"FaceBoundary,omitempty" type:"Struct"`
}

func (s FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributes) GoString() string {
	return s.String()
}

func (s *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributes) SetFaceBoundary(v *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary) *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributes {
	s.FaceBoundary = v
	return s
}

type FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary) String() string {
	return tea.Prettify(s)
}

func (s FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary) GoString() string {
	return s.String()
}

func (s *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary) SetHeight(v int32) *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary {
	s.Height = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary) SetLeft(v int32) *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary {
	s.Left = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary) SetTop(v int32) *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary {
	s.Top = &v
	return s
}

func (s *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary) SetWidth(v int32) *FindSimilarFacesResponseBodyFacesSimilarFacesFaceAttributesFaceBoundary {
	s.Width = &v
	return s
}

type FindSimilarFacesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FindSimilarFacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindSimilarFacesResponse) String() string {
	return tea.Prettify(s)
}

func (s FindSimilarFacesResponse) GoString() string {
	return s.String()
}

func (s *FindSimilarFacesResponse) SetHeaders(v map[string]*string) *FindSimilarFacesResponse {
	s.Headers = v
	return s
}

func (s *FindSimilarFacesResponse) SetStatusCode(v int32) *FindSimilarFacesResponse {
	s.StatusCode = &v
	return s
}

func (s *FindSimilarFacesResponse) SetBody(v *FindSimilarFacesResponseBody) *FindSimilarFacesResponse {
	s.Body = v
	return s
}

type GetImageRequest struct {
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId    *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s GetImageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageRequest) GoString() string {
	return s.String()
}

func (s *GetImageRequest) SetImageUri(v string) *GetImageRequest {
	s.ImageUri = &v
	return s
}

func (s *GetImageRequest) SetProject(v string) *GetImageRequest {
	s.Project = &v
	return s
}

func (s *GetImageRequest) SetSetId(v string) *GetImageRequest {
	s.SetId = &v
	return s
}

type GetImageResponseBody struct {
	Address                      *GetImageResponseBodyAddress              `json:"Address,omitempty" xml:"Address,omitempty" type:"Struct"`
	AddressFailReason            *string                                   `json:"AddressFailReason,omitempty" xml:"AddressFailReason,omitempty"`
	AddressModifyTime            *string                                   `json:"AddressModifyTime,omitempty" xml:"AddressModifyTime,omitempty"`
	AddressStatus                *string                                   `json:"AddressStatus,omitempty" xml:"AddressStatus,omitempty"`
	CreateTime                   *string                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CroppingSuggestion           []*GetImageResponseBodyCroppingSuggestion `json:"CroppingSuggestion,omitempty" xml:"CroppingSuggestion,omitempty" type:"Repeated"`
	CroppingSuggestionFailReason *string                                   `json:"CroppingSuggestionFailReason,omitempty" xml:"CroppingSuggestionFailReason,omitempty"`
	CroppingSuggestionModifyTime *string                                   `json:"CroppingSuggestionModifyTime,omitempty" xml:"CroppingSuggestionModifyTime,omitempty"`
	CroppingSuggestionStatus     *string                                   `json:"CroppingSuggestionStatus,omitempty" xml:"CroppingSuggestionStatus,omitempty"`
	Exif                         *string                                   `json:"Exif,omitempty" xml:"Exif,omitempty"`
	ExternalId                   *string                                   `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	Faces                        []*GetImageResponseBodyFaces              `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	FacesFailReason              *string                                   `json:"FacesFailReason,omitempty" xml:"FacesFailReason,omitempty"`
	FacesModifyTime              *string                                   `json:"FacesModifyTime,omitempty" xml:"FacesModifyTime,omitempty"`
	FacesStatus                  *string                                   `json:"FacesStatus,omitempty" xml:"FacesStatus,omitempty"`
	FileSize                     *int32                                    `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	ImageFormat                  *string                                   `json:"ImageFormat,omitempty" xml:"ImageFormat,omitempty"`
	ImageHeight                  *int32                                    `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	ImageQuality                 *GetImageResponseBodyImageQuality         `json:"ImageQuality,omitempty" xml:"ImageQuality,omitempty" type:"Struct"`
	ImageQualityFailReason       *string                                   `json:"ImageQualityFailReason,omitempty" xml:"ImageQualityFailReason,omitempty"`
	ImageQualityModifyTime       *string                                   `json:"ImageQualityModifyTime,omitempty" xml:"ImageQualityModifyTime,omitempty"`
	ImageQualityStatus           *string                                   `json:"ImageQualityStatus,omitempty" xml:"ImageQualityStatus,omitempty"`
	ImageTime                    *string                                   `json:"ImageTime,omitempty" xml:"ImageTime,omitempty"`
	ImageUri                     *string                                   `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	ImageWidth                   *int32                                    `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
	Location                     *string                                   `json:"Location,omitempty" xml:"Location,omitempty"`
	ModifyTime                   *string                                   `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	OCR                          []*GetImageResponseBodyOCR                `json:"OCR,omitempty" xml:"OCR,omitempty" type:"Repeated"`
	OCRFailReason                *string                                   `json:"OCRFailReason,omitempty" xml:"OCRFailReason,omitempty"`
	OCRModifyTime                *string                                   `json:"OCRModifyTime,omitempty" xml:"OCRModifyTime,omitempty"`
	OCRStatus                    *string                                   `json:"OCRStatus,omitempty" xml:"OCRStatus,omitempty"`
	Orientation                  *string                                   `json:"Orientation,omitempty" xml:"Orientation,omitempty"`
	RemarksA                     *string                                   `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	RemarksArrayA                *string                                   `json:"RemarksArrayA,omitempty" xml:"RemarksArrayA,omitempty"`
	RemarksArrayB                *string                                   `json:"RemarksArrayB,omitempty" xml:"RemarksArrayB,omitempty"`
	RemarksB                     *string                                   `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	RemarksC                     *string                                   `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD                     *string                                   `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	RequestId                    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId                        *string                                   `json:"SetId,omitempty" xml:"SetId,omitempty"`
	SourcePosition               *string                                   `json:"SourcePosition,omitempty" xml:"SourcePosition,omitempty"`
	SourceType                   *string                                   `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SourceUri                    *string                                   `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
	Tags                         []*GetImageResponseBodyTags               `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TagsFailReason               *string                                   `json:"TagsFailReason,omitempty" xml:"TagsFailReason,omitempty"`
	TagsModifyTime               *string                                   `json:"TagsModifyTime,omitempty" xml:"TagsModifyTime,omitempty"`
	TagsStatus                   *string                                   `json:"TagsStatus,omitempty" xml:"TagsStatus,omitempty"`
}

func (s GetImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageResponseBody) SetAddress(v *GetImageResponseBodyAddress) *GetImageResponseBody {
	s.Address = v
	return s
}

func (s *GetImageResponseBody) SetAddressFailReason(v string) *GetImageResponseBody {
	s.AddressFailReason = &v
	return s
}

func (s *GetImageResponseBody) SetAddressModifyTime(v string) *GetImageResponseBody {
	s.AddressModifyTime = &v
	return s
}

func (s *GetImageResponseBody) SetAddressStatus(v string) *GetImageResponseBody {
	s.AddressStatus = &v
	return s
}

func (s *GetImageResponseBody) SetCreateTime(v string) *GetImageResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetImageResponseBody) SetCroppingSuggestion(v []*GetImageResponseBodyCroppingSuggestion) *GetImageResponseBody {
	s.CroppingSuggestion = v
	return s
}

func (s *GetImageResponseBody) SetCroppingSuggestionFailReason(v string) *GetImageResponseBody {
	s.CroppingSuggestionFailReason = &v
	return s
}

func (s *GetImageResponseBody) SetCroppingSuggestionModifyTime(v string) *GetImageResponseBody {
	s.CroppingSuggestionModifyTime = &v
	return s
}

func (s *GetImageResponseBody) SetCroppingSuggestionStatus(v string) *GetImageResponseBody {
	s.CroppingSuggestionStatus = &v
	return s
}

func (s *GetImageResponseBody) SetExif(v string) *GetImageResponseBody {
	s.Exif = &v
	return s
}

func (s *GetImageResponseBody) SetExternalId(v string) *GetImageResponseBody {
	s.ExternalId = &v
	return s
}

func (s *GetImageResponseBody) SetFaces(v []*GetImageResponseBodyFaces) *GetImageResponseBody {
	s.Faces = v
	return s
}

func (s *GetImageResponseBody) SetFacesFailReason(v string) *GetImageResponseBody {
	s.FacesFailReason = &v
	return s
}

func (s *GetImageResponseBody) SetFacesModifyTime(v string) *GetImageResponseBody {
	s.FacesModifyTime = &v
	return s
}

func (s *GetImageResponseBody) SetFacesStatus(v string) *GetImageResponseBody {
	s.FacesStatus = &v
	return s
}

func (s *GetImageResponseBody) SetFileSize(v int32) *GetImageResponseBody {
	s.FileSize = &v
	return s
}

func (s *GetImageResponseBody) SetImageFormat(v string) *GetImageResponseBody {
	s.ImageFormat = &v
	return s
}

func (s *GetImageResponseBody) SetImageHeight(v int32) *GetImageResponseBody {
	s.ImageHeight = &v
	return s
}

func (s *GetImageResponseBody) SetImageQuality(v *GetImageResponseBodyImageQuality) *GetImageResponseBody {
	s.ImageQuality = v
	return s
}

func (s *GetImageResponseBody) SetImageQualityFailReason(v string) *GetImageResponseBody {
	s.ImageQualityFailReason = &v
	return s
}

func (s *GetImageResponseBody) SetImageQualityModifyTime(v string) *GetImageResponseBody {
	s.ImageQualityModifyTime = &v
	return s
}

func (s *GetImageResponseBody) SetImageQualityStatus(v string) *GetImageResponseBody {
	s.ImageQualityStatus = &v
	return s
}

func (s *GetImageResponseBody) SetImageTime(v string) *GetImageResponseBody {
	s.ImageTime = &v
	return s
}

func (s *GetImageResponseBody) SetImageUri(v string) *GetImageResponseBody {
	s.ImageUri = &v
	return s
}

func (s *GetImageResponseBody) SetImageWidth(v int32) *GetImageResponseBody {
	s.ImageWidth = &v
	return s
}

func (s *GetImageResponseBody) SetLocation(v string) *GetImageResponseBody {
	s.Location = &v
	return s
}

func (s *GetImageResponseBody) SetModifyTime(v string) *GetImageResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *GetImageResponseBody) SetOCR(v []*GetImageResponseBodyOCR) *GetImageResponseBody {
	s.OCR = v
	return s
}

func (s *GetImageResponseBody) SetOCRFailReason(v string) *GetImageResponseBody {
	s.OCRFailReason = &v
	return s
}

func (s *GetImageResponseBody) SetOCRModifyTime(v string) *GetImageResponseBody {
	s.OCRModifyTime = &v
	return s
}

func (s *GetImageResponseBody) SetOCRStatus(v string) *GetImageResponseBody {
	s.OCRStatus = &v
	return s
}

func (s *GetImageResponseBody) SetOrientation(v string) *GetImageResponseBody {
	s.Orientation = &v
	return s
}

func (s *GetImageResponseBody) SetRemarksA(v string) *GetImageResponseBody {
	s.RemarksA = &v
	return s
}

func (s *GetImageResponseBody) SetRemarksArrayA(v string) *GetImageResponseBody {
	s.RemarksArrayA = &v
	return s
}

func (s *GetImageResponseBody) SetRemarksArrayB(v string) *GetImageResponseBody {
	s.RemarksArrayB = &v
	return s
}

func (s *GetImageResponseBody) SetRemarksB(v string) *GetImageResponseBody {
	s.RemarksB = &v
	return s
}

func (s *GetImageResponseBody) SetRemarksC(v string) *GetImageResponseBody {
	s.RemarksC = &v
	return s
}

func (s *GetImageResponseBody) SetRemarksD(v string) *GetImageResponseBody {
	s.RemarksD = &v
	return s
}

func (s *GetImageResponseBody) SetRequestId(v string) *GetImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageResponseBody) SetSetId(v string) *GetImageResponseBody {
	s.SetId = &v
	return s
}

func (s *GetImageResponseBody) SetSourcePosition(v string) *GetImageResponseBody {
	s.SourcePosition = &v
	return s
}

func (s *GetImageResponseBody) SetSourceType(v string) *GetImageResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetImageResponseBody) SetSourceUri(v string) *GetImageResponseBody {
	s.SourceUri = &v
	return s
}

func (s *GetImageResponseBody) SetTags(v []*GetImageResponseBodyTags) *GetImageResponseBody {
	s.Tags = v
	return s
}

func (s *GetImageResponseBody) SetTagsFailReason(v string) *GetImageResponseBody {
	s.TagsFailReason = &v
	return s
}

func (s *GetImageResponseBody) SetTagsModifyTime(v string) *GetImageResponseBody {
	s.TagsModifyTime = &v
	return s
}

func (s *GetImageResponseBody) SetTagsStatus(v string) *GetImageResponseBody {
	s.TagsStatus = &v
	return s
}

type GetImageResponseBodyAddress struct {
	AddressLine *string `json:"AddressLine,omitempty" xml:"AddressLine,omitempty"`
	City        *string `json:"City,omitempty" xml:"City,omitempty"`
	Country     *string `json:"Country,omitempty" xml:"Country,omitempty"`
	District    *string `json:"District,omitempty" xml:"District,omitempty"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty"`
	Township    *string `json:"Township,omitempty" xml:"Township,omitempty"`
}

func (s GetImageResponseBodyAddress) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyAddress) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyAddress) SetAddressLine(v string) *GetImageResponseBodyAddress {
	s.AddressLine = &v
	return s
}

func (s *GetImageResponseBodyAddress) SetCity(v string) *GetImageResponseBodyAddress {
	s.City = &v
	return s
}

func (s *GetImageResponseBodyAddress) SetCountry(v string) *GetImageResponseBodyAddress {
	s.Country = &v
	return s
}

func (s *GetImageResponseBodyAddress) SetDistrict(v string) *GetImageResponseBodyAddress {
	s.District = &v
	return s
}

func (s *GetImageResponseBodyAddress) SetProvince(v string) *GetImageResponseBodyAddress {
	s.Province = &v
	return s
}

func (s *GetImageResponseBodyAddress) SetTownship(v string) *GetImageResponseBodyAddress {
	s.Township = &v
	return s
}

type GetImageResponseBodyCroppingSuggestion struct {
	AspectRatio      *string                                                 `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	CroppingBoundary *GetImageResponseBodyCroppingSuggestionCroppingBoundary `json:"CroppingBoundary,omitempty" xml:"CroppingBoundary,omitempty" type:"Struct"`
	Score            *float32                                                `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s GetImageResponseBodyCroppingSuggestion) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyCroppingSuggestion) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyCroppingSuggestion) SetAspectRatio(v string) *GetImageResponseBodyCroppingSuggestion {
	s.AspectRatio = &v
	return s
}

func (s *GetImageResponseBodyCroppingSuggestion) SetCroppingBoundary(v *GetImageResponseBodyCroppingSuggestionCroppingBoundary) *GetImageResponseBodyCroppingSuggestion {
	s.CroppingBoundary = v
	return s
}

func (s *GetImageResponseBodyCroppingSuggestion) SetScore(v float32) *GetImageResponseBodyCroppingSuggestion {
	s.Score = &v
	return s
}

type GetImageResponseBodyCroppingSuggestionCroppingBoundary struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetImageResponseBodyCroppingSuggestionCroppingBoundary) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyCroppingSuggestionCroppingBoundary) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyCroppingSuggestionCroppingBoundary) SetHeight(v int32) *GetImageResponseBodyCroppingSuggestionCroppingBoundary {
	s.Height = &v
	return s
}

func (s *GetImageResponseBodyCroppingSuggestionCroppingBoundary) SetLeft(v int32) *GetImageResponseBodyCroppingSuggestionCroppingBoundary {
	s.Left = &v
	return s
}

func (s *GetImageResponseBodyCroppingSuggestionCroppingBoundary) SetTop(v int32) *GetImageResponseBodyCroppingSuggestionCroppingBoundary {
	s.Top = &v
	return s
}

func (s *GetImageResponseBodyCroppingSuggestionCroppingBoundary) SetWidth(v int32) *GetImageResponseBodyCroppingSuggestionCroppingBoundary {
	s.Width = &v
	return s
}

type GetImageResponseBodyFaces struct {
	Age               *string                                  `json:"Age,omitempty" xml:"Age,omitempty"`
	Attractive        *float32                                 `json:"Attractive,omitempty" xml:"Attractive,omitempty"`
	Emotion           *string                                  `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	EmotionConfidence *float32                                 `json:"EmotionConfidence,omitempty" xml:"EmotionConfidence,omitempty"`
	EmotionDetails    *GetImageResponseBodyFacesEmotionDetails `json:"EmotionDetails,omitempty" xml:"EmotionDetails,omitempty" type:"Struct"`
	FaceAttributes    *GetImageResponseBodyFacesFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" type:"Struct"`
	FaceConfidence    *float32                                 `json:"FaceConfidence,omitempty" xml:"FaceConfidence,omitempty"`
	FaceId            *string                                  `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	FaceQuality       *float32                                 `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	Gender            *string                                  `json:"Gender,omitempty" xml:"Gender,omitempty"`
	GenderConfidence  *float32                                 `json:"GenderConfidence,omitempty" xml:"GenderConfidence,omitempty"`
	GroupId           *string                                  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s GetImageResponseBodyFaces) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyFaces) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyFaces) SetAge(v string) *GetImageResponseBodyFaces {
	s.Age = &v
	return s
}

func (s *GetImageResponseBodyFaces) SetAttractive(v float32) *GetImageResponseBodyFaces {
	s.Attractive = &v
	return s
}

func (s *GetImageResponseBodyFaces) SetEmotion(v string) *GetImageResponseBodyFaces {
	s.Emotion = &v
	return s
}

func (s *GetImageResponseBodyFaces) SetEmotionConfidence(v float32) *GetImageResponseBodyFaces {
	s.EmotionConfidence = &v
	return s
}

func (s *GetImageResponseBodyFaces) SetEmotionDetails(v *GetImageResponseBodyFacesEmotionDetails) *GetImageResponseBodyFaces {
	s.EmotionDetails = v
	return s
}

func (s *GetImageResponseBodyFaces) SetFaceAttributes(v *GetImageResponseBodyFacesFaceAttributes) *GetImageResponseBodyFaces {
	s.FaceAttributes = v
	return s
}

func (s *GetImageResponseBodyFaces) SetFaceConfidence(v float32) *GetImageResponseBodyFaces {
	s.FaceConfidence = &v
	return s
}

func (s *GetImageResponseBodyFaces) SetFaceId(v string) *GetImageResponseBodyFaces {
	s.FaceId = &v
	return s
}

func (s *GetImageResponseBodyFaces) SetFaceQuality(v float32) *GetImageResponseBodyFaces {
	s.FaceQuality = &v
	return s
}

func (s *GetImageResponseBodyFaces) SetGender(v string) *GetImageResponseBodyFaces {
	s.Gender = &v
	return s
}

func (s *GetImageResponseBodyFaces) SetGenderConfidence(v float32) *GetImageResponseBodyFaces {
	s.GenderConfidence = &v
	return s
}

func (s *GetImageResponseBodyFaces) SetGroupId(v string) *GetImageResponseBodyFaces {
	s.GroupId = &v
	return s
}

type GetImageResponseBodyFacesEmotionDetails struct {
	ANGRY     *float32 `json:"ANGRY,omitempty" xml:"ANGRY,omitempty"`
	CALM      *float32 `json:"CALM,omitempty" xml:"CALM,omitempty"`
	DISGUSTED *float32 `json:"DISGUSTED,omitempty" xml:"DISGUSTED,omitempty"`
	HAPPY     *float32 `json:"HAPPY,omitempty" xml:"HAPPY,omitempty"`
	SAD       *float32 `json:"SAD,omitempty" xml:"SAD,omitempty"`
	SCARED    *float32 `json:"SCARED,omitempty" xml:"SCARED,omitempty"`
	SURPRISED *float32 `json:"SURPRISED,omitempty" xml:"SURPRISED,omitempty"`
}

func (s GetImageResponseBodyFacesEmotionDetails) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyFacesEmotionDetails) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyFacesEmotionDetails) SetANGRY(v float32) *GetImageResponseBodyFacesEmotionDetails {
	s.ANGRY = &v
	return s
}

func (s *GetImageResponseBodyFacesEmotionDetails) SetCALM(v float32) *GetImageResponseBodyFacesEmotionDetails {
	s.CALM = &v
	return s
}

func (s *GetImageResponseBodyFacesEmotionDetails) SetDISGUSTED(v float32) *GetImageResponseBodyFacesEmotionDetails {
	s.DISGUSTED = &v
	return s
}

func (s *GetImageResponseBodyFacesEmotionDetails) SetHAPPY(v float32) *GetImageResponseBodyFacesEmotionDetails {
	s.HAPPY = &v
	return s
}

func (s *GetImageResponseBodyFacesEmotionDetails) SetSAD(v float32) *GetImageResponseBodyFacesEmotionDetails {
	s.SAD = &v
	return s
}

func (s *GetImageResponseBodyFacesEmotionDetails) SetSCARED(v float32) *GetImageResponseBodyFacesEmotionDetails {
	s.SCARED = &v
	return s
}

func (s *GetImageResponseBodyFacesEmotionDetails) SetSURPRISED(v float32) *GetImageResponseBodyFacesEmotionDetails {
	s.SURPRISED = &v
	return s
}

type GetImageResponseBodyFacesFaceAttributes struct {
	Beard             *string                                              `json:"Beard,omitempty" xml:"Beard,omitempty"`
	BeardConfidence   *float32                                             `json:"BeardConfidence,omitempty" xml:"BeardConfidence,omitempty"`
	FaceBoundary      *GetImageResponseBodyFacesFaceAttributesFaceBoundary `json:"FaceBoundary,omitempty" xml:"FaceBoundary,omitempty" type:"Struct"`
	Glasses           *string                                              `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	GlassesConfidence *float32                                             `json:"GlassesConfidence,omitempty" xml:"GlassesConfidence,omitempty"`
	HeadPose          *GetImageResponseBodyFacesFaceAttributesHeadPose     `json:"HeadPose,omitempty" xml:"HeadPose,omitempty" type:"Struct"`
	Mask              *string                                              `json:"Mask,omitempty" xml:"Mask,omitempty"`
	MaskConfidence    *float32                                             `json:"MaskConfidence,omitempty" xml:"MaskConfidence,omitempty"`
}

func (s GetImageResponseBodyFacesFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyFacesFaceAttributes) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyFacesFaceAttributes) SetBeard(v string) *GetImageResponseBodyFacesFaceAttributes {
	s.Beard = &v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributes) SetBeardConfidence(v float32) *GetImageResponseBodyFacesFaceAttributes {
	s.BeardConfidence = &v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributes) SetFaceBoundary(v *GetImageResponseBodyFacesFaceAttributesFaceBoundary) *GetImageResponseBodyFacesFaceAttributes {
	s.FaceBoundary = v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributes) SetGlasses(v string) *GetImageResponseBodyFacesFaceAttributes {
	s.Glasses = &v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributes) SetGlassesConfidence(v float32) *GetImageResponseBodyFacesFaceAttributes {
	s.GlassesConfidence = &v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributes) SetHeadPose(v *GetImageResponseBodyFacesFaceAttributesHeadPose) *GetImageResponseBodyFacesFaceAttributes {
	s.HeadPose = v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributes) SetMask(v string) *GetImageResponseBodyFacesFaceAttributes {
	s.Mask = &v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributes) SetMaskConfidence(v float32) *GetImageResponseBodyFacesFaceAttributes {
	s.MaskConfidence = &v
	return s
}

type GetImageResponseBodyFacesFaceAttributesFaceBoundary struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetImageResponseBodyFacesFaceAttributesFaceBoundary) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyFacesFaceAttributesFaceBoundary) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyFacesFaceAttributesFaceBoundary) SetHeight(v int32) *GetImageResponseBodyFacesFaceAttributesFaceBoundary {
	s.Height = &v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributesFaceBoundary) SetLeft(v int32) *GetImageResponseBodyFacesFaceAttributesFaceBoundary {
	s.Left = &v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributesFaceBoundary) SetTop(v int32) *GetImageResponseBodyFacesFaceAttributesFaceBoundary {
	s.Top = &v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributesFaceBoundary) SetWidth(v int32) *GetImageResponseBodyFacesFaceAttributesFaceBoundary {
	s.Width = &v
	return s
}

type GetImageResponseBodyFacesFaceAttributesHeadPose struct {
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	Roll  *float32 `json:"Roll,omitempty" xml:"Roll,omitempty"`
	Yaw   *float32 `json:"Yaw,omitempty" xml:"Yaw,omitempty"`
}

func (s GetImageResponseBodyFacesFaceAttributesHeadPose) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyFacesFaceAttributesHeadPose) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyFacesFaceAttributesHeadPose) SetPitch(v float32) *GetImageResponseBodyFacesFaceAttributesHeadPose {
	s.Pitch = &v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributesHeadPose) SetRoll(v float32) *GetImageResponseBodyFacesFaceAttributesHeadPose {
	s.Roll = &v
	return s
}

func (s *GetImageResponseBodyFacesFaceAttributesHeadPose) SetYaw(v float32) *GetImageResponseBodyFacesFaceAttributesHeadPose {
	s.Yaw = &v
	return s
}

type GetImageResponseBodyImageQuality struct {
	Clarity          *float32 `json:"Clarity,omitempty" xml:"Clarity,omitempty"`
	ClarityScore     *float32 `json:"ClarityScore,omitempty" xml:"ClarityScore,omitempty"`
	Color            *float32 `json:"Color,omitempty" xml:"Color,omitempty"`
	ColorScore       *float32 `json:"ColorScore,omitempty" xml:"ColorScore,omitempty"`
	CompositionScore *float32 `json:"CompositionScore,omitempty" xml:"CompositionScore,omitempty"`
	Contrast         *float32 `json:"Contrast,omitempty" xml:"Contrast,omitempty"`
	ContrastScore    *float32 `json:"ContrastScore,omitempty" xml:"ContrastScore,omitempty"`
	Exposure         *float32 `json:"Exposure,omitempty" xml:"Exposure,omitempty"`
	ExposureScore    *float32 `json:"ExposureScore,omitempty" xml:"ExposureScore,omitempty"`
	OverallScore     *float32 `json:"OverallScore,omitempty" xml:"OverallScore,omitempty"`
}

func (s GetImageResponseBodyImageQuality) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyImageQuality) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImageQuality) SetClarity(v float32) *GetImageResponseBodyImageQuality {
	s.Clarity = &v
	return s
}

func (s *GetImageResponseBodyImageQuality) SetClarityScore(v float32) *GetImageResponseBodyImageQuality {
	s.ClarityScore = &v
	return s
}

func (s *GetImageResponseBodyImageQuality) SetColor(v float32) *GetImageResponseBodyImageQuality {
	s.Color = &v
	return s
}

func (s *GetImageResponseBodyImageQuality) SetColorScore(v float32) *GetImageResponseBodyImageQuality {
	s.ColorScore = &v
	return s
}

func (s *GetImageResponseBodyImageQuality) SetCompositionScore(v float32) *GetImageResponseBodyImageQuality {
	s.CompositionScore = &v
	return s
}

func (s *GetImageResponseBodyImageQuality) SetContrast(v float32) *GetImageResponseBodyImageQuality {
	s.Contrast = &v
	return s
}

func (s *GetImageResponseBodyImageQuality) SetContrastScore(v float32) *GetImageResponseBodyImageQuality {
	s.ContrastScore = &v
	return s
}

func (s *GetImageResponseBodyImageQuality) SetExposure(v float32) *GetImageResponseBodyImageQuality {
	s.Exposure = &v
	return s
}

func (s *GetImageResponseBodyImageQuality) SetExposureScore(v float32) *GetImageResponseBodyImageQuality {
	s.ExposureScore = &v
	return s
}

func (s *GetImageResponseBodyImageQuality) SetOverallScore(v float32) *GetImageResponseBodyImageQuality {
	s.OverallScore = &v
	return s
}

type GetImageResponseBodyOCR struct {
	OCRBoundary   *GetImageResponseBodyOCROCRBoundary `json:"OCRBoundary,omitempty" xml:"OCRBoundary,omitempty" type:"Struct"`
	OCRConfidence *float32                            `json:"OCRConfidence,omitempty" xml:"OCRConfidence,omitempty"`
	OCRContents   *string                             `json:"OCRContents,omitempty" xml:"OCRContents,omitempty"`
}

func (s GetImageResponseBodyOCR) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyOCR) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyOCR) SetOCRBoundary(v *GetImageResponseBodyOCROCRBoundary) *GetImageResponseBodyOCR {
	s.OCRBoundary = v
	return s
}

func (s *GetImageResponseBodyOCR) SetOCRConfidence(v float32) *GetImageResponseBodyOCR {
	s.OCRConfidence = &v
	return s
}

func (s *GetImageResponseBodyOCR) SetOCRContents(v string) *GetImageResponseBodyOCR {
	s.OCRContents = &v
	return s
}

type GetImageResponseBodyOCROCRBoundary struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetImageResponseBodyOCROCRBoundary) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyOCROCRBoundary) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyOCROCRBoundary) SetHeight(v int32) *GetImageResponseBodyOCROCRBoundary {
	s.Height = &v
	return s
}

func (s *GetImageResponseBodyOCROCRBoundary) SetLeft(v int32) *GetImageResponseBodyOCROCRBoundary {
	s.Left = &v
	return s
}

func (s *GetImageResponseBodyOCROCRBoundary) SetTop(v int32) *GetImageResponseBodyOCROCRBoundary {
	s.Top = &v
	return s
}

func (s *GetImageResponseBodyOCROCRBoundary) SetWidth(v int32) *GetImageResponseBodyOCROCRBoundary {
	s.Width = &v
	return s
}

type GetImageResponseBodyTags struct {
	CentricScore  *float32 `json:"CentricScore,omitempty" xml:"CentricScore,omitempty"`
	ParentTagName *string  `json:"ParentTagName,omitempty" xml:"ParentTagName,omitempty"`
	TagConfidence *float32 `json:"TagConfidence,omitempty" xml:"TagConfidence,omitempty"`
	TagLevel      *int32   `json:"TagLevel,omitempty" xml:"TagLevel,omitempty"`
	TagName       *string  `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s GetImageResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyTags) SetCentricScore(v float32) *GetImageResponseBodyTags {
	s.CentricScore = &v
	return s
}

func (s *GetImageResponseBodyTags) SetParentTagName(v string) *GetImageResponseBodyTags {
	s.ParentTagName = &v
	return s
}

func (s *GetImageResponseBodyTags) SetTagConfidence(v float32) *GetImageResponseBodyTags {
	s.TagConfidence = &v
	return s
}

func (s *GetImageResponseBodyTags) SetTagLevel(v int32) *GetImageResponseBodyTags {
	s.TagLevel = &v
	return s
}

func (s *GetImageResponseBodyTags) SetTagName(v string) *GetImageResponseBodyTags {
	s.TagName = &v
	return s
}

type GetImageResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetImageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponse) GoString() string {
	return s.String()
}

func (s *GetImageResponse) SetHeaders(v map[string]*string) *GetImageResponse {
	s.Headers = v
	return s
}

func (s *GetImageResponse) SetStatusCode(v int32) *GetImageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageResponse) SetBody(v *GetImageResponseBody) *GetImageResponse {
	s.Body = v
	return s
}

type GetImageCroppingSuggestionsRequest struct {
	AspectRatios *string `json:"AspectRatios,omitempty" xml:"AspectRatios,omitempty"`
	ImageUri     *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Project      *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s GetImageCroppingSuggestionsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageCroppingSuggestionsRequest) GoString() string {
	return s.String()
}

func (s *GetImageCroppingSuggestionsRequest) SetAspectRatios(v string) *GetImageCroppingSuggestionsRequest {
	s.AspectRatios = &v
	return s
}

func (s *GetImageCroppingSuggestionsRequest) SetImageUri(v string) *GetImageCroppingSuggestionsRequest {
	s.ImageUri = &v
	return s
}

func (s *GetImageCroppingSuggestionsRequest) SetProject(v string) *GetImageCroppingSuggestionsRequest {
	s.Project = &v
	return s
}

type GetImageCroppingSuggestionsResponseBody struct {
	CroppingSuggestions []*GetImageCroppingSuggestionsResponseBodyCroppingSuggestions `json:"CroppingSuggestions,omitempty" xml:"CroppingSuggestions,omitempty" type:"Repeated"`
	ImageUri            *string                                                       `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	RequestId           *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetImageCroppingSuggestionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageCroppingSuggestionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageCroppingSuggestionsResponseBody) SetCroppingSuggestions(v []*GetImageCroppingSuggestionsResponseBodyCroppingSuggestions) *GetImageCroppingSuggestionsResponseBody {
	s.CroppingSuggestions = v
	return s
}

func (s *GetImageCroppingSuggestionsResponseBody) SetImageUri(v string) *GetImageCroppingSuggestionsResponseBody {
	s.ImageUri = &v
	return s
}

func (s *GetImageCroppingSuggestionsResponseBody) SetRequestId(v string) *GetImageCroppingSuggestionsResponseBody {
	s.RequestId = &v
	return s
}

type GetImageCroppingSuggestionsResponseBodyCroppingSuggestions struct {
	AspectRatio      *string                                                                     `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	CroppingBoundary *GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary `json:"CroppingBoundary,omitempty" xml:"CroppingBoundary,omitempty" type:"Struct"`
	Score            *float32                                                                    `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s GetImageCroppingSuggestionsResponseBodyCroppingSuggestions) String() string {
	return tea.Prettify(s)
}

func (s GetImageCroppingSuggestionsResponseBodyCroppingSuggestions) GoString() string {
	return s.String()
}

func (s *GetImageCroppingSuggestionsResponseBodyCroppingSuggestions) SetAspectRatio(v string) *GetImageCroppingSuggestionsResponseBodyCroppingSuggestions {
	s.AspectRatio = &v
	return s
}

func (s *GetImageCroppingSuggestionsResponseBodyCroppingSuggestions) SetCroppingBoundary(v *GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary) *GetImageCroppingSuggestionsResponseBodyCroppingSuggestions {
	s.CroppingBoundary = v
	return s
}

func (s *GetImageCroppingSuggestionsResponseBodyCroppingSuggestions) SetScore(v float32) *GetImageCroppingSuggestionsResponseBodyCroppingSuggestions {
	s.Score = &v
	return s
}

type GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary) String() string {
	return tea.Prettify(s)
}

func (s GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary) GoString() string {
	return s.String()
}

func (s *GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary) SetHeight(v int32) *GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary {
	s.Height = &v
	return s
}

func (s *GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary) SetLeft(v int32) *GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary {
	s.Left = &v
	return s
}

func (s *GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary) SetTop(v int32) *GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary {
	s.Top = &v
	return s
}

func (s *GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary) SetWidth(v int32) *GetImageCroppingSuggestionsResponseBodyCroppingSuggestionsCroppingBoundary {
	s.Width = &v
	return s
}

type GetImageCroppingSuggestionsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetImageCroppingSuggestionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetImageCroppingSuggestionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImageCroppingSuggestionsResponse) GoString() string {
	return s.String()
}

func (s *GetImageCroppingSuggestionsResponse) SetHeaders(v map[string]*string) *GetImageCroppingSuggestionsResponse {
	s.Headers = v
	return s
}

func (s *GetImageCroppingSuggestionsResponse) SetStatusCode(v int32) *GetImageCroppingSuggestionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageCroppingSuggestionsResponse) SetBody(v *GetImageCroppingSuggestionsResponseBody) *GetImageCroppingSuggestionsResponse {
	s.Body = v
	return s
}

type GetImageQualityRequest struct {
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s GetImageQualityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageQualityRequest) GoString() string {
	return s.String()
}

func (s *GetImageQualityRequest) SetImageUri(v string) *GetImageQualityRequest {
	s.ImageUri = &v
	return s
}

func (s *GetImageQualityRequest) SetProject(v string) *GetImageQualityRequest {
	s.Project = &v
	return s
}

type GetImageQualityResponseBody struct {
	ImageQuality *GetImageQualityResponseBodyImageQuality `json:"ImageQuality,omitempty" xml:"ImageQuality,omitempty" type:"Struct"`
	ImageUri     *string                                  `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetImageQualityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageQualityResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageQualityResponseBody) SetImageQuality(v *GetImageQualityResponseBodyImageQuality) *GetImageQualityResponseBody {
	s.ImageQuality = v
	return s
}

func (s *GetImageQualityResponseBody) SetImageUri(v string) *GetImageQualityResponseBody {
	s.ImageUri = &v
	return s
}

func (s *GetImageQualityResponseBody) SetRequestId(v string) *GetImageQualityResponseBody {
	s.RequestId = &v
	return s
}

type GetImageQualityResponseBodyImageQuality struct {
	Clarity          *float32 `json:"Clarity,omitempty" xml:"Clarity,omitempty"`
	ClarityScore     *float32 `json:"ClarityScore,omitempty" xml:"ClarityScore,omitempty"`
	Color            *float32 `json:"Color,omitempty" xml:"Color,omitempty"`
	ColorScore       *float32 `json:"ColorScore,omitempty" xml:"ColorScore,omitempty"`
	CompositionScore *float32 `json:"CompositionScore,omitempty" xml:"CompositionScore,omitempty"`
	Contrast         *float32 `json:"Contrast,omitempty" xml:"Contrast,omitempty"`
	ContrastScore    *float32 `json:"ContrastScore,omitempty" xml:"ContrastScore,omitempty"`
	Exposure         *float32 `json:"Exposure,omitempty" xml:"Exposure,omitempty"`
	ExposureScore    *float32 `json:"ExposureScore,omitempty" xml:"ExposureScore,omitempty"`
	OverallScore     *float32 `json:"OverallScore,omitempty" xml:"OverallScore,omitempty"`
}

func (s GetImageQualityResponseBodyImageQuality) String() string {
	return tea.Prettify(s)
}

func (s GetImageQualityResponseBodyImageQuality) GoString() string {
	return s.String()
}

func (s *GetImageQualityResponseBodyImageQuality) SetClarity(v float32) *GetImageQualityResponseBodyImageQuality {
	s.Clarity = &v
	return s
}

func (s *GetImageQualityResponseBodyImageQuality) SetClarityScore(v float32) *GetImageQualityResponseBodyImageQuality {
	s.ClarityScore = &v
	return s
}

func (s *GetImageQualityResponseBodyImageQuality) SetColor(v float32) *GetImageQualityResponseBodyImageQuality {
	s.Color = &v
	return s
}

func (s *GetImageQualityResponseBodyImageQuality) SetColorScore(v float32) *GetImageQualityResponseBodyImageQuality {
	s.ColorScore = &v
	return s
}

func (s *GetImageQualityResponseBodyImageQuality) SetCompositionScore(v float32) *GetImageQualityResponseBodyImageQuality {
	s.CompositionScore = &v
	return s
}

func (s *GetImageQualityResponseBodyImageQuality) SetContrast(v float32) *GetImageQualityResponseBodyImageQuality {
	s.Contrast = &v
	return s
}

func (s *GetImageQualityResponseBodyImageQuality) SetContrastScore(v float32) *GetImageQualityResponseBodyImageQuality {
	s.ContrastScore = &v
	return s
}

func (s *GetImageQualityResponseBodyImageQuality) SetExposure(v float32) *GetImageQualityResponseBodyImageQuality {
	s.Exposure = &v
	return s
}

func (s *GetImageQualityResponseBodyImageQuality) SetExposureScore(v float32) *GetImageQualityResponseBodyImageQuality {
	s.ExposureScore = &v
	return s
}

func (s *GetImageQualityResponseBodyImageQuality) SetOverallScore(v float32) *GetImageQualityResponseBodyImageQuality {
	s.OverallScore = &v
	return s
}

type GetImageQualityResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetImageQualityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetImageQualityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImageQualityResponse) GoString() string {
	return s.String()
}

func (s *GetImageQualityResponse) SetHeaders(v map[string]*string) *GetImageQualityResponse {
	s.Headers = v
	return s
}

func (s *GetImageQualityResponse) SetStatusCode(v int32) *GetImageQualityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageQualityResponse) SetBody(v *GetImageQualityResponseBody) *GetImageQualityResponse {
	s.Body = v
	return s
}

type GetMediaMetaRequest struct {
	MediaUri *string `json:"MediaUri,omitempty" xml:"MediaUri,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s GetMediaMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaRequest) GoString() string {
	return s.String()
}

func (s *GetMediaMetaRequest) SetMediaUri(v string) *GetMediaMetaRequest {
	s.MediaUri = &v
	return s
}

func (s *GetMediaMetaRequest) SetProject(v string) *GetMediaMetaRequest {
	s.Project = &v
	return s
}

type GetMediaMetaResponseBody struct {
	MediaMeta *GetMediaMetaResponseBodyMediaMeta `json:"MediaMeta,omitempty" xml:"MediaMeta,omitempty" type:"Struct"`
	MediaUri  *string                            `json:"MediaUri,omitempty" xml:"MediaUri,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponseBody) SetMediaMeta(v *GetMediaMetaResponseBodyMediaMeta) *GetMediaMetaResponseBody {
	s.MediaMeta = v
	return s
}

func (s *GetMediaMetaResponseBody) SetMediaUri(v string) *GetMediaMetaResponseBody {
	s.MediaUri = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetRequestId(v string) *GetMediaMetaResponseBody {
	s.RequestId = &v
	return s
}

type GetMediaMetaResponseBodyMediaMeta struct {
	MediaFormat  *GetMediaMetaResponseBodyMediaMetaMediaFormat  `json:"MediaFormat,omitempty" xml:"MediaFormat,omitempty" type:"Struct"`
	MediaStreams *GetMediaMetaResponseBodyMediaMetaMediaStreams `json:"MediaStreams,omitempty" xml:"MediaStreams,omitempty" type:"Struct"`
}

func (s GetMediaMetaResponseBodyMediaMeta) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponseBodyMediaMeta) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponseBodyMediaMeta) SetMediaFormat(v *GetMediaMetaResponseBodyMediaMetaMediaFormat) *GetMediaMetaResponseBodyMediaMeta {
	s.MediaFormat = v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMeta) SetMediaStreams(v *GetMediaMetaResponseBodyMediaMetaMediaStreams) *GetMediaMetaResponseBodyMediaMeta {
	s.MediaStreams = v
	return s
}

type GetMediaMetaResponseBodyMediaMetaMediaFormat struct {
	Address        *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress `json:"Address,omitempty" xml:"Address,omitempty" type:"Struct"`
	Bitrate        *string                                              `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	CreationTime   *string                                              `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Duration       *string                                              `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FormatLongName *string                                              `json:"FormatLongName,omitempty" xml:"FormatLongName,omitempty"`
	FormatName     *string                                              `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	Location       *string                                              `json:"Location,omitempty" xml:"Location,omitempty"`
	NumberPrograms *int32                                               `json:"NumberPrograms,omitempty" xml:"NumberPrograms,omitempty"`
	NumberStreams  *int32                                               `json:"NumberStreams,omitempty" xml:"NumberStreams,omitempty"`
	Size           *string                                              `json:"Size,omitempty" xml:"Size,omitempty"`
	StartTime      *string                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Tag            *GetMediaMetaResponseBodyMediaMetaMediaFormatTag     `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Struct"`
}

func (s GetMediaMetaResponseBodyMediaMetaMediaFormat) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponseBodyMediaMetaMediaFormat) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetAddress(v *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.Address = v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetBitrate(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.Bitrate = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetCreationTime(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.CreationTime = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetDuration(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.Duration = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetFormatLongName(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.FormatLongName = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetFormatName(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.FormatName = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetLocation(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.Location = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetNumberPrograms(v int32) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.NumberPrograms = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetNumberStreams(v int32) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.NumberStreams = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetSize(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.Size = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetStartTime(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.StartTime = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormat) SetTag(v *GetMediaMetaResponseBodyMediaMetaMediaFormatTag) *GetMediaMetaResponseBodyMediaMetaMediaFormat {
	s.Tag = v
	return s
}

type GetMediaMetaResponseBodyMediaMetaMediaFormatAddress struct {
	AddressLine *string `json:"AddressLine,omitempty" xml:"AddressLine,omitempty"`
	City        *string `json:"City,omitempty" xml:"City,omitempty"`
	Country     *string `json:"Country,omitempty" xml:"Country,omitempty"`
	District    *string `json:"District,omitempty" xml:"District,omitempty"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty"`
	Township    *string `json:"Township,omitempty" xml:"Township,omitempty"`
}

func (s GetMediaMetaResponseBodyMediaMetaMediaFormatAddress) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponseBodyMediaMetaMediaFormatAddress) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress) SetAddressLine(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress {
	s.AddressLine = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress) SetCity(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress {
	s.City = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress) SetCountry(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress {
	s.Country = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress) SetDistrict(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress {
	s.District = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress) SetProvince(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress {
	s.Province = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress) SetTownship(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatAddress {
	s.Township = &v
	return s
}

type GetMediaMetaResponseBodyMediaMetaMediaFormatTag struct {
	Album        *string `json:"Album,omitempty" xml:"Album,omitempty"`
	AlbumArtist  *string `json:"AlbumArtist,omitempty" xml:"AlbumArtist,omitempty"`
	Artist       *string `json:"Artist,omitempty" xml:"Artist,omitempty"`
	Composer     *string `json:"Composer,omitempty" xml:"Composer,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Language     *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Performer    *string `json:"Performer,omitempty" xml:"Performer,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetMediaMetaResponseBodyMediaMetaMediaFormatTag) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponseBodyMediaMetaMediaFormatTag) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatTag) SetAlbum(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatTag {
	s.Album = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatTag) SetAlbumArtist(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatTag {
	s.AlbumArtist = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatTag) SetArtist(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatTag {
	s.Artist = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatTag) SetComposer(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatTag {
	s.Composer = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatTag) SetCreationTime(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatTag {
	s.CreationTime = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatTag) SetLanguage(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatTag {
	s.Language = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatTag) SetPerformer(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatTag {
	s.Performer = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaFormatTag) SetTitle(v string) *GetMediaMetaResponseBodyMediaMetaMediaFormatTag {
	s.Title = &v
	return s
}

type GetMediaMetaResponseBodyMediaMetaMediaStreams struct {
	AudioStreams    []*GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams    `json:"AudioStreams,omitempty" xml:"AudioStreams,omitempty" type:"Repeated"`
	SubtitleStreams []*GetMediaMetaResponseBodyMediaMetaMediaStreamsSubtitleStreams `json:"SubtitleStreams,omitempty" xml:"SubtitleStreams,omitempty" type:"Repeated"`
	VideoStreams    []*GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams    `json:"VideoStreams,omitempty" xml:"VideoStreams,omitempty" type:"Repeated"`
}

func (s GetMediaMetaResponseBodyMediaMetaMediaStreams) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponseBodyMediaMetaMediaStreams) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreams) SetAudioStreams(v []*GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) *GetMediaMetaResponseBodyMediaMetaMediaStreams {
	s.AudioStreams = v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreams) SetSubtitleStreams(v []*GetMediaMetaResponseBodyMediaMetaMediaStreamsSubtitleStreams) *GetMediaMetaResponseBodyMediaMetaMediaStreams {
	s.SubtitleStreams = v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreams) SetVideoStreams(v []*GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) *GetMediaMetaResponseBodyMediaMetaMediaStreams {
	s.VideoStreams = v
	return s
}

type GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams struct {
	Bitrate        *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	ChannelLayout  *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	Channels       *int32  `json:"Channels,omitempty" xml:"Channels,omitempty"`
	CodecLongName  *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	CodecName      *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	CodecTag       *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	CodecTimeBase  *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	Duration       *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Frames         *string `json:"Frames,omitempty" xml:"Frames,omitempty"`
	Index          *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Language       *string `json:"Language,omitempty" xml:"Language,omitempty"`
	SampleFormat   *string `json:"SampleFormat,omitempty" xml:"SampleFormat,omitempty"`
	SampleRate     *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeBase       *string `json:"TimeBase,omitempty" xml:"TimeBase,omitempty"`
}

func (s GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetBitrate(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.Bitrate = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetChannelLayout(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.ChannelLayout = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetChannels(v int32) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.Channels = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetCodecLongName(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.CodecLongName = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetCodecName(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.CodecName = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetCodecTag(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.CodecTag = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetCodecTagString(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.CodecTagString = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetCodecTimeBase(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetDuration(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.Duration = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetFrames(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.Frames = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetIndex(v int32) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.Index = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetLanguage(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.Language = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetSampleFormat(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.SampleFormat = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetSampleRate(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.SampleRate = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetStartTime(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.StartTime = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams) SetTimeBase(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsAudioStreams {
	s.TimeBase = &v
	return s
}

type GetMediaMetaResponseBodyMediaMetaMediaStreamsSubtitleStreams struct {
	Index    *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s GetMediaMetaResponseBodyMediaMetaMediaStreamsSubtitleStreams) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponseBodyMediaMetaMediaStreamsSubtitleStreams) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsSubtitleStreams) SetIndex(v int32) *GetMediaMetaResponseBodyMediaMetaMediaStreamsSubtitleStreams {
	s.Index = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsSubtitleStreams) SetLanguage(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsSubtitleStreams {
	s.Language = &v
	return s
}

type GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams struct {
	AverageFrameRate   *string `json:"AverageFrameRate,omitempty" xml:"AverageFrameRate,omitempty"`
	Bitrate            *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	CodecLongName      *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	CodecName          *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	CodecTag           *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	CodecTagString     *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	CodecTimeBase      *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	DisplayAspectRatio *string `json:"DisplayAspectRatio,omitempty" xml:"DisplayAspectRatio,omitempty"`
	Duration           *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FrameRrate         *string `json:"FrameRrate,omitempty" xml:"FrameRrate,omitempty"`
	Frames             *string `json:"Frames,omitempty" xml:"Frames,omitempty"`
	HasBFrames         *int32  `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	Height             *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	Index              *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Language           *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Level              *int32  `json:"Level,omitempty" xml:"Level,omitempty"`
	PixelFormat        *string `json:"PixelFormat,omitempty" xml:"PixelFormat,omitempty"`
	Profile            *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Rotate             *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	SampleAspectRatio  *string `json:"SampleAspectRatio,omitempty" xml:"SampleAspectRatio,omitempty"`
	StartTime          *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeBase           *string `json:"TimeBase,omitempty" xml:"TimeBase,omitempty"`
	Width              *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetAverageFrameRate(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.AverageFrameRate = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetBitrate(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.Bitrate = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetCodecLongName(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.CodecLongName = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetCodecName(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.CodecName = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetCodecTag(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.CodecTag = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetCodecTagString(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.CodecTagString = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetCodecTimeBase(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetDisplayAspectRatio(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.DisplayAspectRatio = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetDuration(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.Duration = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetFrameRrate(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.FrameRrate = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetFrames(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.Frames = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetHasBFrames(v int32) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.HasBFrames = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetHeight(v int32) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.Height = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetIndex(v int32) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.Index = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetLanguage(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.Language = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetLevel(v int32) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.Level = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetPixelFormat(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.PixelFormat = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetProfile(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.Profile = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetRotate(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.Rotate = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetSampleAspectRatio(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.SampleAspectRatio = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetStartTime(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.StartTime = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetTimeBase(v string) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.TimeBase = &v
	return s
}

func (s *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams) SetWidth(v int32) *GetMediaMetaResponseBodyMediaMetaMediaStreamsVideoStreams {
	s.Width = &v
	return s
}

type GetMediaMetaResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMediaMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMediaMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponse) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponse) SetHeaders(v map[string]*string) *GetMediaMetaResponse {
	s.Headers = v
	return s
}

func (s *GetMediaMetaResponse) SetStatusCode(v int32) *GetMediaMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaMetaResponse) SetBody(v *GetMediaMetaResponseBody) *GetMediaMetaResponse {
	s.Body = v
	return s
}

type GetOfficeConversionTaskRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetOfficeConversionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOfficeConversionTaskRequest) GoString() string {
	return s.String()
}

func (s *GetOfficeConversionTaskRequest) SetProject(v string) *GetOfficeConversionTaskRequest {
	s.Project = &v
	return s
}

func (s *GetOfficeConversionTaskRequest) SetTaskId(v string) *GetOfficeConversionTaskRequest {
	s.TaskId = &v
	return s
}

type GetOfficeConversionTaskResponseBody struct {
	CreateTime      *string                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExternalID      *string                                        `json:"ExternalID,omitempty" xml:"ExternalID,omitempty"`
	FailDetail      *GetOfficeConversionTaskResponseBodyFailDetail `json:"FailDetail,omitempty" xml:"FailDetail,omitempty" type:"Struct"`
	FinishTime      *string                                        `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	ImageSpec       *string                                        `json:"ImageSpec,omitempty" xml:"ImageSpec,omitempty"`
	NotifyEndpoint  *string                                        `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	NotifyTopicName *string                                        `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	PageCount       *int32                                         `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	Percent         *int32                                         `json:"Percent,omitempty" xml:"Percent,omitempty"`
	RequestId       *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SrcUri          *string                                        `json:"SrcUri,omitempty" xml:"SrcUri,omitempty"`
	Status          *string                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId          *string                                        `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TgtType         *string                                        `json:"TgtType,omitempty" xml:"TgtType,omitempty"`
	TgtUri          *string                                        `json:"TgtUri,omitempty" xml:"TgtUri,omitempty"`
}

func (s GetOfficeConversionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOfficeConversionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetOfficeConversionTaskResponseBody) SetCreateTime(v string) *GetOfficeConversionTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetExternalID(v string) *GetOfficeConversionTaskResponseBody {
	s.ExternalID = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetFailDetail(v *GetOfficeConversionTaskResponseBodyFailDetail) *GetOfficeConversionTaskResponseBody {
	s.FailDetail = v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetFinishTime(v string) *GetOfficeConversionTaskResponseBody {
	s.FinishTime = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetImageSpec(v string) *GetOfficeConversionTaskResponseBody {
	s.ImageSpec = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetNotifyEndpoint(v string) *GetOfficeConversionTaskResponseBody {
	s.NotifyEndpoint = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetNotifyTopicName(v string) *GetOfficeConversionTaskResponseBody {
	s.NotifyTopicName = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetPageCount(v int32) *GetOfficeConversionTaskResponseBody {
	s.PageCount = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetPercent(v int32) *GetOfficeConversionTaskResponseBody {
	s.Percent = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetRequestId(v string) *GetOfficeConversionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetSrcUri(v string) *GetOfficeConversionTaskResponseBody {
	s.SrcUri = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetStatus(v string) *GetOfficeConversionTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetTaskId(v string) *GetOfficeConversionTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetTgtType(v string) *GetOfficeConversionTaskResponseBody {
	s.TgtType = &v
	return s
}

func (s *GetOfficeConversionTaskResponseBody) SetTgtUri(v string) *GetOfficeConversionTaskResponseBody {
	s.TgtUri = &v
	return s
}

type GetOfficeConversionTaskResponseBodyFailDetail struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetOfficeConversionTaskResponseBodyFailDetail) String() string {
	return tea.Prettify(s)
}

func (s GetOfficeConversionTaskResponseBodyFailDetail) GoString() string {
	return s.String()
}

func (s *GetOfficeConversionTaskResponseBodyFailDetail) SetCode(v string) *GetOfficeConversionTaskResponseBodyFailDetail {
	s.Code = &v
	return s
}

type GetOfficeConversionTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOfficeConversionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOfficeConversionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOfficeConversionTaskResponse) GoString() string {
	return s.String()
}

func (s *GetOfficeConversionTaskResponse) SetHeaders(v map[string]*string) *GetOfficeConversionTaskResponse {
	s.Headers = v
	return s
}

func (s *GetOfficeConversionTaskResponse) SetStatusCode(v int32) *GetOfficeConversionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOfficeConversionTaskResponse) SetBody(v *GetOfficeConversionTaskResponseBody) *GetOfficeConversionTaskResponse {
	s.Body = v
	return s
}

type GetOfficePreviewURLRequest struct {
	Project             *string  `json:"Project,omitempty" xml:"Project,omitempty"`
	SrcType             *string  `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	SrcUri              *string  `json:"SrcUri,omitempty" xml:"SrcUri,omitempty"`
	WatermarkFillStyle  *string  `json:"WatermarkFillStyle,omitempty" xml:"WatermarkFillStyle,omitempty"`
	WatermarkFont       *string  `json:"WatermarkFont,omitempty" xml:"WatermarkFont,omitempty"`
	WatermarkHorizontal *int32   `json:"WatermarkHorizontal,omitempty" xml:"WatermarkHorizontal,omitempty"`
	WatermarkRotate     *float32 `json:"WatermarkRotate,omitempty" xml:"WatermarkRotate,omitempty"`
	WatermarkType       *int32   `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	WatermarkValue      *string  `json:"WatermarkValue,omitempty" xml:"WatermarkValue,omitempty"`
	WatermarkVertical   *int32   `json:"WatermarkVertical,omitempty" xml:"WatermarkVertical,omitempty"`
}

func (s GetOfficePreviewURLRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOfficePreviewURLRequest) GoString() string {
	return s.String()
}

func (s *GetOfficePreviewURLRequest) SetProject(v string) *GetOfficePreviewURLRequest {
	s.Project = &v
	return s
}

func (s *GetOfficePreviewURLRequest) SetSrcType(v string) *GetOfficePreviewURLRequest {
	s.SrcType = &v
	return s
}

func (s *GetOfficePreviewURLRequest) SetSrcUri(v string) *GetOfficePreviewURLRequest {
	s.SrcUri = &v
	return s
}

func (s *GetOfficePreviewURLRequest) SetWatermarkFillStyle(v string) *GetOfficePreviewURLRequest {
	s.WatermarkFillStyle = &v
	return s
}

func (s *GetOfficePreviewURLRequest) SetWatermarkFont(v string) *GetOfficePreviewURLRequest {
	s.WatermarkFont = &v
	return s
}

func (s *GetOfficePreviewURLRequest) SetWatermarkHorizontal(v int32) *GetOfficePreviewURLRequest {
	s.WatermarkHorizontal = &v
	return s
}

func (s *GetOfficePreviewURLRequest) SetWatermarkRotate(v float32) *GetOfficePreviewURLRequest {
	s.WatermarkRotate = &v
	return s
}

func (s *GetOfficePreviewURLRequest) SetWatermarkType(v int32) *GetOfficePreviewURLRequest {
	s.WatermarkType = &v
	return s
}

func (s *GetOfficePreviewURLRequest) SetWatermarkValue(v string) *GetOfficePreviewURLRequest {
	s.WatermarkValue = &v
	return s
}

func (s *GetOfficePreviewURLRequest) SetWatermarkVertical(v int32) *GetOfficePreviewURLRequest {
	s.WatermarkVertical = &v
	return s
}

type GetOfficePreviewURLResponseBody struct {
	AccessToken             *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	AccessTokenExpiredTime  *string `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
	PreviewURL              *string `json:"PreviewURL,omitempty" xml:"PreviewURL,omitempty"`
	RefreshToken            *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	RefreshTokenExpiredTime *string `json:"RefreshTokenExpiredTime,omitempty" xml:"RefreshTokenExpiredTime,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOfficePreviewURLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOfficePreviewURLResponseBody) GoString() string {
	return s.String()
}

func (s *GetOfficePreviewURLResponseBody) SetAccessToken(v string) *GetOfficePreviewURLResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GetOfficePreviewURLResponseBody) SetAccessTokenExpiredTime(v string) *GetOfficePreviewURLResponseBody {
	s.AccessTokenExpiredTime = &v
	return s
}

func (s *GetOfficePreviewURLResponseBody) SetPreviewURL(v string) *GetOfficePreviewURLResponseBody {
	s.PreviewURL = &v
	return s
}

func (s *GetOfficePreviewURLResponseBody) SetRefreshToken(v string) *GetOfficePreviewURLResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *GetOfficePreviewURLResponseBody) SetRefreshTokenExpiredTime(v string) *GetOfficePreviewURLResponseBody {
	s.RefreshTokenExpiredTime = &v
	return s
}

func (s *GetOfficePreviewURLResponseBody) SetRequestId(v string) *GetOfficePreviewURLResponseBody {
	s.RequestId = &v
	return s
}

type GetOfficePreviewURLResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOfficePreviewURLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOfficePreviewURLResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOfficePreviewURLResponse) GoString() string {
	return s.String()
}

func (s *GetOfficePreviewURLResponse) SetHeaders(v map[string]*string) *GetOfficePreviewURLResponse {
	s.Headers = v
	return s
}

func (s *GetOfficePreviewURLResponse) SetStatusCode(v int32) *GetOfficePreviewURLResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOfficePreviewURLResponse) SetBody(v *GetOfficePreviewURLResponseBody) *GetOfficePreviewURLResponse {
	s.Body = v
	return s
}

type GetProjectRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s GetProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectRequest) GoString() string {
	return s.String()
}

func (s *GetProjectRequest) SetProject(v string) *GetProjectRequest {
	s.Project = &v
	return s
}

type GetProjectResponseBody struct {
	BillingType *string `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	CU          *int32  `json:"CU,omitempty" xml:"CU,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Endpoint    *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ModifyTime  *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Project     *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) SetBillingType(v string) *GetProjectResponseBody {
	s.BillingType = &v
	return s
}

func (s *GetProjectResponseBody) SetCU(v int32) *GetProjectResponseBody {
	s.CU = &v
	return s
}

func (s *GetProjectResponseBody) SetCreateTime(v string) *GetProjectResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetProjectResponseBody) SetEndpoint(v string) *GetProjectResponseBody {
	s.Endpoint = &v
	return s
}

func (s *GetProjectResponseBody) SetModifyTime(v string) *GetProjectResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *GetProjectResponseBody) SetProject(v string) *GetProjectResponseBody {
	s.Project = &v
	return s
}

func (s *GetProjectResponseBody) SetRegionId(v string) *GetProjectResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetProjectResponseBody) SetRequestId(v string) *GetProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectResponseBody) SetServiceRole(v string) *GetProjectResponseBody {
	s.ServiceRole = &v
	return s
}

func (s *GetProjectResponseBody) SetType(v string) *GetProjectResponseBody {
	s.Type = &v
	return s
}

type GetProjectResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponse) GoString() string {
	return s.String()
}

func (s *GetProjectResponse) SetHeaders(v map[string]*string) *GetProjectResponse {
	s.Headers = v
	return s
}

func (s *GetProjectResponse) SetStatusCode(v int32) *GetProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectResponse) SetBody(v *GetProjectResponseBody) *GetProjectResponse {
	s.Body = v
	return s
}

type GetSetRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId   *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s GetSetRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSetRequest) GoString() string {
	return s.String()
}

func (s *GetSetRequest) SetProject(v string) *GetSetRequest {
	s.Project = &v
	return s
}

func (s *GetSetRequest) SetSetId(v string) *GetSetRequest {
	s.SetId = &v
	return s
}

type GetSetResponseBody struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FaceCount   *int32  `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	ImageCount  *int32  `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	ModifyTime  *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId       *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	SetName     *string `json:"SetName,omitempty" xml:"SetName,omitempty"`
	VideoCount  *int32  `json:"VideoCount,omitempty" xml:"VideoCount,omitempty"`
	VideoLength *int32  `json:"VideoLength,omitempty" xml:"VideoLength,omitempty"`
}

func (s GetSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSetResponseBody) GoString() string {
	return s.String()
}

func (s *GetSetResponseBody) SetCreateTime(v string) *GetSetResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetSetResponseBody) SetFaceCount(v int32) *GetSetResponseBody {
	s.FaceCount = &v
	return s
}

func (s *GetSetResponseBody) SetImageCount(v int32) *GetSetResponseBody {
	s.ImageCount = &v
	return s
}

func (s *GetSetResponseBody) SetModifyTime(v string) *GetSetResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *GetSetResponseBody) SetRequestId(v string) *GetSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSetResponseBody) SetSetId(v string) *GetSetResponseBody {
	s.SetId = &v
	return s
}

func (s *GetSetResponseBody) SetSetName(v string) *GetSetResponseBody {
	s.SetName = &v
	return s
}

func (s *GetSetResponseBody) SetVideoCount(v int32) *GetSetResponseBody {
	s.VideoCount = &v
	return s
}

func (s *GetSetResponseBody) SetVideoLength(v int32) *GetSetResponseBody {
	s.VideoLength = &v
	return s
}

type GetSetResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSetResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSetResponse) GoString() string {
	return s.String()
}

func (s *GetSetResponse) SetHeaders(v map[string]*string) *GetSetResponse {
	s.Headers = v
	return s
}

func (s *GetSetResponse) SetStatusCode(v int32) *GetSetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSetResponse) SetBody(v *GetSetResponseBody) *GetSetResponse {
	s.Body = v
	return s
}

type GetVideoRequest struct {
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId    *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	VideoUri *string `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
}

func (s GetVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVideoRequest) GoString() string {
	return s.String()
}

func (s *GetVideoRequest) SetProject(v string) *GetVideoRequest {
	s.Project = &v
	return s
}

func (s *GetVideoRequest) SetSetId(v string) *GetVideoRequest {
	s.SetId = &v
	return s
}

func (s *GetVideoRequest) SetVideoUri(v string) *GetVideoRequest {
	s.VideoUri = &v
	return s
}

type GetVideoResponseBody struct {
	CreateTime               *string                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExternalId               *string                          `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	FileSize                 *int32                           `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	ModifyTime               *string                          `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ProcessFailReason        *string                          `json:"ProcessFailReason,omitempty" xml:"ProcessFailReason,omitempty"`
	ProcessModifyTime        *string                          `json:"ProcessModifyTime,omitempty" xml:"ProcessModifyTime,omitempty"`
	ProcessStatus            *string                          `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	RemarksA                 *string                          `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	RemarksB                 *string                          `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	RemarksC                 *string                          `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD                 *string                          `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	RequestId                *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId                    *string                          `json:"SetId,omitempty" xml:"SetId,omitempty"`
	SourcePosition           *string                          `json:"SourcePosition,omitempty" xml:"SourcePosition,omitempty"`
	SourceType               *string                          `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SourceUri                *string                          `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
	VideoDuration            *float32                         `json:"VideoDuration,omitempty" xml:"VideoDuration,omitempty"`
	VideoFacesFailReason     *string                          `json:"VideoFacesFailReason,omitempty" xml:"VideoFacesFailReason,omitempty"`
	VideoFacesModifyTime     *string                          `json:"VideoFacesModifyTime,omitempty" xml:"VideoFacesModifyTime,omitempty"`
	VideoFacesStatus         *string                          `json:"VideoFacesStatus,omitempty" xml:"VideoFacesStatus,omitempty"`
	VideoFormat              *string                          `json:"VideoFormat,omitempty" xml:"VideoFormat,omitempty"`
	VideoFrameTagsFailReason *string                          `json:"VideoFrameTagsFailReason,omitempty" xml:"VideoFrameTagsFailReason,omitempty"`
	VideoFrameTagsModifyTime *string                          `json:"VideoFrameTagsModifyTime,omitempty" xml:"VideoFrameTagsModifyTime,omitempty"`
	VideoFrameTagsStatus     *string                          `json:"VideoFrameTagsStatus,omitempty" xml:"VideoFrameTagsStatus,omitempty"`
	VideoFrames              *int32                           `json:"VideoFrames,omitempty" xml:"VideoFrames,omitempty"`
	VideoHeight              *int32                           `json:"VideoHeight,omitempty" xml:"VideoHeight,omitempty"`
	VideoInfo                *string                          `json:"VideoInfo,omitempty" xml:"VideoInfo,omitempty"`
	VideoOCRFailReason       *string                          `json:"VideoOCRFailReason,omitempty" xml:"VideoOCRFailReason,omitempty"`
	VideoOCRModifyTime       *string                          `json:"VideoOCRModifyTime,omitempty" xml:"VideoOCRModifyTime,omitempty"`
	VideoOCRStatus           *string                          `json:"VideoOCRStatus,omitempty" xml:"VideoOCRStatus,omitempty"`
	VideoSTTFailReason       *string                          `json:"VideoSTTFailReason,omitempty" xml:"VideoSTTFailReason,omitempty"`
	VideoSTTModifyTime       *string                          `json:"VideoSTTModifyTime,omitempty" xml:"VideoSTTModifyTime,omitempty"`
	VideoSTTStatus           *string                          `json:"VideoSTTStatus,omitempty" xml:"VideoSTTStatus,omitempty"`
	VideoTags                []*GetVideoResponseBodyVideoTags `json:"VideoTags,omitempty" xml:"VideoTags,omitempty" type:"Repeated"`
	VideoTagsFailReason      *string                          `json:"VideoTagsFailReason,omitempty" xml:"VideoTagsFailReason,omitempty"`
	VideoTagsModifyTime      *string                          `json:"VideoTagsModifyTime,omitempty" xml:"VideoTagsModifyTime,omitempty"`
	VideoTagsStatus          *string                          `json:"VideoTagsStatus,omitempty" xml:"VideoTagsStatus,omitempty"`
	VideoUri                 *string                          `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
	VideoWidth               *int32                           `json:"VideoWidth,omitempty" xml:"VideoWidth,omitempty"`
}

func (s GetVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVideoResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoResponseBody) SetCreateTime(v string) *GetVideoResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetVideoResponseBody) SetExternalId(v string) *GetVideoResponseBody {
	s.ExternalId = &v
	return s
}

func (s *GetVideoResponseBody) SetFileSize(v int32) *GetVideoResponseBody {
	s.FileSize = &v
	return s
}

func (s *GetVideoResponseBody) SetModifyTime(v string) *GetVideoResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *GetVideoResponseBody) SetProcessFailReason(v string) *GetVideoResponseBody {
	s.ProcessFailReason = &v
	return s
}

func (s *GetVideoResponseBody) SetProcessModifyTime(v string) *GetVideoResponseBody {
	s.ProcessModifyTime = &v
	return s
}

func (s *GetVideoResponseBody) SetProcessStatus(v string) *GetVideoResponseBody {
	s.ProcessStatus = &v
	return s
}

func (s *GetVideoResponseBody) SetRemarksA(v string) *GetVideoResponseBody {
	s.RemarksA = &v
	return s
}

func (s *GetVideoResponseBody) SetRemarksB(v string) *GetVideoResponseBody {
	s.RemarksB = &v
	return s
}

func (s *GetVideoResponseBody) SetRemarksC(v string) *GetVideoResponseBody {
	s.RemarksC = &v
	return s
}

func (s *GetVideoResponseBody) SetRemarksD(v string) *GetVideoResponseBody {
	s.RemarksD = &v
	return s
}

func (s *GetVideoResponseBody) SetRequestId(v string) *GetVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoResponseBody) SetSetId(v string) *GetVideoResponseBody {
	s.SetId = &v
	return s
}

func (s *GetVideoResponseBody) SetSourcePosition(v string) *GetVideoResponseBody {
	s.SourcePosition = &v
	return s
}

func (s *GetVideoResponseBody) SetSourceType(v string) *GetVideoResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetVideoResponseBody) SetSourceUri(v string) *GetVideoResponseBody {
	s.SourceUri = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoDuration(v float32) *GetVideoResponseBody {
	s.VideoDuration = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoFacesFailReason(v string) *GetVideoResponseBody {
	s.VideoFacesFailReason = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoFacesModifyTime(v string) *GetVideoResponseBody {
	s.VideoFacesModifyTime = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoFacesStatus(v string) *GetVideoResponseBody {
	s.VideoFacesStatus = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoFormat(v string) *GetVideoResponseBody {
	s.VideoFormat = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoFrameTagsFailReason(v string) *GetVideoResponseBody {
	s.VideoFrameTagsFailReason = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoFrameTagsModifyTime(v string) *GetVideoResponseBody {
	s.VideoFrameTagsModifyTime = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoFrameTagsStatus(v string) *GetVideoResponseBody {
	s.VideoFrameTagsStatus = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoFrames(v int32) *GetVideoResponseBody {
	s.VideoFrames = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoHeight(v int32) *GetVideoResponseBody {
	s.VideoHeight = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoInfo(v string) *GetVideoResponseBody {
	s.VideoInfo = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoOCRFailReason(v string) *GetVideoResponseBody {
	s.VideoOCRFailReason = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoOCRModifyTime(v string) *GetVideoResponseBody {
	s.VideoOCRModifyTime = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoOCRStatus(v string) *GetVideoResponseBody {
	s.VideoOCRStatus = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoSTTFailReason(v string) *GetVideoResponseBody {
	s.VideoSTTFailReason = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoSTTModifyTime(v string) *GetVideoResponseBody {
	s.VideoSTTModifyTime = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoSTTStatus(v string) *GetVideoResponseBody {
	s.VideoSTTStatus = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoTags(v []*GetVideoResponseBodyVideoTags) *GetVideoResponseBody {
	s.VideoTags = v
	return s
}

func (s *GetVideoResponseBody) SetVideoTagsFailReason(v string) *GetVideoResponseBody {
	s.VideoTagsFailReason = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoTagsModifyTime(v string) *GetVideoResponseBody {
	s.VideoTagsModifyTime = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoTagsStatus(v string) *GetVideoResponseBody {
	s.VideoTagsStatus = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoUri(v string) *GetVideoResponseBody {
	s.VideoUri = &v
	return s
}

func (s *GetVideoResponseBody) SetVideoWidth(v int32) *GetVideoResponseBody {
	s.VideoWidth = &v
	return s
}

type GetVideoResponseBodyVideoTags struct {
	ParentTagName *string  `json:"ParentTagName,omitempty" xml:"ParentTagName,omitempty"`
	TagConfidence *float32 `json:"TagConfidence,omitempty" xml:"TagConfidence,omitempty"`
	TagLevel      *int32   `json:"TagLevel,omitempty" xml:"TagLevel,omitempty"`
	TagName       *string  `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s GetVideoResponseBodyVideoTags) String() string {
	return tea.Prettify(s)
}

func (s GetVideoResponseBodyVideoTags) GoString() string {
	return s.String()
}

func (s *GetVideoResponseBodyVideoTags) SetParentTagName(v string) *GetVideoResponseBodyVideoTags {
	s.ParentTagName = &v
	return s
}

func (s *GetVideoResponseBodyVideoTags) SetTagConfidence(v float32) *GetVideoResponseBodyVideoTags {
	s.TagConfidence = &v
	return s
}

func (s *GetVideoResponseBodyVideoTags) SetTagLevel(v int32) *GetVideoResponseBodyVideoTags {
	s.TagLevel = &v
	return s
}

func (s *GetVideoResponseBodyVideoTags) SetTagName(v string) *GetVideoResponseBodyVideoTags {
	s.TagName = &v
	return s
}

type GetVideoResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetVideoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoResponse) GoString() string {
	return s.String()
}

func (s *GetVideoResponse) SetHeaders(v map[string]*string) *GetVideoResponse {
	s.Headers = v
	return s
}

func (s *GetVideoResponse) SetStatusCode(v int32) *GetVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoResponse) SetBody(v *GetVideoResponseBody) *GetVideoResponse {
	s.Body = v
	return s
}

type GetVideoTaskRequest struct {
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetVideoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVideoTaskRequest) GoString() string {
	return s.String()
}

func (s *GetVideoTaskRequest) SetProject(v string) *GetVideoTaskRequest {
	s.Project = &v
	return s
}

func (s *GetVideoTaskRequest) SetTaskId(v string) *GetVideoTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetVideoTaskRequest) SetTaskType(v string) *GetVideoTaskRequest {
	s.TaskType = &v
	return s
}

type GetVideoTaskResponseBody struct {
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ErrorMessage    *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	Parameters      *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Progress        *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result          *string `json:"Result,omitempty" xml:"Result,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType        *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetVideoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVideoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoTaskResponseBody) SetEndTime(v string) *GetVideoTaskResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetErrorMessage(v string) *GetVideoTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetNotifyEndpoint(v string) *GetVideoTaskResponseBody {
	s.NotifyEndpoint = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetNotifyTopicName(v string) *GetVideoTaskResponseBody {
	s.NotifyTopicName = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetParameters(v string) *GetVideoTaskResponseBody {
	s.Parameters = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetProgress(v int32) *GetVideoTaskResponseBody {
	s.Progress = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetRequestId(v string) *GetVideoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetResult(v string) *GetVideoTaskResponseBody {
	s.Result = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetStartTime(v string) *GetVideoTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetStatus(v string) *GetVideoTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetTaskId(v string) *GetVideoTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetVideoTaskResponseBody) SetTaskType(v string) *GetVideoTaskResponseBody {
	s.TaskType = &v
	return s
}

type GetVideoTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetVideoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVideoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoTaskResponse) GoString() string {
	return s.String()
}

func (s *GetVideoTaskResponse) SetHeaders(v map[string]*string) *GetVideoTaskResponse {
	s.Headers = v
	return s
}

func (s *GetVideoTaskResponse) SetStatusCode(v int32) *GetVideoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoTaskResponse) SetBody(v *GetVideoTaskResponseBody) *GetVideoTaskResponse {
	s.Body = v
	return s
}

type GetWebofficeURLRequest struct {
	File            *string `json:"File,omitempty" xml:"File,omitempty"`
	FileID          *string `json:"FileID,omitempty" xml:"FileID,omitempty"`
	Hidecmb         *bool   `json:"Hidecmb,omitempty" xml:"Hidecmb,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	Permission      *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SrcType         *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	User            *string `json:"User,omitempty" xml:"User,omitempty"`
	Watermark       *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
}

func (s GetWebofficeURLRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebofficeURLRequest) GoString() string {
	return s.String()
}

func (s *GetWebofficeURLRequest) SetFile(v string) *GetWebofficeURLRequest {
	s.File = &v
	return s
}

func (s *GetWebofficeURLRequest) SetFileID(v string) *GetWebofficeURLRequest {
	s.FileID = &v
	return s
}

func (s *GetWebofficeURLRequest) SetHidecmb(v bool) *GetWebofficeURLRequest {
	s.Hidecmb = &v
	return s
}

func (s *GetWebofficeURLRequest) SetNotifyEndpoint(v string) *GetWebofficeURLRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *GetWebofficeURLRequest) SetNotifyTopicName(v string) *GetWebofficeURLRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *GetWebofficeURLRequest) SetPermission(v string) *GetWebofficeURLRequest {
	s.Permission = &v
	return s
}

func (s *GetWebofficeURLRequest) SetProject(v string) *GetWebofficeURLRequest {
	s.Project = &v
	return s
}

func (s *GetWebofficeURLRequest) SetSrcType(v string) *GetWebofficeURLRequest {
	s.SrcType = &v
	return s
}

func (s *GetWebofficeURLRequest) SetUser(v string) *GetWebofficeURLRequest {
	s.User = &v
	return s
}

func (s *GetWebofficeURLRequest) SetWatermark(v string) *GetWebofficeURLRequest {
	s.Watermark = &v
	return s
}

type GetWebofficeURLResponseBody struct {
	AccessToken             *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	AccessTokenExpiredTime  *string `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
	RefreshToken            *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	RefreshTokenExpiredTime *string `json:"RefreshTokenExpiredTime,omitempty" xml:"RefreshTokenExpiredTime,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WebofficeURL            *string `json:"WebofficeURL,omitempty" xml:"WebofficeURL,omitempty"`
}

func (s GetWebofficeURLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWebofficeURLResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebofficeURLResponseBody) SetAccessToken(v string) *GetWebofficeURLResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GetWebofficeURLResponseBody) SetAccessTokenExpiredTime(v string) *GetWebofficeURLResponseBody {
	s.AccessTokenExpiredTime = &v
	return s
}

func (s *GetWebofficeURLResponseBody) SetRefreshToken(v string) *GetWebofficeURLResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *GetWebofficeURLResponseBody) SetRefreshTokenExpiredTime(v string) *GetWebofficeURLResponseBody {
	s.RefreshTokenExpiredTime = &v
	return s
}

func (s *GetWebofficeURLResponseBody) SetRequestId(v string) *GetWebofficeURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWebofficeURLResponseBody) SetWebofficeURL(v string) *GetWebofficeURLResponseBody {
	s.WebofficeURL = &v
	return s
}

type GetWebofficeURLResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetWebofficeURLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWebofficeURLResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWebofficeURLResponse) GoString() string {
	return s.String()
}

func (s *GetWebofficeURLResponse) SetHeaders(v map[string]*string) *GetWebofficeURLResponse {
	s.Headers = v
	return s
}

func (s *GetWebofficeURLResponse) SetStatusCode(v int32) *GetWebofficeURLResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWebofficeURLResponse) SetBody(v *GetWebofficeURLResponseBody) *GetWebofficeURLResponse {
	s.Body = v
	return s
}

type IndexImageRequest struct {
	ExternalId      *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	ImageUri        *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RemarksA        *string `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	RemarksArrayA   *string `json:"RemarksArrayA,omitempty" xml:"RemarksArrayA,omitempty"`
	RemarksArrayB   *string `json:"RemarksArrayB,omitempty" xml:"RemarksArrayB,omitempty"`
	RemarksB        *string `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	RemarksC        *string `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD        *string `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	SetId           *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	SourcePosition  *string `json:"SourcePosition,omitempty" xml:"SourcePosition,omitempty"`
	SourceType      *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SourceUri       *string `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
}

func (s IndexImageRequest) String() string {
	return tea.Prettify(s)
}

func (s IndexImageRequest) GoString() string {
	return s.String()
}

func (s *IndexImageRequest) SetExternalId(v string) *IndexImageRequest {
	s.ExternalId = &v
	return s
}

func (s *IndexImageRequest) SetImageUri(v string) *IndexImageRequest {
	s.ImageUri = &v
	return s
}

func (s *IndexImageRequest) SetNotifyEndpoint(v string) *IndexImageRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *IndexImageRequest) SetNotifyTopicName(v string) *IndexImageRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *IndexImageRequest) SetProject(v string) *IndexImageRequest {
	s.Project = &v
	return s
}

func (s *IndexImageRequest) SetRemarksA(v string) *IndexImageRequest {
	s.RemarksA = &v
	return s
}

func (s *IndexImageRequest) SetRemarksArrayA(v string) *IndexImageRequest {
	s.RemarksArrayA = &v
	return s
}

func (s *IndexImageRequest) SetRemarksArrayB(v string) *IndexImageRequest {
	s.RemarksArrayB = &v
	return s
}

func (s *IndexImageRequest) SetRemarksB(v string) *IndexImageRequest {
	s.RemarksB = &v
	return s
}

func (s *IndexImageRequest) SetRemarksC(v string) *IndexImageRequest {
	s.RemarksC = &v
	return s
}

func (s *IndexImageRequest) SetRemarksD(v string) *IndexImageRequest {
	s.RemarksD = &v
	return s
}

func (s *IndexImageRequest) SetSetId(v string) *IndexImageRequest {
	s.SetId = &v
	return s
}

func (s *IndexImageRequest) SetSourcePosition(v string) *IndexImageRequest {
	s.SourcePosition = &v
	return s
}

func (s *IndexImageRequest) SetSourceType(v string) *IndexImageRequest {
	s.SourceType = &v
	return s
}

func (s *IndexImageRequest) SetSourceUri(v string) *IndexImageRequest {
	s.SourceUri = &v
	return s
}

type IndexImageResponseBody struct {
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExternalId    *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	ImageUri      *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	ModifyTime    *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	RemarksA      *string `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	RemarksArrayA *string `json:"RemarksArrayA,omitempty" xml:"RemarksArrayA,omitempty"`
	RemarksArrayB *string `json:"RemarksArrayB,omitempty" xml:"RemarksArrayB,omitempty"`
	RemarksB      *string `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	RemarksC      *string `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD      *string `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId         *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s IndexImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndexImageResponseBody) GoString() string {
	return s.String()
}

func (s *IndexImageResponseBody) SetCreateTime(v string) *IndexImageResponseBody {
	s.CreateTime = &v
	return s
}

func (s *IndexImageResponseBody) SetExternalId(v string) *IndexImageResponseBody {
	s.ExternalId = &v
	return s
}

func (s *IndexImageResponseBody) SetImageUri(v string) *IndexImageResponseBody {
	s.ImageUri = &v
	return s
}

func (s *IndexImageResponseBody) SetModifyTime(v string) *IndexImageResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *IndexImageResponseBody) SetRemarksA(v string) *IndexImageResponseBody {
	s.RemarksA = &v
	return s
}

func (s *IndexImageResponseBody) SetRemarksArrayA(v string) *IndexImageResponseBody {
	s.RemarksArrayA = &v
	return s
}

func (s *IndexImageResponseBody) SetRemarksArrayB(v string) *IndexImageResponseBody {
	s.RemarksArrayB = &v
	return s
}

func (s *IndexImageResponseBody) SetRemarksB(v string) *IndexImageResponseBody {
	s.RemarksB = &v
	return s
}

func (s *IndexImageResponseBody) SetRemarksC(v string) *IndexImageResponseBody {
	s.RemarksC = &v
	return s
}

func (s *IndexImageResponseBody) SetRemarksD(v string) *IndexImageResponseBody {
	s.RemarksD = &v
	return s
}

func (s *IndexImageResponseBody) SetRequestId(v string) *IndexImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *IndexImageResponseBody) SetSetId(v string) *IndexImageResponseBody {
	s.SetId = &v
	return s
}

type IndexImageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IndexImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndexImageResponse) String() string {
	return tea.Prettify(s)
}

func (s IndexImageResponse) GoString() string {
	return s.String()
}

func (s *IndexImageResponse) SetHeaders(v map[string]*string) *IndexImageResponse {
	s.Headers = v
	return s
}

func (s *IndexImageResponse) SetStatusCode(v int32) *IndexImageResponse {
	s.StatusCode = &v
	return s
}

func (s *IndexImageResponse) SetBody(v *IndexImageResponseBody) *IndexImageResponse {
	s.Body = v
	return s
}

type IndexVideoRequest struct {
	ExternalId      *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RemarksA        *string `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	RemarksB        *string `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	RemarksC        *string `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD        *string `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	SetId           *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	TgtUri          *string `json:"TgtUri,omitempty" xml:"TgtUri,omitempty"`
	VideoUri        *string `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
}

func (s IndexVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s IndexVideoRequest) GoString() string {
	return s.String()
}

func (s *IndexVideoRequest) SetExternalId(v string) *IndexVideoRequest {
	s.ExternalId = &v
	return s
}

func (s *IndexVideoRequest) SetNotifyEndpoint(v string) *IndexVideoRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *IndexVideoRequest) SetNotifyTopicName(v string) *IndexVideoRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *IndexVideoRequest) SetProject(v string) *IndexVideoRequest {
	s.Project = &v
	return s
}

func (s *IndexVideoRequest) SetRemarksA(v string) *IndexVideoRequest {
	s.RemarksA = &v
	return s
}

func (s *IndexVideoRequest) SetRemarksB(v string) *IndexVideoRequest {
	s.RemarksB = &v
	return s
}

func (s *IndexVideoRequest) SetRemarksC(v string) *IndexVideoRequest {
	s.RemarksC = &v
	return s
}

func (s *IndexVideoRequest) SetRemarksD(v string) *IndexVideoRequest {
	s.RemarksD = &v
	return s
}

func (s *IndexVideoRequest) SetSetId(v string) *IndexVideoRequest {
	s.SetId = &v
	return s
}

func (s *IndexVideoRequest) SetTgtUri(v string) *IndexVideoRequest {
	s.TgtUri = &v
	return s
}

func (s *IndexVideoRequest) SetVideoUri(v string) *IndexVideoRequest {
	s.VideoUri = &v
	return s
}

type IndexVideoResponseBody struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExternalId *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	RemarksA   *string `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	RemarksB   *string `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	RemarksC   *string `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD   *string `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId      *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	VideoUri   *string `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
}

func (s IndexVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndexVideoResponseBody) GoString() string {
	return s.String()
}

func (s *IndexVideoResponseBody) SetCreateTime(v string) *IndexVideoResponseBody {
	s.CreateTime = &v
	return s
}

func (s *IndexVideoResponseBody) SetExternalId(v string) *IndexVideoResponseBody {
	s.ExternalId = &v
	return s
}

func (s *IndexVideoResponseBody) SetModifyTime(v string) *IndexVideoResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *IndexVideoResponseBody) SetRemarksA(v string) *IndexVideoResponseBody {
	s.RemarksA = &v
	return s
}

func (s *IndexVideoResponseBody) SetRemarksB(v string) *IndexVideoResponseBody {
	s.RemarksB = &v
	return s
}

func (s *IndexVideoResponseBody) SetRemarksC(v string) *IndexVideoResponseBody {
	s.RemarksC = &v
	return s
}

func (s *IndexVideoResponseBody) SetRemarksD(v string) *IndexVideoResponseBody {
	s.RemarksD = &v
	return s
}

func (s *IndexVideoResponseBody) SetRequestId(v string) *IndexVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *IndexVideoResponseBody) SetSetId(v string) *IndexVideoResponseBody {
	s.SetId = &v
	return s
}

func (s *IndexVideoResponseBody) SetVideoUri(v string) *IndexVideoResponseBody {
	s.VideoUri = &v
	return s
}

type IndexVideoResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IndexVideoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndexVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s IndexVideoResponse) GoString() string {
	return s.String()
}

func (s *IndexVideoResponse) SetHeaders(v map[string]*string) *IndexVideoResponse {
	s.Headers = v
	return s
}

func (s *IndexVideoResponse) SetStatusCode(v int32) *IndexVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *IndexVideoResponse) SetBody(v *IndexVideoResponseBody) *IndexVideoResponse {
	s.Body = v
	return s
}

type ListFaceGroupsRequest struct {
	ExternalId         *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	Limit              *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Marker             *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	Order              *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OrderBy            *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	Project            *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RemarksAQuery      *string `json:"RemarksAQuery,omitempty" xml:"RemarksAQuery,omitempty"`
	RemarksArrayAQuery *string `json:"RemarksArrayAQuery,omitempty" xml:"RemarksArrayAQuery,omitempty"`
	RemarksArrayBQuery *string `json:"RemarksArrayBQuery,omitempty" xml:"RemarksArrayBQuery,omitempty"`
	RemarksBQuery      *string `json:"RemarksBQuery,omitempty" xml:"RemarksBQuery,omitempty"`
	RemarksCQuery      *string `json:"RemarksCQuery,omitempty" xml:"RemarksCQuery,omitempty"`
	RemarksDQuery      *string `json:"RemarksDQuery,omitempty" xml:"RemarksDQuery,omitempty"`
	SetId              *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s ListFaceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFaceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListFaceGroupsRequest) SetExternalId(v string) *ListFaceGroupsRequest {
	s.ExternalId = &v
	return s
}

func (s *ListFaceGroupsRequest) SetLimit(v int32) *ListFaceGroupsRequest {
	s.Limit = &v
	return s
}

func (s *ListFaceGroupsRequest) SetMarker(v string) *ListFaceGroupsRequest {
	s.Marker = &v
	return s
}

func (s *ListFaceGroupsRequest) SetOrder(v string) *ListFaceGroupsRequest {
	s.Order = &v
	return s
}

func (s *ListFaceGroupsRequest) SetOrderBy(v string) *ListFaceGroupsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListFaceGroupsRequest) SetProject(v string) *ListFaceGroupsRequest {
	s.Project = &v
	return s
}

func (s *ListFaceGroupsRequest) SetRemarksAQuery(v string) *ListFaceGroupsRequest {
	s.RemarksAQuery = &v
	return s
}

func (s *ListFaceGroupsRequest) SetRemarksArrayAQuery(v string) *ListFaceGroupsRequest {
	s.RemarksArrayAQuery = &v
	return s
}

func (s *ListFaceGroupsRequest) SetRemarksArrayBQuery(v string) *ListFaceGroupsRequest {
	s.RemarksArrayBQuery = &v
	return s
}

func (s *ListFaceGroupsRequest) SetRemarksBQuery(v string) *ListFaceGroupsRequest {
	s.RemarksBQuery = &v
	return s
}

func (s *ListFaceGroupsRequest) SetRemarksCQuery(v string) *ListFaceGroupsRequest {
	s.RemarksCQuery = &v
	return s
}

func (s *ListFaceGroupsRequest) SetRemarksDQuery(v string) *ListFaceGroupsRequest {
	s.RemarksDQuery = &v
	return s
}

func (s *ListFaceGroupsRequest) SetSetId(v string) *ListFaceGroupsRequest {
	s.SetId = &v
	return s
}

type ListFaceGroupsResponseBody struct {
	FaceGroups []*ListFaceGroupsResponseBodyFaceGroups `json:"FaceGroups,omitempty" xml:"FaceGroups,omitempty" type:"Repeated"`
	NextMarker *string                                 `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFaceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFaceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFaceGroupsResponseBody) SetFaceGroups(v []*ListFaceGroupsResponseBodyFaceGroups) *ListFaceGroupsResponseBody {
	s.FaceGroups = v
	return s
}

func (s *ListFaceGroupsResponseBody) SetNextMarker(v string) *ListFaceGroupsResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListFaceGroupsResponseBody) SetRequestId(v string) *ListFaceGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ListFaceGroupsResponseBodyFaceGroups struct {
	AverageAge     *float32                                            `json:"AverageAge,omitempty" xml:"AverageAge,omitempty"`
	CreateTime     *string                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExternalId     *string                                             `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	FaceCount      *int32                                              `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	Gender         *string                                             `json:"Gender,omitempty" xml:"Gender,omitempty"`
	GroupCoverFace *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace `json:"GroupCoverFace,omitempty" xml:"GroupCoverFace,omitempty" type:"Struct"`
	GroupId        *string                                             `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName      *string                                             `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ImageCount     *int32                                              `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	MaxAge         *float32                                            `json:"MaxAge,omitempty" xml:"MaxAge,omitempty"`
	MinAge         *float32                                            `json:"MinAge,omitempty" xml:"MinAge,omitempty"`
	ModifyTime     *string                                             `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	RemarksA       *string                                             `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	RemarksArrayA  *string                                             `json:"RemarksArrayA,omitempty" xml:"RemarksArrayA,omitempty"`
	RemarksArrayB  *string                                             `json:"RemarksArrayB,omitempty" xml:"RemarksArrayB,omitempty"`
	RemarksB       *string                                             `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	RemarksC       *string                                             `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD       *string                                             `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
}

func (s ListFaceGroupsResponseBodyFaceGroups) String() string {
	return tea.Prettify(s)
}

func (s ListFaceGroupsResponseBodyFaceGroups) GoString() string {
	return s.String()
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetAverageAge(v float32) *ListFaceGroupsResponseBodyFaceGroups {
	s.AverageAge = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetCreateTime(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.CreateTime = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetExternalId(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.ExternalId = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetFaceCount(v int32) *ListFaceGroupsResponseBodyFaceGroups {
	s.FaceCount = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetGender(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.Gender = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetGroupCoverFace(v *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace) *ListFaceGroupsResponseBodyFaceGroups {
	s.GroupCoverFace = v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetGroupId(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.GroupId = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetGroupName(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.GroupName = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetImageCount(v int32) *ListFaceGroupsResponseBodyFaceGroups {
	s.ImageCount = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetMaxAge(v float32) *ListFaceGroupsResponseBodyFaceGroups {
	s.MaxAge = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetMinAge(v float32) *ListFaceGroupsResponseBodyFaceGroups {
	s.MinAge = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetModifyTime(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.ModifyTime = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetRemarksA(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.RemarksA = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetRemarksArrayA(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.RemarksArrayA = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetRemarksArrayB(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.RemarksArrayB = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetRemarksB(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.RemarksB = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetRemarksC(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.RemarksC = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetRemarksD(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.RemarksD = &v
	return s
}

type ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace struct {
	ExternalId   *string                                                         `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	FaceBoundary *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary `json:"FaceBoundary,omitempty" xml:"FaceBoundary,omitempty" type:"Struct"`
	FaceId       *string                                                         `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	ImageHeight  *int64                                                          `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	ImageUri     *string                                                         `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	ImageWidth   *int64                                                          `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
}

func (s ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace) String() string {
	return tea.Prettify(s)
}

func (s ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace) GoString() string {
	return s.String()
}

func (s *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace) SetExternalId(v string) *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace {
	s.ExternalId = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace) SetFaceBoundary(v *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary) *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace {
	s.FaceBoundary = v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace) SetFaceId(v string) *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace {
	s.FaceId = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace) SetImageHeight(v int64) *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace {
	s.ImageHeight = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace) SetImageUri(v string) *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace {
	s.ImageUri = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace) SetImageWidth(v int64) *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFace {
	s.ImageWidth = &v
	return s
}

type ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary) String() string {
	return tea.Prettify(s)
}

func (s ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary) GoString() string {
	return s.String()
}

func (s *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary) SetHeight(v int32) *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary {
	s.Height = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary) SetLeft(v int32) *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary {
	s.Left = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary) SetTop(v int32) *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary {
	s.Top = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary) SetWidth(v int32) *ListFaceGroupsResponseBodyFaceGroupsGroupCoverFaceFaceBoundary {
	s.Width = &v
	return s
}

type ListFaceGroupsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFaceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFaceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFaceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListFaceGroupsResponse) SetHeaders(v map[string]*string) *ListFaceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListFaceGroupsResponse) SetStatusCode(v int32) *ListFaceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFaceGroupsResponse) SetBody(v *ListFaceGroupsResponseBody) *ListFaceGroupsResponse {
	s.Body = v
	return s
}

type ListImagesRequest struct {
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	Limit           *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Marker          *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId           *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s ListImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) SetCreateTimeStart(v string) *ListImagesRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListImagesRequest) SetLimit(v int32) *ListImagesRequest {
	s.Limit = &v
	return s
}

func (s *ListImagesRequest) SetMarker(v string) *ListImagesRequest {
	s.Marker = &v
	return s
}

func (s *ListImagesRequest) SetProject(v string) *ListImagesRequest {
	s.Project = &v
	return s
}

func (s *ListImagesRequest) SetSetId(v string) *ListImagesRequest {
	s.SetId = &v
	return s
}

type ListImagesResponseBody struct {
	Images     []*ListImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	NextMarker *string                         `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId      *string                         `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s ListImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListImagesResponseBody) SetNextMarker(v string) *ListImagesResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImagesResponseBody) SetSetId(v string) *ListImagesResponseBody {
	s.SetId = &v
	return s
}

type ListImagesResponseBodyImages struct {
	Address                      *ListImagesResponseBodyImagesAddress              `json:"Address,omitempty" xml:"Address,omitempty" type:"Struct"`
	AddressFailReason            *string                                           `json:"AddressFailReason,omitempty" xml:"AddressFailReason,omitempty"`
	AddressModifyTime            *string                                           `json:"AddressModifyTime,omitempty" xml:"AddressModifyTime,omitempty"`
	AddressStatus                *string                                           `json:"AddressStatus,omitempty" xml:"AddressStatus,omitempty"`
	CreateTime                   *string                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CroppingSuggestion           []*ListImagesResponseBodyImagesCroppingSuggestion `json:"CroppingSuggestion,omitempty" xml:"CroppingSuggestion,omitempty" type:"Repeated"`
	CroppingSuggestionFailReason *string                                           `json:"CroppingSuggestionFailReason,omitempty" xml:"CroppingSuggestionFailReason,omitempty"`
	CroppingSuggestionModifyTime *string                                           `json:"CroppingSuggestionModifyTime,omitempty" xml:"CroppingSuggestionModifyTime,omitempty"`
	CroppingSuggestionStatus     *string                                           `json:"CroppingSuggestionStatus,omitempty" xml:"CroppingSuggestionStatus,omitempty"`
	Exif                         *string                                           `json:"Exif,omitempty" xml:"Exif,omitempty"`
	ExternalId                   *string                                           `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	Faces                        []*ListImagesResponseBodyImagesFaces              `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	FacesFailReason              *string                                           `json:"FacesFailReason,omitempty" xml:"FacesFailReason,omitempty"`
	FacesModifyTime              *string                                           `json:"FacesModifyTime,omitempty" xml:"FacesModifyTime,omitempty"`
	FacesStatus                  *string                                           `json:"FacesStatus,omitempty" xml:"FacesStatus,omitempty"`
	FileSize                     *int32                                            `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	ImageFormat                  *string                                           `json:"ImageFormat,omitempty" xml:"ImageFormat,omitempty"`
	ImageHeight                  *int32                                            `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	ImageQuality                 *ListImagesResponseBodyImagesImageQuality         `json:"ImageQuality,omitempty" xml:"ImageQuality,omitempty" type:"Struct"`
	ImageQualityFailReason       *string                                           `json:"ImageQualityFailReason,omitempty" xml:"ImageQualityFailReason,omitempty"`
	ImageQualityModifyTime       *string                                           `json:"ImageQualityModifyTime,omitempty" xml:"ImageQualityModifyTime,omitempty"`
	ImageQualityStatus           *string                                           `json:"ImageQualityStatus,omitempty" xml:"ImageQualityStatus,omitempty"`
	ImageTime                    *string                                           `json:"ImageTime,omitempty" xml:"ImageTime,omitempty"`
	ImageUri                     *string                                           `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	ImageWidth                   *int32                                            `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
	Location                     *string                                           `json:"Location,omitempty" xml:"Location,omitempty"`
	ModifyTime                   *string                                           `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	OCR                          []*ListImagesResponseBodyImagesOCR                `json:"OCR,omitempty" xml:"OCR,omitempty" type:"Repeated"`
	OCRFailReason                *string                                           `json:"OCRFailReason,omitempty" xml:"OCRFailReason,omitempty"`
	OCRModifyTime                *string                                           `json:"OCRModifyTime,omitempty" xml:"OCRModifyTime,omitempty"`
	OCRStatus                    *string                                           `json:"OCRStatus,omitempty" xml:"OCRStatus,omitempty"`
	Orientation                  *string                                           `json:"Orientation,omitempty" xml:"Orientation,omitempty"`
	RemarksA                     *string                                           `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	RemarksArrayA                *string                                           `json:"RemarksArrayA,omitempty" xml:"RemarksArrayA,omitempty"`
	RemarksArrayB                *string                                           `json:"RemarksArrayB,omitempty" xml:"RemarksArrayB,omitempty"`
	RemarksB                     *string                                           `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	RemarksC                     *string                                           `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD                     *string                                           `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	SourcePosition               *string                                           `json:"SourcePosition,omitempty" xml:"SourcePosition,omitempty"`
	SourceType                   *string                                           `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SourceUri                    *string                                           `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
	Tags                         []*ListImagesResponseBodyImagesTags               `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TagsFailReason               *string                                           `json:"TagsFailReason,omitempty" xml:"TagsFailReason,omitempty"`
	TagsModifyTime               *string                                           `json:"TagsModifyTime,omitempty" xml:"TagsModifyTime,omitempty"`
	TagsStatus                   *string                                           `json:"TagsStatus,omitempty" xml:"TagsStatus,omitempty"`
}

func (s ListImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImages) SetAddress(v *ListImagesResponseBodyImagesAddress) *ListImagesResponseBodyImages {
	s.Address = v
	return s
}

func (s *ListImagesResponseBodyImages) SetAddressFailReason(v string) *ListImagesResponseBodyImages {
	s.AddressFailReason = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetAddressModifyTime(v string) *ListImagesResponseBodyImages {
	s.AddressModifyTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetAddressStatus(v string) *ListImagesResponseBodyImages {
	s.AddressStatus = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetCreateTime(v string) *ListImagesResponseBodyImages {
	s.CreateTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetCroppingSuggestion(v []*ListImagesResponseBodyImagesCroppingSuggestion) *ListImagesResponseBodyImages {
	s.CroppingSuggestion = v
	return s
}

func (s *ListImagesResponseBodyImages) SetCroppingSuggestionFailReason(v string) *ListImagesResponseBodyImages {
	s.CroppingSuggestionFailReason = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetCroppingSuggestionModifyTime(v string) *ListImagesResponseBodyImages {
	s.CroppingSuggestionModifyTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetCroppingSuggestionStatus(v string) *ListImagesResponseBodyImages {
	s.CroppingSuggestionStatus = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetExif(v string) *ListImagesResponseBodyImages {
	s.Exif = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetExternalId(v string) *ListImagesResponseBodyImages {
	s.ExternalId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetFaces(v []*ListImagesResponseBodyImagesFaces) *ListImagesResponseBodyImages {
	s.Faces = v
	return s
}

func (s *ListImagesResponseBodyImages) SetFacesFailReason(v string) *ListImagesResponseBodyImages {
	s.FacesFailReason = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetFacesModifyTime(v string) *ListImagesResponseBodyImages {
	s.FacesModifyTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetFacesStatus(v string) *ListImagesResponseBodyImages {
	s.FacesStatus = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetFileSize(v int32) *ListImagesResponseBodyImages {
	s.FileSize = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageFormat(v string) *ListImagesResponseBodyImages {
	s.ImageFormat = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageHeight(v int32) *ListImagesResponseBodyImages {
	s.ImageHeight = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageQuality(v *ListImagesResponseBodyImagesImageQuality) *ListImagesResponseBodyImages {
	s.ImageQuality = v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageQualityFailReason(v string) *ListImagesResponseBodyImages {
	s.ImageQualityFailReason = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageQualityModifyTime(v string) *ListImagesResponseBodyImages {
	s.ImageQualityModifyTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageQualityStatus(v string) *ListImagesResponseBodyImages {
	s.ImageQualityStatus = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageTime(v string) *ListImagesResponseBodyImages {
	s.ImageTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageUri(v string) *ListImagesResponseBodyImages {
	s.ImageUri = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageWidth(v int32) *ListImagesResponseBodyImages {
	s.ImageWidth = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetLocation(v string) *ListImagesResponseBodyImages {
	s.Location = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetModifyTime(v string) *ListImagesResponseBodyImages {
	s.ModifyTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetOCR(v []*ListImagesResponseBodyImagesOCR) *ListImagesResponseBodyImages {
	s.OCR = v
	return s
}

func (s *ListImagesResponseBodyImages) SetOCRFailReason(v string) *ListImagesResponseBodyImages {
	s.OCRFailReason = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetOCRModifyTime(v string) *ListImagesResponseBodyImages {
	s.OCRModifyTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetOCRStatus(v string) *ListImagesResponseBodyImages {
	s.OCRStatus = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetOrientation(v string) *ListImagesResponseBodyImages {
	s.Orientation = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetRemarksA(v string) *ListImagesResponseBodyImages {
	s.RemarksA = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetRemarksArrayA(v string) *ListImagesResponseBodyImages {
	s.RemarksArrayA = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetRemarksArrayB(v string) *ListImagesResponseBodyImages {
	s.RemarksArrayB = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetRemarksB(v string) *ListImagesResponseBodyImages {
	s.RemarksB = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetRemarksC(v string) *ListImagesResponseBodyImages {
	s.RemarksC = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetRemarksD(v string) *ListImagesResponseBodyImages {
	s.RemarksD = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetSourcePosition(v string) *ListImagesResponseBodyImages {
	s.SourcePosition = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetSourceType(v string) *ListImagesResponseBodyImages {
	s.SourceType = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetSourceUri(v string) *ListImagesResponseBodyImages {
	s.SourceUri = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetTags(v []*ListImagesResponseBodyImagesTags) *ListImagesResponseBodyImages {
	s.Tags = v
	return s
}

func (s *ListImagesResponseBodyImages) SetTagsFailReason(v string) *ListImagesResponseBodyImages {
	s.TagsFailReason = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetTagsModifyTime(v string) *ListImagesResponseBodyImages {
	s.TagsModifyTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetTagsStatus(v string) *ListImagesResponseBodyImages {
	s.TagsStatus = &v
	return s
}

type ListImagesResponseBodyImagesAddress struct {
	AddressLine *string `json:"AddressLine,omitempty" xml:"AddressLine,omitempty"`
	City        *string `json:"City,omitempty" xml:"City,omitempty"`
	Country     *string `json:"Country,omitempty" xml:"Country,omitempty"`
	District    *string `json:"District,omitempty" xml:"District,omitempty"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty"`
	Township    *string `json:"Township,omitempty" xml:"Township,omitempty"`
}

func (s ListImagesResponseBodyImagesAddress) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesAddress) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesAddress) SetAddressLine(v string) *ListImagesResponseBodyImagesAddress {
	s.AddressLine = &v
	return s
}

func (s *ListImagesResponseBodyImagesAddress) SetCity(v string) *ListImagesResponseBodyImagesAddress {
	s.City = &v
	return s
}

func (s *ListImagesResponseBodyImagesAddress) SetCountry(v string) *ListImagesResponseBodyImagesAddress {
	s.Country = &v
	return s
}

func (s *ListImagesResponseBodyImagesAddress) SetDistrict(v string) *ListImagesResponseBodyImagesAddress {
	s.District = &v
	return s
}

func (s *ListImagesResponseBodyImagesAddress) SetProvince(v string) *ListImagesResponseBodyImagesAddress {
	s.Province = &v
	return s
}

func (s *ListImagesResponseBodyImagesAddress) SetTownship(v string) *ListImagesResponseBodyImagesAddress {
	s.Township = &v
	return s
}

type ListImagesResponseBodyImagesCroppingSuggestion struct {
	AspectRatio      *string                                                         `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	CroppingBoundary *ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary `json:"CroppingBoundary,omitempty" xml:"CroppingBoundary,omitempty" type:"Struct"`
	Score            *float32                                                        `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s ListImagesResponseBodyImagesCroppingSuggestion) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesCroppingSuggestion) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesCroppingSuggestion) SetAspectRatio(v string) *ListImagesResponseBodyImagesCroppingSuggestion {
	s.AspectRatio = &v
	return s
}

func (s *ListImagesResponseBodyImagesCroppingSuggestion) SetCroppingBoundary(v *ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) *ListImagesResponseBodyImagesCroppingSuggestion {
	s.CroppingBoundary = v
	return s
}

func (s *ListImagesResponseBodyImagesCroppingSuggestion) SetScore(v float32) *ListImagesResponseBodyImagesCroppingSuggestion {
	s.Score = &v
	return s
}

type ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) SetHeight(v int32) *ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary {
	s.Height = &v
	return s
}

func (s *ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) SetLeft(v int32) *ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary {
	s.Left = &v
	return s
}

func (s *ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) SetTop(v int32) *ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary {
	s.Top = &v
	return s
}

func (s *ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary) SetWidth(v int32) *ListImagesResponseBodyImagesCroppingSuggestionCroppingBoundary {
	s.Width = &v
	return s
}

type ListImagesResponseBodyImagesFaces struct {
	Age               *int32                                           `json:"Age,omitempty" xml:"Age,omitempty"`
	Attractive        *float32                                         `json:"Attractive,omitempty" xml:"Attractive,omitempty"`
	Emotion           *string                                          `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	EmotionConfidence *float32                                         `json:"EmotionConfidence,omitempty" xml:"EmotionConfidence,omitempty"`
	EmotionDetails    *ListImagesResponseBodyImagesFacesEmotionDetails `json:"EmotionDetails,omitempty" xml:"EmotionDetails,omitempty" type:"Struct"`
	FaceAttributes    *ListImagesResponseBodyImagesFacesFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" type:"Struct"`
	FaceConfidence    *float32                                         `json:"FaceConfidence,omitempty" xml:"FaceConfidence,omitempty"`
	FaceId            *string                                          `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	FaceQuality       *float32                                         `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	Gender            *string                                          `json:"Gender,omitempty" xml:"Gender,omitempty"`
	GenderConfidence  *float32                                         `json:"GenderConfidence,omitempty" xml:"GenderConfidence,omitempty"`
	GroupId           *string                                          `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s ListImagesResponseBodyImagesFaces) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesFaces) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesFaces) SetAge(v int32) *ListImagesResponseBodyImagesFaces {
	s.Age = &v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetAttractive(v float32) *ListImagesResponseBodyImagesFaces {
	s.Attractive = &v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetEmotion(v string) *ListImagesResponseBodyImagesFaces {
	s.Emotion = &v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetEmotionConfidence(v float32) *ListImagesResponseBodyImagesFaces {
	s.EmotionConfidence = &v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetEmotionDetails(v *ListImagesResponseBodyImagesFacesEmotionDetails) *ListImagesResponseBodyImagesFaces {
	s.EmotionDetails = v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetFaceAttributes(v *ListImagesResponseBodyImagesFacesFaceAttributes) *ListImagesResponseBodyImagesFaces {
	s.FaceAttributes = v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetFaceConfidence(v float32) *ListImagesResponseBodyImagesFaces {
	s.FaceConfidence = &v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetFaceId(v string) *ListImagesResponseBodyImagesFaces {
	s.FaceId = &v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetFaceQuality(v float32) *ListImagesResponseBodyImagesFaces {
	s.FaceQuality = &v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetGender(v string) *ListImagesResponseBodyImagesFaces {
	s.Gender = &v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetGenderConfidence(v float32) *ListImagesResponseBodyImagesFaces {
	s.GenderConfidence = &v
	return s
}

func (s *ListImagesResponseBodyImagesFaces) SetGroupId(v string) *ListImagesResponseBodyImagesFaces {
	s.GroupId = &v
	return s
}

type ListImagesResponseBodyImagesFacesEmotionDetails struct {
	ANGRY     *float32 `json:"ANGRY,omitempty" xml:"ANGRY,omitempty"`
	CALM      *float32 `json:"CALM,omitempty" xml:"CALM,omitempty"`
	DISGUSTED *float32 `json:"DISGUSTED,omitempty" xml:"DISGUSTED,omitempty"`
	HAPPY     *float32 `json:"HAPPY,omitempty" xml:"HAPPY,omitempty"`
	SAD       *float32 `json:"SAD,omitempty" xml:"SAD,omitempty"`
	SCARED    *float32 `json:"SCARED,omitempty" xml:"SCARED,omitempty"`
	SURPRISED *float32 `json:"SURPRISED,omitempty" xml:"SURPRISED,omitempty"`
}

func (s ListImagesResponseBodyImagesFacesEmotionDetails) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesFacesEmotionDetails) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesFacesEmotionDetails) SetANGRY(v float32) *ListImagesResponseBodyImagesFacesEmotionDetails {
	s.ANGRY = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesEmotionDetails) SetCALM(v float32) *ListImagesResponseBodyImagesFacesEmotionDetails {
	s.CALM = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesEmotionDetails) SetDISGUSTED(v float32) *ListImagesResponseBodyImagesFacesEmotionDetails {
	s.DISGUSTED = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesEmotionDetails) SetHAPPY(v float32) *ListImagesResponseBodyImagesFacesEmotionDetails {
	s.HAPPY = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesEmotionDetails) SetSAD(v float32) *ListImagesResponseBodyImagesFacesEmotionDetails {
	s.SAD = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesEmotionDetails) SetSCARED(v float32) *ListImagesResponseBodyImagesFacesEmotionDetails {
	s.SCARED = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesEmotionDetails) SetSURPRISED(v float32) *ListImagesResponseBodyImagesFacesEmotionDetails {
	s.SURPRISED = &v
	return s
}

type ListImagesResponseBodyImagesFacesFaceAttributes struct {
	Beard             *string                                                      `json:"Beard,omitempty" xml:"Beard,omitempty"`
	BeardConfidence   *float32                                                     `json:"BeardConfidence,omitempty" xml:"BeardConfidence,omitempty"`
	FaceBoundary      *ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary `json:"FaceBoundary,omitempty" xml:"FaceBoundary,omitempty" type:"Struct"`
	Glasses           *string                                                      `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	GlassesConfidence *float32                                                     `json:"GlassesConfidence,omitempty" xml:"GlassesConfidence,omitempty"`
	HeadPose          *ListImagesResponseBodyImagesFacesFaceAttributesHeadPose     `json:"HeadPose,omitempty" xml:"HeadPose,omitempty" type:"Struct"`
	Mask              *string                                                      `json:"Mask,omitempty" xml:"Mask,omitempty"`
	MaskConfidence    *float32                                                     `json:"MaskConfidence,omitempty" xml:"MaskConfidence,omitempty"`
}

func (s ListImagesResponseBodyImagesFacesFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesFacesFaceAttributes) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributes) SetBeard(v string) *ListImagesResponseBodyImagesFacesFaceAttributes {
	s.Beard = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributes) SetBeardConfidence(v float32) *ListImagesResponseBodyImagesFacesFaceAttributes {
	s.BeardConfidence = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributes) SetFaceBoundary(v *ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) *ListImagesResponseBodyImagesFacesFaceAttributes {
	s.FaceBoundary = v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributes) SetGlasses(v string) *ListImagesResponseBodyImagesFacesFaceAttributes {
	s.Glasses = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributes) SetGlassesConfidence(v float32) *ListImagesResponseBodyImagesFacesFaceAttributes {
	s.GlassesConfidence = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributes) SetHeadPose(v *ListImagesResponseBodyImagesFacesFaceAttributesHeadPose) *ListImagesResponseBodyImagesFacesFaceAttributes {
	s.HeadPose = v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributes) SetMask(v string) *ListImagesResponseBodyImagesFacesFaceAttributes {
	s.Mask = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributes) SetMaskConfidence(v float32) *ListImagesResponseBodyImagesFacesFaceAttributes {
	s.MaskConfidence = &v
	return s
}

type ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) SetHeight(v int32) *ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary {
	s.Height = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) SetLeft(v int32) *ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary {
	s.Left = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) SetTop(v int32) *ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary {
	s.Top = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary) SetWidth(v int32) *ListImagesResponseBodyImagesFacesFaceAttributesFaceBoundary {
	s.Width = &v
	return s
}

type ListImagesResponseBodyImagesFacesFaceAttributesHeadPose struct {
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	Roll  *float32 `json:"Roll,omitempty" xml:"Roll,omitempty"`
	Yaw   *float32 `json:"Yaw,omitempty" xml:"Yaw,omitempty"`
}

func (s ListImagesResponseBodyImagesFacesFaceAttributesHeadPose) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesFacesFaceAttributesHeadPose) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributesHeadPose) SetPitch(v float32) *ListImagesResponseBodyImagesFacesFaceAttributesHeadPose {
	s.Pitch = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributesHeadPose) SetRoll(v float32) *ListImagesResponseBodyImagesFacesFaceAttributesHeadPose {
	s.Roll = &v
	return s
}

func (s *ListImagesResponseBodyImagesFacesFaceAttributesHeadPose) SetYaw(v float32) *ListImagesResponseBodyImagesFacesFaceAttributesHeadPose {
	s.Yaw = &v
	return s
}

type ListImagesResponseBodyImagesImageQuality struct {
	Clarity          *float32 `json:"Clarity,omitempty" xml:"Clarity,omitempty"`
	ClarityScore     *float32 `json:"ClarityScore,omitempty" xml:"ClarityScore,omitempty"`
	Color            *float32 `json:"Color,omitempty" xml:"Color,omitempty"`
	ColorScore       *float32 `json:"ColorScore,omitempty" xml:"ColorScore,omitempty"`
	CompositionScore *float32 `json:"CompositionScore,omitempty" xml:"CompositionScore,omitempty"`
	Contrast         *float32 `json:"Contrast,omitempty" xml:"Contrast,omitempty"`
	ContrastScore    *float32 `json:"ContrastScore,omitempty" xml:"ContrastScore,omitempty"`
	Exposure         *float32 `json:"Exposure,omitempty" xml:"Exposure,omitempty"`
	ExposureScore    *float32 `json:"ExposureScore,omitempty" xml:"ExposureScore,omitempty"`
	OverallScore     *float32 `json:"OverallScore,omitempty" xml:"OverallScore,omitempty"`
}

func (s ListImagesResponseBodyImagesImageQuality) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesImageQuality) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesImageQuality) SetClarity(v float32) *ListImagesResponseBodyImagesImageQuality {
	s.Clarity = &v
	return s
}

func (s *ListImagesResponseBodyImagesImageQuality) SetClarityScore(v float32) *ListImagesResponseBodyImagesImageQuality {
	s.ClarityScore = &v
	return s
}

func (s *ListImagesResponseBodyImagesImageQuality) SetColor(v float32) *ListImagesResponseBodyImagesImageQuality {
	s.Color = &v
	return s
}

func (s *ListImagesResponseBodyImagesImageQuality) SetColorScore(v float32) *ListImagesResponseBodyImagesImageQuality {
	s.ColorScore = &v
	return s
}

func (s *ListImagesResponseBodyImagesImageQuality) SetCompositionScore(v float32) *ListImagesResponseBodyImagesImageQuality {
	s.CompositionScore = &v
	return s
}

func (s *ListImagesResponseBodyImagesImageQuality) SetContrast(v float32) *ListImagesResponseBodyImagesImageQuality {
	s.Contrast = &v
	return s
}

func (s *ListImagesResponseBodyImagesImageQuality) SetContrastScore(v float32) *ListImagesResponseBodyImagesImageQuality {
	s.ContrastScore = &v
	return s
}

func (s *ListImagesResponseBodyImagesImageQuality) SetExposure(v float32) *ListImagesResponseBodyImagesImageQuality {
	s.Exposure = &v
	return s
}

func (s *ListImagesResponseBodyImagesImageQuality) SetExposureScore(v float32) *ListImagesResponseBodyImagesImageQuality {
	s.ExposureScore = &v
	return s
}

func (s *ListImagesResponseBodyImagesImageQuality) SetOverallScore(v float32) *ListImagesResponseBodyImagesImageQuality {
	s.OverallScore = &v
	return s
}

type ListImagesResponseBodyImagesOCR struct {
	OCRBoundary   *ListImagesResponseBodyImagesOCROCRBoundary `json:"OCRBoundary,omitempty" xml:"OCRBoundary,omitempty" type:"Struct"`
	OCRConfidence *float32                                    `json:"OCRConfidence,omitempty" xml:"OCRConfidence,omitempty"`
	OCRContents   *string                                     `json:"OCRContents,omitempty" xml:"OCRContents,omitempty"`
}

func (s ListImagesResponseBodyImagesOCR) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesOCR) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesOCR) SetOCRBoundary(v *ListImagesResponseBodyImagesOCROCRBoundary) *ListImagesResponseBodyImagesOCR {
	s.OCRBoundary = v
	return s
}

func (s *ListImagesResponseBodyImagesOCR) SetOCRConfidence(v float32) *ListImagesResponseBodyImagesOCR {
	s.OCRConfidence = &v
	return s
}

func (s *ListImagesResponseBodyImagesOCR) SetOCRContents(v string) *ListImagesResponseBodyImagesOCR {
	s.OCRContents = &v
	return s
}

type ListImagesResponseBodyImagesOCROCRBoundary struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListImagesResponseBodyImagesOCROCRBoundary) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesOCROCRBoundary) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesOCROCRBoundary) SetHeight(v int32) *ListImagesResponseBodyImagesOCROCRBoundary {
	s.Height = &v
	return s
}

func (s *ListImagesResponseBodyImagesOCROCRBoundary) SetLeft(v int32) *ListImagesResponseBodyImagesOCROCRBoundary {
	s.Left = &v
	return s
}

func (s *ListImagesResponseBodyImagesOCROCRBoundary) SetTop(v int32) *ListImagesResponseBodyImagesOCROCRBoundary {
	s.Top = &v
	return s
}

func (s *ListImagesResponseBodyImagesOCROCRBoundary) SetWidth(v int32) *ListImagesResponseBodyImagesOCROCRBoundary {
	s.Width = &v
	return s
}

type ListImagesResponseBodyImagesTags struct {
	CentricScore  *float32 `json:"CentricScore,omitempty" xml:"CentricScore,omitempty"`
	ParentTagName *string  `json:"ParentTagName,omitempty" xml:"ParentTagName,omitempty"`
	TagConfidence *float32 `json:"TagConfidence,omitempty" xml:"TagConfidence,omitempty"`
	TagLevel      *int32   `json:"TagLevel,omitempty" xml:"TagLevel,omitempty"`
	TagName       *string  `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s ListImagesResponseBodyImagesTags) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesTags) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesTags) SetCentricScore(v float32) *ListImagesResponseBodyImagesTags {
	s.CentricScore = &v
	return s
}

func (s *ListImagesResponseBodyImagesTags) SetParentTagName(v string) *ListImagesResponseBodyImagesTags {
	s.ParentTagName = &v
	return s
}

func (s *ListImagesResponseBodyImagesTags) SetTagConfidence(v float32) *ListImagesResponseBodyImagesTags {
	s.TagConfidence = &v
	return s
}

func (s *ListImagesResponseBodyImagesTags) SetTagLevel(v int32) *ListImagesResponseBodyImagesTags {
	s.TagLevel = &v
	return s
}

func (s *ListImagesResponseBodyImagesTags) SetTagName(v string) *ListImagesResponseBodyImagesTags {
	s.TagName = &v
	return s
}

type ListImagesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponse) GoString() string {
	return s.String()
}

func (s *ListImagesResponse) SetHeaders(v map[string]*string) *ListImagesResponse {
	s.Headers = v
	return s
}

func (s *ListImagesResponse) SetStatusCode(v int32) *ListImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
	s.Body = v
	return s
}

type ListOfficeConversionTaskRequest struct {
	Marker  *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxKeys *int32  `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s ListOfficeConversionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeConversionTaskRequest) GoString() string {
	return s.String()
}

func (s *ListOfficeConversionTaskRequest) SetMarker(v string) *ListOfficeConversionTaskRequest {
	s.Marker = &v
	return s
}

func (s *ListOfficeConversionTaskRequest) SetMaxKeys(v int32) *ListOfficeConversionTaskRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListOfficeConversionTaskRequest) SetProject(v string) *ListOfficeConversionTaskRequest {
	s.Project = &v
	return s
}

type ListOfficeConversionTaskResponseBody struct {
	NextMarker *string                                      `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks      []*ListOfficeConversionTaskResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ListOfficeConversionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeConversionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListOfficeConversionTaskResponseBody) SetNextMarker(v string) *ListOfficeConversionTaskResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBody) SetRequestId(v string) *ListOfficeConversionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBody) SetTasks(v []*ListOfficeConversionTaskResponseBodyTasks) *ListOfficeConversionTaskResponseBody {
	s.Tasks = v
	return s
}

type ListOfficeConversionTaskResponseBodyTasks struct {
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExternalID      *string `json:"ExternalID,omitempty" xml:"ExternalID,omitempty"`
	FinishTime      *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	ImageSpec       *string `json:"ImageSpec,omitempty" xml:"ImageSpec,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	PageCount       *int32  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	Percent         *int32  `json:"Percent,omitempty" xml:"Percent,omitempty"`
	SrcUri          *string `json:"SrcUri,omitempty" xml:"SrcUri,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TgtType         *string `json:"TgtType,omitempty" xml:"TgtType,omitempty"`
	TgtUri          *string `json:"TgtUri,omitempty" xml:"TgtUri,omitempty"`
}

func (s ListOfficeConversionTaskResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeConversionTaskResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetCreateTime(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.CreateTime = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetExternalID(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.ExternalID = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetFinishTime(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.FinishTime = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetImageSpec(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.ImageSpec = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetNotifyEndpoint(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.NotifyEndpoint = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetNotifyTopicName(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.NotifyTopicName = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetPageCount(v int32) *ListOfficeConversionTaskResponseBodyTasks {
	s.PageCount = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetPercent(v int32) *ListOfficeConversionTaskResponseBodyTasks {
	s.Percent = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetSrcUri(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.SrcUri = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetStatus(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetTaskId(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetTgtType(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.TgtType = &v
	return s
}

func (s *ListOfficeConversionTaskResponseBodyTasks) SetTgtUri(v string) *ListOfficeConversionTaskResponseBodyTasks {
	s.TgtUri = &v
	return s
}

type ListOfficeConversionTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOfficeConversionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOfficeConversionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeConversionTaskResponse) GoString() string {
	return s.String()
}

func (s *ListOfficeConversionTaskResponse) SetHeaders(v map[string]*string) *ListOfficeConversionTaskResponse {
	s.Headers = v
	return s
}

func (s *ListOfficeConversionTaskResponse) SetStatusCode(v int32) *ListOfficeConversionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOfficeConversionTaskResponse) SetBody(v *ListOfficeConversionTaskResponseBody) *ListOfficeConversionTaskResponse {
	s.Body = v
	return s
}

type ListProjectsRequest struct {
	Marker  *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxKeys *int32  `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) SetMarker(v string) *ListProjectsRequest {
	s.Marker = &v
	return s
}

func (s *ListProjectsRequest) SetMaxKeys(v int32) *ListProjectsRequest {
	s.MaxKeys = &v
	return s
}

type ListProjectsResponseBody struct {
	NextMarker *string                             `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	Projects   []*ListProjectsResponseBodyProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) SetNextMarker(v string) *ListProjectsResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListProjectsResponseBody) SetProjects(v []*ListProjectsResponseBodyProjects) *ListProjectsResponseBody {
	s.Projects = v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

type ListProjectsResponseBodyProjects struct {
	BillingType *string `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	CU          *int32  `json:"CU,omitempty" xml:"CU,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Endpoint    *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ModifyTime  *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Project     *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListProjectsResponseBodyProjects) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyProjects) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyProjects) SetBillingType(v string) *ListProjectsResponseBodyProjects {
	s.BillingType = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetCU(v int32) *ListProjectsResponseBodyProjects {
	s.CU = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetCreateTime(v string) *ListProjectsResponseBodyProjects {
	s.CreateTime = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetEndpoint(v string) *ListProjectsResponseBodyProjects {
	s.Endpoint = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetModifyTime(v string) *ListProjectsResponseBodyProjects {
	s.ModifyTime = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetProject(v string) *ListProjectsResponseBodyProjects {
	s.Project = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetRegionId(v string) *ListProjectsResponseBodyProjects {
	s.RegionId = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetServiceRole(v string) *ListProjectsResponseBodyProjects {
	s.ServiceRole = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetType(v string) *ListProjectsResponseBodyProjects {
	s.Type = &v
	return s
}

type ListProjectsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectsResponse) SetHeaders(v map[string]*string) *ListProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectsResponse) SetStatusCode(v int32) *ListProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectsResponse) SetBody(v *ListProjectsResponseBody) *ListProjectsResponse {
	s.Body = v
	return s
}

type ListSetTagsRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId   *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s ListSetTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSetTagsRequest) GoString() string {
	return s.String()
}

func (s *ListSetTagsRequest) SetProject(v string) *ListSetTagsRequest {
	s.Project = &v
	return s
}

func (s *ListSetTagsRequest) SetSetId(v string) *ListSetTagsRequest {
	s.SetId = &v
	return s
}

type ListSetTagsResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId     *string                        `json:"SetId,omitempty" xml:"SetId,omitempty"`
	Tags      []*ListSetTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListSetTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSetTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSetTagsResponseBody) SetRequestId(v string) *ListSetTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSetTagsResponseBody) SetSetId(v string) *ListSetTagsResponseBody {
	s.SetId = &v
	return s
}

func (s *ListSetTagsResponseBody) SetTags(v []*ListSetTagsResponseBodyTags) *ListSetTagsResponseBody {
	s.Tags = v
	return s
}

type ListSetTagsResponseBodyTags struct {
	TagCount *int32  `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	TagLevel *int32  `json:"TagLevel,omitempty" xml:"TagLevel,omitempty"`
	TagName  *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s ListSetTagsResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s ListSetTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListSetTagsResponseBodyTags) SetTagCount(v int32) *ListSetTagsResponseBodyTags {
	s.TagCount = &v
	return s
}

func (s *ListSetTagsResponseBodyTags) SetTagLevel(v int32) *ListSetTagsResponseBodyTags {
	s.TagLevel = &v
	return s
}

func (s *ListSetTagsResponseBodyTags) SetTagName(v string) *ListSetTagsResponseBodyTags {
	s.TagName = &v
	return s
}

type ListSetTagsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSetTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSetTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSetTagsResponse) GoString() string {
	return s.String()
}

func (s *ListSetTagsResponse) SetHeaders(v map[string]*string) *ListSetTagsResponse {
	s.Headers = v
	return s
}

func (s *ListSetTagsResponse) SetStatusCode(v int32) *ListSetTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSetTagsResponse) SetBody(v *ListSetTagsResponseBody) *ListSetTagsResponse {
	s.Body = v
	return s
}

type ListSetsRequest struct {
	Marker  *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s ListSetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSetsRequest) GoString() string {
	return s.String()
}

func (s *ListSetsRequest) SetMarker(v string) *ListSetsRequest {
	s.Marker = &v
	return s
}

func (s *ListSetsRequest) SetProject(v string) *ListSetsRequest {
	s.Project = &v
	return s
}

type ListSetsResponseBody struct {
	NextMarker *string                     `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	RequestId  *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sets       []*ListSetsResponseBodySets `json:"Sets,omitempty" xml:"Sets,omitempty" type:"Repeated"`
}

func (s ListSetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSetsResponseBody) SetNextMarker(v string) *ListSetsResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListSetsResponseBody) SetRequestId(v string) *ListSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSetsResponseBody) SetSets(v []*ListSetsResponseBodySets) *ListSetsResponseBody {
	s.Sets = v
	return s
}

type ListSetsResponseBodySets struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FaceCount   *int32  `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	ImageCount  *int32  `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	ModifyTime  *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	SetId       *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	SetName     *string `json:"SetName,omitempty" xml:"SetName,omitempty"`
	VideoCount  *int32  `json:"VideoCount,omitempty" xml:"VideoCount,omitempty"`
	VideoLength *int32  `json:"VideoLength,omitempty" xml:"VideoLength,omitempty"`
}

func (s ListSetsResponseBodySets) String() string {
	return tea.Prettify(s)
}

func (s ListSetsResponseBodySets) GoString() string {
	return s.String()
}

func (s *ListSetsResponseBodySets) SetCreateTime(v string) *ListSetsResponseBodySets {
	s.CreateTime = &v
	return s
}

func (s *ListSetsResponseBodySets) SetFaceCount(v int32) *ListSetsResponseBodySets {
	s.FaceCount = &v
	return s
}

func (s *ListSetsResponseBodySets) SetImageCount(v int32) *ListSetsResponseBodySets {
	s.ImageCount = &v
	return s
}

func (s *ListSetsResponseBodySets) SetModifyTime(v string) *ListSetsResponseBodySets {
	s.ModifyTime = &v
	return s
}

func (s *ListSetsResponseBodySets) SetSetId(v string) *ListSetsResponseBodySets {
	s.SetId = &v
	return s
}

func (s *ListSetsResponseBodySets) SetSetName(v string) *ListSetsResponseBodySets {
	s.SetName = &v
	return s
}

func (s *ListSetsResponseBodySets) SetVideoCount(v int32) *ListSetsResponseBodySets {
	s.VideoCount = &v
	return s
}

func (s *ListSetsResponseBodySets) SetVideoLength(v int32) *ListSetsResponseBodySets {
	s.VideoLength = &v
	return s
}

type ListSetsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSetsResponse) GoString() string {
	return s.String()
}

func (s *ListSetsResponse) SetHeaders(v map[string]*string) *ListSetsResponse {
	s.Headers = v
	return s
}

func (s *ListSetsResponse) SetStatusCode(v int32) *ListSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSetsResponse) SetBody(v *ListSetsResponseBody) *ListSetsResponse {
	s.Body = v
	return s
}

type ListVideoTasksRequest struct {
	Marker   *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxKeys  *int32  `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListVideoTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVideoTasksRequest) GoString() string {
	return s.String()
}

func (s *ListVideoTasksRequest) SetMarker(v string) *ListVideoTasksRequest {
	s.Marker = &v
	return s
}

func (s *ListVideoTasksRequest) SetMaxKeys(v int32) *ListVideoTasksRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListVideoTasksRequest) SetProject(v string) *ListVideoTasksRequest {
	s.Project = &v
	return s
}

func (s *ListVideoTasksRequest) SetTaskType(v string) *ListVideoTasksRequest {
	s.TaskType = &v
	return s
}

type ListVideoTasksResponseBody struct {
	NextMarker *string                            `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks      []*ListVideoTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ListVideoTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVideoTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListVideoTasksResponseBody) SetNextMarker(v string) *ListVideoTasksResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListVideoTasksResponseBody) SetRequestId(v string) *ListVideoTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVideoTasksResponseBody) SetTasks(v []*ListVideoTasksResponseBodyTasks) *ListVideoTasksResponseBody {
	s.Tasks = v
	return s
}

type ListVideoTasksResponseBodyTasks struct {
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ErrorMessage    *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	Parameters      *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Progress        *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Result          *string `json:"Result,omitempty" xml:"Result,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType        *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListVideoTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListVideoTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListVideoTasksResponseBodyTasks) SetEndTime(v string) *ListVideoTasksResponseBodyTasks {
	s.EndTime = &v
	return s
}

func (s *ListVideoTasksResponseBodyTasks) SetErrorMessage(v string) *ListVideoTasksResponseBodyTasks {
	s.ErrorMessage = &v
	return s
}

func (s *ListVideoTasksResponseBodyTasks) SetNotifyEndpoint(v string) *ListVideoTasksResponseBodyTasks {
	s.NotifyEndpoint = &v
	return s
}

func (s *ListVideoTasksResponseBodyTasks) SetNotifyTopicName(v string) *ListVideoTasksResponseBodyTasks {
	s.NotifyTopicName = &v
	return s
}

func (s *ListVideoTasksResponseBodyTasks) SetParameters(v string) *ListVideoTasksResponseBodyTasks {
	s.Parameters = &v
	return s
}

func (s *ListVideoTasksResponseBodyTasks) SetProgress(v int32) *ListVideoTasksResponseBodyTasks {
	s.Progress = &v
	return s
}

func (s *ListVideoTasksResponseBodyTasks) SetResult(v string) *ListVideoTasksResponseBodyTasks {
	s.Result = &v
	return s
}

func (s *ListVideoTasksResponseBodyTasks) SetStartTime(v string) *ListVideoTasksResponseBodyTasks {
	s.StartTime = &v
	return s
}

func (s *ListVideoTasksResponseBodyTasks) SetStatus(v string) *ListVideoTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListVideoTasksResponseBodyTasks) SetTaskId(v string) *ListVideoTasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ListVideoTasksResponseBodyTasks) SetTaskType(v string) *ListVideoTasksResponseBodyTasks {
	s.TaskType = &v
	return s
}

type ListVideoTasksResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListVideoTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVideoTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVideoTasksResponse) GoString() string {
	return s.String()
}

func (s *ListVideoTasksResponse) SetHeaders(v map[string]*string) *ListVideoTasksResponse {
	s.Headers = v
	return s
}

func (s *ListVideoTasksResponse) SetStatusCode(v int32) *ListVideoTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVideoTasksResponse) SetBody(v *ListVideoTasksResponseBody) *ListVideoTasksResponse {
	s.Body = v
	return s
}

type ListVideosRequest struct {
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	Marker          *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId           *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s ListVideosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVideosRequest) GoString() string {
	return s.String()
}

func (s *ListVideosRequest) SetCreateTimeStart(v string) *ListVideosRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListVideosRequest) SetMarker(v string) *ListVideosRequest {
	s.Marker = &v
	return s
}

func (s *ListVideosRequest) SetProject(v string) *ListVideosRequest {
	s.Project = &v
	return s
}

func (s *ListVideosRequest) SetSetId(v string) *ListVideosRequest {
	s.SetId = &v
	return s
}

type ListVideosResponseBody struct {
	NextMarker *string                         `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId      *string                         `json:"SetId,omitempty" xml:"SetId,omitempty"`
	Videos     []*ListVideosResponseBodyVideos `json:"Videos,omitempty" xml:"Videos,omitempty" type:"Repeated"`
}

func (s ListVideosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVideosResponseBody) GoString() string {
	return s.String()
}

func (s *ListVideosResponseBody) SetNextMarker(v string) *ListVideosResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListVideosResponseBody) SetRequestId(v string) *ListVideosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVideosResponseBody) SetSetId(v string) *ListVideosResponseBody {
	s.SetId = &v
	return s
}

func (s *ListVideosResponseBody) SetVideos(v []*ListVideosResponseBodyVideos) *ListVideosResponseBody {
	s.Videos = v
	return s
}

type ListVideosResponseBodyVideos struct {
	CreateTime          *string                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExternalId          *string                                  `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	FileSize            *int32                                   `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	ModifyTime          *string                                  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ProcessFailReason   *string                                  `json:"ProcessFailReason,omitempty" xml:"ProcessFailReason,omitempty"`
	ProcessModifyTime   *string                                  `json:"ProcessModifyTime,omitempty" xml:"ProcessModifyTime,omitempty"`
	ProcessStatus       *string                                  `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	RemarksA            *string                                  `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	RemarksB            *string                                  `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	RemarksC            *string                                  `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD            *string                                  `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	SourcePosition      *string                                  `json:"SourcePosition,omitempty" xml:"SourcePosition,omitempty"`
	SourceType          *string                                  `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SourceUri           *string                                  `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
	VideoDuration       *float32                                 `json:"VideoDuration,omitempty" xml:"VideoDuration,omitempty"`
	VideoFormat         *string                                  `json:"VideoFormat,omitempty" xml:"VideoFormat,omitempty"`
	VideoFrames         *int32                                   `json:"VideoFrames,omitempty" xml:"VideoFrames,omitempty"`
	VideoHeight         *int32                                   `json:"VideoHeight,omitempty" xml:"VideoHeight,omitempty"`
	VideoTags           []*ListVideosResponseBodyVideosVideoTags `json:"VideoTags,omitempty" xml:"VideoTags,omitempty" type:"Repeated"`
	VideoTagsFailReason *string                                  `json:"VideoTagsFailReason,omitempty" xml:"VideoTagsFailReason,omitempty"`
	VideoTagsModifyTime *string                                  `json:"VideoTagsModifyTime,omitempty" xml:"VideoTagsModifyTime,omitempty"`
	VideoTagsStatus     *string                                  `json:"VideoTagsStatus,omitempty" xml:"VideoTagsStatus,omitempty"`
	VideoUri            *string                                  `json:"VideoUri,omitempty" xml:"VideoUri,omitempty"`
	VideoWidth          *int32                                   `json:"VideoWidth,omitempty" xml:"VideoWidth,omitempty"`
}

func (s ListVideosResponseBodyVideos) String() string {
	return tea.Prettify(s)
}

func (s ListVideosResponseBodyVideos) GoString() string {
	return s.String()
}

func (s *ListVideosResponseBodyVideos) SetCreateTime(v string) *ListVideosResponseBodyVideos {
	s.CreateTime = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetExternalId(v string) *ListVideosResponseBodyVideos {
	s.ExternalId = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetFileSize(v int32) *ListVideosResponseBodyVideos {
	s.FileSize = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetModifyTime(v string) *ListVideosResponseBodyVideos {
	s.ModifyTime = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetProcessFailReason(v string) *ListVideosResponseBodyVideos {
	s.ProcessFailReason = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetProcessModifyTime(v string) *ListVideosResponseBodyVideos {
	s.ProcessModifyTime = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetProcessStatus(v string) *ListVideosResponseBodyVideos {
	s.ProcessStatus = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetRemarksA(v string) *ListVideosResponseBodyVideos {
	s.RemarksA = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetRemarksB(v string) *ListVideosResponseBodyVideos {
	s.RemarksB = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetRemarksC(v string) *ListVideosResponseBodyVideos {
	s.RemarksC = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetRemarksD(v string) *ListVideosResponseBodyVideos {
	s.RemarksD = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetSourcePosition(v string) *ListVideosResponseBodyVideos {
	s.SourcePosition = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetSourceType(v string) *ListVideosResponseBodyVideos {
	s.SourceType = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetSourceUri(v string) *ListVideosResponseBodyVideos {
	s.SourceUri = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetVideoDuration(v float32) *ListVideosResponseBodyVideos {
	s.VideoDuration = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetVideoFormat(v string) *ListVideosResponseBodyVideos {
	s.VideoFormat = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetVideoFrames(v int32) *ListVideosResponseBodyVideos {
	s.VideoFrames = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetVideoHeight(v int32) *ListVideosResponseBodyVideos {
	s.VideoHeight = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetVideoTags(v []*ListVideosResponseBodyVideosVideoTags) *ListVideosResponseBodyVideos {
	s.VideoTags = v
	return s
}

func (s *ListVideosResponseBodyVideos) SetVideoTagsFailReason(v string) *ListVideosResponseBodyVideos {
	s.VideoTagsFailReason = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetVideoTagsModifyTime(v string) *ListVideosResponseBodyVideos {
	s.VideoTagsModifyTime = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetVideoTagsStatus(v string) *ListVideosResponseBodyVideos {
	s.VideoTagsStatus = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetVideoUri(v string) *ListVideosResponseBodyVideos {
	s.VideoUri = &v
	return s
}

func (s *ListVideosResponseBodyVideos) SetVideoWidth(v int32) *ListVideosResponseBodyVideos {
	s.VideoWidth = &v
	return s
}

type ListVideosResponseBodyVideosVideoTags struct {
	ParentTagName *string  `json:"ParentTagName,omitempty" xml:"ParentTagName,omitempty"`
	TagConfidence *float32 `json:"TagConfidence,omitempty" xml:"TagConfidence,omitempty"`
	TagLevel      *int32   `json:"TagLevel,omitempty" xml:"TagLevel,omitempty"`
	TagName       *string  `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s ListVideosResponseBodyVideosVideoTags) String() string {
	return tea.Prettify(s)
}

func (s ListVideosResponseBodyVideosVideoTags) GoString() string {
	return s.String()
}

func (s *ListVideosResponseBodyVideosVideoTags) SetParentTagName(v string) *ListVideosResponseBodyVideosVideoTags {
	s.ParentTagName = &v
	return s
}

func (s *ListVideosResponseBodyVideosVideoTags) SetTagConfidence(v float32) *ListVideosResponseBodyVideosVideoTags {
	s.TagConfidence = &v
	return s
}

func (s *ListVideosResponseBodyVideosVideoTags) SetTagLevel(v int32) *ListVideosResponseBodyVideosVideoTags {
	s.TagLevel = &v
	return s
}

func (s *ListVideosResponseBodyVideosVideoTags) SetTagName(v string) *ListVideosResponseBodyVideosVideoTags {
	s.TagName = &v
	return s
}

type ListVideosResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListVideosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVideosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVideosResponse) GoString() string {
	return s.String()
}

func (s *ListVideosResponse) SetHeaders(v map[string]*string) *ListVideosResponse {
	s.Headers = v
	return s
}

func (s *ListVideosResponse) SetStatusCode(v int32) *ListVideosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVideosResponse) SetBody(v *ListVideosResponseBody) *ListVideosResponse {
	s.Body = v
	return s
}

type OpenImmServiceRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s OpenImmServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenImmServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenImmServiceRequest) SetOwnerId(v int64) *OpenImmServiceRequest {
	s.OwnerId = &v
	return s
}

type OpenImmServiceResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenImmServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenImmServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenImmServiceResponseBody) SetOrderId(v string) *OpenImmServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenImmServiceResponseBody) SetRequestId(v string) *OpenImmServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenImmServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenImmServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenImmServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenImmServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenImmServiceResponse) SetHeaders(v map[string]*string) *OpenImmServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenImmServiceResponse) SetStatusCode(v int32) *OpenImmServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenImmServiceResponse) SetBody(v *OpenImmServiceResponseBody) *OpenImmServiceResponse {
	s.Body = v
	return s
}

type PutProjectRequest struct {
	Project     *string `json:"Project,omitempty" xml:"Project,omitempty"`
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
}

func (s PutProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PutProjectRequest) GoString() string {
	return s.String()
}

func (s *PutProjectRequest) SetProject(v string) *PutProjectRequest {
	s.Project = &v
	return s
}

func (s *PutProjectRequest) SetServiceRole(v string) *PutProjectRequest {
	s.ServiceRole = &v
	return s
}

type PutProjectResponseBody struct {
	BillingType *string `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	CU          *int32  `json:"CU,omitempty" xml:"CU,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Endpoint    *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ModifyTime  *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Project     *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PutProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PutProjectResponseBody) SetBillingType(v string) *PutProjectResponseBody {
	s.BillingType = &v
	return s
}

func (s *PutProjectResponseBody) SetCU(v int32) *PutProjectResponseBody {
	s.CU = &v
	return s
}

func (s *PutProjectResponseBody) SetCreateTime(v string) *PutProjectResponseBody {
	s.CreateTime = &v
	return s
}

func (s *PutProjectResponseBody) SetEndpoint(v string) *PutProjectResponseBody {
	s.Endpoint = &v
	return s
}

func (s *PutProjectResponseBody) SetModifyTime(v string) *PutProjectResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *PutProjectResponseBody) SetProject(v string) *PutProjectResponseBody {
	s.Project = &v
	return s
}

func (s *PutProjectResponseBody) SetRegionId(v string) *PutProjectResponseBody {
	s.RegionId = &v
	return s
}

func (s *PutProjectResponseBody) SetRequestId(v string) *PutProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutProjectResponseBody) SetServiceRole(v string) *PutProjectResponseBody {
	s.ServiceRole = &v
	return s
}

func (s *PutProjectResponseBody) SetType(v string) *PutProjectResponseBody {
	s.Type = &v
	return s
}

type PutProjectResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PutProjectResponse) GoString() string {
	return s.String()
}

func (s *PutProjectResponse) SetHeaders(v map[string]*string) *PutProjectResponse {
	s.Headers = v
	return s
}

func (s *PutProjectResponse) SetStatusCode(v int32) *PutProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PutProjectResponse) SetBody(v *PutProjectResponseBody) *PutProjectResponse {
	s.Body = v
	return s
}

type RefreshOfficePreviewTokenRequest struct {
	AccessToken  *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	Project      *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
}

func (s RefreshOfficePreviewTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshOfficePreviewTokenRequest) GoString() string {
	return s.String()
}

func (s *RefreshOfficePreviewTokenRequest) SetAccessToken(v string) *RefreshOfficePreviewTokenRequest {
	s.AccessToken = &v
	return s
}

func (s *RefreshOfficePreviewTokenRequest) SetProject(v string) *RefreshOfficePreviewTokenRequest {
	s.Project = &v
	return s
}

func (s *RefreshOfficePreviewTokenRequest) SetRefreshToken(v string) *RefreshOfficePreviewTokenRequest {
	s.RefreshToken = &v
	return s
}

type RefreshOfficePreviewTokenResponseBody struct {
	AccessToken             *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	AccessTokenExpiredTime  *string `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
	RefreshToken            *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	RefreshTokenExpiredTime *string `json:"RefreshTokenExpiredTime,omitempty" xml:"RefreshTokenExpiredTime,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshOfficePreviewTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshOfficePreviewTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshOfficePreviewTokenResponseBody) SetAccessToken(v string) *RefreshOfficePreviewTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *RefreshOfficePreviewTokenResponseBody) SetAccessTokenExpiredTime(v string) *RefreshOfficePreviewTokenResponseBody {
	s.AccessTokenExpiredTime = &v
	return s
}

func (s *RefreshOfficePreviewTokenResponseBody) SetRefreshToken(v string) *RefreshOfficePreviewTokenResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *RefreshOfficePreviewTokenResponseBody) SetRefreshTokenExpiredTime(v string) *RefreshOfficePreviewTokenResponseBody {
	s.RefreshTokenExpiredTime = &v
	return s
}

func (s *RefreshOfficePreviewTokenResponseBody) SetRequestId(v string) *RefreshOfficePreviewTokenResponseBody {
	s.RequestId = &v
	return s
}

type RefreshOfficePreviewTokenResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RefreshOfficePreviewTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshOfficePreviewTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshOfficePreviewTokenResponse) GoString() string {
	return s.String()
}

func (s *RefreshOfficePreviewTokenResponse) SetHeaders(v map[string]*string) *RefreshOfficePreviewTokenResponse {
	s.Headers = v
	return s
}

func (s *RefreshOfficePreviewTokenResponse) SetStatusCode(v int32) *RefreshOfficePreviewTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshOfficePreviewTokenResponse) SetBody(v *RefreshOfficePreviewTokenResponseBody) *RefreshOfficePreviewTokenResponse {
	s.Body = v
	return s
}

type RefreshWebofficeTokenRequest struct {
	AccessToken  *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	Project      *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
}

func (s RefreshWebofficeTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshWebofficeTokenRequest) GoString() string {
	return s.String()
}

func (s *RefreshWebofficeTokenRequest) SetAccessToken(v string) *RefreshWebofficeTokenRequest {
	s.AccessToken = &v
	return s
}

func (s *RefreshWebofficeTokenRequest) SetProject(v string) *RefreshWebofficeTokenRequest {
	s.Project = &v
	return s
}

func (s *RefreshWebofficeTokenRequest) SetRefreshToken(v string) *RefreshWebofficeTokenRequest {
	s.RefreshToken = &v
	return s
}

type RefreshWebofficeTokenResponseBody struct {
	AccessToken             *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	AccessTokenExpiredTime  *string `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
	RefreshToken            *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	RefreshTokenExpiredTime *string `json:"RefreshTokenExpiredTime,omitempty" xml:"RefreshTokenExpiredTime,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshWebofficeTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshWebofficeTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshWebofficeTokenResponseBody) SetAccessToken(v string) *RefreshWebofficeTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *RefreshWebofficeTokenResponseBody) SetAccessTokenExpiredTime(v string) *RefreshWebofficeTokenResponseBody {
	s.AccessTokenExpiredTime = &v
	return s
}

func (s *RefreshWebofficeTokenResponseBody) SetRefreshToken(v string) *RefreshWebofficeTokenResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *RefreshWebofficeTokenResponseBody) SetRefreshTokenExpiredTime(v string) *RefreshWebofficeTokenResponseBody {
	s.RefreshTokenExpiredTime = &v
	return s
}

func (s *RefreshWebofficeTokenResponseBody) SetRequestId(v string) *RefreshWebofficeTokenResponseBody {
	s.RequestId = &v
	return s
}

type RefreshWebofficeTokenResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RefreshWebofficeTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshWebofficeTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshWebofficeTokenResponse) GoString() string {
	return s.String()
}

func (s *RefreshWebofficeTokenResponse) SetHeaders(v map[string]*string) *RefreshWebofficeTokenResponse {
	s.Headers = v
	return s
}

func (s *RefreshWebofficeTokenResponse) SetStatusCode(v int32) *RefreshWebofficeTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshWebofficeTokenResponse) SetBody(v *RefreshWebofficeTokenResponseBody) *RefreshWebofficeTokenResponse {
	s.Body = v
	return s
}

type UpdateFaceGroupRequest struct {
	ExternalId       *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	GroupCoverFaceId *string `json:"GroupCoverFaceId,omitempty" xml:"GroupCoverFaceId,omitempty"`
	GroupId          *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName        *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Project          *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RemarksA         *string `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	RemarksArrayA    *string `json:"RemarksArrayA,omitempty" xml:"RemarksArrayA,omitempty"`
	RemarksArrayB    *string `json:"RemarksArrayB,omitempty" xml:"RemarksArrayB,omitempty"`
	RemarksB         *string `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	RemarksC         *string `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD         *string `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	ResetItems       *string `json:"ResetItems,omitempty" xml:"ResetItems,omitempty"`
	SetId            *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s UpdateFaceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateFaceGroupRequest) SetExternalId(v string) *UpdateFaceGroupRequest {
	s.ExternalId = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetGroupCoverFaceId(v string) *UpdateFaceGroupRequest {
	s.GroupCoverFaceId = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetGroupId(v string) *UpdateFaceGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetGroupName(v string) *UpdateFaceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetProject(v string) *UpdateFaceGroupRequest {
	s.Project = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetRemarksA(v string) *UpdateFaceGroupRequest {
	s.RemarksA = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetRemarksArrayA(v string) *UpdateFaceGroupRequest {
	s.RemarksArrayA = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetRemarksArrayB(v string) *UpdateFaceGroupRequest {
	s.RemarksArrayB = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetRemarksB(v string) *UpdateFaceGroupRequest {
	s.RemarksB = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetRemarksC(v string) *UpdateFaceGroupRequest {
	s.RemarksC = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetRemarksD(v string) *UpdateFaceGroupRequest {
	s.RemarksD = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetResetItems(v string) *UpdateFaceGroupRequest {
	s.ResetItems = &v
	return s
}

func (s *UpdateFaceGroupRequest) SetSetId(v string) *UpdateFaceGroupRequest {
	s.SetId = &v
	return s
}

type UpdateFaceGroupResponseBody struct {
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId     *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s UpdateFaceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFaceGroupResponseBody) SetGroupId(v string) *UpdateFaceGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *UpdateFaceGroupResponseBody) SetRequestId(v string) *UpdateFaceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFaceGroupResponseBody) SetSetId(v string) *UpdateFaceGroupResponseBody {
	s.SetId = &v
	return s
}

type UpdateFaceGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateFaceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFaceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateFaceGroupResponse) SetHeaders(v map[string]*string) *UpdateFaceGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateFaceGroupResponse) SetStatusCode(v int32) *UpdateFaceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFaceGroupResponse) SetBody(v *UpdateFaceGroupResponseBody) *UpdateFaceGroupResponse {
	s.Body = v
	return s
}

type UpdateImageRequest struct {
	ExternalId     *string                    `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	Faces          []*UpdateImageRequestFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	ImageUri       *string                    `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Project        *string                    `json:"Project,omitempty" xml:"Project,omitempty"`
	RemarksA       *string                    `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	RemarksArrayA  *string                    `json:"RemarksArrayA,omitempty" xml:"RemarksArrayA,omitempty"`
	RemarksArrayB  *string                    `json:"RemarksArrayB,omitempty" xml:"RemarksArrayB,omitempty"`
	RemarksB       *string                    `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	RemarksC       *string                    `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD       *string                    `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	SetId          *string                    `json:"SetId,omitempty" xml:"SetId,omitempty"`
	SourcePosition *string                    `json:"SourcePosition,omitempty" xml:"SourcePosition,omitempty"`
	SourceType     *string                    `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SourceUri      *string                    `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
}

func (s UpdateImageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageRequest) SetExternalId(v string) *UpdateImageRequest {
	s.ExternalId = &v
	return s
}

func (s *UpdateImageRequest) SetFaces(v []*UpdateImageRequestFaces) *UpdateImageRequest {
	s.Faces = v
	return s
}

func (s *UpdateImageRequest) SetImageUri(v string) *UpdateImageRequest {
	s.ImageUri = &v
	return s
}

func (s *UpdateImageRequest) SetProject(v string) *UpdateImageRequest {
	s.Project = &v
	return s
}

func (s *UpdateImageRequest) SetRemarksA(v string) *UpdateImageRequest {
	s.RemarksA = &v
	return s
}

func (s *UpdateImageRequest) SetRemarksArrayA(v string) *UpdateImageRequest {
	s.RemarksArrayA = &v
	return s
}

func (s *UpdateImageRequest) SetRemarksArrayB(v string) *UpdateImageRequest {
	s.RemarksArrayB = &v
	return s
}

func (s *UpdateImageRequest) SetRemarksB(v string) *UpdateImageRequest {
	s.RemarksB = &v
	return s
}

func (s *UpdateImageRequest) SetRemarksC(v string) *UpdateImageRequest {
	s.RemarksC = &v
	return s
}

func (s *UpdateImageRequest) SetRemarksD(v string) *UpdateImageRequest {
	s.RemarksD = &v
	return s
}

func (s *UpdateImageRequest) SetSetId(v string) *UpdateImageRequest {
	s.SetId = &v
	return s
}

func (s *UpdateImageRequest) SetSourcePosition(v string) *UpdateImageRequest {
	s.SourcePosition = &v
	return s
}

func (s *UpdateImageRequest) SetSourceType(v string) *UpdateImageRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateImageRequest) SetSourceUri(v string) *UpdateImageRequest {
	s.SourceUri = &v
	return s
}

type UpdateImageRequestFaces struct {
	FaceId  *string `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s UpdateImageRequestFaces) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageRequestFaces) GoString() string {
	return s.String()
}

func (s *UpdateImageRequestFaces) SetFaceId(v string) *UpdateImageRequestFaces {
	s.FaceId = &v
	return s
}

func (s *UpdateImageRequestFaces) SetGroupId(v string) *UpdateImageRequestFaces {
	s.GroupId = &v
	return s
}

type UpdateImageShrinkRequest struct {
	ExternalId     *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	FacesShrink    *string `json:"Faces,omitempty" xml:"Faces,omitempty"`
	ImageUri       *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Project        *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RemarksA       *string `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	RemarksArrayA  *string `json:"RemarksArrayA,omitempty" xml:"RemarksArrayA,omitempty"`
	RemarksArrayB  *string `json:"RemarksArrayB,omitempty" xml:"RemarksArrayB,omitempty"`
	RemarksB       *string `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	RemarksC       *string `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD       *string `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	SetId          *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	SourcePosition *string `json:"SourcePosition,omitempty" xml:"SourcePosition,omitempty"`
	SourceType     *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SourceUri      *string `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
}

func (s UpdateImageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageShrinkRequest) SetExternalId(v string) *UpdateImageShrinkRequest {
	s.ExternalId = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetFacesShrink(v string) *UpdateImageShrinkRequest {
	s.FacesShrink = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetImageUri(v string) *UpdateImageShrinkRequest {
	s.ImageUri = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetProject(v string) *UpdateImageShrinkRequest {
	s.Project = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetRemarksA(v string) *UpdateImageShrinkRequest {
	s.RemarksA = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetRemarksArrayA(v string) *UpdateImageShrinkRequest {
	s.RemarksArrayA = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetRemarksArrayB(v string) *UpdateImageShrinkRequest {
	s.RemarksArrayB = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetRemarksB(v string) *UpdateImageShrinkRequest {
	s.RemarksB = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetRemarksC(v string) *UpdateImageShrinkRequest {
	s.RemarksC = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetRemarksD(v string) *UpdateImageShrinkRequest {
	s.RemarksD = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetSetId(v string) *UpdateImageShrinkRequest {
	s.SetId = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetSourcePosition(v string) *UpdateImageShrinkRequest {
	s.SourcePosition = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetSourceType(v string) *UpdateImageShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetSourceUri(v string) *UpdateImageShrinkRequest {
	s.SourceUri = &v
	return s
}

type UpdateImageResponseBody struct {
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExternalId    *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	ImageUri      *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	ModifyTime    *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	RemarksA      *string `json:"RemarksA,omitempty" xml:"RemarksA,omitempty"`
	RemarksArrayA *string `json:"RemarksArrayA,omitempty" xml:"RemarksArrayA,omitempty"`
	RemarksArrayB *string `json:"RemarksArrayB,omitempty" xml:"RemarksArrayB,omitempty"`
	RemarksB      *string `json:"RemarksB,omitempty" xml:"RemarksB,omitempty"`
	RemarksC      *string `json:"RemarksC,omitempty" xml:"RemarksC,omitempty"`
	RemarksD      *string `json:"RemarksD,omitempty" xml:"RemarksD,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId         *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
}

func (s UpdateImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageResponseBody) SetCreateTime(v string) *UpdateImageResponseBody {
	s.CreateTime = &v
	return s
}

func (s *UpdateImageResponseBody) SetExternalId(v string) *UpdateImageResponseBody {
	s.ExternalId = &v
	return s
}

func (s *UpdateImageResponseBody) SetImageUri(v string) *UpdateImageResponseBody {
	s.ImageUri = &v
	return s
}

func (s *UpdateImageResponseBody) SetModifyTime(v string) *UpdateImageResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *UpdateImageResponseBody) SetRemarksA(v string) *UpdateImageResponseBody {
	s.RemarksA = &v
	return s
}

func (s *UpdateImageResponseBody) SetRemarksArrayA(v string) *UpdateImageResponseBody {
	s.RemarksArrayA = &v
	return s
}

func (s *UpdateImageResponseBody) SetRemarksArrayB(v string) *UpdateImageResponseBody {
	s.RemarksArrayB = &v
	return s
}

func (s *UpdateImageResponseBody) SetRemarksB(v string) *UpdateImageResponseBody {
	s.RemarksB = &v
	return s
}

func (s *UpdateImageResponseBody) SetRemarksC(v string) *UpdateImageResponseBody {
	s.RemarksC = &v
	return s
}

func (s *UpdateImageResponseBody) SetRemarksD(v string) *UpdateImageResponseBody {
	s.RemarksD = &v
	return s
}

func (s *UpdateImageResponseBody) SetRequestId(v string) *UpdateImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateImageResponseBody) SetSetId(v string) *UpdateImageResponseBody {
	s.SetId = &v
	return s
}

type UpdateImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateImageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageResponse) GoString() string {
	return s.String()
}

func (s *UpdateImageResponse) SetHeaders(v map[string]*string) *UpdateImageResponse {
	s.Headers = v
	return s
}

func (s *UpdateImageResponse) SetStatusCode(v int32) *UpdateImageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateImageResponse) SetBody(v *UpdateImageResponseBody) *UpdateImageResponse {
	s.Body = v
	return s
}

type UpdateProjectRequest struct {
	NewCU          *int32  `json:"NewCU,omitempty" xml:"NewCU,omitempty"`
	NewServiceRole *string `json:"NewServiceRole,omitempty" xml:"NewServiceRole,omitempty"`
	Project        *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) SetNewCU(v int32) *UpdateProjectRequest {
	s.NewCU = &v
	return s
}

func (s *UpdateProjectRequest) SetNewServiceRole(v string) *UpdateProjectRequest {
	s.NewServiceRole = &v
	return s
}

func (s *UpdateProjectRequest) SetProject(v string) *UpdateProjectRequest {
	s.Project = &v
	return s
}

type UpdateProjectResponseBody struct {
	CU          *int32  `json:"CU,omitempty" xml:"CU,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifyTime  *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Project     *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBody) SetCU(v int32) *UpdateProjectResponseBody {
	s.CU = &v
	return s
}

func (s *UpdateProjectResponseBody) SetCreateTime(v string) *UpdateProjectResponseBody {
	s.CreateTime = &v
	return s
}

func (s *UpdateProjectResponseBody) SetModifyTime(v string) *UpdateProjectResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *UpdateProjectResponseBody) SetProject(v string) *UpdateProjectResponseBody {
	s.Project = &v
	return s
}

func (s *UpdateProjectResponseBody) SetRegionId(v string) *UpdateProjectResponseBody {
	s.RegionId = &v
	return s
}

func (s *UpdateProjectResponseBody) SetRequestId(v string) *UpdateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectResponseBody) SetServiceRole(v string) *UpdateProjectResponseBody {
	s.ServiceRole = &v
	return s
}

func (s *UpdateProjectResponseBody) SetType(v string) *UpdateProjectResponseBody {
	s.Type = &v
	return s
}

type UpdateProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponse) SetHeaders(v map[string]*string) *UpdateProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectResponse) SetStatusCode(v int32) *UpdateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectResponse) SetBody(v *UpdateProjectResponseBody) *UpdateProjectResponse {
	s.Body = v
	return s
}

type UpdateSetRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	SetId   *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	SetName *string `json:"SetName,omitempty" xml:"SetName,omitempty"`
}

func (s UpdateSetRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSetRequest) GoString() string {
	return s.String()
}

func (s *UpdateSetRequest) SetProject(v string) *UpdateSetRequest {
	s.Project = &v
	return s
}

func (s *UpdateSetRequest) SetSetId(v string) *UpdateSetRequest {
	s.SetId = &v
	return s
}

func (s *UpdateSetRequest) SetSetName(v string) *UpdateSetRequest {
	s.SetName = &v
	return s
}

type UpdateSetResponseBody struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SetId      *string `json:"SetId,omitempty" xml:"SetId,omitempty"`
	SetName    *string `json:"SetName,omitempty" xml:"SetName,omitempty"`
}

func (s UpdateSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSetResponseBody) SetCreateTime(v string) *UpdateSetResponseBody {
	s.CreateTime = &v
	return s
}

func (s *UpdateSetResponseBody) SetModifyTime(v string) *UpdateSetResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *UpdateSetResponseBody) SetRequestId(v string) *UpdateSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSetResponseBody) SetSetId(v string) *UpdateSetResponseBody {
	s.SetId = &v
	return s
}

func (s *UpdateSetResponseBody) SetSetName(v string) *UpdateSetResponseBody {
	s.SetName = &v
	return s
}

type UpdateSetResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSetResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSetResponse) GoString() string {
	return s.String()
}

func (s *UpdateSetResponse) SetHeaders(v map[string]*string) *UpdateSetResponse {
	s.Headers = v
	return s
}

func (s *UpdateSetResponse) SetStatusCode(v int32) *UpdateSetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSetResponse) SetBody(v *UpdateSetResponseBody) *UpdateSetResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-beijing-gov-1": tea.String("imm-vpc.cn-beijing-gov-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("imm"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompareImageFacesWithOptions(request *CompareImageFacesRequest, runtime *util.RuntimeOptions) (_result *CompareImageFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FaceIdA)) {
		query["FaceIdA"] = request.FaceIdA
	}

	if !tea.BoolValue(util.IsUnset(request.FaceIdB)) {
		query["FaceIdB"] = request.FaceIdB
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUriA)) {
		query["ImageUriA"] = request.ImageUriA
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUriB)) {
		query["ImageUriB"] = request.ImageUriB
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.SetId)) {
		query["SetId"] = request.SetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CompareImageFaces"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CompareImageFacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompareImageFaces(request *CompareImageFacesRequest) (_result *CompareImageFacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompareImageFacesResponse{}
	_body, _err := client.CompareImageFacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConvertOfficeFormatWithOptions(request *ConvertOfficeFormatRequest, runtime *util.RuntimeOptions) (_result *ConvertOfficeFormatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndPage)) {
		query["EndPage"] = request.EndPage
	}

	if !tea.BoolValue(util.IsUnset(request.FitToPagesTall)) {
		query["FitToPagesTall"] = request.FitToPagesTall
	}

	if !tea.BoolValue(util.IsUnset(request.FitToPagesWide)) {
		query["FitToPagesWide"] = request.FitToPagesWide
	}

	if !tea.BoolValue(util.IsUnset(request.Hidecomments)) {
		query["Hidecomments"] = request.Hidecomments
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSheetCol)) {
		query["MaxSheetCol"] = request.MaxSheetCol
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSheetCount)) {
		query["MaxSheetCount"] = request.MaxSheetCount
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSheetRow)) {
		query["MaxSheetRow"] = request.MaxSheetRow
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["ModelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.PdfVector)) {
		query["PdfVector"] = request.PdfVector
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.SheetOnePage)) {
		query["SheetOnePage"] = request.SheetOnePage
	}

	if !tea.BoolValue(util.IsUnset(request.SrcType)) {
		query["SrcType"] = request.SrcType
	}

	if !tea.BoolValue(util.IsUnset(request.SrcUri)) {
		query["SrcUri"] = request.SrcUri
	}

	if !tea.BoolValue(util.IsUnset(request.StartPage)) {
		query["StartPage"] = request.StartPage
	}

	if !tea.BoolValue(util.IsUnset(request.TgtFilePages)) {
		query["TgtFilePages"] = request.TgtFilePages
	}

	if !tea.BoolValue(util.IsUnset(request.TgtFilePrefix)) {
		query["TgtFilePrefix"] = request.TgtFilePrefix
	}

	if !tea.BoolValue(util.IsUnset(request.TgtFileSuffix)) {
		query["TgtFileSuffix"] = request.TgtFileSuffix
	}

	if !tea.BoolValue(util.IsUnset(request.TgtType)) {
		query["TgtType"] = request.TgtType
	}

	if !tea.BoolValue(util.IsUnset(request.TgtUri)) {
		query["TgtUri"] = request.TgtUri
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConvertOfficeFormat"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConvertOfficeFormatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConvertOfficeFormat(request *ConvertOfficeFormatRequest) (_result *ConvertOfficeFormatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConvertOfficeFormatResponse{}
	_body, _err := client.ConvertOfficeFormatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGrabFrameTaskWithOptions(request *CreateGrabFrameTaskRequest, runtime *util.RuntimeOptions) (_result *CreateGrabFrameTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomMessage)) {
		query["CustomMessage"] = request.CustomMessage
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyEndpoint)) {
		query["NotifyEndpoint"] = request.NotifyEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.TargetList)) {
		query["TargetList"] = request.TargetList
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUri)) {
		query["VideoUri"] = request.VideoUri
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGrabFrameTask"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGrabFrameTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGrabFrameTask(request *CreateGrabFrameTaskRequest) (_result *CreateGrabFrameTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGrabFrameTaskResponse{}
	_body, _err := client.CreateGrabFrameTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGroupFacesJobWithOptions(request *CreateGroupFacesJobRequest, runtime *util.RuntimeOptions) (_result *CreateGroupFacesJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NotifyEndpoint)) {
		query["NotifyEndpoint"] = request.NotifyEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.SetId)) {
		query["SetId"] = request.SetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGroupFacesJob"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupFacesJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGroupFacesJob(request *CreateGroupFacesJobRequest) (_result *CreateGroupFacesJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGroupFacesJobResponse{}
	_body, _err := client.CreateGroupFacesJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMergeFaceGroupsJobWithOptions(request *CreateMergeFaceGroupsJobRequest, runtime *util.RuntimeOptions) (_result *CreateMergeFaceGroupsJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomMessage)) {
		query["CustomMessage"] = request.CustomMessage
	}

	if !tea.BoolValue(util.IsUnset(request.GroupIdFrom)) {
		query["GroupIdFrom"] = request.GroupIdFrom
	}

	if !tea.BoolValue(util.IsUnset(request.GroupIdTo)) {
		query["GroupIdTo"] = request.GroupIdTo
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyEndpoint)) {
		query["NotifyEndpoint"] = request.NotifyEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.SetId)) {
		query["SetId"] = request.SetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMergeFaceGroupsJob"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMergeFaceGroupsJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMergeFaceGroupsJob(request *CreateMergeFaceGroupsJobRequest) (_result *CreateMergeFaceGroupsJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMergeFaceGroupsJobResponse{}
	_body, _err := client.CreateMergeFaceGroupsJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOfficeConversionTaskWithOptions(request *CreateOfficeConversionTaskRequest, runtime *util.RuntimeOptions) (_result *CreateOfficeConversionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayDpi)) {
		query["DisplayDpi"] = request.DisplayDpi
	}

	if !tea.BoolValue(util.IsUnset(request.EndPage)) {
		query["EndPage"] = request.EndPage
	}

	if !tea.BoolValue(util.IsUnset(request.FitToPagesTall)) {
		query["FitToPagesTall"] = request.FitToPagesTall
	}

	if !tea.BoolValue(util.IsUnset(request.FitToPagesWide)) {
		query["FitToPagesWide"] = request.FitToPagesWide
	}

	if !tea.BoolValue(util.IsUnset(request.Hidecomments)) {
		query["Hidecomments"] = request.Hidecomments
	}

	if !tea.BoolValue(util.IsUnset(request.IdempotentToken)) {
		query["IdempotentToken"] = request.IdempotentToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSheetCol)) {
		query["MaxSheetCol"] = request.MaxSheetCol
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSheetCount)) {
		query["MaxSheetCount"] = request.MaxSheetCount
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSheetRow)) {
		query["MaxSheetRow"] = request.MaxSheetRow
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["ModelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyEndpoint)) {
		query["NotifyEndpoint"] = request.NotifyEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.PdfVector)) {
		query["PdfVector"] = request.PdfVector
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.SheetOnePage)) {
		query["SheetOnePage"] = request.SheetOnePage
	}

	if !tea.BoolValue(util.IsUnset(request.SrcType)) {
		query["SrcType"] = request.SrcType
	}

	if !tea.BoolValue(util.IsUnset(request.SrcUri)) {
		query["SrcUri"] = request.SrcUri
	}

	if !tea.BoolValue(util.IsUnset(request.StartPage)) {
		query["StartPage"] = request.StartPage
	}

	if !tea.BoolValue(util.IsUnset(request.TgtFilePages)) {
		query["TgtFilePages"] = request.TgtFilePages
	}

	if !tea.BoolValue(util.IsUnset(request.TgtFilePrefix)) {
		query["TgtFilePrefix"] = request.TgtFilePrefix
	}

	if !tea.BoolValue(util.IsUnset(request.TgtFileSuffix)) {
		query["TgtFileSuffix"] = request.TgtFileSuffix
	}

	if !tea.BoolValue(util.IsUnset(request.TgtType)) {
		query["TgtType"] = request.TgtType
	}

	if !tea.BoolValue(util.IsUnset(request.TgtUri)) {
		query["TgtUri"] = request.TgtUri
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		query["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOfficeConversionTask"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOfficeConversionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOfficeConversionTask(request *CreateOfficeConversionTaskRequest) (_result *CreateOfficeConversionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOfficeConversionTaskResponse{}
	_body, _err := client.CreateOfficeConversionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSetWithOptions(request *CreateSetRequest, runtime *util.RuntimeOptions) (_result *CreateSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.SetId)) {
		query["SetId"] = request.SetId
	}

	if !tea.BoolValue(util.IsUnset(request.SetName)) {
		query["SetName"] = request.SetName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSet"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSet(request *CreateSetRequest) (_result *CreateSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSetResponse{}
	_body, _err := client.CreateSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVideoCompressTaskWithOptions(request *CreateVideoCompressTaskRequest, runtime *util.RuntimeOptions) (_result *CreateVideoCompressTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomMessage)) {
		query["CustomMessage"] = request.CustomMessage
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyEndpoint)) {
		query["NotifyEndpoint"] = request.NotifyEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.TargetList)) {
		query["TargetList"] = request.TargetList
	}

	if !tea.BoolValue(util.IsUnset(request.TargetSegment)) {
		query["TargetSegment"] = request.TargetSegment
	}

	if !tea.BoolValue(util.IsUnset(request.TargetSubtitle)) {
		query["TargetSubtitle"] = request.TargetSubtitle
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUri)) {
		query["VideoUri"] = request.VideoUri
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVideoCompressTask"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVideoCompressTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVideoCompressTask(request *CreateVideoCompressTaskRequest) (_result *CreateVideoCompressTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVideoCompressTaskResponse{}
	_body, _err := client.CreateVideoCompressTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DecodeBlindWatermarkWithOptions(request *DecodeBlindWatermarkRequest, runtime *util.RuntimeOptions) (_result *DecodeBlindWatermarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageQuality)) {
		query["ImageQuality"] = request.ImageQuality
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUri)) {
		query["ImageUri"] = request.ImageUri
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		query["Model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalImageUri)) {
		query["OriginalImageUri"] = request.OriginalImageUri
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.TargetUri)) {
		query["TargetUri"] = request.TargetUri
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DecodeBlindWatermark"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DecodeBlindWatermarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DecodeBlindWatermark(request *DecodeBlindWatermarkRequest) (_result *DecodeBlindWatermarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DecodeBlindWatermarkResponse{}
	_body, _err := client.DecodeBlindWatermarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImageWithOptions(request *DeleteImageRequest, runtime *util.RuntimeOptions) (_result *DeleteImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUri)) {
		query["ImageUri"] = request.ImageUri
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.SetId)) {
		query["SetId"] = request.SetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteImage"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImage(request *DeleteImageRequest) (_result *DeleteImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImageResponse{}
	_body, _err := client.DeleteImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteOfficeConversionTaskWithOptions(request *DeleteOfficeConversionTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteOfficeConversionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteOfficeConversionTask"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteOfficeConversionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteOfficeConversionTask(request *DeleteOfficeConversionTaskRequest) (_result *DeleteOfficeConversionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteOfficeConversionTaskResponse{}
	_body, _err := client.DeleteOfficeConversionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProjectWithOptions(request *DeleteProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProject"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProject(request *DeleteProjectRequest) (_result *DeleteProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSetWithOptions(request *DeleteSetRequest, runtime *util.RuntimeOptions) (_result *DeleteSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.SetId)) {
		query["SetId"] = request.SetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSet"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSet(request *DeleteSetRequest) (_result *DeleteSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSetResponse{}
	_body, _err := client.DeleteSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVideoWithOptions(request *DeleteVideoRequest, runtime *util.RuntimeOptions) (_result *DeleteVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.Resources)) {
		query["Resources"] = request.Resources
	}

	if !tea.BoolValue(util.IsUnset(request.SetId)) {
		query["SetId"] = request.SetId
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUri)) {
		query["VideoUri"] = request.VideoUri
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVideo"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVideoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVideo(request *DeleteVideoRequest) (_result *DeleteVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVideoResponse{}
	_body, _err := client.DeleteVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVideoTaskWithOptions(request *DeleteVideoTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteVideoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVideoTask"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVideoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVideoTask(request *DeleteVideoTaskRequest) (_result *DeleteVideoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVideoTaskResponse{}
	_body, _err := client.DeleteVideoTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectImageBodiesWithOptions(request *DetectImageBodiesRequest, runtime *util.RuntimeOptions) (_result *DetectImageBodiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUri)) {
		query["ImageUri"] = request.ImageUri
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectImageBodies"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectImageBodiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectImageBodies(request *DetectImageBodiesRequest) (_result *DetectImageBodiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectImageBodiesResponse{}
	_body, _err := client.DetectImageBodiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectImageFacesWithOptions(request *DetectImageFacesRequest, runtime *util.RuntimeOptions) (_result *DetectImageFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUri)) {
		query["ImageUri"] = request.ImageUri
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectImageFaces"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectImageFacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectImageFaces(request *DetectImageFacesRequest) (_result *DetectImageFacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectImageFacesResponse{}
	_body, _err := client.DetectImageFacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectImageQRCodesWithOptions(request *DetectImageQRCodesRequest, runtime *util.RuntimeOptions) (_result *DetectImageQRCodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUri)) {
		query["ImageUri"] = request.ImageUri
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectImageQRCodes"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectImageQRCodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectImageQRCodes(request *DetectImageQRCodesRequest) (_result *DetectImageQRCodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectImageQRCodesResponse{}
	_body, _err := client.DetectImageQRCodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectImageTagsWithOptions(request *DetectImageTagsRequest, runtime *util.RuntimeOptions) (_result *DetectImageTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUri)) {
		query["ImageUri"] = request.ImageUri
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectImageTags"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectImageTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectImageTags(request *DetectImageTagsRequest) (_result *DetectImageTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectImageTagsResponse{}
	_body, _err := client.DetectImageTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectQRCodesWithOptions(request *DetectQRCodesRequest, runtime *util.RuntimeOptions) (_result *DetectQRCodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.SrcUris)) {
		query["SrcUris"] = request.SrcUris
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectQRCodes"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectQRCodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectQRCodes(request *DetectQRCodesRequest) (_result *DetectQRCodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectQRCodesResponse{}
	_body, _err := client.DetectQRCodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EncodeBlindWatermarkWithOptions(request *EncodeBlindWatermarkRequest, runtime *util.RuntimeOptions) (_result *EncodeBlindWatermarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ImageQuality)) {
		query["ImageQuality"] = request.ImageQuality
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUri)) {
		query["ImageUri"] = request.ImageUri
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		query["Model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.TargetImageType)) {
		query["TargetImageType"] = request.TargetImageType
	}

	if !tea.BoolValue(util.IsUnset(request.TargetUri)) {
		query["TargetUri"] = request.TargetUri
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkUri)) {
		query["WatermarkUri"] = request.WatermarkUri
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EncodeBlindWatermark"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EncodeBlindWatermarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EncodeBlindWatermark(request *EncodeBlindWatermarkRequest) (_result *EncodeBlindWatermarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EncodeBlindWatermarkResponse{}
	_body, _err := client.EncodeBlindWatermarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindImagesWithOptions(request *FindImagesRequest, runtime *util.RuntimeOptions) (_result *FindImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressLineContentsMatch)) {
		query["AddressLineContentsMatch"] = request.AddressLineContentsMatch
	}

	if !tea.BoolValue(util.IsUnset(request.AgeRange)) {
		query["AgeRange"] = request.AgeRange
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeRange)) {
		query["CreateTimeRange"] = request.CreateTimeRange
	}

	if !tea.BoolValue(util.IsUnset(request.Emotion)) {
		query["Emotion"] = request.Emotion
	}

	if !tea.BoolValue(util.IsUnset(request.ExternalId)) {
		query["ExternalId"] = request.ExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.FacesModifyTimeRange)) {
		query["FacesModifyTimeRange"] = request.FacesModifyTimeRange
	}

	if !tea.BoolValue(util.IsUnset(request.Gender)) {
		query["Gender"] = request.Gender
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageSizeRange)) {
		query["ImageSizeRange"] = request.ImageSizeRange
	}

	if !tea.BoolValue(util.IsUnset(request.ImageTimeRange)) {
		query["ImageTimeRange"] = request.ImageTimeRange
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.LocationBoundary)) {
		query["LocationBoundary"] = request.LocationBoundary
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyTimeRange)) {
		query["ModifyTimeRange"] = request.ModifyTimeRange
	}

	if !tea.BoolValue(util.IsUnset(request.OCRContentsMatch)) {
		query["OCRContentsMatch"] = request.OCRContentsMatch
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksAPrefix)) {
		query["RemarksAPrefix"] = request.RemarksAPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksArrayAIn)) {
		query["RemarksArrayAIn"] = request.RemarksArrayAIn
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksArrayBIn)) {
		query["RemarksArrayBIn"] = request.RemarksArrayBIn
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksBPrefix)) {
		query["RemarksBPrefix"] = request.RemarksBPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksCPrefix)) {
		query["RemarksCPrefix"] = request.RemarksCPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksDPrefix)) {
		query["RemarksDPrefix"] = request.RemarksDPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.SetId)) {
		query["SetId"] = request.SetId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceUriPrefix)) {
		query["SourceUriPrefix"] = request.SourceUriPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.TagNames)) {
		query["TagNames"] = request.TagNames
	}

	if !tea.BoolValue(util.IsUnset(request.TagsModifyTimeRange)) {
		query["TagsModifyTimeRange"] = request.TagsModifyTimeRange
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FindImages"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FindImages(request *FindImagesRequest) (_result *FindImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindImagesResponse{}
	_body, _err := client.FindImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindSimilarFacesWithOptions(request *FindSimilarFacesRequest, runtime *util.RuntimeOptions) (_result *FindSimilarFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FaceId)) {
		query["FaceId"] = request.FaceId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUri)) {
		query["ImageUri"] = request.ImageUri
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.MinSimilarity)) {
		query["MinSimilarity"] = request.MinSimilarity
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.ResponseFormat)) {
		query["ResponseFormat"] = request.ResponseFormat
	}

	if !tea.BoolValue(util.IsUnset(request.SetId)) {
		query["SetId"] = request.SetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FindSimilarFaces"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindSimilarFacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FindSimilarFaces(request *FindSimilarFacesRequest) (_result *FindSimilarFacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindSimilarFacesResponse{}
	_body, _err := client.FindSimilarFacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetImageWithOptions(request *GetImageRequest, runtime *util.RuntimeOptions) (_result *GetImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUri)) {
		query["ImageUri"] = request.ImageUri
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.SetId)) {
		query["SetId"] = request.SetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetImage"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetImage(request *GetImageRequest) (_result *GetImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetImageResponse{}
	_body, _err := client.GetImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetImageCroppingSuggestionsWithOptions(request *GetImageCroppingSuggestionsRequest, runtime *util.RuntimeOptions) (_result *GetImageCroppingSuggestionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AspectRatios)) {
		query["AspectRatios"] = request.AspectRatios
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUri)) {
		query["ImageUri"] = request.ImageUri
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetImageCroppingSuggestions"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetImageCroppingSuggestionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetImageCroppingSuggestions(request *GetImageCroppingSuggestionsRequest) (_result *GetImageCroppingSuggestionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetImageCroppingSuggestionsResponse{}
	_body, _err := client.GetImageCroppingSuggestionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetImageQualityWithOptions(request *GetImageQualityRequest, runtime *util.RuntimeOptions) (_result *GetImageQualityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUri)) {
		query["ImageUri"] = request.ImageUri
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetImageQuality"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetImageQualityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetImageQuality(request *GetImageQualityRequest) (_result *GetImageQualityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetImageQualityResponse{}
	_body, _err := client.GetImageQualityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMediaMetaWithOptions(request *GetMediaMetaRequest, runtime *util.RuntimeOptions) (_result *GetMediaMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MediaUri)) {
		query["MediaUri"] = request.MediaUri
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMediaMeta"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMediaMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMediaMeta(request *GetMediaMetaRequest) (_result *GetMediaMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMediaMetaResponse{}
	_body, _err := client.GetMediaMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOfficeConversionTaskWithOptions(request *GetOfficeConversionTaskRequest, runtime *util.RuntimeOptions) (_result *GetOfficeConversionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOfficeConversionTask"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOfficeConversionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOfficeConversionTask(request *GetOfficeConversionTaskRequest) (_result *GetOfficeConversionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOfficeConversionTaskResponse{}
	_body, _err := client.GetOfficeConversionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOfficePreviewURLWithOptions(request *GetOfficePreviewURLRequest, runtime *util.RuntimeOptions) (_result *GetOfficePreviewURLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.SrcType)) {
		query["SrcType"] = request.SrcType
	}

	if !tea.BoolValue(util.IsUnset(request.SrcUri)) {
		query["SrcUri"] = request.SrcUri
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkFillStyle)) {
		query["WatermarkFillStyle"] = request.WatermarkFillStyle
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkFont)) {
		query["WatermarkFont"] = request.WatermarkFont
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkHorizontal)) {
		query["WatermarkHorizontal"] = request.WatermarkHorizontal
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkRotate)) {
		query["WatermarkRotate"] = request.WatermarkRotate
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkType)) {
		query["WatermarkType"] = request.WatermarkType
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkValue)) {
		query["WatermarkValue"] = request.WatermarkValue
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkVertical)) {
		query["WatermarkVertical"] = request.WatermarkVertical
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOfficePreviewURL"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOfficePreviewURLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOfficePreviewURL(request *GetOfficePreviewURLRequest) (_result *GetOfficePreviewURLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOfficePreviewURLResponse{}
	_body, _err := client.GetOfficePreviewURLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectWithOptions(request *GetProjectRequest, runtime *util.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProject"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProject(request *GetProjectRequest) (_result *GetProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSetWithOptions(request *GetSetRequest, runtime *util.RuntimeOptions) (_result *GetSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.SetId)) {
		query["SetId"] = request.SetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSet"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSet(request *GetSetRequest) (_result *GetSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSetResponse{}
	_body, _err := client.GetSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVideoWithOptions(request *GetVideoRequest, runtime *util.RuntimeOptions) (_result *GetVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.SetId)) {
		query["SetId"] = request.SetId
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUri)) {
		query["VideoUri"] = request.VideoUri
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVideo"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVideoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVideo(request *GetVideoRequest) (_result *GetVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVideoResponse{}
	_body, _err := client.GetVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVideoTaskWithOptions(request *GetVideoTaskRequest, runtime *util.RuntimeOptions) (_result *GetVideoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVideoTask"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVideoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVideoTask(request *GetVideoTaskRequest) (_result *GetVideoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVideoTaskResponse{}
	_body, _err := client.GetVideoTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWebofficeURLWithOptions(request *GetWebofficeURLRequest, runtime *util.RuntimeOptions) (_result *GetWebofficeURLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.File)) {
		query["File"] = request.File
	}

	if !tea.BoolValue(util.IsUnset(request.FileID)) {
		query["FileID"] = request.FileID
	}

	if !tea.BoolValue(util.IsUnset(request.Hidecmb)) {
		query["Hidecmb"] = request.Hidecmb
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyEndpoint)) {
		query["NotifyEndpoint"] = request.NotifyEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.Permission)) {
		query["Permission"] = request.Permission
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.SrcType)) {
		query["SrcType"] = request.SrcType
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	if !tea.BoolValue(util.IsUnset(request.Watermark)) {
		query["Watermark"] = request.Watermark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWebofficeURL"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWebofficeURLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWebofficeURL(request *GetWebofficeURLRequest) (_result *GetWebofficeURLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWebofficeURLResponse{}
	_body, _err := client.GetWebofficeURLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndexImageWithOptions(request *IndexImageRequest, runtime *util.RuntimeOptions) (_result *IndexImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExternalId)) {
		query["ExternalId"] = request.ExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUri)) {
		query["ImageUri"] = request.ImageUri
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyEndpoint)) {
		query["NotifyEndpoint"] = request.NotifyEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksA)) {
		query["RemarksA"] = request.RemarksA
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksArrayA)) {
		query["RemarksArrayA"] = request.RemarksArrayA
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksArrayB)) {
		query["RemarksArrayB"] = request.RemarksArrayB
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksB)) {
		query["RemarksB"] = request.RemarksB
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksC)) {
		query["RemarksC"] = request.RemarksC
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksD)) {
		query["RemarksD"] = request.RemarksD
	}

	if !tea.BoolValue(util.IsUnset(request.SetId)) {
		query["SetId"] = request.SetId
	}

	if !tea.BoolValue(util.IsUnset(request.SourcePosition)) {
		query["SourcePosition"] = request.SourcePosition
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceUri)) {
		query["SourceUri"] = request.SourceUri
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("IndexImage"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &IndexImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndexImage(request *IndexImageRequest) (_result *IndexImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IndexImageResponse{}
	_body, _err := client.IndexImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndexVideoWithOptions(request *IndexVideoRequest, runtime *util.RuntimeOptions) (_result *IndexVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExternalId)) {
		query["ExternalId"] = request.ExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyEndpoint)) {
		query["NotifyEndpoint"] = request.NotifyEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksA)) {
		query["RemarksA"] = request.RemarksA
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksB)) {
		query["RemarksB"] = request.RemarksB
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksC)) {
		query["RemarksC"] = request.RemarksC
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksD)) {
		query["RemarksD"] = request.RemarksD
	}

	if !tea.BoolValue(util.IsUnset(request.SetId)) {
		query["SetId"] = request.SetId
	}

	if !tea.BoolValue(util.IsUnset(request.TgtUri)) {
		query["TgtUri"] = request.TgtUri
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUri)) {
		query["VideoUri"] = request.VideoUri
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("IndexVideo"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &IndexVideoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndexVideo(request *IndexVideoRequest) (_result *IndexVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IndexVideoResponse{}
	_body, _err := client.IndexVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFaceGroupsWithOptions(request *ListFaceGroupsRequest, runtime *util.RuntimeOptions) (_result *ListFaceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExternalId)) {
		query["ExternalId"] = request.ExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksAQuery)) {
		query["RemarksAQuery"] = request.RemarksAQuery
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksArrayAQuery)) {
		query["RemarksArrayAQuery"] = request.RemarksArrayAQuery
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksArrayBQuery)) {
		query["RemarksArrayBQuery"] = request.RemarksArrayBQuery
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksBQuery)) {
		query["RemarksBQuery"] = request.RemarksBQuery
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksCQuery)) {
		query["RemarksCQuery"] = request.RemarksCQuery
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksDQuery)) {
		query["RemarksDQuery"] = request.RemarksDQuery
	}

	if !tea.BoolValue(util.IsUnset(request.SetId)) {
		query["SetId"] = request.SetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFaceGroups"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFaceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFaceGroups(request *ListFaceGroupsRequest) (_result *ListFaceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFaceGroupsResponse{}
	_body, _err := client.ListFaceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListImagesWithOptions(request *ListImagesRequest, runtime *util.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateTimeStart)) {
		query["CreateTimeStart"] = request.CreateTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.SetId)) {
		query["SetId"] = request.SetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListImages"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOfficeConversionTaskWithOptions(request *ListOfficeConversionTaskRequest, runtime *util.RuntimeOptions) (_result *ListOfficeConversionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxKeys)) {
		query["MaxKeys"] = request.MaxKeys
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOfficeConversionTask"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOfficeConversionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOfficeConversionTask(request *ListOfficeConversionTaskRequest) (_result *ListOfficeConversionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOfficeConversionTaskResponse{}
	_body, _err := client.ListOfficeConversionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectsWithOptions(request *ListProjectsRequest, runtime *util.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxKeys)) {
		query["MaxKeys"] = request.MaxKeys
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjects"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProjects(request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSetTagsWithOptions(request *ListSetTagsRequest, runtime *util.RuntimeOptions) (_result *ListSetTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.SetId)) {
		query["SetId"] = request.SetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSetTags"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSetTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSetTags(request *ListSetTagsRequest) (_result *ListSetTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSetTagsResponse{}
	_body, _err := client.ListSetTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSetsWithOptions(request *ListSetsRequest, runtime *util.RuntimeOptions) (_result *ListSetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSets"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSets(request *ListSetsRequest) (_result *ListSetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSetsResponse{}
	_body, _err := client.ListSetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVideoTasksWithOptions(request *ListVideoTasksRequest, runtime *util.RuntimeOptions) (_result *ListVideoTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxKeys)) {
		query["MaxKeys"] = request.MaxKeys
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVideoTasks"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVideoTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVideoTasks(request *ListVideoTasksRequest) (_result *ListVideoTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVideoTasksResponse{}
	_body, _err := client.ListVideoTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVideosWithOptions(request *ListVideosRequest, runtime *util.RuntimeOptions) (_result *ListVideosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateTimeStart)) {
		query["CreateTimeStart"] = request.CreateTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.SetId)) {
		query["SetId"] = request.SetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVideos"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVideosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVideos(request *ListVideosRequest) (_result *ListVideosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVideosResponse{}
	_body, _err := client.ListVideosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenImmServiceWithOptions(request *OpenImmServiceRequest, runtime *util.RuntimeOptions) (_result *OpenImmServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenImmService"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenImmServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenImmService(request *OpenImmServiceRequest) (_result *OpenImmServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenImmServiceResponse{}
	_body, _err := client.OpenImmServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutProjectWithOptions(request *PutProjectRequest, runtime *util.RuntimeOptions) (_result *PutProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceRole)) {
		query["ServiceRole"] = request.ServiceRole
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutProject"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutProject(request *PutProjectRequest) (_result *PutProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutProjectResponse{}
	_body, _err := client.PutProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshOfficePreviewTokenWithOptions(request *RefreshOfficePreviewTokenRequest, runtime *util.RuntimeOptions) (_result *RefreshOfficePreviewTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.RefreshToken)) {
		query["RefreshToken"] = request.RefreshToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshOfficePreviewToken"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshOfficePreviewTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshOfficePreviewToken(request *RefreshOfficePreviewTokenRequest) (_result *RefreshOfficePreviewTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshOfficePreviewTokenResponse{}
	_body, _err := client.RefreshOfficePreviewTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshWebofficeTokenWithOptions(request *RefreshWebofficeTokenRequest, runtime *util.RuntimeOptions) (_result *RefreshWebofficeTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.RefreshToken)) {
		query["RefreshToken"] = request.RefreshToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshWebofficeToken"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshWebofficeTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshWebofficeToken(request *RefreshWebofficeTokenRequest) (_result *RefreshWebofficeTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshWebofficeTokenResponse{}
	_body, _err := client.RefreshWebofficeTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFaceGroupWithOptions(request *UpdateFaceGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateFaceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExternalId)) {
		query["ExternalId"] = request.ExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupCoverFaceId)) {
		query["GroupCoverFaceId"] = request.GroupCoverFaceId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksA)) {
		query["RemarksA"] = request.RemarksA
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksArrayA)) {
		query["RemarksArrayA"] = request.RemarksArrayA
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksArrayB)) {
		query["RemarksArrayB"] = request.RemarksArrayB
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksB)) {
		query["RemarksB"] = request.RemarksB
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksC)) {
		query["RemarksC"] = request.RemarksC
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksD)) {
		query["RemarksD"] = request.RemarksD
	}

	if !tea.BoolValue(util.IsUnset(request.ResetItems)) {
		query["ResetItems"] = request.ResetItems
	}

	if !tea.BoolValue(util.IsUnset(request.SetId)) {
		query["SetId"] = request.SetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFaceGroup"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFaceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFaceGroup(request *UpdateFaceGroupRequest) (_result *UpdateFaceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFaceGroupResponse{}
	_body, _err := client.UpdateFaceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateImageWithOptions(tmpReq *UpdateImageRequest, runtime *util.RuntimeOptions) (_result *UpdateImageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateImageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Faces)) {
		request.FacesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Faces, tea.String("Faces"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExternalId)) {
		query["ExternalId"] = request.ExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.FacesShrink)) {
		query["Faces"] = request.FacesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUri)) {
		query["ImageUri"] = request.ImageUri
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksA)) {
		query["RemarksA"] = request.RemarksA
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksArrayA)) {
		query["RemarksArrayA"] = request.RemarksArrayA
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksArrayB)) {
		query["RemarksArrayB"] = request.RemarksArrayB
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksB)) {
		query["RemarksB"] = request.RemarksB
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksC)) {
		query["RemarksC"] = request.RemarksC
	}

	if !tea.BoolValue(util.IsUnset(request.RemarksD)) {
		query["RemarksD"] = request.RemarksD
	}

	if !tea.BoolValue(util.IsUnset(request.SetId)) {
		query["SetId"] = request.SetId
	}

	if !tea.BoolValue(util.IsUnset(request.SourcePosition)) {
		query["SourcePosition"] = request.SourcePosition
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceUri)) {
		query["SourceUri"] = request.SourceUri
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateImage"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateImage(request *UpdateImageRequest) (_result *UpdateImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateImageResponse{}
	_body, _err := client.UpdateImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProjectWithOptions(request *UpdateProjectRequest, runtime *util.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewCU)) {
		query["NewCU"] = request.NewCU
	}

	if !tea.BoolValue(util.IsUnset(request.NewServiceRole)) {
		query["NewServiceRole"] = request.NewServiceRole
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProject"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProject(request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSetWithOptions(request *UpdateSetRequest, runtime *util.RuntimeOptions) (_result *UpdateSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.SetId)) {
		query["SetId"] = request.SetId
	}

	if !tea.BoolValue(util.IsUnset(request.SetName)) {
		query["SetName"] = request.SetName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSet"),
		Version:     tea.String("2017-09-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSet(request *UpdateSetRequest) (_result *UpdateSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSetResponse{}
	_body, _err := client.UpdateSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
