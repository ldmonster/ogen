// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"io"
	"mime/multipart"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

func encodeUserPassloginPostRequest(
	req UserPassloginPostReq,
	span trace.Span,
) (
	data io.Reader,
	contentType string,
	rerr error,
) {
	q := uri.NewQueryEncoder()
	{
		// Encode "passcode" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "passcode",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(req.Passcode))
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	e := new(bytes.Buffer)
	w := multipart.NewWriter(e)
	defer func() {
		cerr := w.Close()
		if rerr == nil {
			rerr = cerr
		}
	}()
	if err := q.WriteMultipart(w); err != nil {
		return e, "", errors.Wrap(err, "write multipart")
	}
	return e, w.FormDataContentType(), nil
}
func encodeUserPostingPostRequest(
	req UserPostingPostReq,
	span trace.Span,
) (
	data io.Reader,
	contentType string,
	rerr error,
) {
	q := uri.NewQueryEncoder()
	{
		// Encode "board" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "board",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(req.Board))
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "captcha_type" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "captcha_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(string(req.CaptchaType)))
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "comment" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "comment",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := req.Comment.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "email" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "email",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := req.Email.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "file[]" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "file[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeArray(func(e uri.Encoder) error {
				for i, item := range req.File {
					if err := func() error {
						return e.EncodeValue(conv.StringToString(item))
					}(); err != nil {
						return errors.Wrapf(err, "[%d]", i)
					}
				}
				return nil
			})
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "icon" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "icon",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := req.Icon.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "name" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "name",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := req.Name.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "op_mark" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "op_mark",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := req.OpMark.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "subject" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "subject",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := req.Subject.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "tags" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "tags",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := req.Tags.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "thread" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "thread",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := req.Thread.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	e := new(bytes.Buffer)
	w := multipart.NewWriter(e)
	defer func() {
		cerr := w.Close()
		if rerr == nil {
			rerr = cerr
		}
	}()
	if err := q.WriteMultipart(w); err != nil {
		return e, "", errors.Wrap(err, "write multipart")
	}
	return e, w.FormDataContentType(), nil
}
func encodeUserReportPostRequest(
	req UserReportPostReq,
	span trace.Span,
) (
	data io.Reader,
	contentType string,
	rerr error,
) {
	q := uri.NewQueryEncoder()
	{
		// Encode "board" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "board",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(req.Board))
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "comment" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "comment",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(req.Comment))
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "post" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "post",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeArray(func(e uri.Encoder) error {
				for i, item := range req.Post {
					if err := func() error {
						return e.EncodeValue(conv.IntToString(item))
					}(); err != nil {
						return errors.Wrapf(err, "[%d]", i)
					}
				}
				return nil
			})
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "thread" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "thread",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.IntToString(req.Thread))
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	e := new(bytes.Buffer)
	w := multipart.NewWriter(e)
	defer func() {
		cerr := w.Close()
		if rerr == nil {
			rerr = cerr
		}
	}()
	if err := q.WriteMultipart(w); err != nil {
		return e, "", errors.Wrap(err, "write multipart")
	}
	return e, w.FormDataContentType(), nil
}