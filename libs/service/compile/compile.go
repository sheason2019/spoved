package compile_service

import (
	"context"
	"strings"
	"time"

	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/libs/dbc"
)

func Compile(ctx context.Context, image, nextVersion, branch string, proj *ent.Project, usr *ent.User) (*ent.CompileRecord, error) {
	statusCode := 0

	client := dbc.GetClient()
	record, err := client.CompileRecord.Create().
		SetBranch(branch).
		SetCreatedAt(time.Now()).
		SetImage(image).
		SetStatusCode(statusCode).
		SetVersion(nextVersion).
		AddOperator(usr).
		AddProject(proj).
		Save(ctx)
	if err != nil {
		return record, err
	}

	output, err := CompileRun(image, nextVersion, branch, proj, usr.Username)

	if err != nil {
		statusCode = -1
	} else {
		statusCode = 1
	}

	record, err = client.CompileRecord.UpdateOne(record).
		SetStatusCode(statusCode).
		SetOutput(strings.Join(output, "\n")).
		Save(ctx)

	return record, err
}
