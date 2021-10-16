package _0211014_20211017

import (
	"encoding/json"
	_05_CDS "github.com/alex-shpak/hugo-book/homework/20211014-20211017/105-CDS"
	"io/ioutil"
	"sort"
	"strings"
	"testing"
)

func TestIP(t *testing.T) {
	ipData, _ := ioutil.ReadFile("ip.txt")
	data := strings.Split(string(ipData), "\r\n")
	in := make([]string, 0, len(data))
	out := make([][]string, 0, len(data))
	for _, v := range data {
		tmp := strings.Split(v, " ")
		if len(tmp) == 2 {
			in = append(in, tmp[0])
			var arr []string
			_ = json.Unmarshal([]byte(tmp[1]), &arr)
			out = append(out, arr)
		}
	}

	t.Run("105 CDS", func(t *testing.T) {
		registry := _05_CDS.NewIPV4()
		for i := 0; i < len(in); i++ {
			res := registry.IPFilter(in[i])
			if len(res) != len(out[i]) {
				t.Fatalf("in=%s out=%s got=%s", in[i], out[i], res)
			}
			sort.Strings(res)
			sort.Strings(out[i])
			for j, v := range res {
				if v != out[i][j] {
					t.Fatalf("in=%s out=%s got=%s", in[i], out[i], registry.IPFilter(in[i]))
				}
			}

		}
	})
}
