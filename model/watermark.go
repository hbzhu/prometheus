// Copyright 2013 Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	dto "github.com/prometheus/prometheus/model/generated"
	"time"
)

// Watermark provides a representation of dto.MetricHighWatermark with
// associated business logic methods attached to it to enhance code readability.
type Watermark struct {
	time.Time
}

// NewWatermarkFromHighWatermarkDTO builds Watermark from the provided
// dto.MetricHighWatermark object.
func NewWatermarkFromHighWatermarkDTO(d *dto.MetricHighWatermark) Watermark {
	return Watermark{
		time.Unix(*d.Timestamp, 0),
	}
}
