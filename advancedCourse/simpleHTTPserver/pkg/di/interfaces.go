package di

// IStatRepository описывает запись перехода по ID ссылки.
type IStatRepository interface {
	AddClick(linkId uint)
}
