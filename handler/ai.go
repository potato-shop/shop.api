package handler

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/gin-gonic/gin"
)

type ProductAnalysis struct {
	Title       string `json:"title"`
	Price       string `json:"price"`
	Description string `json:"description"`
}

func AnalyzeImage(ctx *gin.Context) {
	client := anthropic.NewClient(
		option.WithAPIKey(os.Getenv("CLAUDE_API_KEY")),
	)

	file, err := ctx.FormFile("UploadedFile")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// 讀取檔案並轉成 base64
	src, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	defer src.Close()

	fileBytes, err := io.ReadAll(src)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	base64Str := base64.StdEncoding.EncodeToString(fileBytes)
	mediaType := http.DetectContentType(fileBytes)

	message, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		Model:     anthropic.ModelClaude_3_Haiku_20240307,
		MaxTokens: 1024,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(
				anthropic.NewImageBlockBase64(mediaType, base64Str),
				anthropic.NewTextBlock(`請分析這張商品圖片，回傳 JSON 格式：
{
  "title": "商品名稱(繁體中文)",
  "price": "建議售價(只要數字不要單位)",
  "description": "商品介紹(繁體中文)"
}
只回傳 JSON，不要其他文字。`),
			),
		},
	})
	if err != nil {
		panic(err)
	}

	// 顯示 token 用量
	fmt.Printf("輸入: %d tokens, 輸出: %d tokens\n",
		message.Usage.InputTokens, message.Usage.OutputTokens)

	// 截取 { ... } 之間的 JSON（避免 Claude 包在 markdown 代碼塊裡）
	text := message.Content[0].AsText().Text
	start := strings.Index(text, "{")
	end := strings.LastIndex(text, "}")
	if start == -1 || end == -1 {
		panic("回應中找不到 JSON: " + text)
	}
	text = text[start : end+1]

	// 解析結果
	var result ProductAnalysis
	if err := json.Unmarshal([]byte(text), &result); err != nil {
		panic(err)
	}

	log.Println("result: ", result)

	fmt.Printf("傷品名稱 %s\n", result.Description)
	fmt.Printf("售價: %s\n", result.Price)
	fmt.Printf("商品描述: %s\n", result.Description)

	ctx.JSON(http.StatusOK, result)
}
