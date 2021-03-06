/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package bindings

type CloudV3 struct {
	*RequestSchema
	/** skip_ticks */
	SkipTicks bool `json:"skip_ticks"`

	/** version */
	Version string `json:"version"`

	/** branch_name */
	BranchName string `json:"branch_name"`

	/** build_number */
	BuildNumber string `json:"build_number"`

	/** Node index number cloud status is collected from (zero-based) */
	NodeIdx int32 `json:"node_idx"`

	/** cloud_name */
	CloudName string `json:"cloud_name"`

	/** cloud_size */
	CloudSize int32 `json:"cloud_size"`

	/** cloud_uptime_millis */
	CloudUptimeMillis int64 `json:"cloud_uptime_millis"`

	/** cloud_healthy */
	CloudHealthy bool `json:"cloud_healthy"`

	/** Nodes reporting unhealthy */
	BadNodes int32 `json:"bad_nodes"`

	/** Cloud voting is stable */
	Consensus bool `json:"consensus"`

	/** Cloud is accepting new members or not */
	Locked bool `json:"locked"`

	/** Cloud is in client mode. */
	IsClient bool `json:"is_client"`

	/** nodes */
	Nodes []*NodeV3 `json:"nodes"`

	/* INHERITED: Comma-separated list of JSON field paths to exclude from the result, used like: "/3/Frames?_exclude_fields=frames/frame_id/URL,__meta"
	ExcludeFields string: "" `json:"_exclude_fields"`
	*/
}

func NewCloudV3() *CloudV3 {
	return &CloudV3{
		SkipTicks:         false,
		Version:           "",
		BranchName:        "",
		BuildNumber:       "",
		NodeIdx:           0,
		CloudName:         "",
		CloudSize:         0,
		CloudUptimeMillis: 0,
		CloudHealthy:      false,
		BadNodes:          0,
		Consensus:         false,
		Locked:            false,
		IsClient:          false,
		Nodes:             nil,
		RequestSchema: &RequestSchema{
			ExcludeFields: "",
		},
	}
}
