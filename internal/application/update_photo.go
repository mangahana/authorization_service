package application

import (
	"authorization_service/internal/core/cerror"
	"authorization_service/internal/core/models"
	"bytes"
	"context"
	"image"
	"image/jpeg"
	"image/png"
	"math"
	"net/http"
	"slices"

	"golang.org/x/image/draw"
	"golang.org/x/image/webp"
)

var supportedFormats = []string{"image/jpeg", "image/png", "image/webp"}

func (u *useCase) UpdatePhoto(c context.Context, user *models.UserSession, file []byte) error {
	mime := http.DetectContentType(file)

	if !slices.Contains(supportedFormats, mime) {
		return cerror.New(cerror.UNSUPPORTED_FORMAT, "UNSUPPORTED_FORMAT")
	}

	resized, err := thumbnail(file, 512)
	if err != nil {
		return err
	}

	filename, err := u.s3.Put(c, resized)
	if err != nil {
		return err
	}

	if err := u.repo.UpdatePhoto(c, user.ID, filename); err != nil {
		return err
	}

	input := models.UpdateUserEvent{
		ID:       user.ID,
		Username: user.Username,
		Photo:    filename,
	}
	u.amqp.SendUserUpdateEvent(input)

	return nil
}

func thumbnail(file []byte, width int) ([]byte, error) {
	var src image.Image
	var output bytes.Buffer
	var err error

	mimetype := http.DetectContentType(file)

	r := bytes.NewReader(file)

	switch mimetype {
	case "image/jpeg":
		src, err = jpeg.Decode(r)
	case "image/png":
		src, err = png.Decode(r)
	case "image/webp":
		src, err = webp.Decode(r)
	}

	if err != nil {
		return []byte{}, err
	}

	if src.Bounds().Dx() <= 512 {
		return file, nil
	}

	ratio := (float64)(src.Bounds().Max.Y) / (float64)(src.Bounds().Max.X)
	height := int(math.Round(float64(width) * ratio))

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	err = jpeg.Encode(&output, dst, nil)
	if err != nil {
		return []byte{}, err
	}

	return output.Bytes(), err
}
