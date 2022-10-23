package template

type ShardTemplateFactory interface {
	Ding(sa *DingSA, info *SendInfo, ops ...FactoryOption) DingSender
}

var _ ShardTemplateFactory = &shardTemplateFactory{}

type shardTemplateFactory struct {
	Title string
}

type FactoryOption func(*shardTemplateFactory) *shardTemplateFactory

func (s *shardTemplateFactory) Ding(sa *DingSA, info *SendInfo, ops ...FactoryOption) DingSender {
	sf := &shardTemplateFactory{}
	if len(ops) > 0 {
		for _, opt := range ops {
			sf = opt(s)
		}
	}
	return NewDingSender(sa, info, sf.Title)
}

func NewTemplateFactory(title string) ShardTemplateFactory {
	return &shardTemplateFactory{
		Title: title,
	}
}
