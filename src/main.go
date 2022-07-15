package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
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

	// Create Variables
	V_Hoje := time.Now()

	//Gera arquivo
	V_Arquivo := "Cloud-Wizard-" + V_Projeto + ".html" 
	f, err := os.Create(V_Arquivo)
    	check(err)

    	w := bufio.NewWriter(f)
	//Head
    	w.WriteString("<html><head><title>Cloud Wizard</title></head>")
    	w.WriteString("<style>")
    	w.WriteString(".recurso {background-color: lightblue; text-align: center;}")
    	w.WriteString("</style>")
    	w.WriteString("</head>")
	//Body
		w.WriteString("<html><head><title>Cloud Wizard</title></head><body bgcolor=\"#E7E7E7\">")
		w.WriteString("<center><table border=\"0\" width=\"90%\"><tr><td width=\"25%\"><img src=\"https://github.com/brunorusso/cloud-wizard/blob/develop/img/Cloud-Wizard-Logo.png\"></td>")
		w.WriteString("<td><H1><center>Cloud Wizard</H1></center></td></tr></table></center>")
		w.WriteString("<br>")
    	w.WriteString(fmt.Sprintf("<H1>Projeto: <b> %s </b> </H1><br>", V_Projeto))
    	w.WriteString("<br><br><hr>")
	//Content
    	w.WriteString("<table border=\"1\" width=\"80%\">")
		w.WriteString("<tr>")
		w.WriteString("<th width=\"20%\">Camada</th>")
		w.WriteString("<th width=\"80%\" colspan=\"3\">Recursos</th>")
		w.WriteString("</tr>")
    	w.WriteString("<tr>")
		w.WriteString("<td>CDN e Aceleração de Conteúdo</td>")
		w.WriteString(fmt.Sprintf("<td> <div class=\"recurso\"><a href=\"%s\" target=\"_blank\"><img src=\"Arch_Amazon-CloudFront_32.png\"> <br>CloudFront</a></div></td>", L_Front))
		w.WriteString(fmt.Sprintf("<td> <div class=\"recurso\"><a href=\"%s\" target=\"_blank\"><img src=\"Arch_Amazon-Simple-Storage-Service_32.png\"><br>S3</a></div</td>"))
		w.WriteString(fmt.Sprintf("<td> <div class=\"recurso\"><a href=\"%s\" target=\"_blank\"><img src=\"\"><br>APIGateway</a></div></td>"))
    	w.WriteString("</tr>")
    	w.WriteString("<tr>")
		w.WriteString("<td>Computação</td>")
		w.WriteString(fmt.Sprintf("<td> <div class=\"recurso\"><a href=\"%s\" target=\"_blank\"><img src=\"Arch_Amazon-Elastic-Container-Kubernetes_32.png\"><br>EKS</a></div</td>"))
		w.WriteString(fmt.Sprintf("<td> <div class=\"recurso\"><a href=\"%s\" target=\"_blank\"><img src=\"Arch_Amazon-Elastic-Container-Service_32.png\"><br>ECS</a></div</td>"))
		w.WriteString(fmt.Sprintf("<td> <div class=\"recurso\"><a href=\"%s\" target=\"_blank\"><img src=\"Arch_AWS-Lambda_32.png\"><br>Lambda</a></div</td>"))
    	w.WriteString("</tr>")
    	w.WriteString("<tr>")
		w.WriteString("<td>Banco de Dados e Cache</td>")
		w.WriteString(fmt.Sprintf("<td> <div class=\"recurso\"><a href=\"%s\" target=\"_blank\"><img src=\"Arch_Amazon-RDS_32.png\"><br>RDS</a></div</td>"))
		w.WriteString(fmt.Sprintf("<td> <div class=\"recurso\"><a href=\"%s\" target=\"_blank\"><img src=\"Arch_Amazon-DynamoDB_32.png\"><br>DynamoDB</a></div</td>"))
		w.WriteString(fmt.Sprintf("<td> <div class=\"recurso\"><a href=\"%s\" target=\"_blank\"><img src=\"Arch_Amazon-ElastiCache_32.png\"><br>Cache</a></div</td>"))
    	w.WriteString("</tr>")
    	w.WriteString("<tr>")
		w.WriteString("<td>Integração</td>")
		w.WriteString(fmt.Sprintf("<td> <div class=\"recurso\"><a href=\"%s\" target=\"_blank\"><img src=\"\"></a></div</td>"))
		w.WriteString(fmt.Sprintf("<td> <div class=\"recurso\"><a href=\"%s\" target=\"_blank\"><img src=\"\"></a></div</td>"))
		w.WriteString(fmt.Sprintf("<td> <div class=\"recurso\"><a href=\"%s\" target=\"_blank\"><img src=\"\"></a></div</td>"))
    	w.WriteString("</tr>")
    	w.WriteString("<tr>")
		w.WriteString("<td>Governança</td>")
		w.WriteString(fmt.Sprintf("<td> <div class=\"recurso\"><a href=\"%s\" target=\"_blank\"><img src=\"\"></a></div</td>"))
		w.WriteString(fmt.Sprintf("<td> <div class=\"recurso\"><a href=\"%s\" target=\"_blank\"><img src=\"\"></a></div</td>"))
		w.WriteString(fmt.Sprintf("<td> <div class=\"recurso\"><a href=\"%s\" target=\"_blank\"><img src=\"\"></a></div</td>"))
    	w.WriteString("</tr>")
    	w.WriteString("</table>")

		//Footer
		w.WriteString(fmt.Sprintf("<center><h5>Wizard criado em: %s<h5></center>", V_Hoje))

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

