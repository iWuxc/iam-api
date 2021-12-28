// Copyright 2021 iwuxc <wuxc.eng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

type Template struct {

	metav1.ObjectMeta `json:"metadata,omitempty"`

	TemplateUse string `json:"templateUse"`

	Template string `json:"template"`

	TemplateType string `json:"templateType"`

}

// TableName maps to mysql table name.
func (s *Template) TableName() string {
	return "template"
}
