package interfaces

type (
	KiwiDatalayer interface {
		SetValueWithKey(key string, value string) error
		GetValueByKey(key string) string
	}
)
