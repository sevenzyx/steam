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

type ColSpecifierV3 struct {
	*Schema
	/** Name of the column */
	ColumnName string `json:"column_name,omitempty"`
	/** List of fields which specify columns that must contain this column */
	IsMemberOfFrames []string `json:"is_member_of_frames,omitempty"`
}

func NewColSpecifierV3() *ColSpecifierV3 {
	return &ColSpecifierV3{
		ColumnName:       "",
		IsMemberOfFrames: nil,
		Schema:           &Schema{},
	}
}
