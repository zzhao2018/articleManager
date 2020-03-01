package wxutil

import (
	"log"
	"sync"
)

type SafeMap struct {
	loginMap map[string]bool
	mapLock sync.RWMutex
}

func NewSafeMap()(*SafeMap){
	return &SafeMap{
		loginMap: make(map[string]bool),
		mapLock:  sync.RWMutex{},
	}
}

func(s *SafeMap)Get(name string)(bool,bool){
	s.mapLock.RLock()
	defer s.mapLock.RUnlock()
	loginStatus, ok :=s.loginMap[name]
	return loginStatus,ok
}

func(s *SafeMap)Put(name string,value bool){
	s.mapLock.Lock()
	defer s.mapLock.Unlock()
	s.loginMap[name]=value
}

//定期清理
//此处可优化
func(s *SafeMap)ClearPair(){
	log.Printf("===========begin clear map:%+v===========\n",s.loginMap)
	newMap:=make(map[string]bool,len(s.loginMap))
	s.mapLock.RLock()
	for k,v:=range s.loginMap{
		newMap[k]=v
	}
	s.mapLock.RUnlock()
	s.loginMap=newMap
	log.Printf("++++++++++++++finish clear map:%+v+++++++++++++\n",s.loginMap)
}