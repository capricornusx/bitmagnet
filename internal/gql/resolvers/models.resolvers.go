package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"

	"github.com/bitmagnet-io/bitmagnet/internal/gql"
	"github.com/bitmagnet-io/bitmagnet/internal/gql/gqlmodel"
	"github.com/bitmagnet-io/bitmagnet/internal/model"
)

// OriginalLanguage is the resolver for the originalLanguage field.
func (r *contentResolver) OriginalLanguage(ctx context.Context, obj *model.Content) (*model.Language, error) {
	var language *model.Language
	if obj.OriginalLanguage.Valid {
		language = &obj.OriginalLanguage.Language
	}
	return language, nil
}

// Sources is the resolver for the sources field.
func (r *torrentResolver) Sources(ctx context.Context, obj *model.Torrent) ([]gqlmodel.TorrentSource, error) {
	return gqlmodel.TorrentSourcesFromTorrent(*obj), nil
}

// Content returns gql.ContentResolver implementation.
func (r *Resolver) Content() gql.ContentResolver { return &contentResolver{r} }

// Torrent returns gql.TorrentResolver implementation.
func (r *Resolver) Torrent() gql.TorrentResolver { return &torrentResolver{r} }

type contentResolver struct{ *Resolver }
type torrentResolver struct{ *Resolver }