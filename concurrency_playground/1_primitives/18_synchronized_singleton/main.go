package main

import "sync"

type Singleton struct{}

var once sync.Once
var instance *Singleton

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})

	return instance
}

// Lesson #4 video, time: 00:00:00
