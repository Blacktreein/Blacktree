// entry point of worker 
package main

import (
	"log"
	"github.com/Blacktreein/Blacktree/build-worker/internal/utils"
)


func main(){
	log.Println("🔄 Starting build worker")
	
	// Load env variables
	if err := utils.EnvInit("./.env"); err != nil {
		log.Fatalf("❌ Failed to load env variables: %v", err)
	}

	log.Println()
	
	// listening to the port 

	
}