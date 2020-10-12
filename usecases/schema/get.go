//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package schema

import (
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
)

// GetSchema retrieves a locally cached copy of the schema
func (m *Manager) GetSchema(principal *models.Principal) (schema.Schema, error) {
	err := m.authorizer.Authorize(principal, "list", "schema/*")
	if err != nil {
		return schema.Schema{}, err
	}

	return schema.Schema{
		Actions: m.state.ActionSchema,
		Things:  m.state.ThingSchema,
	}, nil
}

// GetSchemaSkipAuth can never be used as a response to a user request as it
// could leak the schema to an unauthorized user, is intended to be used for
// non-user triggered processes, such as regular updates / maintenance / etc
func (m *Manager) GetSchemaSkipAuth() schema.Schema {
	return schema.Schema{
		Actions: m.state.ActionSchema,
		Things:  m.state.ThingSchema,
	}
}

func (m *Manager) Indexed(className, propertyName string) bool {
	s := schema.Schema{
		Actions: m.state.ActionSchema,
		Things:  m.state.ThingSchema,
	}
	class := s.FindClassByName(schema.ClassName(className))
	if class == nil {
		return false
	}

	for _, prop := range class.Properties {
		if prop.Name == propertyName {
			if prop.Index == nil {
				return true
			}

			return *prop.Index
		}
	}

	return false
}

func (m *Manager) VectorizeClassName(className string) bool {
	s := schema.Schema{
		Actions: m.state.ActionSchema,
		Things:  m.state.ThingSchema,
	}
	class := s.FindClassByName(schema.ClassName(className))
	if class == nil {
		return false
	}

	return VectorizeClassName(class)
}

func (m *Manager) VectorizePropertyName(className, propertyName string) bool {
	s := schema.Schema{
		Actions: m.state.ActionSchema,
		Things:  m.state.ThingSchema,
	}
	class := s.FindClassByName(schema.ClassName(className))
	if class == nil {
		return false
	}

	for _, prop := range class.Properties {
		if prop.Name == propertyName {
			return prop.VectorizePropertyName
		}
	}

	return false
}
