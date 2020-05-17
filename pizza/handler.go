package main

import (
	"github.com/botogonia/flowbot"
)

func chatHandler3(ch *flowbot.Chat) {

	defer ch.Close()

	_, _ = ch.WaitUpdate()

	chiceId1, data1 := ch.Choice(0,
	//_, data1 := ch.Choice(0,
		"Выберите кнопку",
		&flowbot.Kbrd{
			{{Text: "1", Data: "1"}, {Text: "2", Data: "2"}, {Text: "3", Data: "3"}},
			{{Text: "4", Data: "4"}, {Text: "5", Data: "5"}, {Text: "6", Data: "6"}},
		},
		"Ожидается нажатие соответствующей кнопки",
	)

	//ch.SendText("Вы выбрали: " + data1)

	//_, txt := ch.Prompt(0,  "Введите текст","Ожидается ввод текста")
	_, txt := ch.Prompt(chiceId1, "Вы выбрали: "+data1+"\nТеперь введите текст", "Ожидается ввод текста")
	//ch.SendMsg(clb.Message.MessageID, "Вы выбрали: "+clb.Data, nil)

	//log.Println(runtime.NumGoroutine())

	//txt := ch.WaitText("Ожидается ввод текста")

	ch.SendText("вы ввели: " + txt)

	ch.SendText("Чат завершен")

	//log.Println(runtime.NumGoroutine())

	//}
}
