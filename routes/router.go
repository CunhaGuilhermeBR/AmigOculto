package routes

import (
	"net/http"
	"src/amigOculto/controller"
)

func LoadRoutes() {
	//Rota padrão
	http.HandleFunc("/", controller.Index)
	//Rota que redireciona para a tela de criação
	http.HandleFunc("/create", controller.CreateOccultFriend)
	//Rota que salva o registro no banco de dados
	http.HandleFunc("/register", controller.Register)
	//Rota que redireciona para a tela de participação
	http.HandleFunc("/draw", controller.DrawFriend)
	//Rota que sortea alguém
	http.HandleFunc("/pick", controller.GetPick)
	//Rota que cofirma o sorteado
	http.HandleFunc("/confirme", controller.Confirme)
	http.HandleFunc("/faq", controller.Faq)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
}
