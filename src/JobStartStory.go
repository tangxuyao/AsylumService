package asylum

import (
	"../../MongoData"
	"gopkg.in/mgo.v2/bson"
	"time"
	"log"
)

type JobStartStory struct {
	M     *mongo.MongoDB
	Token string
}

func (h *JobStartStory) Do() error {
	actor := mongo.Charactor{
		ID: bson.ObjectIdHex(h.Token),
	}

	if err := actor.UpdatePlace(h.M, mongo.PLACE_ASYLUM); err != nil {
		log.Fatalf("更新Actor.Place属性失败:%s", err)
		return err
	}

	actorID := actor.ID.Hex()
	now := time.Now()

	t := mongo.Todo{
		ID:         bson.NewObjectId(),
		ActorToken: actorID,
		Place:      mongo.PLACE_ASYLUM,
		Task:       0,
		CreateTime: now,
		StartTime:  now,
		Duration:   time.Minute,
	}

	t.ToDB(h.M)
	return nil
}
