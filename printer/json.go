package printer

import (
	"encoding/json"
	"fmt"

	"github.com/wata727/tflint/issue"
)

func (p *Printer) JSONPrint(issues []*issue.Issue) {
	result, err := json.Marshal(issues)
	if err != nil {
		fmt.Fprint(p.stderr, err)
	}
	fmt.Fprint(p.stdout, string(result))
}
