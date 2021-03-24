package adapter

type ArchiveAdpater interface {
	ProvideServices() []*ArchiveService
}
