package asylum

import (
	"../../MongoData"
	"proto/asylum"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
	"time"
	"log"
)

type AsylumService struct {
	M *mongo.MongoDB
}

func (s *AsylumService) TakeActor(c context.Context, req *asylum_api.TakeActorReq, rsp *asylum_api.TakeActorRsp) error {
	actor := mongo.Charactor{ID: bson.ObjectIdHex(req.Token)}
	if err := actor.RemoveByID(s.M, mongo.DB_GLOBAL); err != nil {
		log.Fatalf("从%s移除时发生异常: %s\n", mongo.DB_GLOBAL, err)
		return err
	}

	if err := actor.ToDB(s.M, mongo.DB_ASYLUM); err != nil {
		log.Fatalf("保存到%s出错: %s\n", mongo.DB_ASYLUM, err)
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
	log.Println("TakeActor Successful!!!")
	return nil
}

func (s *AsylumService) AsylumPing(ctx context.Context, in *asylum_api.AsylumPingReq, out *asylum_api.AsylumPingRsp) error {
	return nil
}
