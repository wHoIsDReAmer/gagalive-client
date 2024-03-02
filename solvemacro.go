package client

import (
	"fmt"
	"strings"
)

type SolveMacro struct {
	macroMap map[string]int
}

func NewSolveMacro() *SolveMacro {
	sm := &SolveMacro{
		macroMap: make(map[string]int),
	}
	sm.setupMacroLists()
	return sm
}

func (sm *SolveMacro) setupMacroLists() {
	// setup with individual symbols
	sm.setup(strings.Split("⓪ ① ② ③ ④ ⑤ ⑥ ⑦ ⑧ ⑨", " "), 0)
	sm.setup(strings.Split("➀ ➁ ➂ ➃ ➄ ➅ ➆ ➇ ➈", " "), 1)
	sm.setup(strings.Split("➊ ➋ ➌ ➍ ➎ ➏ ➐ ➑ ➒", " "), 1)
	sm.setup(strings.Split("⓵ ⓶ ⓷ ⓸ ⓹ ⓺ ⓻ ⓼ ⓽", " "), 1)
	sm.setup(strings.Split("⓿ ❶ ❷ ❸ ❹ ❺ ❻ ❼ ❽ ❾", " "), 0)
	sm.setup(strings.Split("0⃝ 1⃝ 2⃝ 3⃝ 4⃝ 5⃝ 6⃝ 7⃝ 8⃝ 9⃝", " "), 0)
	sm.setup(strings.Split("⁰ ¹ ² ³ ⁴ ⁵ ⁶ ⁷ ⁸ ⁹", " "), 0)
	sm.setup(strings.Split("₀ ₁ ₂ ₃ ₄ ₅ ₆ ₇ ₈ ₉", " "), 0)
	sm.setup(strings.Split("⒈ ⒉ ⒊ ⒋ ⒌ ⒍ ⒎ ⒏ ⒐", " "), 1)
	sm.setup(strings.Split("⑴ ⑵ ⑶ ⑷ ⑸ ⑹ ⑺ ⑻ ⑼", " "), 1)
	sm.setup(strings.Split("０ １ ２ ３ ４ ５ ６ ７ ８ ９", " "), 0)
	sm.setup(strings.Split("영 일 이 삼 사 오 육 칠 팔 구", " "), 0)
	sm.setup(strings.Split("(0) (1) (2) (3) (4) (5) (6) (7) (8) (9)", " "), 0)
	sm.setup(strings.Split("(숫자ㅇㅕㅇ) (숫자ㅇㅣㄹ) (숫자ㅇㅣ) (숫자ㅅㅏㅁ) (숫자ㅅㅏ) (숫자ㅇㅗ) (숫자ㅇㅠㄱ) (숫자ㅊㅣㄹ) (숫자ㅍㅏㄹ) (숫자ㄱㅜ)", " "), 0)

	sm.setup(strings.Split("𝟎 𝟏 𝟐 𝟑 𝟒 𝟓 𝟔 𝟕 𝟖 𝟗", " "), 0)
	sm.setup(strings.Split("𝟘 𝟙 𝟚 𝟛 𝟜 𝟝 𝟞 𝟟 𝟠 𝟡", " "), 0)
	sm.setup(strings.Split("𝟬 𝟭 𝟮 𝟯 𝟰 𝟱 𝟲 𝟳 𝟴 𝟵", " "), 0)
	sm.setup(strings.Split("𝟶 𝟷 𝟸 𝟹 𝟺 𝟻 𝟼 𝟽 𝟾 𝟿", " "), 0)

	sm.macroMap["ᅵ"] = 1
	sm.macroMap["ㅣ"] = 1
	sm.macroMap["|"] = 1
	sm.macroMap["l"] = 1
	sm.macroMap["৪"] = 8
	sm.macroMap["О"] = 0
	sm.macroMap["o"] = 0
	sm.macroMap["ㅇ"] = 0
	sm.macroMap["੦"] = 0
	sm.macroMap["O"] = 0
}

func (sm *SolveMacro) setup(list []string, start int) {
	for _, symbol := range list {
		sm.macroMap[symbol] = start
		start++
	}
}

func (sm *SolveMacro) Solve(prb string) string {
	var sb strings.Builder
	index := 0
	for index < len(prb) {
		matched := false
		for key, value := range sm.macroMap {
			if strings.HasPrefix(prb[index:], key) {
				sb.WriteString(fmt.Sprintf("%d", value))
				index += len(key)
				matched = true
				break
			}
		}
		if !matched {
			sb.WriteByte(prb[index])
			index++
		}
	}
	return sb.String()
}
