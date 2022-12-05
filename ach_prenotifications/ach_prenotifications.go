package ach_prenotifications

import "increase/core"
import "fmt"

type ACHPrenotifications struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewACHPrenotifications(requster core.Requester) (r *ACHPrenotifications) {
	r = &ACHPrenotifications{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *ACHPrenotifications) Create(body ACHPrenotificationsCreateParams, opts ...*core.RequestOpts) (res ACHPrenotification, err error) {
	err = r.post(
		"/ach_prenotifications",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *ACHPrenotifications) Retrieve(ach_prenotification_id string, opts ...*core.RequestOpts) (res ACHPrenotification, err error) {
	err = r.get(
		fmt.Sprintf("/ach_prenotifications/%s", ach_prenotification_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *ACHPrenotifications) List(query *ACHPrenotificationsListQuery, opts ...*core.RequestOpts) (res ACHPrenotificationsListResponse, err error) {
	err = r.get(
		"/ach_prenotifications",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}
