package plugin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestHeader(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ok := func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	}
	h := new(Header)
	h.Set = map[string]string{"Name": "Jon"}
	h.Add = map[string]string{"Name": "Joe"}
	h.Del = []string{"Delete"}
	rec.Header().Set("Delete", "me")

	h.Initialize()
	h.Process(ok)(c)

	assert.Equal(t, "Jon", rec.Header().Get("Name"))                    // Set
	assert.EqualValues(t, []string{"Jon", "Joe"}, rec.Header()["Name"]) // Add
	assert.Equal(t, "", rec.Header().Get("Delete"))                     // Del
}
