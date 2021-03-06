package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
	"strings"
)

func main() {
	//Input data to process
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

	fmt.Print("Possui Cache? (s/n): ")
	V_Cache, _ := reader.ReadString('\n')

	fmt.Print("Possui Lambda? (s/n): ")
	V_Lambda, _ := reader.ReadString('\n')

	fmt.Print("Possui SQS? (s/n): ")
	V_SQS, _ := reader.ReadString('\n')

	fmt.Print("Possui SNS? (s/n): ")
	V_SNS, _ := reader.ReadString('\n')

	// Check Data
	V_Front = check_resposta(V_Front)
	V_API = check_resposta(V_API)
	V_EKS = check_resposta(V_EKS)
	V_ECS = check_resposta(V_ECS)
	V_RDS = check_resposta(V_RDS)
	V_DynamoDB = check_resposta(V_DynamoDB)
	V_Cache = check_resposta(V_Cache)
	V_Lambda = check_resposta(V_Lambda)
	V_SQS = check_resposta(V_SQS)
	V_SNS = check_resposta(V_SNS)
	

	// Docs links
	L_Front := "https://docs.aws.amazon.com/AmazonS3/latest/userguide/WebsiteHosting.html"
	L_S3 := "https://docs.aws.amazon.com/AmazonS3/latest/userguide/WebsiteHosting.html"
	L_API := "https://docs.aws.amazon.com/apigateway/latest/developerguide/index.html"
	L_EKS := "https://docs.aws.amazon.com/eks/latest/userguide/index.html"
	L_ECS := "https://docs.aws.amazon.com/AmazonECS/latest/developerguide/index.html"
	L_RDS := "https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/index.html"
	L_DynamoDB := "https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/index.html"
	L_Lambda := "https://docs.aws.amazon.com/lambda/latest/dg/index.html"
	L_SQS := "https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/welcome.html"
	L_SNS := "https://docs.aws.amazon.com/sns/latest/dg/index.html"
	L_Cache := "https://docs.aws.amazon.com/sns/latest/dg/index.html"

	// Image links
	IMG_Front := "https://brunorusso.com.br/icons/Arch_Amazon-CloudFront_32.png"
	IMG_S3 := "https://brunorusso.com.br/icons/Arch_Amazon-Simple-Storage-Service_32.png"
	IMG_API := "https://brunorusso.com.br/icons/"
	IMG_EKS := "https://brunorusso.com.br/icons/Arch_Amazon-Elastic-Container-Kubernetes_32.png"
	IMG_ECS := "https://brunorusso.com.br/icons/Arch_Amazon-Elastic-Container-Service_32.png"
	IMG_RDS := "https://brunorusso.com.br/icons/Arch_Amazon-RDS_32.png"
	IMG_DynamoDB := "https://brunorusso.com.br/icons/Arch_Amazon-DynamoDB_32.png"
	IMG_Lambda := "https://brunorusso.com.br/icons/Arch_AWS-Lambda_32.png"
	IMG_SQS := "https://brunorusso.com.br/icons/"
	IMG_SNS := "https://brunorusso.com.br/icons/"
	IMG_Cache := "https://brunorusso.com.br/icons/Arch_Amazon-ElastiCache_32.png"

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
   	w.WriteString(".table {background-color: #EAECEE; border: 1px solid; border-color: #566573; border-radius: 5px; border-spacing: 0;}")
   	w.WriteString(".par {background-color: #D4E6F1; text-align: center; color: black; border: 1px solid; border-color: #566573;}")
   	w.WriteString(".impar {background-color: #D1F2EB; text-align: center; color: black; border: 1px solid; border-color: #566573;}")
   	w.WriteString("a.ativo a.ativo:link, a.ativo:visited, a.ativo:hover, a.ativo:active {color: red;}")
   	w.WriteString("a.inativo a.inativo:link, a.inativo:visited, a.inativo:hover, a.inativo:active {color: black;}")
   	w.WriteString("</style>")
   	w.WriteString("</head>")
	//Body
	w.WriteString("<html><head><title>Cloud Wizard</title></head><body bgcolor=\"#E7E7E7\">")
	w.WriteString("<center><table border=\"0\" width=\"90%\"><tr><td width=\"25%\"><img src=\"https://github.com/brunorusso/cloud-wizard/blob/develop/img/Cloud-Wizard-Logo.png\"></td>")
	w.WriteString(fmt.Sprintf("<td><H1><center>Projeto: %s </H1></center></td></tr></table></center>", V_Projeto))    	
	w.WriteString("<br><br><hr>")
	//Content
	w.WriteString("<table class=\"table\" align=\"center\">")
	w.WriteString("<tr>")
	w.WriteString("<th width=\"20%\">Camada</th>")
	w.WriteString("<th width=\"80%\" colspan=\"3\">Recursos</th>")
	w.WriteString("</tr>")
    w.WriteString("<tr>")
	w.WriteString("<td div class=\"par\">CDN e Aceleração de Conteúdo</div></td>")
	w.WriteString(fmt.Sprintf("<td class=\"par\"> <div><a class=\"%s\" href=\"%s\" target=\"_blank\"><img src=\"%s\"> <br>CloudFront</a></div></td>", V_Front, L_Front, IMG_Front))
	w.WriteString(fmt.Sprintf("<td class=\"par\"> <div><a class=\"%s\" href=\"%s\" target=\"_blank\"><img src=\"%s\"><br>S3</a></div</td>", V_Front, L_S3, IMG_S3))
	w.WriteString(fmt.Sprintf("<td class=\"par\"> <div><a class=\"%s\" href=\"%s\" target=\"_blank\"><img src=\"%s\"><br>APIGateway</a></div></td>", V_API, L_API, IMG_API))
    w.WriteString("</tr>")
    w.WriteString("<tr>")
	w.WriteString("<td div class=\"impar\">Computação</div></td>")
	w.WriteString(fmt.Sprintf("<td class=\"impar\"> <div><a class=\"%s\" href=\"%s\" target=\"_blank\"><img src=\"%s\"><br>EKS</a></div</td>", V_EKS, L_EKS, IMG_EKS))
	w.WriteString(fmt.Sprintf("<td class=\"impar\"> <div><a class=\"%s\" href=\"%s\" target=\"_blank\"><img src=\"%s\"><br>ECS</a></div</td>", V_ECS, L_ECS, IMG_ECS))
	w.WriteString(fmt.Sprintf("<td class=\"impar\"> <div><a class=\"%s\" href=\"%s\" target=\"_blank\"><img src=\"%s\"><br>Lambda</a></div</td>", V_Lambda, L_Lambda, IMG_Lambda))
    w.WriteString("</tr>")
    w.WriteString("<tr>")
	w.WriteString("<td div class=\"par\">Banco de Dados e Cache</div></td>")
	w.WriteString(fmt.Sprintf("<td class=\"par\"> <div><a class=\"%s\" href=\"%s\" target=\"_blank\"><img src=\"%s\"><br>RDS</a></div</td>", V_RDS, L_RDS, IMG_RDS))
	w.WriteString(fmt.Sprintf("<td class=\"par\"> <div><a class=\"%s\" href=\"%s\" target=\"_blank\"><img src=\"%s\"><br>DynamoDB</a></div</td>", V_DynamoDB, L_DynamoDB, IMG_DynamoDB))
	w.WriteString(fmt.Sprintf("<td class=\"par\"> <div><a class=\"%s\" href=\"%s\" target=\"_blank\"><img src=\"%s\"><br>Cache</a></div</td>", V_Cache, L_Cache, IMG_Cache))
    w.WriteString("</tr>")
    w.WriteString("<tr>")
	w.WriteString("<td div class=\"impar\">Integração</div></td>")
	w.WriteString(fmt.Sprintf("<td class=\"impar\"> <div><a class=\"%s\" href=\"%s\" target=\"_blank\"><img src=\"%s\"><br>SQS</a></div</td>", V_SQS, L_SQS, IMG_SQS))
	w.WriteString(fmt.Sprintf("<td class=\"impar\"> <div><a class=\"%s\" href=\"%s\" target=\"_blank\"><img src=\"%s\"><br>SNS</a></div</td>", V_SNS, L_SNS, IMG_SNS))
	w.WriteString(fmt.Sprintf("<td class=\"impar\"> <div><a class=\"%s\" href=\"%s\" target=\"_blank\"><img src=\"%s\"></a></div</td>"))
    w.WriteString("</tr>")
    w.WriteString("<tr>")
	w.WriteString("<td div class=\"par\">Governança</div></td>")
	w.WriteString(fmt.Sprintf("<td class=\"par\"> <div><a class=\"%s\" href=\"%s\" target=\"_blank\"><img src=\"\"></a></div</td>"))
	w.WriteString(fmt.Sprintf("<td class=\"par\"> <div><a class=\"%s\" href=\"%s\" target=\"_blank\"><img src=\"\"></a></div</td>"))
	w.WriteString(fmt.Sprintf("<td class=\"par\"> <div><a class=\"%s\" href=\"%s\" target=\"_blank\"><img src=\"\"></a></div</td>"))
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


func check_resposta(x string) string {
	if strings.Contains(x, "s"){
		V_ativo := "ativo"
		return V_ativo
	}
	V_ativo := "inativo"
	return V_ativo
}