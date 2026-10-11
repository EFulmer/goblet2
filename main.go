package goblet2

import "errors"

func OneRepMax(weight, reps int) (float32, error) {
	// Epley
	if weight < 1 {
		return 0, errors.New("received non-positive weight value")
	} else if reps < 1 {
		return 0, errors.New("received non-positive rep count")
	}
	return float32(weight) * (1 + float32(reps)/30), nil
}

func StrengthWeightRatio(orm float32, weight int) (float32, error) {
	if weight < 1 {
		return 0, errors.New("received non-positive weight value")
	} else if orm <= 0 {
		return 0, errors.New("received non-positive ORM")
	}
	return orm / float32(weight), nil
}
