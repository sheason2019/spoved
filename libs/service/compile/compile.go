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
		SetOutput("").
		AddOperator(usr).
		AddProject(proj).
		Save(ctx)

	go func() {
		output, err := CompileRun(image, nextVersion, branch, proj, usr.Username)

		if err != nil {
			statusCode = -1
		} else {
			statusCode = 1
		}

		client.CompileRecord.UpdateOne(record).
			SetStatusCode(statusCode).
			SetOutput(strings.Join(output, "\n")).
			SaveX(context.Background())
	}()

	return record, err
}
