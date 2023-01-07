package service

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

func RawToMap(c *gin.Context) map[string]string {
	b, _ := c.GetRawData()
	m := make(map[string]interface{})
	_ = json.Unmarshal(b, &m)
	m2 := make(map[string]string)
	for k, v := range m {
		s, _ := v.(string)
		m2[k] = s
	}
	return m2
}
