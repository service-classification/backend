package handlers

import (
	"backend/internal/models"
	_ "embed"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"github.com/xuri/excelize/v2"
)

// BuildReport godoc
//
//	@Summary		Build fiscal report
//	@Description	Generates a fiscal report in Excel format and returns it as a downloadable file.
//	@Tags			Reports
//	@Produce		application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
//	@Success		200	{file}		application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/report [get]
func (h *Handler) BuildReport(c *gin.Context) {
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=report_profit_2024.xlsx")
	c.Header("Access-Control-Expose-Headers", "*")

	f := excelize.NewFile()

	// Create styles
	titleStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 14,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#D9E1F2"},
		},
	})
	if err != nil {
		log.Printf("Failed to create title style: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	subtitleStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 12,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#FCE4D6"},
		},
	})
	if err != nil {
		log.Printf("Failed to create subtitle style: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write the titles
	err = f.MergeCell("Sheet1", "A1", "A2")
	if err != nil {
		log.Printf("Failed to merge cells: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = f.SetCellValue("Sheet1", "A1", "Financial Class")
	if err != nil {
		log.Printf("Failed to set cell value: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = f.SetCellStyle("Sheet1", "A1", "A2", titleStyle)
	if err != nil {
		log.Printf("Failed to set cell style: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = f.MergeCell("Sheet1", "B1", "F1")
	if err != nil {
		log.Printf("Failed to merge cells: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = f.SetCellValue("Sheet1", "B1", "2024")
	if err != nil {
		log.Printf("Failed to set cell value: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = f.SetCellStyle("Sheet1", "B1", "F1", titleStyle)
	if err != nil {
		log.Printf("Failed to set cell style: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Add quartal headers
	headers := []string{"1 Quarter", "2 Quarter", "3 Quarter", "4 Quarter", "Total"}
	for i, header := range headers {
		col := string(rune('A' + i + 1)) // Convert index to column letter
		cell := col + "2"
		f.SetCellValue("Sheet1", cell, header)
		f.SetCellStyle("Sheet1", cell, cell, subtitleStyle)
	}
	err = f.SetColWidth("Sheet1", "A", "F", 20)
	if err != nil {
		log.Printf("Failed to set column width: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	services, err := h.ServiceRepo.List(0, 10_000)
	if err != nil {
		log.Printf("Failed to list services: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var classes []*models.Class
	for _, service := range services {
		if service.ApprovedAt != nil {
			classes = append(classes, service.Class)
		}
	}

	classTitleStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})
	if err != nil {
		log.Printf("Failed to create class title style: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	titles := make([]string, 0, len(classes))
	maxLength := 0
	for _, class := range classes {
		titles = append(titles, class.Title)
		if len(class.Title) > maxLength {
			maxLength = len(class.Title)
		}
	}
	err = f.SetSheetCol("Sheet1", "A3", &titles)
	if err != nil {
		log.Printf("Failed to set sheet column: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Set column widths for better appearance
	err = f.SetColWidth("Sheet1", "A", "A", float64(maxLength))
	if err != nil {
		log.Printf("Failed to set column width: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = f.SetCellStyle("Sheet1", "A3", "A"+strconv.Itoa(len(classes)+2), classTitleStyle)
	if err != nil {
		log.Printf("Failed to set cell style: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	customNumberFormat := `"$"#,##0.00`
	financialStyle, err := f.NewStyle(&excelize.Style{
		NumFmt:       164,
		CustomNumFmt: &customNumberFormat,
		Font: &excelize.Font{
			Bold: false,
			Size: 11,
		},
	})
	if err != nil {
		log.Printf("Failed to create financial style: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:  true,
			Size:  11,
			Color: "#FFFFFF",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FF5733"},
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "right",
			Vertical:   "center",
		},
	})
	if err != nil {
		log.Printf("Failed to create total style: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for i := range classes {
		row := generateRandomMoneyRow(4)
		err := f.SetSheetRow("Sheet1", "B"+strconv.Itoa(i+3), &row)
		if err != nil {
			log.Printf("Failed to set sheet row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		f.SetCellFormula("Sheet1", "F"+strconv.Itoa(i+3), "SUM(B"+strconv.Itoa(i+3)+":E"+strconv.Itoa(i+3)+")")
		f.SetCellStyle("Sheet1", "B"+strconv.Itoa(i+3), "E"+strconv.Itoa(i+3), financialStyle)
		f.SetCellStyle("Sheet1", "F"+strconv.Itoa(i+3), "F"+strconv.Itoa(i+3), totalStyle)
	}

	_, err = f.WriteTo(c.Writer)
	if err != nil {
		log.Printf("Failed to write to response: %v", err)
	}
	c.Status(http.StatusOK)
}

func generateRandomMoneyRow(size int) []float64 {
	type price struct {
		Amount float64 `faker:"amount"`
	}

	row := make([]float64, size+1)
	for i := range row {
		c := price{}
		faker.FakeData(&c)
		row[i] = c.Amount
	}

	return row
}
