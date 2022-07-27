package pubsub

type MessageHeaders map[string]string

func (h MessageHeaders) Add(key, value string) {
	h[key] = value
}
