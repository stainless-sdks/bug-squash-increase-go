package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"net/url"
)

type DocumentService struct {
	Options []options.RequestOption
}

func NewDocumentService(opts ...options.RequestOption) (r *DocumentService) {
	r = &DocumentService{}
	r.Options = opts
	return
}

// Retrieve a Document
func (r *DocumentService) Get(ctx context.Context, document_id string, opts ...options.RequestOption) (res *types.Document, err error) {
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("documents/%s", document_id))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "GET", u, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// List Documents
func (r *DocumentService) List(ctx context.Context, query *types.DocumentListParams, opts ...options.RequestOption) (res *types.DocumentsPage, err error) {
	u, err := url.Parse(fmt.Sprintf("documents"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg := options.NewRequestConfig(ctx, "GET", u, opts...)
	res = &types.DocumentsPage{
		Page: &pagination.Page[types.Document]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}
