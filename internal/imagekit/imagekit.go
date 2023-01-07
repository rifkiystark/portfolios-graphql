package imagekit

import (
	"context"

	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
)

type ImageKit interface {
	Upload(ctx context.Context, file interface{}, fileName string) (string, string, error)
	Delete(ctx context.Context, fileId string) error
}

type ImageKitImpl struct {
	ik *imagekit.ImageKit
}

func InitImageKit(publicKey string, privateKey string, url string) ImageKit {
	ik := imagekit.NewFromParams(imagekit.NewParams{
		PublicKey:   publicKey,
		PrivateKey:  privateKey,
		UrlEndpoint: url,
	})

	return &ImageKitImpl{ik: ik}
}

func (i *ImageKitImpl) Upload(ctx context.Context, file interface{}, fileName string) (string, string, error) {
	resp, err := i.ik.Uploader.Upload(ctx, file, uploader.UploadParam{
		FileName: fileName,
	})
	if err != nil {
		return "", "", err
	}

	return resp.Data.FileId, resp.Data.Url, nil
}

func (i *ImageKitImpl) Delete(ctx context.Context, fileId string) error {
	_, err := i.ik.Media.DeleteFile(ctx, fileId)
	if err != nil {
		return err
	}
	return nil
}
