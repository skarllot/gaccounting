/*
 * Copyright (C) 2015 Fabrício Godoy <skarllot@gmail.com>
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
 */

package controllers

import (
	"fmt"
	"net/http"

	rqhttp "github.com/raiqub/http"
	"gopkg.in/mgo.v2"
)

// parseIdError returns a not found ID error when aplicable;
// otherwise returns a internal server error.
func parseIdError(id string, err error) rqhttp.JsonError {
	var jerr rqhttp.JsonError
	if err == mgo.ErrNotFound {
		jerr = rqhttp.NewJsonErrorFromError(
			http.StatusNotFound, NotFoundId(id))
	} else {
		jerr = rqhttp.NewJsonErrorFromError(
			http.StatusInternalServerError, err)
	}

	return jerr
}

type InvalidId string

func (e InvalidId) Error() string {
	return fmt.Sprintf("Invalid Id for '%s'", string(e))
}

type NotFoundId string

func (e NotFoundId) Error() string {
	return fmt.Sprintf("The Id '%s' was not found", string(e))
}

type InvalidLevel string

func (e InvalidLevel) Error() string {
	return fmt.Sprintf("Invalid level for '%s'", string(e))
}
