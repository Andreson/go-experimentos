package memory

import (
	boot "br.com.hdi.cross.cache/boostrap"
	"br.com.hdi.cross.cache/cache/domain"
)

type MemoryProvider struct {
	ConneError error
}

func (mp *MemoryProvider) Get(key string) (resp domain.CacheResponse) {
	byteData, err := mp.Build().Get(key)
	if err != nil {
		resp.SetStatus(err)
		boot.Log.Debugf("Erro recuperar cache in memory: %s", resp)
		return
	}
	resp.SetData(string(byteData))
	boot.Log.Debugf("Ger data cache in memory: %s", resp)
	return

}
func (mp *MemoryProvider) Set(key string, data string) domain.CacheResponse {

	err := mp.Build().Set(key, []byte(data))
	resp := domain.CacheResponse{}

	if err != nil {
		resp.SetStatus(err)
		boot.Log.Debugf("Erro add data cache in memory: %s", resp)
	}
	return resp
}
func (mp *MemoryProvider) Del(key string) (resp domain.CacheResponse) {
	err := mp.Build().Delete(key)
	resp.SetStatus(err)
	if err != nil {
		boot.Log.Debugf("Erro ao deletar data cache in memory: %s", resp)
	}
	return
}
func (rp *MemoryProvider) Connect() domain.CacheConnection {

	return domain.CacheConnection{}
}

func (rp *MemoryProvider) Error() error {
	return nil
}
