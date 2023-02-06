package compile_service

import "github.com/pkg/errors"

func FindNextVersionForProject(projectId int, variant string) (string, error) {
	lastRecord, err := FindLastRecordByProjectId(projectId)
	if err != nil {
		return "", errors.WithStack(err)
	}

	currentVersion := "0.0.0"
	if lastRecord != nil {
		currentVersion = lastRecord.Version
	}

	nv, err := nextVersion(currentVersion, variant)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return nv, nil
}
