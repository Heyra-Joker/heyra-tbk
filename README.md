# heyra-tbk

### download
`go get github.com/Heyra-Joker/heyra-tbk`

### have old version, update
`go get github.com/Heyra-Joker/heyra-tbk@releases`

#### like
`go get github.com/Heyra-Joker/heyra-tbk@v1.0.0`

### Usage

```go
package main

import (
	"fmt"
	"github.com/Heyra-Joker/heyra-tbk/api"
)

func main() {
	
	// init
	var rest = api.Rest{}
	var req = api.TbkDgOptimusPromotionRequest{}
	
	rest.RestUrl = "xxx"
	rest.AppSecret = "xxx"
	rest.AppKey = "xxxx"
	
	// rest.Other = map[string]string{"xxx":"xxx"} if public params not in Rest.
	
	req.AdzoneId = "xxxx"
	req.PromotionId = "xxxx"
	
	// get response
	data, err := req.GetResponse(rest)
	
	fmt.Println(data, err)
}
```




### Methods

- tbkCouponGetRequest.go
- tbkDgOptimusMaterialRequest.go
- tbkDgVegasTljInstanceReportRequest.go
- tbkShopRecommendGetRequest.go
- juItemsSearchRequest.go
- tbkDgMaterialOptionalRequest.go
- tbkDgOptimusPromotionRequest.go
- tbkItemInfoGetRequest.go
- tbkTpwdCreateRequest.go
- tbkActivityInfoGetRequest.go
- tbkDgOptimusMaterial.go
- tbkDgVegasTljCreateRequest.go
- tbkShopGetRequest.go

