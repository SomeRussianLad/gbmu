package ppu

import (
	"gbmu/emulation/controllers"
	"gbmu/emulation/memory"
)

//	Архитекутра и планирование структуры Fetcher
//	Он должен:
//		1. Искать по запросу разные типы тайлов (BG, Win, Obj)
//		2. Ре-инициализировать свое состояние на поиск другого объекта по первому запросу
//		3. Возвращать по update() массив из 8-и структур, хранящих ID цвета и константу палитры из LCD (BG, Win, Obj)
//
//	Методы:
//		init(frameX, frameY uint8, tileType int) - метод должен сообщать структуре, какой тип тайла ищется (BG, Win, Obj),
//			и координаты дисплея для того, чтобы рассчитать offset для TileMap
//

// TODO(somerussianlad): Проведи рефактор

const (
	FET_STATE_READ_TILE_ID = iota // all of the states take 2 dots to complete
	FET_STATE_READ_TILE_DATA_0
	FET_STATE_READ_TILE_DATA_1
	FET_STATE_PUSH // except for pushing as fetcher does it every 1 dot
)

const (
	TILE_TYPE_BG = iota
	TILE_TYPE_WIN
	TILE_TYPE_OBJ
)

type fetcher struct {
	memory memory.Memory
	p      pusher

	lcd *controllers.LCD

	cycles int
	state  int

	tileMapX uint8
	tileMapY uint8

	tileType  int
	tileID    uint8
	tileData0 uint8
	tileData1 uint8
	tileData  []uint8
}

func newFetcher(memory memory.Memory, lcd *controllers.LCD, p pusher) *fetcher {
	fetcher := &fetcher{
		memory: memory,
		p:      p,

		lcd: lcd,

		tileType: TILE_TYPE_BG,
		tileData: make([]uint8, 0, 8),
	}

	return fetcher
}

func (f *fetcher) init(frameX, frameY uint8, tileType int) {
	f.tileType = tileType
	f.tileData = f.tileData[:0]

	switch f.tileType {
	case TILE_TYPE_BG:
		f.tileMapX = ((frameX + f.lcd.GetSCX()) / 8) % 32
		f.tileMapY = ((frameY + f.lcd.GetSCY()) / 8) % 32

	case TILE_TYPE_WIN:
		f.tileMapX = ((frameX - f.lcd.GetWX() + 7) / 8) % 32
		f.tileMapY = ((frameY - f.lcd.GetWY()) / 8) % 32
	}

	f.cycles = 0
	f.state = FET_STATE_READ_TILE_ID
}

func (f *fetcher) nextTile() {
	f.tileMapX++
	if f.tileMapX > 31 {
		f.tileMapX = 0
		// f.tileMapY++
	}
	// if f.tileMapY > 31 {
	// 	f.tileMapY = 0
	// }
}

func (f *fetcher) performReadTileID() {
	var tileMapAddr uint16

	switch f.tileType {
	case TILE_TYPE_BG:
		tileMapAddr = f.lcd.GetBGTileMapAddr()

	case TILE_TYPE_WIN:
		tileMapAddr = f.lcd.GetWinTileMapAddr()
	}

	addr := tileMapAddr + uint16(f.tileMapX) + uint16(f.tileMapY)*32
	f.nextTile()

	f.tileID = f.memory.Read(addr)
	f.state = FET_STATE_READ_TILE_DATA_0
}

func (f *fetcher) performReadTileData0() {
	tileDataAddr, signedAddressing := f.lcd.GetBGAndWinTileDataAddr()

	if signedAddressing && f.tileID > 127 {
		tileDataAddr = memory.MEM_VRAM_TILE_DATA_1
		f.tileID -= 128
	}

	lineOffset := ((uint16(f.lcd.GetLY()) + uint16(f.lcd.GetSCY())) % 8) * 2
	addr := tileDataAddr + uint16(f.tileID)*16 + lineOffset

	f.tileData0 = f.memory.Read(addr)
	f.state = FET_STATE_READ_TILE_DATA_1
}

func (f *fetcher) performReadTileData1() {
	tileDataAddr, signedAddressing := f.lcd.GetBGAndWinTileDataAddr()

	if signedAddressing && f.tileID > 127 {
		tileDataAddr = memory.MEM_VRAM_TILE_DATA_1
		f.tileID -= 128
	}

	lineOffset := ((uint16(f.lcd.GetLY()) + uint16(f.lcd.GetSCY())) % 8) * 2
	addr := tileDataAddr + uint16(f.tileID)*16 + lineOffset + 1

	f.tileData1 = f.memory.Read(addr)
	f.state = FET_STATE_PUSH
}

func (f *fetcher) performPush() {
	if len(f.tileData) == 0 {
		for i := 0; i < 8; i++ {
			hi := (f.tileData1 >> (7 - i)) & 1
			lo := (f.tileData0 >> (7 - i)) & 1

			colorID := hi<<1 | lo

			f.tileData = append(f.tileData, colorID)
		}
	}

	if f.p.push(f.tileData) {
		f.tileData = f.tileData[:0]
		f.state = FET_STATE_READ_TILE_ID
	}
}

func (f *fetcher) update() {
	f.cycles += PPU_DOT_LEN

	if f.state != FET_STATE_PUSH && f.cycles < 2*PPU_DOT_LEN {
		return
	}
	f.cycles = 0

	switch f.state {
	case FET_STATE_READ_TILE_ID:
		f.performReadTileID()

	case FET_STATE_READ_TILE_DATA_0:
		f.performReadTileData0()

	case FET_STATE_READ_TILE_DATA_1:
		f.performReadTileData1()

	case FET_STATE_PUSH:
		f.performPush()
	}
}
