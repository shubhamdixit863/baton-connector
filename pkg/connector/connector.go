package connector

import (
	"context"
	"io"

	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/connectorbuilder"

	"github.com/conductorone/baton-demo/pkg/client"
)

// TODO: Define the name of the struct that you want  to use in place of Demo use the proper name

type Demo struct {
	client *client.Client
}

// ResourceSyncers returns a ResourceSyncer for each resource type that should be synced from the upstream service.
func (d *Demo) ResourceSyncers(ctx context.Context) []connectorbuilder.ResourceSyncer {
	return []connectorbuilder.ResourceSyncer{}
}

// Asset takes an input AssetRef and attempts to fetch it using the connector's authenticated http client
// It streams a response, always starting with a metadata object, following by chunked payloads for the asset.
func (d *Demo) Asset(ctx context.Context, asset *v2.AssetRef) (string, io.ReadCloser, error) {
	return "", nil, nil
}

// Metadata returns metadata about the connector.
func (d *Demo) Metadata(ctx context.Context) (*v2.ConnectorMetadata, error) {
	return &v2.ConnectorMetadata{
		DisplayName: "Demo",
		Description: "A demo connector",
	}, nil
}

// Validate is called to ensure that the connector is properly configured. It should exercise any API credentials
// to be sure that they are valid.
func (d *Demo) Validate(ctx context.Context) (annotations.Annotations, error) {
	return nil, nil
}

func (d *Demo) Close() error {
	return d.client.Close()
}

// New returns a new instance of the Demo connector.
func New(ctx context.Context, dbFileName string, initDB bool) (*Demo, error) {
	cli, err := client.NewClient(dbFileName, initDB)
	if err != nil {
		return nil, err
	}
	demo := &Demo{
		client: cli,
	}

	return demo, nil
}
