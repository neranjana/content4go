package content

// Service that handles Content operations
type Service interface {
	RetrieveCurrentVersionByID(ID string, currentTimestamp int64) (retrievedContent Content, errorThrown error)
	RetrieveByIDAndVersion(ID string, version int) (retrievedContent Content, errorThrown error)
	RetrieveAllVersionsByID(ID string) (allVersions []Content, errorThrown error)
	Store(newContent Content) (version string, errorThrown error)
	DeleteVersion(ID string, versionToBeDeleted int) (deletedVersion string, errorThrown error)
	UpdateVersion(contentToBeUpdated Content) (version string, errorThrown error)
}
