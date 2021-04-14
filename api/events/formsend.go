// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

//// FormSend defines the event instance for config sending
//// form data, Form and FormValues are required for
////dispatch.
//type FormSend struct {
//	event  *event
//	Form   *domain.Form
//	Values domain.FormValues
//}
//
//// FormSend
////
//// Creates a new FormSend.
//func NewFormSend(d *deps.Deps) *FormSend {
//	return &FormSend{
//		event: &event{
//			deps:      d,
//			subject:   SubjectPrefix + "Reset Password",
//			template:  "form-send",
//			plaintext: "newMailer form submission",
//		},
//	}
//}
//
//// Dispatch
////
//// Dispatches the FormSend event.
//func (r *FormSend) Dispatch(data interface{}, recipients []string, attachments mail.Attachments) error {
//	const op = "Events.ChangedPassword.Dispatch"
//
//	fs, ok := data.(FormSend)
//	if !ok {
//		return &errors.Error{Code: errors.INTERNAL, Message: "FormSend should be passed to dispatch", Operation: op, Err: WrongTypeErr}
//	}
//
//	if fs.Form == nil {
//		return &errors.Error{Code: errors.INTERNAL, Message: "Form cannot be nil", Operation: op, Err: fmt.Errorf("form is nil")}
//	}
//
//	fv := make(domain.FormValues)
//	for _, v := range fs.Form.Fields {
//		val, ok := fs.Values[v.Key]
//		if !ok {
//			continue
//		}
//		if v.Type != "file" {
//			fv[v.Label.String()] = val
//		}
//	}
//	fs.Values = fv
//
//	err := r.event.send(fs, recipients, attachments)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

// OLD
//
//// Send the verify email event.
//func (e *FormSend) Send(f *FormSendData, attachments forms.Attachments) error {
//	const op = "events.VerifyEmail.Send"
//
//	fv := make(domain.FormValues)
//	for _, v := range f.Form.Fields {
//		val, ok := f.Values[v.Key]
//		if !ok {
//			continue
//		}
//		if v.Type != "file" {
//			fv[v.Label.String()] = val
//		}
//	}
//	f.Values = fv
//
//	html, err := e.mailer.ExecuteHTML("form-send.html", &f)
//	if err != nil {
//		return err
//	}
//
//	tm := mailer.Sender{
//		To:          f.Form.GetRecipients(),
//		subject:     f.Form.EmailSubject,
//		HTML:        html,
//		Attachments: attachments,
//	}
//
//	e.mailer.Send(&tm)
//
//	return nil
//}
