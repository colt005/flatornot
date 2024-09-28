package handlers

import (
	"fmt"
	"html/template"

	"github.com/colt005/flatornot/common"
	"github.com/labstack/echo/v4"
)

func (h *Handler) HandleIndex(e echo.Context) error {
	tmpl, err := template.New("index.html").Funcs(template.FuncMap{
		"percentage": common.CalculatePercentage,
		"mul":        common.Mul,
	}).ParseFiles("assets/index.html")
	if err != nil {

		return err
	}
	pollData := h.service.GetPollData()
	fmt.Println(pollData)
	err = tmpl.Execute(e.Response().Writer, pollData)
	if err != nil {
		return err
	}

	return nil

}
