package watcher

import "testing"

func TestIsChangeRequired(t *testing.T) {
	// test the isChangeRequired function
	// if the current metrics value is greater than the threshold value
	// it should return true
	// if the current metrics value is less than the threshold value
	// it should return true
	// if the current metrics value is equal to the threshold value
	// it should return false
	change := isChangeRequired(0.9, 0.8)
	if change != true {
		t.Errorf("expected true, got %v", change)
	}
	change = isChangeRequired(0.7, 0.8)
	if change != true {
		t.Errorf("expected true, got %v", change)
	}
}

func TestCalculateDesiredReplica(t *testing.T) {
	// test the calculateDesiredReplica function
	// if the current metrics value is zero
	// it should return error
	// if the threshold value is zero
	// it should return error
	// if the current metrics value is greater than the threshold value
	// it should return the desired replica
	// if the current metrics value is less than the threshold value
	// it should return the desired replica
	desiredReplicas, err := calculateDesiredReplica(1, 0.8, 0.9)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if desiredReplicas != 2 {
		t.Errorf("expected 2, got %v", desiredReplicas)
	}
	desiredReplicas, err = calculateDesiredReplica(1, 0.8, 0.7)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if desiredReplicas != 1 {
		t.Errorf("expected 1, got %v", desiredReplicas)
	}
	_, err = calculateDesiredReplica(1, 0.8, 0)
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	_, err = calculateDesiredReplica(1, 0, 0.9)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
