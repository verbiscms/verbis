// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

//// VerifyEmail defines the event instance for config
//// verifying email addresses when a user signs
//// up. Token and User are required for
//// dispatch.
//type VerifyEmail struct {
//	event *event
//	Token string
//	User  domain.UserPart
//}
//
//// NewVerifyEmail
////
//// Creates a new VerifyEmail.
//func NewVerifyEmail(mail Mailer) *VerifyEmail {
//	return &VerifyEmail{
//		event: &event{
//			mailer:    mail,
//			subject:   SubjectPrefix + "Verify Email",
//			template:  "verify-email",
//			plaintext: "Thanks for signing up! Please verify your email address with a Verbis site.",
//		},
//	}
//}
//
//// Dispatch
////
//// Dispatches the VerifyEmail event.
//func (r *VerifyEmail) Dispatch(data interface{}, recipients []string, attachments mail.Attachments) error {
//	const op = "Events.ResetPassword.Dispatch"
//
//	rp, ok := data.(VerifyEmail)
//	if !ok {
//		return &errors.Error{Code: errors.INTERNAL, Message: "VerifyEmail should be passed to dispatch", Operation: op, Err: WrongTypeErr}
//	}
//
//	rp.Token = encryption.MD5Hash(strconv.Itoa(rp.User.Id) + rp.User.Email)
//
//	err := r.event.send(rp, recipients, attachments)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
