package di

type IStatRepository interface {
	AddClick(linkId uint) error
}
