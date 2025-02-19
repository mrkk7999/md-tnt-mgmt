package implementation

// HeartBeat
func (s *service) HeartBeat() map[string]string {
	return s.repo.HeartBeat()
}
