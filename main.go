package main

import (
	"context"
	"fmt"
	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
)

//var appID = "cli_a3d526d772bc1013"
//var appSecret = "MRAqULW5PfU1CjOklJ0TYekZHHaHBKTh"

//var client = lark.NewClient(appID, appSecret,
//	lark.WithLogLevel(larkcore.LogLevelDebug),
//	lark.WithReqTimeout(3*time.Second),
//	lark.WithEnableTokenCache(true),
//	lark.WithHelpdeskCredential("id", "token"),
//	lark.WithHttpClient(http.DefaultClient))

func main()  {
	// 发起请求
	//resp, err := client.Docx.Document.Create(context.Background(), larkdocx.NewCreateDocumentReqBuilder().
	//	Body(larkdocx.NewCreateDocumentReqBodyBuilder().
	//		FolderToken("token").
	//		Title("title").
	//		Build()).
	//	Build())

	//req := larkcontact.NewGetUserReqBuilder().UserId(`8g3fgebb`).UserIdType(`user_id`).Build()
	//
	//resp, err := client.Contact.User.Get(context.Background(), req)
	//
	////处理错误
	//if err != nil {
	//	// 处理err
	//	return
	//}
	//
	//// 服务端错误处理
	//if !resp.Success() {
	//	fmt.Println(resp.Code, resp.Msg, resp.RequestId())
	//	return
	//}
	//
	//// 业务数据处理
	//fmt.Println(larkcore.Prettify(resp.Data))

	client := lark.NewClient(appID, appSecret)
	// 创建请求对象
	req := larkcontact.NewGetUserReqBuilder().
		UserId("5fe5a4e9").
		UserIdType("user_id").
		Build()
	// 发起请求
	resp, err := client.Contact.User.Get(context.Background(), req)

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return
	}

	// 业务处理
	fmt.Println(larkcore.Prettify(resp))

}
