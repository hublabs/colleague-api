package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hublabs/colleague-api/colleagues"
	"github.com/hublabs/colleague-api/factory"

	"github.com/labstack/echo"
	"github.com/pangpanglabs/goutils/test"
)

func Test_ColleagueApiController_GetColleagueById(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/v1/colleagues/:id", nil)

	c, rec := SetContext(req)
	dbSession := factory.DB(c.Request().Context())
	dbSession.Begin()
	defer func() {
		factory.DB(c.Request().Context()).Close()
		factory.DB(c.Request().Context()).Rollback()
	}()

	c.SetParamNames("id")
	c.SetParamValues("1")

	test.Ok(t, ColleagueApiController{}.GetColleagueById(c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  colleagues.Colleague   `json:"result"`
		Success bool                   `json:"success"`
		Errors  map[string]interface{} `json:"error"`
	}

	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.Id, int64(1))
	test.Equals(t, v.Result.Name, "系统管理员")
}
