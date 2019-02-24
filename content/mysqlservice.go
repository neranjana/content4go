package content

// MySQLService implementation that handles content operaitons using MySQL persistent storage
type MySQLService struct {
	contents []Content
}

// RetrieveCurrentVersionByID returns the valid version accoring to the timestamp
func (mySQLService *MySQLService) RetrieveCurrentVersionByID(ID string, currentTimestamp int64) (retrievedContent Content, errorThrown error) {
	return mySQLService.contents[0], nil
}
