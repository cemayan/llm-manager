package utils

type CommonParams struct {
	Temp      float64
	StopWords []string
}

// ConfigureParams configures some params such as stopWords and temp
func (c *CommonParams) ConfigureParams(params map[string]interface{}) {

	if val, ok := params["stopWords"]; ok {
		if val2, ok2 := val.([]string); ok2 {
			c.StopWords = append(c.StopWords, val2...)
		}
	}

	if val, ok := params["temp"]; ok {
		if val2, ok2 := val.(float64); ok2 {
			c.Temp = val2
		}
	} else {
		c.Temp = 0.8 // default temp
	}
}
