---

# hCaptcha Solver

This is a Go package for solving hCaptcha challenges automaticly with https://hcoptcha.online. It allows you to create tasks to solve hCaptcha challenges and retrieve the solutions.

## Installation

To use this package, you need to have Go installed. You can install it using the following command:

```bash
go get github.com/CirtuitSavage/Hcoptcha-GO/hcaptcha
```

Replace `yourusername` with your GitHub username.

## Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/yourusername/hCaptcha/hcaptcha"
)

func main() {
	// Replace "yourAPIKey" with your actual hCaptcha API key
	apiKey := "yourAPIKey"

	// Replace these values with your proxy, site key, and optional rqdata
	proxy := "yourProxy"
	siteKey := "yourSiteKey"
	rqdata := "" // Set to your rqdata if needed
	url := "https://example.com" // Set the URL here

	// Create a new hCaptcha client
	client := hcaptcha.NewClient(apiKey)

	// Create a captcha solving task
	taskID, err := client.CreateTask(proxy, siteKey, rqdata, url)
	if err != nil {
		fmt.Println("Error creating task:", err)
		return
	}

	fmt.Println("Task created successfully. Task ID:", taskID)

	// Loop to continuously check the task status until it's no longer processing
	for {
		// Get task data
		taskData, err := client.GetTaskData(taskID)
		if err != nil {
			fmt.Println("Error getting task data:", err)
			return
		}

		// Check if the task is still processing
		state, ok := taskData["task"].(map[string]interface{})["state"].(string)
		if !ok {
			fmt.Println("Unable to get task state")
			return
		}
		if state != "processing" {
			// Task is no longer processing, retrieve the solution
			captchaKey, ok := taskData["task"].(map[string]interface{})["captcha_key"].(string)
			if !ok {
				fmt.Println("Unable to get captcha key")
				return
			}
			fmt.Println("Captcha solution:", captchaKey)
			return
		}

		// Sleep for a short duration before checking again
		time.Sleep(1 * time.Second)
	}
}
```

Replace `yourAPIKey`, `yourProxy`, `yourSiteKey` and `Url` with your actual hCaptcha API key, proxy, and site key.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
