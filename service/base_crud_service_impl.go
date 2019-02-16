package service

import (
    . "github.com/Danceiny/dict-service/persistence/entity"
)


type BaseCrudServiceImpl struct {
	RepoServ RepositoryService
}

func (impl *BaseCrudServiceImpl) Delete(entity EntityIfc) {

}
func (impl *BaseCrudServiceImpl) Update(entity EntityIfc) {

}
func (impl *BaseCrudServiceImpl) Add(entity EntityIfc) {
	impl.RepoServ.Add(entity)
}
func (impl *BaseCrudServiceImpl) MultiGet(t DictTypeEnum, bids []BID, simple bool) []EntityIfc {
	return nil
}
func (impl *BaseCrudServiceImpl) Get(t DictTypeEnum, bid BID) EntityIfc {
	return nil
}
func (impl *BaseCrudServiceImpl) GetEntity(t DictTypeEnum, bid BID,
	simple bool, skipCache bool, withTrashed bool) EntityIfc {
	return nil
}
func (impl *BaseCrudServiceImpl) Exist(t DictTypeEnum, bid BID) bool {
	return true
}
