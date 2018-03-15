package asylum

import (
	"proto/asylum"
	"golang.org/x/net/context"
	"../../MongoData"
	"gopkg.in/mgo.v2/bson"
	"time"
	"log"
)

type AsylumService struct {
	M *mongo.MongoDB
}

func (s *AsylumService) TakeActor(c context.Context, req *asylum_api.TakeActorReq, rsp *asylum_api.TakeActorRsp) error {
	log.Println("enter take actor scope..")
	actor := mongo.Charactor{ID: bson.ObjectId(req.Token)}
	if err := actor.RemoveByID(s.M, mongo.DB_GLOBAL); err != nil {
		return err
	}

	if err := actor.ToDB(s.M, mongo.DB_ASYLUM); err != nil {
		return err
	}

	actorID := string(actor.ID)

	t := mongo.Todo{
		ID:         bson.NewObjectId(),
		ActorToken: actorID,
		Place:      0,
		Task:       0,
		CreateTime: time.Now(),
		StartTime:  time.Now(),
		Duration:   time.Minute,
	}

	t.ToDB(s.M)
	return nil
}

func (s *AsylumService) AsylumPing(ctx context.Context, in *asylum_api.AsylumPingReq, out *asylum_api.AsylumPingRsp) error {
	return nil
}
