package handlers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"

	"github.com/colt005/flatornot/common"
	"github.com/labstack/echo/v4"
)

func (h *Handler) HandleVote(e echo.Context) error {
	v := make(map[string]string)
	if err := e.Bind(&v); err != nil {
		return err
	}
	fmt.Println(v)

	vote := v["type"]

	h.service.AddVote(vote)

	tmpl, err := template.New("poll-results").Funcs(template.FuncMap{
		"percentage": common.CalculatePercentage,
		"mul":        common.Mul,
	}).ParseFiles("assets/index.html")

	if err != nil {
		return err
	}

	pollData := h.service.GetPollData()

	buf := &bytes.Buffer{}

	if err := tmpl.ExecuteTemplate(buf, "poll-results", pollData); err != nil {
		log.Println("Error executing template:", err)
	}

	sEnc := base64.StdEncoding.EncodeToString(buf.Bytes())

	h.service.BroadcastVotes(sEnc)

	pollData.LatestPun = common.P.Random(vote)

	tmpl.ExecuteTemplate(e.Response().Writer, "poll-results", pollData)

	return nil

}