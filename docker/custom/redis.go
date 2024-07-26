package custom

type Redis struct{}

func (redis *Redis) Ports() map[string]string {
	return map[string]string{
		"6379": "6379",
	}
}

func (redis *Redis) Volumes() map[string]string {
	return nil
}

func (redis *Redis) Environments() map[string]string {
	return nil
}

func (redis *Redis) Others() []string {
	return nil
}
