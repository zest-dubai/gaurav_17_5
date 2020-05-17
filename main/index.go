package main

import(
	"fmt"
	"context"
	"main/kafka"
	"github.com/gin-gonic/gin"
)


func runProducer(){

	err := kafka.Push(context.Background(), nil, []byte("Transaction Done"))
	fmt.Println(err);
}

func runConnection(){

	kafkaBrokersUrls:=[]string{"localhost:19092","localhost:29092" , "localhost:39092" }
	var clientId string ="first_consumer"
	var foo string = "foo"
	fmt.Println(kafkaBrokersUrls);
	var w,error = kafka.Configure(kafkaBrokersUrls,clientId,foo);
	
	fmt.Println(w,error);
	//runProducer();
  
}

type Request_body struct{
		
		 Request_id string `json: request_id`
		 Topic_name string `json:topic_name`
		 Message_body string `json:message_body`
		 Transaction_id string `json:transaction_id`
		 Email string `json:email`
		 Phone string `json:phone`
		 Customer_id string `json:Customer_id`	
		 Key string `json:key`

}


func main(){

	router := gin.Default()

	router.POST("/done", func(c *gin.Context) {
	  
		var body Request_body
		c.BindJSON(&body)
          


        if body.Message_body=="successfull" {
			
			c.JSON(200, gin.H{
				"status":  "Transaction Successful",
				
			})

		   runConnection()

		}else{
			    c.JSON(200, gin.H{
				"Status":"Your transaction failed",
			  })
		}

	})

	router.Run(":8080")

}
