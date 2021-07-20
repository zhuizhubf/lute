// Lute - 一款结构化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package test

import (
	"github.com/88250/lute/ast"
	"testing"

	"github.com/88250/lute"
)

var md2BlockDOMTests = []parseTest{

	{"14", "foo<span data-type=\"vlink\">bar</span>baz\n", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">foo<span data-type=\"vlink\">bar</span>baz</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"13", "1. foo\n\n    1. bar\n", "<div data-subtype=\"o\" data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeList\" class=\"list\"><div data-marker=\"1.\" data-subtype=\"o\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeListItem\" class=\"li\"><div class=\"protyle-action protyle-action--order\" contenteditable=\"false\" draggable=\"true\">1.</div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">foo</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-subtype=\"o\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeList\" class=\"list\"><div data-marker=\"1.\" data-subtype=\"o\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeListItem\" class=\"li\"><div class=\"protyle-action protyle-action--order\" contenteditable=\"false\" draggable=\"true\">1.</div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">bar</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"12", "![foo](bar)", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\"><span contenteditable=\"false\" data-type=\"img\" class=\"img\"><span> </span><span><span class=\"protyle-action protyle-icons\"><span><svg class=\"svg\"><use xlink:href=\"#iconMore\"></use></svg></span></span><img src=\"/siyuan/0/测试笔记/bar\" data-src=\"bar\" alt=\"foo\" /><span class=\"protyle-action__drag\"></span><span class=\"protyle-action__title\"></span></span><span> </span></span></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"11", "| col1 | col2 | col3 |\n| ------ | ------ | ------ |\n|      |      |      |\n|      |      |      |\n\n<video controls=\"controls\" src=\"assets/RPReplay_Final16206653001-20210513170900-i8k37se.mp4\" data-src=\"assets/RPReplay_Final16206653001-20210513170900-i8k37se.mp4\"></video>\n\n11 ((20210513163831-8k81fw8))\n\n{{select * from blocks where id='20210513163831-8k81fw8'}}\n```js\nvar a = 1\n```\n\n![CleanShot_2021-05-07_at_18.30.11.gif](assets/CleanShot_2021-05-07_at_18.30.11-20210513163234-77jawpg.gif \"title\")\n\n<audio controls=\"controls\" src=\"assets/record1620894762009-20210513163242-toqc85e.wav\" data-src=\"assets/record1620894762009-20210513163242-toqc85e.wav\"></audio>\n", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeTable\" class=\"table\"><div contenteditable=\"true\" spellcheck=\"false\"><table><thead><tr><th>col1</th><th>col2</th><th>col3</th></tr></thead><tbody><tr><td></td><td></td><td></td></tr><tr><td></td><td></td><td></td></tr></tbody></table></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"2\" data-type=\"NodeVideo\" class=\"iframe\"><div class=\"iframe-content\"><video controls=\"controls\" src=\"/siyuan/0/测试笔记/assets/RPReplay_Final16206653001-20210513170900-i8k37se.mp4\" data-src=\"assets/RPReplay_Final16206653001-20210513170900-i8k37se.mp4\"></video><span class=\"protyle-action__drag\" contenteditable=\"false\"></span>\u200b</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"3\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">11 <span data-type=\"block-ref\" data-id=\"20210513163831-8k81fw8\" data-anchor=\"\"></span></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-content=\"select * from blocks where id='20210513163831-8k81fw8'\" data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"4\" data-type=\"NodeBlockQueryEmbed\" class=\"render-node\"><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"5\" data-type=\"NodeCodeBlock\" class=\"code-block\"><div class=\"protyle-action protyle-icons\"><span class=\"protyle-action__language\" contenteditable=\"false\">js</span><span class=\"protyle-action__copy\"></span></div><div contenteditable=\"true\" spellcheck=\"false\">var a = 1\n</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"6\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\"><span contenteditable=\"false\" data-type=\"img\" class=\"img\"><span> </span><span><span class=\"protyle-action protyle-icons\"><span><svg class=\"svg\"><use xlink:href=\"#iconMore\"></use></svg></span></span><img src=\"/siyuan/0/测试笔记/assets/CleanShot_2021-05-07_at_18.30.11-20210513163234-77jawpg.gif\" data-src=\"assets/CleanShot_2021-05-07_at_18.30.11-20210513163234-77jawpg.gif\" alt=\"CleanShot_2021-05-07_at_18.30.11.gif\" title=\"title\" /><span class=\"protyle-action__drag\"></span><span class=\"protyle-action__title\">title</span></span><span> </span></span></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"7\" data-type=\"NodeAudio\" class=\"iframe\"><div class=\"iframe-content\"><audio controls=\"controls\" src=\"/siyuan/0/测试笔记/assets/record1620894762009-20210513163242-toqc85e.wav\" data-src=\"assets/record1620894762009-20210513163242-toqc85e.wav\"></audio>\u200b</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"10", "((20210510191408-b2n8h2c \"foo\"))", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\"><span data-type=\"block-ref\" data-id=\"20210510191408-b2n8h2c\" data-anchor=\"foo\">foo</span></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"9", "{{{\n\nfoo\n\n{{{\n\n* bar\n\n  {{{\n  \n  bar2\n  \n  bar3\n  \n  }}}\n\nbar4\n\n}}}\n\nbaz\n\n}}}\n\nbazz\n\n", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeSuperBlock\" class=\"sb\" data-sb-layout=\"row\"><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">foo</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeSuperBlock\" class=\"sb\" data-sb-layout=\"row\"><div data-subtype=\"u\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeList\" class=\"list\"><div data-marker=\"*\" data-subtype=\"u\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeListItem\" class=\"li\"><div class=\"protyle-action\" draggable=\"true\"><svg><use xlink:href=\"#iconDot\"></use></svg></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">bar</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeSuperBlock\" class=\"sb\" data-sb-layout=\"row\"><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">bar2</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">bar3</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">bar4</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">baz</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"2\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">bazz</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"8", "{{{\n\nfoo\n\n{{{\n\n* bar \n\nba2\n\n}}}\n\nbaz\n\n}}}\n\nbazz\n\n", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeSuperBlock\" class=\"sb\" data-sb-layout=\"row\"><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">foo</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeSuperBlock\" class=\"sb\" data-sb-layout=\"row\"><div data-subtype=\"u\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeList\" class=\"list\"><div data-marker=\"*\" data-subtype=\"u\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeListItem\" class=\"li\"><div class=\"protyle-action\" draggable=\"true\"><svg><use xlink:href=\"#iconDot\"></use></svg></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">bar</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">ba2</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">baz</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"2\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">bazz</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"7", "{{{\n\n* baz\n\n}}}\n\nbazz\n\n", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeSuperBlock\" class=\"sb\" data-sb-layout=\"row\"><div data-subtype=\"u\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeList\" class=\"list\"><div data-marker=\"*\" data-subtype=\"u\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeListItem\" class=\"li\"><div class=\"protyle-action\" draggable=\"true\"><svg><use xlink:href=\"#iconDot\"></use></svg></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">baz</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"2\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">bazz</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"6", "{{{col\n\n{{{\nfoo\n\nbar\n}}}\n\nbaz\n\n}}}\n\nbazz\n\n", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeSuperBlock\" class=\"sb\" data-sb-layout=\"col\"><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeSuperBlock\" class=\"sb\" data-sb-layout=\"row\"><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">foo</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">bar</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">baz</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"2\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">bazz</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"5", "{{{col\n\n{{{\nfoo\n}}}\n\n}}}\n\nbar\n\n", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeSuperBlock\" class=\"sb\" data-sb-layout=\"col\"><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeSuperBlock\" class=\"sb\" data-sb-layout=\"row\"><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">foo</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"2\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">bar</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"4", "<u>foo</u>", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\"><u>foo</u></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"3", "* [x] foo", "<div data-subtype=\"t\" data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeList\" class=\"list\"><div data-marker=\"*\" data-subtype=\"t\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeListItem\" class=\"li protyle-task--done\"><div class=\"protyle-action protyle-action--task\"><svg><use xlink:href=\"#iconCheck\"></use></svg></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">foo</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"2", "{{select * from blocks}}", "<div data-content=\"select * from blocks\" data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeBlockQueryEmbed\" class=\"render-node\"><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"1", "<kbd>foo</kbd>", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\"><kbd>foo</kbd></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"0", "<audio src=\"assets/foo\"></audio>", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeAudio\" class=\"iframe\"><div class=\"iframe-content\"><audio src=\"/siyuan/0/测试笔记/assets/foo\" data-src=\"assets/foo\"></audio>\u200b</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
}

func TestMd2BlockDOM(t *testing.T) {
	luteEngine := lute.New()
	luteEngine.SetProtyleWYSIWYG(true)
	luteEngine.SetToC(true)
	luteEngine.ParseOptions.BlockRef = true
	luteEngine.ParseOptions.KramdownBlockIAL = true
	luteEngine.RenderOptions.KramdownBlockIAL = true
	luteEngine.ParseOptions.Tag = true
	luteEngine.ParseOptions.SuperBlock = true
	luteEngine.SetLinkBase("/siyuan/0/测试笔记/")
	luteEngine.ParseOptions.GitConflict = true
	luteEngine.ParseOptions.LinkRef = false
	luteEngine.SetAutoSpace(true)
	luteEngine.SetParagraphBeginningSpace(true)

	ast.Testing = true
	for _, test := range md2BlockDOMTests {
		result := luteEngine.Md2BlockDOM(test.from)
		if test.to != result {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal html\n\t%q", test.name, test.to, result, test.from)
		}
	}
	ast.Testing = false
}
