package book

import (
	"encoding/json"

	"github.com/go-redis/redis"
)

// RedisAdapter ...
type RedisAdapter struct {
	Client *redis.Client
}

// Find ...
func (r RedisAdapter) Find(id string) (Book, error) {
	var b Book
	v, err := r.Client.Get(id).Result()
	if err == redis.Nil {
		return b, nil
	} else if err != nil {
		return b, err
	}
	err = json.Unmarshal([]byte(v), &b)
	if err != nil {
		return b, err
	}
	return b, nil
}

// Create ...
func (r RedisAdapter) Create(b Book) error {
	data, err := json.Marshal(b)
	if err != nil {
		return err
	}
	err = r.Client.Set(b.ID, string(data), 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// Delete ...
func (r RedisAdapter) Delete(id string) error {
	err := r.Client.Del(id).Err()
	if err != nil {
		return err
	}
	return nil
}

// Update ...
func (r RedisAdapter) Update(id string, b Book) error {
	data, err := json.Marshal(b)
	if err != nil {
		return err
	}
	err = r.Client.Set(id, string(data), 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// NewRedisAdapter ...
func NewRedisAdapter(r *redis.Client) RedisAdapter {
	return RedisAdapter{Client: r}
}
