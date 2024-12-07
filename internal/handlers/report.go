package handlers

import (
	_ "embed"

	"github.com/gin-gonic/gin"
)

//go:embed report/report_profit_2022.pdf
var report []byte

// BuildReport godoc
//
//	@Summary		Build fiscal report
//	@Description	Generates a fiscal report in PDF format and returns it as a downloadable file.
//	@Tags			Reports
//	@Produce		application/pdf
//	@Success		200	{file}		application/pdf
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/report [get]
func (h *Handler) BuildReport(c *gin.Context) {
	//todo implement

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=report_profit_2022.pdf")
	c.Data(200, "application/pdf", report)
}
