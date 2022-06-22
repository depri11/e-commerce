package helper

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func UploadImages(ext string, file multipart.File, handle *multipart.FileHeader) (*uploader.UploadResult, error) {
	cld, err := cloudinary.NewFromParams("dlyp1s66j", "424186382842538", "nsffCHZoLCYSbEkgn-TptP2EL9c")
	if err != nil {
		return nil, err
	}

	var ctx = context.Background()
	result, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: "e-commerce/" + ext + "/" + handle.Filename,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
