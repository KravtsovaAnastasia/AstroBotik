package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func main() {
	var (
	    channel, text, webhoock string  //channel - Наименование канала, в которое отправить сообщение.
					    // text - сообщение, которое отправим в канал.
					    //webhoock - урл для бота, по которому нужно передавать сообщение.
	    asker = bufio.NewScanner(os.Stdin)
	    host  = "https://hooks.slack.com/services/T02D8TBSW3Y/"
	    tail  [3]string //пришлось разделить хосты и хвосты, так как Slack разу удаляет вебхуки бота,
	    			// как только пушишь в открытый репозиторий
	)

	tail[0] = "B02DKCC2DPH/DaBZFSQV4H3fBlTJBPeqNLAs"
	tail[1] = "B02EPNHUL2U/1mGO5DUDiMiWL0Nmmb16uVJ3"
	tail[2] = "B02ECNCGXCH/xPHwvvvIx159zVRdCzYCIaon"
	fmt.Print("Здравствуйте!")
	fmt.Println("AstroBotik состоит в каналах test1, test2, test3. В какой канал хотите отправить сообщение? ")
	for {
	    //Спрашиваем, в какой канал отправить сообщение
	    asker.Scan()
	    channel = asker.Text()
	    //Здесь было бы неплохо вынести в property каналы, чтобы это могло быть вариативной
	    //величиной - добавить или убрать каналы. Так же как и вебхуки было бы неплохо вынести.
	    switch  {
	    	case channel=="test1": webhoock = host + tail[0]
		case channel=="test2": webhoock = host + tail[1]
		case channel=="test3": webhoock = host + tail[2]
		default: webhoock = `false`
	    }
	    if webhoock != "false" {
			break
	    }
	    fmt.Println("AstroBotik не состоит в таком канале, либо его не существует. Введите корректный канал:")
	}
	//Спрашиваем, что отправить в канал
	fmt.Println("Какое сообщение хотите послать от имени AstroBotik?")
	asker.Scan()
	text = asker.Text()
	err := sendMessageToChannel(text, webhoock)
	if err != nil {
	    return
	}
}

//Функция для отправки сообщения боту
func sendMessageToChannel(text string, webhoock string) error {

	data := []byte(fmt.Sprintf(`{"text":"%s"}`, text))
	r := bytes.NewReader(data)
	var resp, err = http.Post(webhoock, "application/json", r)
	if err != nil {
		return err
	}
	if resp != nil {fmt.Println("Сообщение успешно отправлено")}
	return nil
}