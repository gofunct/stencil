//+build wireinject

package app

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"go.opencensus.io/trace"
	"gocloud.dev/blob"
	"gocloud.dev/blob/fileblob"
	"gocloud.dev/requestlog"
	"gocloud.dev/runtimevar"
	"gocloud.dev/runtimevar/filevar"
	"gocloud.dev/server"
)

// This file wires the generic interfaces up to local implementations. It won't
// be directly included in the final binary, since it includes a Wire injector
// template function (setupLocal), but the declarations will be copied into
// wire_gen.go when Wire is run.

// setupLocal is a Wire injector function that sets up the App using
// local implementations.
func setupLocal(ctx context.Context, a *App) (*App, func(), error) {
	// This will be filled in by Wire with providers from the provider sets in
	// wire.Build.
	wire.Build(
		wire.InterfaceValue(new(requestlog.Logger), requestlog.Logger(nil)),
		wire.InterfaceValue(new(trace.Exporter), trace.Exporter(nil)),
		server.Set,
		applicationSet,
		dialLocalSQL,
		localBucket,
		localRuntimeVar,
	)
	return nil, nil, nil
}

// localBucket is a Wire provider function that returns a directory-based bucket
// based on the command-line a.
func localBucket(a *App) (*blob.Bucket, error) {
	return fileblob.OpenBucket(a.Values.Bucket, nil)
}

// dialLocalSQL is a Wire provider function that connects to a MySQL database
// (usually on localhost).
func dialLocalSQL(a *App) (*sql.DB, error) {

	cfg := &mysql.Config{
		User:   a.Values.Bucket,
		Passwd: a.Values.DbPassword,
		Net:    "tcp",
		Addr:   a.Values.DbHost,
		DBName: a.Values.DbName,
	}
	return sql.Open("mysql", cfg.FormatDSN())
}

// localRuntimeVar is a Wire provider function that returns the Message of the
// Day variable based on a local file.
func localRuntimeVar(a *App) (*runtimevar.Variable, func(), error) {
	dur, _ := time.ParseDuration(a.Values.RunVarWait)
	v, err := filevar.New(a.Values.RunVar, runtimevar.StringDecoder, &filevar.Options{
		WaitDuration: dur,
	})
	if err != nil {
		return nil, nil, err
	}
	return v, func() { v.Close() }, nil
}
