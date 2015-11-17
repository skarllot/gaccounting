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

type AccountsController struct {
	col *mgo.Collection
}

func NewAccountsController(db *mgo.Database) *AccountsController {
	return &AccountsController{db.C(models.C_ACCOUNTS_NAME)}
}

func (s *AccountsController) GetAccount(c *gin.Context) {
	sMajor := c.Param("major")
	iMajor, err := strconv.Atoi(sMajor)
	if err != nil || iMajor < 1 || iMajor > 999 {
		jerr := rqhttp.NewJsonErrorFromError(
			http.StatusGone, InvalidId(sMajor))
		c.JSON(jerr.Status, jerr)
		return
	}

	account := models.Account{}

	// db.accounts.find({_id: major})
	err = s.col.
		Find(bson.M{"_id": iMajor}).
		Select(bson.M{"minors": 0}).
		One(&account)
	if err != nil {
		jerr := parseIdError(sMajor, err)
		c.JSON(jerr.Status, jerr)
		return
	}

	c.JSON(http.StatusOK, account)
}

func (s *AccountsController) GetAccountList(c *gin.Context) {
	level, err := s.getQueryLevel(c)
	if err != nil {
		jerr := rqhttp.NewJsonErrorFromError(http.StatusGone, err)
		c.JSON(jerr.Status, jerr)
		return
	}

	var filter bson.M
	switch level {
	case -1:
		filter = nil
	case 1:
		// db.accounts.find({_id: { $lte: 9 } })
		filter = bson.M{"_id": bson.M{"$lte": 9}}
	case 2:
		// db.accounts.find({_id: { $gte: 10, $lte: 99 } })
		filter = bson.M{"_id": bson.M{"$gte": 10, "$lte": 99}}
	case 3:
		// db.accounts.find({_id: { $gte: 100, $lte: 999 } })
		filter = bson.M{"_id": bson.M{"$gte": 100, "$lte": 999}}
	}

	list := make([]models.Account, 0)
	err = s.col.
		Find(filter).
		Select(bson.M{"minors": 0}).
		All(&list)
	if err != nil {
		jerr := rqhttp.NewJsonErrorFromError(
			http.StatusInternalServerError, err)
		c.JSON(jerr.Status, jerr)
		return
	}

	c.JSON(http.StatusOK, list)
}

func (s *AccountsController) getQueryLevel(c *gin.Context) (int, error) {
	sLevel := c.Query("level")
	if sLevel == "" {
		return -1, nil
	}

	iLevel, err := strconv.Atoi(sLevel)
	if err != nil || iLevel < 1 || iLevel > 3 {
		return -1, InvalidLevel(sLevel)
	}

	return iLevel, nil
}

func (s *AccountsController) SetRoutes(router gin.IRouter) gin.IRouter {
	g := router.Group("/accounts")
	g.GET("/", s.GetAccountList)
	g.GET("/:major", s.GetAccount)

	return g
}
