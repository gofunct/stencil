package cloud

import (
	"context"
	"github.com/gofunct/stencil/app"
	"github.com/gofunct/stencil/app/server"
	"time"

	awsclient "github.com/aws/aws-sdk-go/aws/client"
	"github.com/google/wire"
	"gocloud.dev/aws/awscloud"
	"gocloud.dev/blob"
	"gocloud.dev/blob/s3blob"
	"gocloud.dev/mysql/rdsmysql"
	"gocloud.dev/runtimevar"
	"gocloud.dev/runtimevar/paramstore"
)

// setupAWS is a Wire injector function that sets up the application using AWS.
func setupAWS(ctx context.Context, f ...app.AppFunc) (*Api, func(), error) {
	// This will be filled in by Wire with providers from the provider sets in
	// wire.Build.
	wire.Build(
		app.NewApp,
		awscloud.AWS,
		rdsmysql.Open,
		server.NewServer,
		awsBucket,
		awsMOTDVar,
		awsSQLParams,
	)
	return nil, nil, nil
}

// awsBucket is a Wire provider function that returns the S3 bucket based on the
// command-line flags.
func awsBucket(ctx context.Context, cp awsclient.ConfigProvider, a *app.App) (*blob.Bucket, error) {
	return s3blob.OpenBucket(ctx, cp, a.Values.DbHost, nil)
}

// awsSQLParams is a Wire provider function that returns the RDS SQL connection
// parameters based on the command-line flags. Other providers inside
// awscloud.AWS use the parameters to construct a *sql.DB.
func awsSQLParams(a *app.App) *rdsmysql.Params {
	return &rdsmysql.Params{
		Endpoint: a.Values.DbHost,
		Database: a.Values.DbName,
		User:     a.Values.DbUser,
		Password: a.Values.DbPassword,
	}
}

// awsMOTDVar is a Wire provider function that returns the Message of the Day
// variable from SSM Parameter Store.
func awsMOTDVar(ctx context.Context, sess awsclient.ConfigProvider, a *app.App) (*runtimevar.Variable, error) {
	dur, _ := time.ParseDuration(a.Values.RunVarWait)
	return paramstore.NewVariable(sess, a.Values.State.RunVarName.Value, runtimevar.StringDecoder, &paramstore.Options{
		WaitDuration: dur,
	})
}
