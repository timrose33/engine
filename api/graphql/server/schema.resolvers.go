package server

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"fmt"

	"github.com/owasp-amass/engine/api/graphql/server/model"
)

// CreateSession is the resolver for the createSession field.
func (r *mutationResolver) CreateSession(ctx context.Context, input model.CreateSessionInput) (*model.Session, error) {
	testSession := &model.Session{
		Token: "00000000-0000-0000-0000-000000000000", //?
	}
	return testSession, nil
}

// CreateEvent is the resolver for the createEvent field.
func (r *mutationResolver) CreateEvent(ctx context.Context, input model.CreateEventInput) (*model.Event, error) {
	//event := &events.Event{Name: "NewSession", Type: events.EventTypeLog, Priority: 1}
	fmt.Println(fmt.Sprintf("%#v", input))

	//r.scheduler.Schedule(event)

	testSession := &model.Event{
		ID: "00000000-0000-0000-0000-000000000000",
	}
	return testSession, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }