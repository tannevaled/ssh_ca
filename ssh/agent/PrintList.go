package SSHAgent

import (
	"fmt"
	"os"
	"strings"
	"time"

	strftime "github.com/itchyny/timefmt-go"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"golang.org/x/crypto/ssh"
)

/**
 *
 * Lists fingerprints of all identities currently represented by the agent.
 *
 */
func (self SSHAgentStruct) PrintList() {
	__key_array, err := self.agent.List()
	if err != nil {
		fmt.Println([]byte(err.Error()))
		os.Exit(1)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	//t.AppendHeader(table.Row{"AFTER", "BEFORE", "PRINCIPALS"})
	t.AppendHeader(table.Row{"VALID UNTIL", "ALLOWED PRINCIPALS"})
	//t.AppendRows([]table.Row{
	//	{1, "Arya", "Stark", 3000},
	//	{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
	//})
	//t.AppendSeparator()
	//t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
	//t.AppendFooter(table.Row{"", "", "Total", 10000})

	var __str_valid_before string
	for _, __ssh_key := range __key_array {
		/**
		 *
		 * x/crypto: document how to unmarshal ssh certs
		 *
		 * https://github.com/golang/go/issues/22046
		 *
		 * https://cs.opensource.google/go/x/crypto/+/refs/tags/v0.10.0:ssh/keys.go;l=271
		 *
		 */
		__public_key, err := ssh.ParsePublicKey(__ssh_key.Blob)
		if err != nil {
			fmt.Println([]byte(err.Error()))
			os.Exit(1)
		}
		if __ssh_certificate, ok := __public_key.(*ssh.Certificate); ok {
			__now := time.Now() //.Unix()
			__valid_before := time.Unix(int64(__ssh_certificate.ValidBefore), 0)
			if __valid_before.After(__now) {
				__str_valid_before = text.BgGreen.Sprint(
					strftime.Format(
						__valid_before,
						"%Y/%m/%d %H:%M:%S",
					),
				)
			} else {
				__str_valid_before = text.BgRed.Sprint(strftime.Format(__valid_before, "%Y/%m/%d %H:%M:%S"))
			}
			// strftime.Format(time.Unix(int64(__ssh_certificate.ValidAfter), 0), "%Y/%m/%d %H:%M:%S"),
			t.AppendRow(table.Row{
				__str_valid_before,
				strings.Join(__ssh_certificate.ValidPrincipals, ", "),
			})
			//str_valid_after := fmt.Sprintf(
			//	"%s",
			//	strftime.Format(time.Unix(int64(__ssh_certificate.ValidAfter), 0), "%Y/%m/%d %H:%M:%S"),
			//)
			//str_valid_before := fmt.Sprintf(
			//	"%s",
			//	strftime.Format(time.Unix(int64(__ssh_certificate.ValidBefore), 0), "%Y/%m/%d %H:%M:%S"),
			//)
			//str_principals := fmt.Sprintf("<Principals>%#v</Principals>", __ssh_certificate.ValidPrincipals)
			//str_extensions := fmt.Sprintf("<Extensions>%#v</Extensions>", __ssh_certificate.Extensions)
			//fmt.Printf(
			//	"<SSHCertificate>\n %s %s\n\t%s\n\t%s\n</SSHCertificate>\n",
			//	str_valid_after,
			//	str_valid_before,
			//	str_principals,
			//	str_extensions,
			//)
		}
	}
	t.Render()
}
