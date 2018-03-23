package asylum

import (
	"../../MongoData"
)

type JobRunStory struct {
	M     *mongo.MongoDB
	Token string
}

func (job *JobRunStory) Do() error {



	return nil
}
