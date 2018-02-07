package attachment

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/astaxie/beego"
	"github.com/missdeer/kelly/modules/models"
	"github.com/missdeer/kelly/modules/utils"
	"github.com/missdeer/kelly/setting"
	"github.com/missdeer/kelly/upyun"
	"github.com/nfnt/resize"
)

func uploadToUpyun(imageLocalPath string, imageRemotePath string, result chan error) {
	if setting.UpYunEnabled {
		beego.Info("start uploading to upyun ", imageLocalPath)
		var upyunio *upyun.UpYun
		upyunio = upyun.NewUpYun(setting.UpYunBucketName, setting.UpYunUsername, setting.UpYunPassword)
		f, err := os.OpenFile(imageLocalPath, os.O_RDONLY, 0644)
		if err != nil {
			beego.Error("opening local saved path failed: ", err)
			return
		}
		defer f.Close()
		err = upyunio.WriteFile(imageRemotePath, f, true)
		if err != nil {
			beego.Error("writing file to UpYun failed: ", err)
		}
		result <- err
	} else {
		result <- nil
	}
	beego.Info("uploaded to upyun ", imageRemotePath)
}

func SaveImage(m *models.Image, r io.ReadSeeker, mime string, filename string, created time.Time) error {
	var ext string

	// test image mime type
	switch mime {
	case "image/jpeg":
		ext = ".jpg"

	case "image/png":
		ext = ".png"

	case "image/gif":
		ext = ".gif"

	default:
		ext = filepath.Ext(filename)
		switch ext {
		case ".jpg", ".png", ".gif":
		default:
			return fmt.Errorf("unsupport image format `%s`", filename)
		}
	}

	// decode image
	var img image.Image
	var err error
	switch ext {
	case ".jpg":
		m.Ext = 1
		img, err = jpeg.Decode(r)
	case ".png":
		m.Ext = 2
		img, err = png.Decode(r)
	case ".gif":
		m.Ext = 3
		img, err = gif.Decode(r)
	}

	if err != nil {
		return err
	}

	m.Width = img.Bounds().Dx()
	m.Height = img.Bounds().Dy()
	m.Created = created

	if err := m.Insert(); err != nil || m.Id <= 0 {
		return err
	}

	path := GenImagePath(m)
	os.MkdirAll(path, 0755)

	fullPath := GenImageFilePath(m, 0)
	var smallPath, middlePath string
	if _, err := r.Seek(0, 0); err != nil {
		return err
	}

	var file *os.File
	if f, err := os.OpenFile(fullPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644); err != nil {
		return err
	} else {
		file = f
	}
	defer file.Close()

	if _, err := io.Copy(file, r); err != nil {
		os.RemoveAll(fullPath)
		return err
	}

	var key = "upload" + m.LinkFull()

	uploadChannels := make(chan error, 6)

	go uploadToUpyun(fullPath, "/"+key, uploadChannels)
	var result error = nil

	if ext != ".gif" {
		if m.Width > setting.ImageSizeSmall {
			if err := ImageResize(m, img, setting.ImageSizeSmall); err != nil {
				<-uploadChannels
				<-uploadChannels
				os.RemoveAll(fullPath)
				return err
			}
			smallPath = GenImageFilePath(m, setting.ImageSizeSmall)
			key = "upload" + m.LinkSmall()
			go uploadToUpyun(smallPath, "/"+key, uploadChannels)
		}

		if m.Width > setting.ImageSizeMiddle {
			if err := ImageResize(m, img, setting.ImageSizeMiddle); err != nil {
				<-uploadChannels
				<-uploadChannels
				<-uploadChannels
				<-uploadChannels
				os.RemoveAll(fullPath)
				os.RemoveAll(smallPath)
				return err
			}
			middlePath = GenImageFilePath(m, setting.ImageSizeMiddle)
			key = "upload" + m.LinkMiddle()
			go uploadToUpyun(middlePath, "/"+key, uploadChannels)
		}
	}
	if len(smallPath) > 0 {
		for i := 0; i < 2; i++ {
			if err := <-uploadChannels; err != nil {
				result = err
			}
		}
		os.RemoveAll(smallPath)
	}
	for i := 0; i < 2; i++ {
		if err := <-uploadChannels; err != nil {
			result = err
		}
	}
	os.RemoveAll(fullPath)

	if len(middlePath) > 0 {
		for i := 0; i < 2; i++ {
			if err := <-uploadChannels; err != nil {
				result = err
			}
		}
		os.RemoveAll(middlePath)
	}

	return result
}

func ImageResize(img *models.Image, im image.Image, width int) error {
	savePath := GenImageFilePath(img, width)
	im = resize.Resize(uint(width), 0, im, resize.Bilinear)

	var file *os.File
	if f, err := os.OpenFile(savePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644); err != nil {
		return err
	} else {
		file = f
	}
	defer file.Close()

	var err error
	switch img.Ext {
	case 1:
		err = jpeg.Encode(file, im, &jpeg.Options{90})
	case 2:
		err = png.Encode(file, im)
	default:
		return fmt.Errorf("<ImageResize> unsupport image format")
	}

	return err
}

func GenImagePath(img *models.Image) string {
	return "upload/img/" + beego.Date(img.Created, "y/m/d/s/") + utils.ToStr(img.Id) + "/"
}

func GenImageFilePath(img *models.Image, width int) string {
	var size string
	if width == 0 {
		size = "full"
	} else {
		size = utils.ToStr(width)
	}
	return GenImagePath(img) + size + img.GetExt()
}
