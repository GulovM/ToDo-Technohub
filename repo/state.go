package repo

var Users = make(map[string]User)

var Tasks = make(map[int]Task)

var UserTaskIDs = make(map[int]int)

var nextUserID = 0
var nextTaskID = 0
