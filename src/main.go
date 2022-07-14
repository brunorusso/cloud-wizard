package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Cloud Wizard")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Informe o nome do Projeto: ")
	V_Projeto, _ := reader.ReadString('\n')

	fmt.Println("Para as peguntas a seguir, responda s para sim ou n para Não")

	fmt.Print("Possui Front com conteúdo estático? (s/n): ")
	V_Front, _ := reader.ReadString('\n')

	fmt.Print("Possui API? (s/n): ")
	V_API, _ := reader.ReadString('\n')

	fmt.Print("Possui EKS? (s/n): ")
	V_EKS, _ := reader.ReadString('\n')

	fmt.Print("Possui ECS? (s/n): ")
	V_ECS, _ := reader.ReadString('\n')

	fmt.Print("Possui RDS? (s/n): ")
	V_RDS, _ := reader.ReadString('\n')

	fmt.Print("Possui DynamoDB? (s/n): ")
	V_DynamoDB, _ := reader.ReadString('\n')

	fmt.Print("Possui Lambda? (s/n): ")
	V_Lambda, _ := reader.ReadString('\n')

	fmt.Print("Possui SQS? (s/n): ")
	V_SQS, _ := reader.ReadString('\n')

	fmt.Print("Possui SNS? (s/n): ")
	V_SNS, _ := reader.ReadString('\n')


	// Links
	L_Front := "https://docs.aws.amazon.com/AmazonS3/latest/userguide/WebsiteHosting.html"
	L_API := "https://docs.aws.amazon.com/apigateway/latest/developerguide/index.html"
	L_EKS := "https://docs.aws.amazon.com/eks/latest/userguide/index.html"
	L_ECS := "https://docs.aws.amazon.com/AmazonECS/latest/developerguide/index.html"
	L_RDS := "https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/index.html"
	L_DynamoDB := "https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/index.html"
	L_Lambda := "https://docs.aws.amazon.com/lambda/latest/dg/index.html"
	L_SQS := "https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/welcome.html"
	L_SNS := "https://docs.aws.amazon.com/sns/latest/dg/index.html"



	// Exibe na tela 
	fmt.Sprintf("<b>Projeto: %s</b><br>", V_Projeto)
	fmt.Sprintf("<br><b>Site Estático: %s - %s</b>", V_Front, L_Front)
	fmt.Sprintf("<br><b>API: %s - %s:</b>", V_API, L_API)
	fmt.Sprintf("<br><b>EKS: %s - %s: </b>", V_EKS, L_EKS)
	fmt.Sprintf("<br><b>ECS: %s - %s: </b>", V_ECS, L_ECS)
	fmt.Sprintf("<br><b>RDS: %s - %s: </b>", V_RDS, L_RDS)
	fmt.Sprintf("<br><b>DynamoDB: %s - %s: </b>", V_DynamoDB, L_DynamoDB)
	fmt.Sprintf("<br><b>Lambda: %s - %s: </b>", V_Lambda, L_Lambda)
	fmt.Sprintf("<br><b>SQS: %s - %s: </b>", V_SQS, L_SQS)
	fmt.Sprintf("<br><b>SNS: %s - %s: </b>", V_SNS, L_SNS)


	//Gera arquivo
	V_Arquivo := "Cloud-Wizard-" + V_Projeto + ".html" 
	f, err := os.Create(V_Arquivo)
    	check(err)

    	w := bufio.NewWriter(f)
	//Head
    	w.WriteString("<html><head><title>Cloud Wizard</title></head><body bgcolor=\"#E7E7E7\">")
    	w.WriteString("<center><H1>Cloud Wizard</H1></center><br><br>")
	//Body
    	w.WriteString("<img src=\"https://github.com/brunorusso/cloud-wizard/blob/develop/img/Cloud-Wizard-Logo.png\"><br>")
    	w.WriteString(fmt.Sprintf("<b>Projeto: %s </b><br>", V_Projeto))
    	w.WriteString("<br><br>")
    	w.WriteString("<br><br>")
    	w.WriteString("<br><br>")
	//Content
    	w.WriteString("<table border=\"1\">")
    	w.WriteString("<tr colspan=\"3\">")
	w.WriteString("<td> cel1 </td>")
	w.WriteString("<td> cel2 </td>")
	w.WriteString("<td> cel3 </td>")
    	w.WriteString("</tr>")
    	w.WriteString("</table>")


  	check(err)

    	w.WriteString("</body></html>")
    	check(err)

    	defer f.Close()
    	w.Flush()


}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

