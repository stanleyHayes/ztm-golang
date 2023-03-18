package exercise

import "testing"

func TestAddEnergy(t *testing.T) {
	player := Player{Name: "Mario", Health: MaxHealth, Energy: MaxEnergy}

	actual := uint(100)
	expected := player.Energy
	if actual != expected {
		t.Errorf("Expected %v but received %v", actual, expected)
	}

	player.AddEnergy(120)
	actual = uint(100)
	expected = player.Energy

	if actual != expected {
		t.Errorf("Expected %v but received %v", actual, expected)
	}
}

func TestReduceEnergy(t *testing.T) {
	player := Player{Name: "Mario", Health: MaxHealth, Energy: MaxEnergy}

	actual := uint(100)
	expected := player.Energy
	if actual != expected {
		t.Errorf("Expected %v but received %v", actual, expected)
	}

	player.ReduceEnergy(40)
	actual = uint(60)
	expected = player.Energy

	if actual != expected {
		t.Errorf("Expected %v but received %v", actual, expected)
	}
}

func TestReduceHealth(t *testing.T) {
	player := Player{Name: "Mario", Health: MaxHealth, Energy: MaxEnergy}

	actual := uint(100)
	expected := player.Health
	if actual != expected {
		t.Errorf("Expected %v but received %v", actual, expected)
	}

	player.ReduceHealth(40)
	actual = uint(60)
	expected = player.Health

	if actual != expected {
		t.Errorf("Expected %v but received %v", actual, expected)
	}
}

func TestAddHealth(t *testing.T) {
	player := Player{Name: "Mario", Health: MaxHealth, Energy: MaxEnergy}

	actual := uint(100)
	expected := player.Energy
	if actual != expected {
		t.Errorf("Expected %v but received %v", actual, expected)
	}

	player.AddHealth(120)
	actual = uint(100)
	expected = player.Health

	if actual != expected {
		t.Errorf("Expected %v but received %v", actual, expected)
	}
}
