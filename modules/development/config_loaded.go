package development

type ILoaded interface {
	AfterLoaded() error
}

type ILoadedPriority interface {
	AfterLoadedPriority() int
}

type afterLoadedPriority struct {
	Priority int
	Func     func() error
}

func getAfterLoaded(v configUnit) (afterLoadedPriority, bool) {
	loaded, ok := v.model.(ILoaded)
	if !ok {
		return afterLoadedPriority{}, false
	}

	priority := 0
	if priorityF, ok := v.model.(ILoadedPriority); ok {
		priority = priorityF.AfterLoadedPriority()
	}

	return afterLoadedPriority{
		Priority: priority,
		Func:     loaded.AfterLoaded,
	}, true
}
