package infra

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"

	"github.com/prokosna/medusa_synapse/domain"
	"github.com/prokosna/medusa_synapse/exception"
)

type ValidatorJson struct {
}

func NewValidatorJson() *ValidatorJson {
	return &ValidatorJson{}
}

func (v *ValidatorJson) Validate(img domain.Image) error {
	err := validateProperties(img)
	if err != nil {
		return err
	}
	return validateDataFormat(img)
}

func validateProperties(img domain.Image) error {
	if img.CameraId == "" {
		return exception.NewBadRequestError("camera_id can not be empty")
	}
	if img.ImageId <= 0 {
		return exception.NewBadRequestError("image_id must be greater than 1")
	}
	if img.Data == "" {
		return exception.NewBadRequestError("data can not be empty")
	}
	if img.Timestamp.IsZero() {
		return exception.NewBadRequestError("timestamp can not be empty")
	}
	return nil
}

func validateDataFormat(img domain.Image) error {
	result, err := base64.StdEncoding.DecodeString(img.Data)
	if err != nil {
		return exception.NewBadRequestError(err.Error())
	}
	_, format, err := image.DecodeConfig(bytes.NewReader(result))
	if format != "jpeg" {
		return exception.NewBadRequestError(fmt.Sprintf("jpeg format is expected as data, actual: %s", format))
	}
	return nil
}
