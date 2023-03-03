package key

import (
	"fmt"
	"io"
)

var (
	Nil               = byte(0x00)
	ReleaseAll  Input = Input{Nil, Nil, Nil, Nil, Nil, Nil, Nil, Nil}
	MuteKey     Key   = Key(0x7f)
	VoluemUpKey Key   = Key(0x7f)

	AKey     Key = Key(0x04) // Keyboard a and A
	BKey     Key = Key(0x05) // Keyboard b and B
	CKey     Key = Key(0x06) // Keyboard c and C
	DKey     Key = Key(0x07) // Keyboard d and D
	EKey     Key = Key(0x08) // Keyboard e and E
	FKey     Key = Key(0x09) // Keyboard f and F
	GKey     Key = Key(0x0a) // Keyboard g and G
	HKey     Key = Key(0x0b) // Keyboard h and H
	IKey     Key = Key(0x0c) // Keyboard i and I
	JKey     Key = Key(0x0d) // Keyboard j and J
	KKey     Key = Key(0x0e) // Keyboard k and K
	LKey     Key = Key(0x0f) // Keyboard l and L
	MKey     Key = Key(0x10) // Keyboard m and M
	NKey     Key = Key(0x11) // Keyboard n and N
	OKey     Key = Key(0x12) // Keyboard o and O
	PKey     Key = Key(0x13) // Keyboard p and P
	QKey     Key = Key(0x14) // Keyboard q and Q
	RKey     Key = Key(0x15) // Keyboard r and R
	SKey     Key = Key(0x16) // Keyboard s and S
	TKey     Key = Key(0x17) // Keyboard t and T
	UKey     Key = Key(0x18) // Keyboard u and U
	VKey     Key = Key(0x19) // Keyboard v and V
	WKey     Key = Key(0x1a) // Keyboard w and W
	XKey     Key = Key(0x1b) // Keyboard x and X
	YKey     Key = Key(0x1c) // Keyboard y and Y
	ZKey     Key = Key(0x1d) // Keyboard z and Z
	OneKey   Key = Key(0x1e) // Keyboard 1 and !
	TwoKey   Key = Key(0x1f) // Keyboard 2 and @
	ThreeKey Key = Key(0x20) // Keyboard 3 and #
	FourKey  Key = Key(0x21) // Keyboard 4 and $
	FiveKey  Key = Key(0x22) // Keyboard 5 and %
	SixKey   Key = Key(0x23) // Keyboard 6 and ^
	SevenKey Key = Key(0x24) // Keyboard 7 and &
	EightKey Key = Key(0x25) // Keyboard 8 and *
	NineKey  Key = Key(0x26) // Keyboard 9 and (
	ZeroKey  Key = Key(0x27) // Keyboard 0 and )

	ENTERKey      Key = Key(0x28) // Keyboard Return (ENTER)
	ESCKey        Key = Key(0x29) // Keyboard ESCAPE
	BACKSPACEKey  Key = Key(0x2a) // Keyboard DELETE (Backspace)
	TABKey        Key = Key(0x2b) // Keyboard Tab
	SPACEKey      Key = Key(0x2c) // Keyboard Spacebar
	MINUSKey      Key = Key(0x2d) // Keyboard - and _
	EQUALKey      Key = Key(0x2e) // Keyboard = and +
	LEFTBRACEKey  Key = Key(0x2f) // Keyboard [ and {
	RIGHTBRACEKey Key = Key(0x30) // Keyboard ] and }
	BACKSLASHKey  Key = Key(0x31) // Keyboard \ and |
	HASHTILDEKey  Key = Key(0x32) // Keyboard Non-US # and ~
	SEMICOLONKey  Key = Key(0x33) // Keyboard ; and :
	APOSTROPHEKey Key = Key(0x34) // Keyboard ' and "
	GRAVEKey      Key = Key(0x35) // Keyboard ` and ~
	COMMAKey      Key = Key(0x36) // Keyboard , and <
	DOTKey        Key = Key(0x37) // Keyboard . and >
	SLASHKey      Key = Key(0x38) // Keyboard / and ?
	CAPSLOCKKey   Key = Key(0x39) // Keyboard Caps Lock

	F1Key  Key = Key(0x3a) // Keyboard F1
	F2Key  Key = Key(0x3b) // Keyboard F2
	F3Key  Key = Key(0x3c) // Keyboard F3
	F4Key  Key = Key(0x3d) // Keyboard F4
	F5Key  Key = Key(0x3e) // Keyboard F5
	F6Key  Key = Key(0x3f) // Keyboard F6
	F7Key  Key = Key(0x40) // Keyboard F7
	F8Key  Key = Key(0x41) // Keyboard F8
	F9Key  Key = Key(0x42) // Keyboard F9
	F10Key Key = Key(0x43) // Keyboard F10
	F11Key Key = Key(0x44) // Keyboard F11
	F12Key Key = Key(0x45) // Keyboard F12

	SYSRQKey      Key = Key(0x46) // Keyboard Print Screen
	SCROLLLOCKKey Key = Key(0x47) // Keyboard Scroll Lock
	PAUSEKey      Key = Key(0x48) // Keyboard Pause
	INSERTKey     Key = Key(0x49) // Keyboard Insert
	HOMEKey       Key = Key(0x4a) // Keyboard Home
	PAGEUPKey     Key = Key(0x4b) // Keyboard Page Up
	DELETEKey     Key = Key(0x4c) // Keyboard Delete Forward
	ENDKey        Key = Key(0x4d) // Keyboard End
	PAGEDOWNKey   Key = Key(0x4e) // Keyboard Page Down
	RIGHTKey      Key = Key(0x4f) // Keyboard Right Arrow
	LEFTKey       Key = Key(0x50) // Keyboard Left Arrow
	DOWNKey       Key = Key(0x51) // Keyboard Down Arrow
	UPKey         Key = Key(0x52) // Keyboard Up Arrow

	NUMLOCKKey    Key = Key(0x53) // Keyboard Num Lock and Clear
	KPSLASHKey    Key = Key(0x54) // Keypad /
	KPASTERISKKey Key = Key(0x55) // Keypad *
	KPMINUSKey    Key = Key(0x56) // Keypad -
	KPPLUSKey     Key = Key(0x57) // Keypad +
	KPENTERKey    Key = Key(0x58) // Keypad ENTER
	KP1Key        Key = Key(0x59) // Keypad 1 and End
	KP2Key        Key = Key(0x5a) // Keypad 2 and Down Arrow
	KP3Key        Key = Key(0x5b) // Keypad 3 and PageDn
	KP4Key        Key = Key(0x5c) // Keypad 4 and Left Arrow
	KP5Key        Key = Key(0x5d) // Keypad 5
	KP6Key        Key = Key(0x5e) // Keypad 6 and Right Arrow
	KP7Key        Key = Key(0x5f) // Keypad 7 and Home
	KP8Key        Key = Key(0x60) // Keypad 8 and Up Arrow
	KP9Key        Key = Key(0x61) // Keypad 9 and Page Up
	KP0Key        Key = Key(0x62) // Keypad 0 and Insert
	KPDOTKey      Key = Key(0x63) // Keypad . and Delete

	// 102NDKey Key = Key(0x64)// Keyboard Non-US \ and |
	COMPOSEKey Key = Key(0x65) // Keyboard Application
	POWERKey   Key = Key(0x66) // Keyboard Power
	KPEQUALKey Key = Key(0x67) // Keypad =

	F13Key Key = Key(0x68) // Keyboard F13
	F14Key Key = Key(0x69) // Keyboard F14
	F15Key Key = Key(0x6a) // Keyboard F15
	F16Key Key = Key(0x6b) // Keyboard F16
	F17Key Key = Key(0x6c) // Keyboard F17
	F18Key Key = Key(0x6d) // Keyboard F18
	F19Key Key = Key(0x6e) // Keyboard F19
	F20Key Key = Key(0x6f) // Keyboard F20
	F21Key Key = Key(0x70) // Keyboard F21
	F22Key Key = Key(0x71) // Keyboard F22
	F23Key Key = Key(0x72) // Keyboard F23
	F24Key Key = Key(0x73) // Keyboard F24

	OPENKey       Key = Key(0x74) // Keyboard Execute
	HELPKey       Key = Key(0x75) // Keyboard Help
	PROPSKey      Key = Key(0x76) // Keyboard Menu
	FRONTKey      Key = Key(0x77) // Keyboard Select
	STOPKey       Key = Key(0x78) // Keyboard Stop
	AGAINKey      Key = Key(0x79) // Keyboard Again
	UNDOKey       Key = Key(0x7a) // Keyboard Undo
	CUTKey        Key = Key(0x7b) // Keyboard Cut
	COPYKey       Key = Key(0x7c) // Keyboard Copy
	PASTEKey      Key = Key(0x7d) // Keyboard Paste
	FINDKey       Key = Key(0x7e) // Keyboard Find
	MUTEKey       Key = Key(0x7f) // Keyboard Mute
	VOLUMEUPKey   Key = Key(0x80) // Keyboard Volume Up
	VOLUMEDOWNKey Key = Key(0x81) // Keyboard Volume Down
	// 0x82  Keyboard Locking Caps Lock
	// 0x83  Keyboard Locking Num Lock
	// 0x84  Keyboard Locking Scroll Lock
	KPCOMMAKey Key = Key(0x85) // Keypad Comma
	// 0x86  Keypad Equal Sign
	ROKey               Key = Key(0x87) // Keyboard International1
	KATAKANAHIRAGANAKey Key = Key(0x88) // Keyboard International2
	YENKey              Key = Key(0x89) // Keyboard International3
	HENKANKey           Key = Key(0x8a) // Keyboard International4
	MUHENKANKey         Key = Key(0x8b) // Keyboard International5
	KPJPCOMMAKey        Key = Key(0x8c) // Keyboard International6
	// 0x8d  Keyboard International7
	// 0x8e  Keyboard International8
	// 0x8f  Keyboard International9
	HANGEULKey        Key = Key(0x90) // Keyboard LANG1
	HANJAKey          Key = Key(0x91) // Keyboard LANG2
	KATAKANAKey       Key = Key(0x92) // Keyboard LANG3
	HIRAGANAKey       Key = Key(0x93) // Keyboard LANG4
	ZENKAKUHANKAKUKey Key = Key(0x94) // Keyboard LANG5
	// 0x95  Keyboard LANG6
	// 0x96  Keyboard LANG7
	// 0x97  Keyboard LANG8
	// 0x98  Keyboard LANG9
	// 0x99  Keyboard Alternate Erase
	// 0x9a  Keyboard SysReq/Attention
	// 0x9b  Keyboard Cancel
	// 0x9c  Keyboard Clear
	// 0x9d  Keyboard Prior
	// 0x9e  Keyboard Return
	// 0x9f  Keyboard Separator
	// 0xa0  Keyboard Out
	// 0xa1  Keyboard Oper
	// 0xa2  Keyboard Clear/Again
	// 0xa3  Keyboard CrSel/Props
	// 0xa4  Keyboard ExSel

	// 0xb0  Keypad 00
	// 0xb1  Keypad 000
	// 0xb2  Thousands Separator
	// 0xb3  Decimal Separator
	// 0xb4  Currency Unit
	// 0xb5  Currency Sub-unit
	KPLEFTPARENKey  Key = Key(0xb6) // Keypad (
	KPRIGHTPARENKey Key = Key(0xb7) // Keypad )
	// 0xb8  Keypad {
	// 0xb9  Keypad }
	// 0xba  Keypad Tab
	// 0xbb  Keypad Backspace
	// 0xbc  Keypad A
	// 0xbd  Keypad B
	// 0xbe  Keypad C
	// 0xbf  Keypad D
	// 0xc0  Keypad E
	// 0xc1  Keypad F
	// 0xc2  Keypad XOR
	// 0xc3  Keypad ^
	// 0xc4  Keypad %
	// 0xc5  Keypad <
	// 0xc6  Keypad >
	// 0xc7  Keypad &
	// 0xc8  Keypad &&
	// 0xc9  Keypad |
	// 0xca  Keypad ||
	// 0xcb  Keypad :
	// 0xcc  Keypad #
	// 0xcd  Keypad Space
	// 0xce  Keypad @
	// 0xcf  Keypad !
	// 0xd0  Keypad Memory Store
	// 0xd1  Keypad Memory Recall
	// 0xd2  Keypad Memory Clear
	// 0xd3  Keypad Memory Add
	// 0xd4  Keypad Memory Subtract
	// 0xd5  Keypad Memory Multiply
	// 0xd6  Keypad Memory Divide
	// 0xd7  Keypad +/-
	// 0xd8  Keypad Clear
	// 0xd9  Keypad Clear Entry
	// 0xda  Keypad Binary
	// 0xdb  Keypad Octal
	// 0xdc  Keypad Decimal
	// 0xdd  Keypad Hexadecimal

	LEFTCTRLKey   Key = Key(0xe0) // Keyboard Left Control
	LEFTSHIFTKey  Key = Key(0xe1) // Keyboard Left Shift
	LEFTALTKey    Key = Key(0xe2) // Keyboard Left Alt
	LEFTMETAKey   Key = Key(0xe3) // Keyboard Left GUI
	RIGHTCTRLKey  Key = Key(0xe4) // Keyboard Right Control
	RIGHTSHIFTKey Key = Key(0xe5) // Keyboard Right Shift
	RIGHTALTKey   Key = Key(0xe6) // Keyboard Right Alt
	RIGHTMETAKey  Key = Key(0xe7) // Keyboard Right GUI

	MEDIA_PLAYPAUSEKey    Key = Key(0xe8) //
	MEDIA_STOPCDKey       Key = Key(0xe9) //
	MEDIA_PREVIOUSSONGKey Key = Key(0xea) //
	MEDIA_NEXTSONGKey     Key = Key(0xeb) //
	MEDIA_EJECTCDKey      Key = Key(0xec) //
	MEDIA_VOLUMEUPKey     Key = Key(0xed) //
	MEDIA_VOLUMEDOWNKey   Key = Key(0xee) //
	MEDIA_MUTEKey         Key = Key(0xef) //
	MEDIA_WWWKey          Key = Key(0xf0) //
	MEDIA_BACKKey         Key = Key(0xf1) //
	MEDIA_FORWARDKey      Key = Key(0xf2) //
	MEDIA_STOPKey         Key = Key(0xf3) //
	MEDIA_FINDKey         Key = Key(0xf4) //
	MEDIA_SCROLLUPKey     Key = Key(0xf5) //
	MEDIA_SCROLLDOWNKey   Key = Key(0xf6) //
	MEDIA_EDITKey         Key = Key(0xf7) //
	MEDIA_SLEEPKey        Key = Key(0xf8) //
	MEDIA_COFFEEKey       Key = Key(0xf9) //
	MEDIA_REFRESHKey      Key = Key(0xfa) //
	MEDIA_CALCKey         Key = Key(0xfb) //

	// Modifiers
	NoMod  Modifier = Modifier(0x00)
	LCtrl  Modifier = Modifier(0x01)
	LShift Modifier = Modifier(0x02)
	LAlt   Modifier = Modifier(0x04)
	LMeta  Modifier = Modifier(0x08)
	RCtrl  Modifier = Modifier(0x10)
	RShift Modifier = Modifier(0x20)
	RAlt   Modifier = Modifier(0x40)
	RMeta  Modifier = Modifier(0x80)
)

type Key byte
type Modifier byte
type Input []byte

func GetReleaseAllKeys() Input {
	return ReleaseAll
}

func GetPressKey(key Key, mod Modifier) []Input {
	k := Input{byte(mod), Nil, byte(key), Nil, Nil, Nil, Nil, Nil}
	fmt.Printf("KEY:%X, %X -> %+v\n", key, mod, k)
	return []Input{k}
}

func GetPressAndReleaseKey(key Key, mod Modifier) []Input {
	k := GetPressKey(key, mod)
	return append(k, ReleaseAll)
}

func PressAndHoldKey(w io.Writer, key Key, mod Modifier) {
	RunInput(w, GetPressKey(MuteKey, NoMod))
}

func PressAndReleaseKey(w io.Writer, key Key, mod Modifier) {
	RunInput(w, GetPressAndReleaseKey(MuteKey, NoMod))
}

func ReleaseAllKeys(w io.Writer, key Key, mod Modifier) {
	RunInput(w, []Input{ReleaseAll})
}

// TODO: figure out how to sleep
func RunInput(w io.Writer, in []Input) {
	for _, i := range in {
		n, err := w.Write(i)
		if err != nil {
			fmt.Println("err:", err.Error())
		}
		fmt.Println("Write:", n)
	}
}
