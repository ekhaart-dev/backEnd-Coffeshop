package pkg

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/spf13/viper"
)

func CloudInary(file interface{}) (string, error) {
	name := viper.GetString("cloudinary.name")
	key := viper.GetString("cloudinary.key")
	secret := viper.GetString("cloudinary.sec")

	cld, _ := cloudinary.NewFromParams(name, key, secret)

	result, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{})
	if err != nil {
		return "", err
	}

	return result.URL, nil
}
