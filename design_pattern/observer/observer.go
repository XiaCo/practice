package observer

type Observer interface {
	Update(subject Subject)
}

type Subject interface {
	Notice()
	GetID() string
}

type Publisher struct {
	observers []Observer
	ID        string
}

func (s Publisher) Notice() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

func (s Publisher) GetID() string {
	return s.ID
}

type Subscriber struct {
	ID string
}

func (s *Subscriber) Update(subject Subject) {
	s.ID = subject.GetID()
	return
}
