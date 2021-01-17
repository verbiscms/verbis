package resolve

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// post
//
// Resolves a post from the given value.
// Returns the domain.PostData if it was found and no error occurred.
// Returns errors.INVALID if the domain.FieldValue could not be cast to an integer.
func (v *Value) post(value domain.FieldValue) (interface{}, error) {
	const op = "FieldResolver.post"

	id, err := value.Int()
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to cast post ID to an integer", Operation: op, Err: err}
	}

	post, err := v.store.Posts.GetById(id)
	if err != nil {
		return nil, err
	}

	formatPost, err := v.store.Posts.Format(post)
	if err != nil {
		return nil, err
	}

	return formatPost, nil
}
