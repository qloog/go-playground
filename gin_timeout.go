package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// timeout middleware wraps the request context with a timeout
func timeoutMiddleware(timeout time.Duration) func(c *gin.Context) {
	fmt.Println(1111)
	return func(c *gin.Context) {
		fmt.Println(c.FullPath())
		fmt.Println(3333)
		//是否超时，都会执行，进行收尾
		// wrap the request context with a timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		defer func() {
			// check if context timeout was reached
			if ctx.Err() == context.DeadlineExceeded {
				//这里超时，才会执行
				fmt.Println(4444)
				// write response and abort the request
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.Abort()
			}
		}()
		fmt.Println(5555)
		//是否超时，都会执行
		// replace request with context wrapped request
		c.Request = c.Request.WithContext(ctx)
		c.Next() //实际调用具体的handler处理业务，实际还在这个方法中，所以业务执行结束会回到中间件中执行中间件中的defer函数
	}
}

func timedHandler(duration time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {

		fmt.Println(6666)

		// get the underlying request context
		ctx := c.Request.Context()

		// create the response data type to use as a channel type
		type responseData struct {
			status int
			body   map[string]interface{}
		}

		// create a done channel to tell the request it's done
		doneChan := make(chan responseData)

		// here you put the actual work needed for the request
		// and then send the doneChan with the status and body
		// to finish the request by writing the response
		go func() {
			time.Sleep(duration)
			doneChan <- responseData{
				status: 200,
				body:   gin.H{"hello": "world"},
			}
		}()

		// non-blocking select on two channels see if the request
		// times out or finishes
		select {

		// if the context is done it timed out or was cancelled
		// so don't return anything
		case <-ctx.Done():
			return

			// if the request finished then finish the request by
			// writing the response
		case res := <-doneChan:
			c.JSON(res.status, res.body)
		}
	}
}

func main() {
	// create new gin without any middleware
	engine := gin.New()

	// add timeout middleware with 2 second duration
	engine.Use(timeoutMiddleware(time.Second * 2))

	// create a handler that will last 1 seconds
	engine.GET("/short", timedHandler(time.Second))
	//结果
	//3333
	//5555
	//22222
	// create a route that will last 5 seconds
	engine.GET("/long", timedHandler(time.Second*5))
	//3333
	//5555
	//4444
	//22222
	// run the server
	log.Fatal(engine.Run(":8089"))

}
