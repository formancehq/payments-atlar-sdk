// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
)

func encodeFinancialDataV2EntitiesPostRequest(
	req *CreateEntityRequest,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeIamV2betaOAuth2TokenPostRequest(
	req *IamV2betaOAuth2TokenPostReq,
	r *http.Request,
) error {
	const contentType = "application/x-www-form-urlencoded"
	request := req

	q := uri.NewFormEncoder(map[string]string{})
	{
		// Encode "grant_type" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "grant_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(string(request.GrantType)))
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	encoded := q.Values().Encode()
	ht.SetBody(r, strings.NewReader(encoded), contentType)
	return nil
}

func encodePaymentsV2CreditTransferBatchesIDApprovePostRequest(
	req *ApproveBatchRequest,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodePaymentsV2CreditTransferBatchesIDRejectPostRequest(
	req *RejectPaymentRequest,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodePaymentsV2betaCreditTransferBatchesPostRequest(
	req *PaymentsV2betaCreditTransferBatchesPostReqMultipartFormData,
	r *http.Request,
) error {
	const contentType = "multipart/form-data"
	request := req

	q := uri.NewFormEncoder(map[string]string{})
	{
		// Encode "treatment" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "treatment",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(string(request.Treatment)))
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "fileContentType" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "fileContentType",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(string(request.FileContentType)))
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	body, boundary := ht.CreateMultipartBody(func(w *multipart.Writer) error {
		if err := request.File.WriteMultipart("file", w); err != nil {
			return errors.Wrap(err, "write \"file\"")
		}
		if err := q.WriteMultipart(w); err != nil {
			return errors.Wrap(err, "write multipart")
		}
		return nil
	})
	ht.SetCloserBody(r, body, mime.FormatMediaType(contentType, map[string]string{"boundary": boundary}))
	return nil
}