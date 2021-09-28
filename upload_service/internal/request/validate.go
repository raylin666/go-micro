package request

type ValidatorStreamUploadFileRequest struct {
	Stream []byte `label:"stream" validate:"required"`
}
