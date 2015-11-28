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

const (
	C_ACCOUNTS_NAME = "accounts"
)

type (
	Account struct {
		AccountAtom `bson:",inline"`
		Inverted    bool           `bson:"inverted,omitempty" json:"inverted,omitempty"`
		Minors      []AccountMinor `bson:"minors,omitempty" json:"minors,omitempty"`
	}

	AccountMinor struct {
		AccountAtom `bson:",inline"`
		Inverted    bool          `bson:"inverted,omitempty" json:"inverted,omitempty"`
		Subs        []AccountAtom `bson:"subs,omitempty" json:"subs,omitempty"`
	}

	AccountAtom struct {
		Id    int    `bson:"_id" json:"id"`
		Name  string `bson:"name" json:"name"`
		Notes string `bson:"notes,omitempty" json:"notes,omitempty"`
	}
)
