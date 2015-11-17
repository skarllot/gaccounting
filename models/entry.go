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

package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	Entry struct {
		Id     bson.ObjectId   `bson:"_id,omitempty" json:"id,omitempty"`
		Debit  []EntrySide     `bson:"debit" json:"debit"`
		Credit []EntrySide     `bson:"credit" json:"credit"`
		Date   time.Time       `bson:"date" json:"date"`
		Value  int64           `bson:"value" json:"value"`
		Tags   []string        `bson:"tags" json:"tags"`
		Notes  string          `bson:"notes" json:"notes"`
		Files  []bson.ObjectId `bson:"files" json:"files"`
	}

	EntrySide struct {
		Account AccountRef `bson:"account" json:"account"`
		Value   int64      `bson:"value,omitempty" json:"value,omitempty"`
	}

	AccountRef struct {
		Major int `bson:"major" json:"major"`
		Minor int `bson:"minor" json:"minor"`
		Ref   int `bson:"ref" json:"ref"`
	}
)
