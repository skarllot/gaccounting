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

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skarllot/gaccounting/controllers"
)

type Router struct {
	*gin.Engine
}

func NewRouter(session *Session) *Router {
	router := gin.Default()
	router.RedirectTrailingSlash = true

	g := router.Group("/v1")
	db := session.DB("")

	// Resources
	gAccounts := controllers.
		NewAccountsController(db).
		SetRoutes(g)
	controllers.
		NewAccountMinorsController(db).
		SetRoutes(gAccounts)

	return &Router{router}
}
