package mp

import (
	"fmt"
)

type Mp struct {
	body map[string]int64
}

func NewMap() *Mp {
	return &Mp{body: nil}
}

func (m *Mp) Add(key string, data int64) {
	// добавить обработку ошибки с одинаковыми ключами
	if m.body == nil {
		m.body = make(map[string]int64)
	}
	m.body[key] = data
}

func (m *Mp) Print() {
	if m.body == nil {
		fmt.Println("no data")
		return
	}
	fmt.Println(m.body)
}

// Len возвращает длину списка
func (m *Mp) Len() (length int64) {
	if m.body == nil {
		fmt.Println("no data")
		return
	}
	length = int64(len(m.body))
	return length
}

// RemoveByIndex удаляет элемент из списка по индексу
func (m *Mp) RemoveByIndex(id string) {
	if m.body == nil {
		fmt.Println("no data")
		return
	}
	delete(m.body, id)
}

// RemoveByValue удаляет элемент из списка по значению
func (m *Mp) RemoveByValue(value int64) bool {
	if m.body == nil {
		fmt.Println("no data")
		return false
	}
	for key := range m.body {
		if m.body[key] == value {
			delete(m.body, key)
			return true
		}
	}
	return false
}

func (m *Mp) RemoveAllByValue(value int64) {
	if m.body == nil {
		fmt.Println("no data")
		return
	}
	for key := range m.body {
		if m.body[key] == value {
			delete(m.body, key)
		}
	}
}

func (m *Mp) GetByIndex(id string) (value int64, ok bool) {
	if m.body == nil {
		fmt.Println("no data")
		return 0, false
	}
	i, ok := m.body[id]
	return i, ok
}

func (m *Mp) GetByValue(value int64) (index string, ok bool) {
	if m.body == nil {
		fmt.Println("no data")
		return "", false
	}
	for key := range m.body {
		if m.body[key] == value {
			return key, true
		}
	}
	return "", false
}

func (m *Mp) GetAllByValue(value int64) (ids []string, ok bool) {
	if m.body == nil {
		fmt.Println("no data")
		return []string{}, false
	}
	for key := range m.body {
		if m.body[key] == value {
			ids = append(ids, key)
		}
	}
	if len(ids) > 0 {
		return ids, true
	}
	return []string{}, false
}

func (m *Mp) GetAll() (values []int64, ok bool) {
	if m.body == nil {
		fmt.Println("no data")
		return []int64{}, false
	}
	for key := range m.body {
		values = append(values, m.body[key])
	}
	if len(values) > 0 {
		return values, true
	}
	return []int64{}, false
}

// Clear очищает список
func (m *Mp) Clear() {
	m.body = nil
}

