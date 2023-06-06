package internal

type SearchService struct {
	Search    Search
	MovieRepo MovieRepo
	Signer    Signer
}

func NewSearchService(se Search, mr MovieRepo, si Signer) *SearchService {
	notNil(se, "SearchService.Search")
	notNil(mr, "SearchService.MovieRepo")
	notNil(si, "SearchService.Signer")
	return &SearchService{se, mr, si}
}

func (ss *SearchService) Resync(signature string) error {
	if err := ss.Signer.Compare("resync", signature); err != nil {
		return err
	}
	if err := ss.Search.Clear(); err != nil {
		return err
	}
	mch, err := ss.MovieRepo.GetAll()
	if err != nil {
		return err
	}
	for m := range mch {
		err := ss.Search.Insert(&SearchInsertItem{
			ID:       m.ID.String(),
			Title:    m.Title,
			Provider: m.Provider,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
