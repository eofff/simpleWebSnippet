package services

import (
	"encoding/json"
	"gosample/redisApi"
	"gosample/repository"
	"strconv"
)

func SampleGet() ([](*repository.Sample), error) {
	return repository.SampleGetAll()
}

func SampleGetById(id int) (*repository.Sample, error) {
	e, err := redisApi.Exists(strconv.Itoa(id))

	if err != nil {
		return nil, err
	}
	if e {
		sampleStr, err := redisApi.Get(strconv.Itoa(id))

		if err != nil {
			return nil, err
		}

		var sample repository.Sample
		err = json.Unmarshal([]byte(sampleStr), &sample)

		if err != nil {
			return nil, err
		} else {
			return &sample, nil
		}
	} else {
		res, err := repository.SampleGetById(id)
		if err != nil {
			return nil, err
		}

		cache, err := json.Marshal(*res)
		if err != nil {
			return nil, err
		}

		err = redisApi.Set(strconv.Itoa(id), string(cache))
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func SampleInsert(sample *repository.Sample) error {
	return repository.SampleInsert(sample)
}
