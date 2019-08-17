package status

import (
)

// InfoT provides runtime information.
type Info struct {
	Services  int `json:"services"`
	PipelineRunning  int `json:"pipeline_running"`
	PipelineComplete  int `json:"pipeline_complete"`
}

func (i *Info) ServiceAdd() {
}

func (i *Info) ServiceComplate() {
}

func (i *Info) PipelineAdd() {
}

func (i *Info) PipelineComplate() {
}
