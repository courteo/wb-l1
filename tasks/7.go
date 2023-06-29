package tasks

import "sync"

type syncMap struct {
	m     map[int]int
	mutex sync.Mutex   				// добавили мьютекс
}

func (s *syncMap) Add(key, val int) {
	s.mutex.Lock()					// лоичим мьютекс, чтобы потокобезопасно записать значение
	s.m[key] = val
	s.mutex.Unlock()				// анлочим, чтобы другие горутины могли использовать
}

func (s *syncMap) Get(key int) int {
	s.mutex.Lock()					// лоичим мьютекс, чтобы потокобезопасно получить значение
	res := s.m[key]
	s.mutex.Unlock()				// анлочим, чтобы другие горутины могли использовать
	return res
}