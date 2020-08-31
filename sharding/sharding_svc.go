package sharding

type Service struct {
}

func (svc *Service) ShardingItems() []int32 {
	return []int32{1, 2, 3, 4, 5, 6}
}
