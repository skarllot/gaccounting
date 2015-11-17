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
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	rqhttp "github.com/raiqub/http"
	"github.com/skarllot/gaccounting/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AccountMinorsController struct {
	col *mgo.Collection
}

func NewAccountMinorsController(db *mgo.Database) *AccountMinorsController {
	return &AccountMinorsController{db.C(models.C_ACCOUNTS_NAME)}
}

func (s *AccountMinorsController) GetMinorList(c *gin.Context) {
	sMajor := c.Param("major")
	iMajor, err := strconv.Atoi(sMajor)
	if err != nil {
		jerr := rqhttp.NewJsonErrorFromError(http.StatusGone, InvalidId(sMajor))
		c.JSON(jerr.Status, jerr)
		return
	}

	account := models.Account{}

	err = s.col.
		Find(bson.M{"_id": iMajor}).
		Select(bson.M{"minors.refs": 0}).
		One(&account)
	if err != nil {
		jerr := parseIdError(sMajor, err)
		c.JSON(jerr.Status, jerr)
		return
	}

	c.JSON(http.StatusOK, account)
}

func (s *AccountMinorsController) SetRoutes(router gin.IRouter) gin.IRouter {
	g := router.Group("/:major/minors")
	g.GET("/", s.GetMinorList)

	return router
}
