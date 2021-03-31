// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/ainsleyclark/verbis/api/mail/events"
	service "github.com/ainsleyclark/verbis/api/services/forms"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

// Send
//
// Returns http.StatusOK if the form was deleted.
// Returns http.StatusInternalServerError if there was an error deleting the form.
// Returns http.StatusBadRequest if the the form wasn't found or no ID was passed.
func (f *Forms) Send(ctx *gin.Context) {
	const op = "FormHandler.Send"

	uniq, err := uuid.Parse(ctx.Param("uuid"))
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Error parsing UUID", err)
		return
	}

	form, err := f.Store.Forms.FindByUUID(uniq)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	err = ctx.ShouldBind(form.Body)
	if err != nil {
		// If file has an empty value, no validation data is returned.
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	values, attachments, err := service.NewReader(&form, f.Paths.Storage).Values()
	if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	sub := domain.FormSubmission{
		FormId:    form.Id,
		Fields:    values,
		IPAddress: ctx.ClientIP(),
		UserAgent: ctx.Request.UserAgent(),
	}

	if form.StoreDB {
		err = f.Store.Forms.Submit(sub)
		if err != nil {
			api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
			return
		}
	}

	if form.EmailSend {
		fs, err := events.NewFormSend()
		if err != nil {
			api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
			return
		}

		data := &events.FormSendData{
			//Site:   s.siteModel.GetGlobalConfig(),
			Form:   &form,
			Values: values,
		}

		err = fs.Send(data, attachments)

		if err != nil {
			api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
			return
		}
	}

	api.Respond(ctx, http.StatusOK, "Successfully sent form with ID: "+strconv.Itoa(form.Id), nil)
}
