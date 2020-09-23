// Lute - 一款对中文语境优化的 Markdown 引擎，支持 Go 和 JavaScript
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
	"testing"

	"github.com/88250/lute"
)

var spinVditorIRBlockDOMTests = []*parseTest{

	{"22", "<p data-block=\"0\" data-node-id=\"20200923134039-7dpbj1g\" data-type=\"p\">foo<span data-type=\"inline-node\" class=\"vditor-ir__node\"><span color=\"red\"><span class=\"vditor-ir__marker\">&lt;span color=\"red\"&gt;</span></span>bar</span><span data-type=\"inline-node\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker\">&lt;/span&gt;<wbr></span></span></p>", "<p data-block=\"0\" data-node-id=\"20200923134039-7dpbj1g\" data-type=\"p\">foo<span color=\"red\"><span data-type=\"inline-node\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker\">&lt;span color=&quot;red&quot;&gt;</span></span>bar</span><span data-type=\"inline-node\" class=\"vditor-ir__node vditor-ir__node--expand\"><span class=\"vditor-ir__marker\">&lt;/span&gt;</span></span><wbr></p>"},
	{"21", "<p data-block=\"0\" data-node-id=\"20200923142109-gavrnnd\" data-type=\"p\"><span data-type=\"inline-node\" class=\"vditor-ir__node\"><code class=\"vditor-ir__marker\">&lt;font&gt;</code></span>bar&lt;/font&gt;<wbr></p>", "<p data-block=\"0\" data-node-id=\"20200923142109-gavrnnd\" data-type=\"p\"><font><span data-type=\"inline-node\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker\">&lt;font&gt;</span></span>bar</font><span data-type=\"inline-node\" class=\"vditor-ir__node vditor-ir__node--expand\"><span class=\"vditor-ir__marker\">&lt;/font&gt;</span></span><wbr></p>"},
	{"20", "<p data-block=\"0\" data-node-id=\"20200923134039-7dpbj1g\" data-type=\"p\">foo<span data-type=\"html-inline\"><code class=\"vditor-ir__marker\">&lt;br&gt;</code></span>bar<wbr></p>", "<p data-block=\"0\" data-node-id=\"20200923134039-7dpbj1g\" data-type=\"p\">foo<span data-type=\"html-inline\"><span class=\"vditor-ir__marker vditor-ir__br\">&lt;br&gt;</span></span>bar<wbr></p>"},
	{"19", "<div data-block=\"0\" data-node-id=\"20200923125212-ar1r1h2\" data-type=\"code-block\" class=\"vditor-ir__node\"><wbr><br><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code>\n</code></pre><pre class=\"vditor-ir__preview\" data-render=\"1\"><div class=\"vditor-copy\"><textarea></textarea><span aria-label=\"复制\" onmouseover=\"this.setAttribute('aria-label', '复制')\" class=\"b3-tooltips b3-tooltips__w\" onclick=\"this.previousElementSibling.select();document.execCommand('copy');this.setAttribute('aria-label', '已复制')\"><svg><use xlink:href=\"#iconCopy\"></use></svg></span></div><code class=\"hljs\"></code></pre><span data-type=\"code-block-close-marker\">```</span></div>", "<p data-block=\"0\" data-node-id=\"20200923125212-ar1r1h2\" data-type=\"p\"><wbr></p>"},
	{"18", "<p data-block=\"0\" data-node-id=\"20200923122329-80i6cd5\" data-type=\"p\">&lt;!--<wbr></p>", "<div data-block=\"0\" data-node-id=\"20200923122329-80i6cd5\" data-type=\"html-block\"><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code data-type=\"html-block\">&lt;!--<wbr></code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><!--</pre></div>"},
	{"17", "<p data-block=\"0\" data-node-id=\"20200922082313-rs3d9fh\" data-type=\"p\"><span data-type=\"mark\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker\">==</span><mark data-newline=\"1\">foo</mark><span class=\"vditor-ir__marker\">==\nb<wbr></span></span></p>", "<p data-block=\"0\" data-node-id=\"20200922082313-rs3d9fh\" data-type=\"p\"><span data-type=\"mark\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker\">==</span><mark data-newline=\"1\">foo</mark><span class=\"vditor-ir__marker\">==</span></span>\nb<wbr></p>"},
	{"16", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20200921114017-e8lmtfv\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"\">foo<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20200921114021-wesqaq8\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"\">bar</li><ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20200921114050-wu4fp2r\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"\">baz</li></ul><li><wbr><br></li></ul></li></ul>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20200921114017-e8lmtfv\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"\">foo<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20200921114050-wu4fp2r\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"\">bar<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"\">baz</li></ul></li></ul><ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20200921114021-wesqaq8\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"\"><wbr></li></ul></li></ul>"},
	{"15", "<p data-block=\"0\" data-node-id=\"20200921085450-aw0ce3a\" data-type=\"p\">!((20200920232128-n6x9jj7<wbr>))</p>", "<div data-block=\"0\" data-node-id=\"20200921085450-aw0ce3a\" data-type=\"block-ref-embed\" data-text=\"0\" class=\"vditor-ir__node vditor-ir__node--expand\"><span class=\"vditor-ir__marker\">!</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200920232128-n6x9jj7<wbr></span><span class=\"vditor-ir__marker\"> </span><span class=\"vditor-ir__marker\">\"</span><span class=\"vditor-ir__blockref\"></span><span class=\"vditor-ir__marker\">\"</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><div data-block-def-id=\"20200920232128-n6x9jj7<wbr>\" data-render=\"2\" data-type=\"block-render\"></div></div>"},
	{"14", "<p data-block=\"0\" data-node-id=\"20200921084558-fwzglr4\" data-type=\"p\">!((20200920232128-n6x9jj7))<wbr></p>", "<div data-block=\"0\" data-node-id=\"20200921084558-fwzglr4\" data-type=\"block-ref-embed\" data-text=\"0\" class=\"vditor-ir__node vditor-ir__node--expand\"><span class=\"vditor-ir__marker\">!</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200920232128-n6x9jj7</span><span class=\"vditor-ir__marker\"> </span><span class=\"vditor-ir__marker\">\"</span><span class=\"vditor-ir__blockref\"></span><span class=\"vditor-ir__marker\">\"</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><wbr><div data-block-def-id=\"20200920232128-n6x9jj7\" data-render=\"2\" data-type=\"block-render\"></div></div>"},
	{"13", "<p data-block=\"0\" data-node-id=\"20200921084143-fakzf28\" data-type=\"p\">​!<span data-type=\"block-ref\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200920232128-n6x9jj7</span><span class=\"vditor-ir__marker\"> </span><span class=\"vditor-ir__marker\">\"<wbr></span><span class=\"vditor-ir__marker\">\"</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span></p>", "<div data-block=\"0\" data-node-id=\"20200921084143-fakzf28\" data-type=\"block-ref-embed\" data-text=\"0\" class=\"vditor-ir__node vditor-ir__node--expand\"><span class=\"vditor-ir__marker\">!</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200920232128-n6x9jj7</span><span class=\"vditor-ir__marker\"> </span><span class=\"vditor-ir__marker\">\"</span><span class=\"vditor-ir__blockref\"><wbr></span><span class=\"vditor-ir__marker\">\"</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><div data-block-def-id=\"20200920232128-n6x9jj7\" data-render=\"2\" data-type=\"block-render\"></div></div>"},
	{"12", "<div data-block=\"0\" data-node-id=\"\" data-type=\"block-ref-embed\" data-text=\"0\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker\">!</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">id</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><div data-block-def-id=\"id\" data-render=\"2\" data-type=\"block-render\"></div></div>", "<div data-block=\"0\" data-node-id=\"\" data-type=\"block-ref-embed\" data-text=\"0\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker\">!</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">id</span><span class=\"vditor-ir__marker\"> </span><span class=\"vditor-ir__marker\">\"</span><span class=\"vditor-ir__blockref\"></span><span class=\"vditor-ir__marker\">\"</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><div data-block-def-id=\"id\" data-render=\"2\" data-type=\"block-render\"></div></div>"},
	{"11", "<p data-node-id=\"20200920115443-34yds2f\" data-block=\"0\">​\n```\n$$\n\\boldsymbol{r}\n$$\n```<wbr></p>", "<div data-block=\"0\" data-node-id=\"20200920115443-34yds2f\" data-type=\"code-block\" class=\"vditor-ir__node vditor-ir__node--expand\"><span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">\u200b</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code>$$\n\\boldsymbol{r}\n$$<wbr>\n</code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code>$$\n\\boldsymbol{r}\n$$</code></pre><span data-type=\"code-block-close-marker\">```</span></div>"},
	{"10", "<p data-block=\"0\" data-node-id=\"20200919110303-ki7u3yc\" data-type=\"p\">&lt;div&gt;<wbr></p>", "<div data-block=\"0\" data-node-id=\"20200919110303-ki7u3yc\" data-type=\"html-block\"><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code data-type=\"html-block\">&lt;div&gt;<wbr></code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><div></pre></div>"},
	{"9", "<div data-block=\"0\" data-node-id=\"20200919105707-ofvauia\" data-type=\"code-block\" class=\"vditor-ir__node\"><span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">​f<wbr></span><pre class=\"vditor-ir__marker--pre\"><code>\n</code></pre><span data-type=\"code-block-close-marker\">```</span></div>", "<div data-block=\"0\" data-node-id=\"20200919105707-ofvauia\" data-type=\"code-block\" class=\"vditor-ir__node vditor-ir__node--expand\"><span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">\u200bf<wbr></span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code class=\"language-f\">\n</code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code class=\"language-f\">\n</code></pre><span data-type=\"code-block-close-marker\">```</span></div>"},
	{"8", "<ul data-marker=\"*\" data-block=\"0\" data-node-id=\"20200919102109-s3a5xwo\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"\"><p data-block=\"0\" data-node-id=\"20200919105157-ucv97d9\" data-type=\"p\">foo</p><div data-block=\"0\" data-node-id=\"20200919105319-ie5iyfq\" data-type=\"math-block\" class=\"vditor-ir__node\"><span data-type=\"math-block-open-marker\">$$</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code data-type=\"math-block\" class=\"language-math\">bar<wbr></code></pre></div></li></ul>", "<ul data-marker=\"*\" data-block=\"0\" data-node-id=\"20200919102109-s3a5xwo\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"\"><p data-block=\"0\" data-node-id=\"20200919105157-ucv97d9\" data-type=\"p\">foo</p><div data-block=\"0\" data-node-id=\"20200919105319-ie5iyfq\" data-type=\"math-block\" class=\"vditor-ir__node vditor-ir__node--expand\"><span data-type=\"math-block-open-marker\">$$</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code data-type=\"math-block\" class=\"language-math\">bar<wbr></code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code data-type=\"math-block\" class=\"language-math\">bar</code></pre><span data-type=\"math-block-close-marker\">$$</span></div></li></ul>"},
	{"7", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20200919102109-s3a5xwo\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"\">foo<div data-block=\"0\" data-node-id=\"20200919102144-q0l1bco\" data-type=\"code-block\" class=\"vditor-ir__node\"><span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">​</span><pre class=\"vditor-ir__marker--pre\"><code>bar<wbr></code></pre></div></li></ul>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20200919102109-s3a5xwo\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"\">foo<div data-block=\"0\" data-node-id=\"20200919102144-q0l1bco\" data-type=\"code-block\" class=\"vditor-ir__node vditor-ir__node--expand\"><span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">\u200b</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code>bar<wbr>\n</code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code>bar\n</code></pre><span data-type=\"code-block-close-marker\">```</span></div></li></ul>"},
	{"6", "<blockquote data-block=\"0\" data-node-id=\"20200919091557-m4ivx3j\" data-type=\"blockquote\"><p data-block=\"0\" data-node-id=\"20200919091706-6f175t8\" data-type=\"p\"><wbr><br></p></blockquote>", "<blockquote data-block=\"0\" data-node-id=\"20200919091557-m4ivx3j\" data-type=\"blockquote\"><p data-block=\"0\" data-node-id=\"20200919091706-6f175t8\" data-type=\"p\"><wbr></p></blockquote>"},
	{"5", "<p data-block=\"0\" data-node-id=\"20200916084414-0hk6l8v\" data-type=\"p\">![foo中文.png](assets/foo中文<wbr>.png)</p>", "<p data-block=\"0\" data-node-id=\"20200916084414-0hk6l8v\" data-type=\"p\"><span class=\"vditor-ir__node vditor-ir__node--expand\" data-type=\"img\"><span class=\"vditor-ir__marker\">!</span><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">[</span><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">foo 中文.png</span><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">]</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">assets/foo中文<wbr>.png</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><img src=\" http://127.0.0.1:6807/webdav/q0fk7yv/测试笔记/assets/foo中文.png\" alt=\"foo 中文.png\" /></span></p>"},
	{"4", "<h1 data-block=\"0\" class=\"vditor-ir__node\" data-node-id=\"test\" bookmark=\"bookmark\" data-type=\"h\" id=\"ir-foo\" data-marker=\"#\"><span class=\"vditor-ir__marker vditor-ir__marker--heading\" data-type=\"heading-marker\"># </span>foo</h1>", "<h1 data-block=\"0\" class=\"vditor-ir__node\" data-node-id=\"test\" bookmark=\"bookmark\" data-type=\"h\" id=\"ir-foo\" data-marker=\"#\"><span class=\"vditor-ir__marker vditor-ir__marker--heading\" data-type=\"heading-marker\"># </span>foo</h1>"},
	{"3", "<div data-block=\"0\" data-node-id=\"test\" bookmark=\"bookmark\" data-type=\"code-block\" class=\"vditor-ir__node\"><span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">\u200b</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code>foo\n</code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code>foo\n</code></pre><span data-type=\"code-block-close-marker\">```</span></div>", "<div data-block=\"0\" data-node-id=\"test\" bookmark=\"bookmark\" data-type=\"code-block\" class=\"vditor-ir__node\"><span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">\u200b</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code>foo\n</code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code>foo\n</code></pre><span data-type=\"code-block-close-marker\">```</span></div>"},
	{"2", "<p data-block=\"0\" data-node-id=\"20200915173154-1wi2p2h\" data-type=\"p\">$$<wbr></p>", "<div data-block=\"0\" data-node-id=\"20200915173154-1wi2p2h\" data-type=\"math-block\" class=\"vditor-ir__node vditor-ir__node--expand\"><span data-type=\"math-block-open-marker\">$$</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code data-type=\"math-block\" class=\"language-math\"><wbr>\n</code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code data-type=\"math-block\" class=\"language-math\"></code></pre><span data-type=\"math-block-close-marker\">$$</span></div>"},
	{"1", "<p data-block=\"0\" data-node-id=\"20200915172226-iexs3bo\" data-type=\"p\">```<wbr></p>", "<div data-block=\"0\" data-node-id=\"20200915172226-iexs3bo\" data-type=\"code-block\" class=\"vditor-ir__node vditor-ir__node--expand\"><span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">\u200b<wbr></span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code>\n</code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code></code></pre><span data-type=\"code-block-close-marker\">```</span></div>"},
	{"0", "<p data-block=\"0\" data-node-id=\"20200914181352-laa3jyd\" data-type=\"p\">foo<wbr></p>", "<p data-block=\"0\" data-node-id=\"20200914181352-laa3jyd\" data-type=\"p\">foo<wbr></p>"},
}

func TestSpinVditorIRBlockDOM(t *testing.T) {
	luteEngine := lute.New()
	luteEngine.Mark = true
	luteEngine.BlockRef = true
	luteEngine.KramdownIAL = true
	luteEngine.SetLinkBase(" http://127.0.0.1:6807/webdav/q0fk7yv/测试笔记/")

	for _, test := range spinVditorIRBlockDOMTests {
		html := luteEngine.SpinVditorIRBlockDOM(test.from)

		if test.to != html {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal html\n\t%q", test.name, test.to, html, test.from)
		}
	}
}
