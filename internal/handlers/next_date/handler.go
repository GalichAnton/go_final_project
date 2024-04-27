package next_date

import (
	"fmt"
	"net/http"
	"time"

	"github.com/GalichAnton/go_final_project/internal/utils"
)

func Handle(w http.ResponseWriter, req *http.Request) {
	nowParam := req.URL.Query().Get("now")
	dateParam := req.URL.Query().Get("date")
	repeatParam := req.URL.Query().Get("repeat")

	now, err := time.Parse("20060102", nowParam)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid 'now' parameter: %v", err), http.StatusBadRequest)
		return
	}

	nextDate, err := utils.NextDate(now, dateParam, repeatParam)
	if err != nil {
		http.Error(w, fmt.Sprintf("error finding next date: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(nextDate))
}
