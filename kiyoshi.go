package main

import "sync"

const Zun = "ズン"
const Doko = "ドコ"
const Kiyoshi = "キ・ヨ・シ！"
const NG = "・・・"

var kmtx sync.Mutex

var kiyoshi map[string]int = make(map[string]int)

func Judge(from string, msg string) string {
	kmtx.Lock()
	defer kmtx.Unlock()

	ng := func() string {
		kiyoshi[from] = 0
		return NG
	}

	switch kiyoshi[from] {
	case 0, 1, 2, 3:
		if msg == Zun {
			kiyoshi[from]++
			return ""
		}
		return ng()
	case 4:
		if msg == Doko {
			kiyoshi[from] = 0
			return Kiyoshi
		}
		return ng()
	}
	return ""
}
