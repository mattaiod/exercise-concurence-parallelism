package main

// Implémente un flag atomique
type AtomicFlag struct {
	state int32
}

func (f *AtomicFlag) SetTrue() {
	// Définit l'état à true
}

func (f *AtomicFlag) IsTrue() bool {
	// Vérifie si l'état est true
	return false
}
