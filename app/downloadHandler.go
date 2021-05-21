package app

import (
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
	file,filename, err := d.xlsxService.Generate()
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

	d.fileService.Save(b.Bytes())

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=" + *filename)

	c.Data(http.StatusOK, "application/octet-stream", b.Bytes())
}
