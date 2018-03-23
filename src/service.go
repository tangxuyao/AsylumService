package asylum

import (
	"../../MongoData"
	"proto/asylum"
	"golang.org/x/net/context"
	"github.com/tangxuyao/Frameworks"
)

type AsylumService struct {
	M *mongo.MongoDB
}

func (s *AsylumService) StartStory(c context.Context, req *asylum_api.StartStoryReq, rsp *asylum_api.StartStoryRsp) error {
	job := &JobStartStory{M: s.M, Token: req.Token}
	Frameworks.JobQueue <- job
	return nil
}

func (s *AsylumService) RunStory(ctx context.Context, req *asylum_api.RunStoryReq, rsp *asylum_api.RunStoryRsp) error {
	job := &JobRunStory{M: s.M, Token: req.Token}
	Frameworks.JobQueue <- job
	return nil
}

func (s *AsylumService) AsylumPing(ctx context.Context, in *asylum_api.AsylumPingReq, out *asylum_api.AsylumPingRsp) error {
	return nil
}
