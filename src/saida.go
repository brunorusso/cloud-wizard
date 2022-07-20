package main

import (
    "os"
    "bufio"
    "fmt"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func main() {
    f, err := os.Create("Cloud-Wizard-Report.html")
    check(err)

    Projeto := "Cloud1"

    w := bufio.NewWriter(f)
    w.WriteString("<html><head><title>Cloud Wizard</title></head><body bgcolor=\"#E7E7E7\">")
    w.WriteString("<center><H1>Cloud Wizard</H1></center><br><br>")
    w.WriteString("<img src=\"https://github.com/brunorusso/cloud-wizard/blob/develop/img/Cloud-Wizard-Logo.png\"><br>")
    w.WriteString(fmt.Sprintf("<b>Projeto: %s </b><br>", Projeto))


    check(err)

    w.WriteString("</body></html>")
    check(err)

    defer f.Close()
    w.Flush()

}
