package app

import (
	"awesomeProject/domain"
	errs "awesomeProject/errors"
	"awesomeProject/service"
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DownloadHandler struct {
	xlsxService service.XlsxService
	fileService service.FileService
}

func (d *DownloadHandler) downloadXlsx(c *gin.Context) {
	file, filename, err := d.xlsxService.Generate()
	if err != nil {
		c.JSON(err.Code, err.AsMessage())
		return
	}

	var b bytes.Buffer
	if err := file.Write(&b); err != nil {
		appError := errs.NewUnexpectedError("Unexpected XLSX error")
		c.JSON(appError.Code, appError.AsMessage())
		return
	}

	metadata := &domain.FileMetadata{
		FsId:     "fefefef",
		Filename: "egd",
		Date:     "27.10.10.",
		Len:      nil,
	}

	_, appError := d.fileService.Save(b.Bytes(), metadata)
	if appError != nil {
		c.JSON(appError.Code, appError.AsMessage())
		return
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+*filename)

	c.Data(http.StatusOK, "application/octet-stream", b.Bytes())
}

func (d *DownloadHandler) downloadById(c *gin.Context) {

	buffer, appError := d.fileService.DownloadById("fefefef")

	if appError != nil {
		c.JSON(appError.Code, appError.AsMessage())
		return
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+"prase")

	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())
}
