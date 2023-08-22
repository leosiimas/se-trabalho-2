package metrics

type SensorData struct {
	Text     string       `json:"Text"`
	Min      string       `json:"Min"`
	Value    string       `json:"Value"`
	Max      string       `json:"Max"`
	Children []SensorData `json:"Children"`
}

type MetricInfo struct {
	Text  string `json:"Text"`
	Min   string `json:"Min"`
	Value string `json:"Value"`
	Max   string `json:"Max"`
}

type Metric struct {
	Text       string       `json:"Text"`
	MetricInfo []MetricInfo `json:"MetricInfo"`
}

type Hardware struct {
	Hardware string   `json:"Hardware"`
	Metric   []Metric `json:"Metric"`
}

type HardwareData struct {
	Time     string     `json:"Time"`
	Hardware []Hardware `json:"Hardware"`
}

type DataExport struct {
	Hardware   string       `json:"Hardware"`
	Metric     string       `json:"Metric"`
	MetricInfo string       `json:"MetricInfo"`
	DataValues []DataValues `json:"DataValues"`
}

type DataValues struct {
	Time  int64   `json:"Time"`
	Min   float64 `json:"Min"`
	Value float64 `json:"Value"`
	Max   float64 `json:"Max"`
}

type MetricInfoData struct {
	Hardware string  `json:"Hardware"`
	Metric   string  `json:"Metric"`
	Sensor   string  `json:"Sensor"`
	Value    float64 `json:"Value"`
	Measure  string  `json:"measure"`
}

type DataEntry struct {
	Language string           `json:"Language"`
	Time     string           `json:"time"`
	Data     []MetricInfoData `json:"data"`
}

type PathValue struct {
	Text       string
	OutputPath string
	Addrs      string
	Language   string
}
