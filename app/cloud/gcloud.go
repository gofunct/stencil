//+build wireinject

package cloud

import (
	"context"
	"github.com/gofunct/stencil/app"
	"github.com/gofunct/stencil/app/server"
	"time"

	"github.com/google/wire"
	"gocloud.dev/blob"
	"gocloud.dev/blob/gcsblob"
	"gocloud.dev/gcp"
	"gocloud.dev/gcp/gcpcloud"
	"gocloud.dev/mysql/cloudmysql"
	"gocloud.dev/runtimevar"
	"gocloud.dev/runtimevar/runtimeconfigurator"
	pb "google.golang.org/genproto/googleapis/cloud/runtimeconfig/v1beta1"
)

// This file wires the generic interfaces up to Google Cloud Platform (GCP). It
// won't be directly included in the final binary, since it includes a Wire
// injector template function (setupGCP), but the declarations will be copied
// into wire_gen.go when Wire is run.

// setupGCP is a Wire injector function that sets up the App using GCP.
func setupGCP(ctx context.Context, f ...app.AppFunc) (*Api, func(), error) {
	// This will be filled in by Wire with providers from the provider sets in
	// wire.Build.
	wire.Build(
		app.NewApp,
		gcpcloud.GCP,
		server.NewServer,
		cloudmysql.Open,
		gcpBucket,
		gcpMOTDVar,
		gcpSQLParams,
	)
	return nil, nil, nil
}

// gcpBucket is a Wire provider function that returns the GCS bucket based on
// the command-line a.
func gcpBucket(ctx context.Context, a *app.App, client *gcp.HTTPClient) (*blob.Bucket, error) {
	return gcsblob.OpenBucket(ctx, client, a.Values.Bucket, nil)
}

// gcpSQLParams is a Wire provider function that returns the Cloud SQL
// connection parameters based on the command-line a. Other providers inside
// gcpcloud.GCP use the parameters to construct a *sql.DB.
func gcpSQLParams(id gcp.ProjectID, a *app.App) *cloudmysql.Params {
	return &cloudmysql.Params{
		ProjectID: string(id),
		Region:    a.Values.State.DatabaseRegion.Value,
		Instance:  a.Values.State.DatabaseInstance.Value,
		Database:  a.Values.DbName,
		User:      a.Values.DbUser,
		Password:  a.Values.DbPassword,
	}
}

// gcpMOTDVar is a Wire provider function that returns the Message of the Day
// variable from Runtime Configurator.
func gcpMOTDVar(ctx context.Context, client pb.RuntimeConfigManagerClient, project gcp.ProjectID, a *app.App) (*runtimevar.Variable, func(), error) {
	name := runtimeconfigurator.ResourceName{
		ProjectID: string(project),
		Config:    a.Values.State.RUnVarConfig.Value,
		Variable:  a.Values.RunVar,
	}
	dur, _ := time.ParseDuration(a.Values.RunVarWait)
	v, err := runtimeconfigurator.NewVariable(client, name, runtimevar.StringDecoder, &runtimeconfigurator.Options{
		WaitDuration: dur,
	})
	if err != nil {
		return nil, nil, err
	}
	return v, func() { v.Close() }, nil
}
