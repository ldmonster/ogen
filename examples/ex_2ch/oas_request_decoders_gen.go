// Code generated by ogen, DO NOT EDIT.

package api

import (
	"mime"
	"net/http"
	"net/url"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

func (s *Server) decodeUserPassloginPostRequest(r *http.Request, span trace.Span) (req UserPassloginPostReq, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "multipart/form-data":
		if r.ContentLength == 0 {
			return req, nil
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, errors.Wrap(err, "parse multipart form")
		}
		form := url.Values(r.MultipartForm.Value)

		if len(form) == 0 {
			return req, nil
		}

		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "passcode",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					req.Passcode = c
					return nil
				}); err != nil {
					return req, errors.Wrap(err, "decode \"passcode\"")
				}
			} else {
				return req, errors.Wrap(err, "query")
			}
		}

		return req, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeUserPostingPostRequest(r *http.Request, span trace.Span) (req UserPostingPostReq, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "multipart/form-data":
		if r.ContentLength == 0 {
			return req, nil
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, errors.Wrap(err, "parse multipart form")
		}
		form := url.Values(r.MultipartForm.Value)

		if len(form) == 0 {
			return req, nil
		}

		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "board",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					req.Board = c
					return nil
				}); err != nil {
					return req, errors.Wrap(err, "decode \"board\"")
				}
			} else {
				return req, errors.Wrap(err, "query")
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "captcha_type",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					req.CaptchaType = CaptchaType(c)
					return nil
				}); err != nil {
					return req, errors.Wrap(err, "decode \"captcha_type\"")
				}
				if err := func() error {
					if err := req.CaptchaType.Validate(); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return req, errors.Wrap(err, "validate")
				}
			} else {
				return req, errors.Wrap(err, "query")
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "comment",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var reqDotCommentVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						reqDotCommentVal = c
						return nil
					}(); err != nil {
						return err
					}
					req.Comment.SetTo(reqDotCommentVal)
					return nil
				}); err != nil {
					return req, errors.Wrap(err, "decode \"comment\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "email",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var reqDotEmailVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						reqDotEmailVal = c
						return nil
					}(); err != nil {
						return err
					}
					req.Email.SetTo(reqDotEmailVal)
					return nil
				}); err != nil {
					return req, errors.Wrap(err, "decode \"email\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "file[]",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					return d.DecodeArray(func(d uri.Decoder) error {
						var reqDotFileVal string
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							reqDotFileVal = c
							return nil
						}(); err != nil {
							return err
						}
						req.File = append(req.File, reqDotFileVal)
						return nil
					})
				}); err != nil {
					return req, errors.Wrap(err, "decode \"file[]\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "icon",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var reqDotIconVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						reqDotIconVal = c
						return nil
					}(); err != nil {
						return err
					}
					req.Icon.SetTo(reqDotIconVal)
					return nil
				}); err != nil {
					return req, errors.Wrap(err, "decode \"icon\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "name",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var reqDotNameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						reqDotNameVal = c
						return nil
					}(); err != nil {
						return err
					}
					req.Name.SetTo(reqDotNameVal)
					return nil
				}); err != nil {
					return req, errors.Wrap(err, "decode \"name\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "op_mark",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var reqDotOpMarkVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						reqDotOpMarkVal = c
						return nil
					}(); err != nil {
						return err
					}
					req.OpMark.SetTo(reqDotOpMarkVal)
					return nil
				}); err != nil {
					return req, errors.Wrap(err, "decode \"op_mark\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "subject",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var reqDotSubjectVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						reqDotSubjectVal = c
						return nil
					}(); err != nil {
						return err
					}
					req.Subject.SetTo(reqDotSubjectVal)
					return nil
				}); err != nil {
					return req, errors.Wrap(err, "decode \"subject\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "tags",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var reqDotTagsVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						reqDotTagsVal = c
						return nil
					}(); err != nil {
						return err
					}
					req.Tags.SetTo(reqDotTagsVal)
					return nil
				}); err != nil {
					return req, errors.Wrap(err, "decode \"tags\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "thread",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var reqDotThreadVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						reqDotThreadVal = c
						return nil
					}(); err != nil {
						return err
					}
					req.Thread.SetTo(reqDotThreadVal)
					return nil
				}); err != nil {
					return req, errors.Wrap(err, "decode \"thread\"")
				}
			}
		}

		return req, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeUserReportPostRequest(r *http.Request, span trace.Span) (req UserReportPostReq, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "multipart/form-data":
		if r.ContentLength == 0 {
			return req, nil
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, errors.Wrap(err, "parse multipart form")
		}
		form := url.Values(r.MultipartForm.Value)

		if len(form) == 0 {
			return req, nil
		}

		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "board",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					req.Board = c
					return nil
				}); err != nil {
					return req, errors.Wrap(err, "decode \"board\"")
				}
			} else {
				return req, errors.Wrap(err, "query")
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "comment",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					req.Comment = c
					return nil
				}); err != nil {
					return req, errors.Wrap(err, "decode \"comment\"")
				}
			} else {
				return req, errors.Wrap(err, "query")
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "post",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					return d.DecodeArray(func(d uri.Decoder) error {
						var reqDotPostVal int
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToInt(val)
							if err != nil {
								return err
							}

							reqDotPostVal = c
							return nil
						}(); err != nil {
							return err
						}
						req.Post = append(req.Post, reqDotPostVal)
						return nil
					})
				}); err != nil {
					return req, errors.Wrap(err, "decode \"post\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "thread",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					req.Thread = c
					return nil
				}); err != nil {
					return req, errors.Wrap(err, "decode \"thread\"")
				}
			} else {
				return req, errors.Wrap(err, "query")
			}
		}

		return req, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}