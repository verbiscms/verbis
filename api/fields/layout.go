package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/fields/layout"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

// GetGroup
//
func (s *Service) GetLayout(name string, args ...interface{}) domain.Field {
	l, err := layout.ByName(name, s.handleLayoutArgs(args))
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error()
		return domain.Field{}
	}
	return l
}

// GetGroups
//
func (s *Service) GetLayouts(args ...interface{}) []domain.FieldGroup {
	return s.handleLayoutArgs(args)
}

// handleLayoutArgs
//
func (s *Service) handleLayoutArgs(args []interface{}) []domain.FieldGroup {
	switch len(args) {
	case 1:
		layout := s.getLayoutByPost(args[0])
		return layout
	default:
		return s.layout
	}
}

// getLayoutByPost
//
// Returns the layout by Post with the given ID.
// Logs errors.INVALID if the id failed to be cast to an int.
// Logs if the post if was not found or there was an error obtaining/formatting the post.
func (s *Service) getLayoutByPost(id interface{}) []domain.FieldGroup {
	const op = "FieldsService.getFieldsByPost"

	i, err := cast.ToIntE(id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": &errors.Error{Code: errors.INVALID, Message: "Unable to cast Post ID to integer", Operation: op, Err: err},
		}).Error()
		return nil
	}

	post, err := s.store.Posts.GetById(i, true)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error()
		return nil
	}

	return post.Layout
}
